package models

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// MongoResponse - structure of the Mongo document sent by the db
type MongoResponse struct {
	MongoId    string `json:"mongo_id" bson:"_id"`
	Name       string `json:"name" bson:"_name"`
	ApiId      int    `json:"api_id" bson:"id"`
	ApiVersion string `json:"api_version" bson:"_version"`
	OASVersion string `json:"oas_version"`
	Commits    int    `json:"commits_n" bson:"commits"`
	Latest     bool   `json:"is_latest" bson:"latest"`

	Specification bson.Raw `json:"-" bson:"api"`
	NameAlt       string   `json:"-" bson:"api_title"`
	ApiVersionAlt string   `json:"-" bson:"api_version"`
}

// MongoResponseWithApi - structure containing both the mongo document and the embedding created by the backend
type MongoResponseWithApi struct {
	MongoResponse MongoResponse `json:"metadata"`
	Specification string        `json:"specification"`
}

// InitObject - function to fix the initiated object
func (b *MongoResponse) InitObject() {
	oasOpenapi := b.Specification.Lookup("openapi").String()
	oasSwagger := b.Specification.Lookup("swagger").String()
	oasOpenapi = strings.Trim(oasOpenapi, `\\"`)
	oasSwagger = strings.Trim(oasSwagger, `\\"`)

	if strings.Compare(oasOpenapi, "") != 0 {
		b.OASVersion = oasOpenapi
	} else {
		b.OASVersion = oasSwagger
	}

	if strings.Compare(b.NameAlt, "") != 0 {
		b.Name = b.NameAlt
		b.ApiVersion = b.ApiVersionAlt
	}
}
