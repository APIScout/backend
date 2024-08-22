package controller

import (
	"backend/app/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetDSL godoc
//
//	@Summary		Retrieve DSL Specifications
//	@Description	Retrieve the filtering DSL specifications
//	@Tags			DSL
//	@Produce		json
//	@Param			page	query		string	false	"parameter name"
//	@Success		200		{object}	[]string
//	@Failure		400		{object}	models.HTTPError
//	@Failure		500		{object}	models.HTTPError
//	@Router			/search [post]
func GetDSL(ctx *gin.Context) {
	var parameter = ctx.Query("parameter")

	ctx.Header("Access-Control-Allow-Origin", "*")

	if parameter == "" {
		ctx.JSON(http.StatusOK, models.TypesMap)
		return
	} else {
		if param, ok := models.TypesMap[parameter]; ok {
			ctx.JSON(http.StatusOK, models.TypeToOperatorsMap[param])
			return
		}
	}

	NewHTTPError(ctx, http.StatusBadRequest, "Malformed request")
	return
}
