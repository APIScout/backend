package controller

import (
	"backend/app/internal/elastic"
	"backend/app/internal/embedding"
	"backend/app/internal/models"
	"backend/app/internal/mongodb"
	"backend/app/internal/retrieval"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SearchHandler godoc
//
//	@Summary		Search OpenAPI specifications
//	@Description	Retrieve OpenAPI specifications matching the given query
//	@Tags			search
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int						false	"page number"		minimum(1)	default(1)
//	@Param			pageSize	query		int						false	"size of the page"	minimum(1)	maximum(100)	default(10)
//	@Param			k			query		int						false	"knn's k"			minimum(1)	maximum(100)	default(100)
//	@Param			fragment	body		models.EmbeddingRequest	true	"search query"
//	@Success		200			{object}	[]models.SpecificationWithApi
//	@Failure		400			{object}	models.HTTPError
//	@Failure		500			{object}	models.HTTPError
//	@Router			/search [post]
func SearchHandler(mongoClient *mongo.Client, elasticClient *elasticsearch.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var body models.EmbeddingRequest
		pageSize, err := GetQueryValueAndValidate(ctx, "size")
		page, err := GetQueryValueAndValidate(ctx, "page")
		k, err := GetQueryValueAndValidate(ctx, "k")

		if err != nil {
			NewHTTPError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		err = ctx.BindJSON(&body)

		if err != nil {
			NewHTTPError(ctx, http.StatusBadRequest, "The query has not been correctly formatted")
			return
		}

		embeddings, _, err := embedding.PerformPipeline([]string{body.Fragment}, true)

		if err != nil {
			NewHTTPError(ctx, http.StatusInternalServerError, "The embedding backend is not reachable")
			return
		}

		filters, err := retrieval.ParseDSLRequest(body.DSL)

		if err != nil {
			NewHTTPError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		if strings.Compare(body.Fragment, "") != 0 && page * pageSize > k {
			NewHTTPError(ctx, http.StatusBadRequest, "page * size must not be greater than k")
			return
		}

		var query string

		if strings.Compare(body.Fragment, "") == 0 {
			query = retrieval.CreateKnnQuery(nil, *filters, pageSize, page, k)
		} else {
			query = retrieval.CreateKnnQuery(embeddings.Predictions[0], *filters, pageSize, page, k)
		}

		response, err := elastic.SearchDocument(elasticClient, query, "apis")

		if err != nil {
			NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		var jsonMaps []models.SpecificationWithApi
		db := mongoClient.Database("apis")

		for _, item := range response.Hits.Hits {
			objId, err := primitive.ObjectIDFromHex(item.Document.Metadata.MongoId)
			document, err := mongodb.RetrieveDocument(db, bson.M{"_id": objId}, "specifications")

			if err != nil {
				NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
				return
			}

			var specObj models.SpecificationWithApi
			var jsonMap models.MongoResponse
			err = bson.Unmarshal(document, &jsonMap)
			specObj.MongoDocument = jsonMap.InitObject()
			specObj.MongoDocument.Length = item.Document.Metadata.Length
			specObj.MongoDocument.Score = item.Score
			specObj.MongoDocument.Metrics = item.Document.Metadata.Metrics
			specObj.Specification = document.Lookup("api").String()

			if err != nil {
				NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
				return
			}

			if len(body.Fields) > 0 {
				specObj, err = retrieval.FilterFields(specObj, body.Fields)

				if err != nil {
					NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
					return
				}
			}

			jsonMaps = append(jsonMaps, specObj)
		}

		ctx.JSON(http.StatusOK, jsonMaps)
	}

	return fn
}

func GetEmbedding(ctx *gin.Context) {
	res := struct {
		Query string `json:"query"`
	}{}

	err := ctx.BindJSON(&res)

	if err != nil {
		NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res.Query = embedding.PreprocessFragment([]string{res.Query}, false)[0]

	ctx.JSON(http.StatusOK, res)
}

func GetQueryValueAndValidate(ctx *gin.Context, key string) (int, error) {
	if defaultValues, in := models.ParametersMap[key]; in {
		if value, present := ctx.GetQuery(key); present {
			valueInt, err := strconv.Atoi(value)

			if (err != nil || valueInt < defaultValues[1] || valueInt > defaultValues[2]) && strings.Compare(key, "page") != 0 {
				return 0, errors.New(fmt.Sprintf("%s must be a number, >= %d, and <= %d",
					key, defaultValues[1], defaultValues[2]))
			} else {
				return valueInt, nil
			}
		} else {
			return defaultValues[0], nil
		}
	}

	return 0, errors.New("the passed key is invalid")
}
