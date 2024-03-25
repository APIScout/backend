package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"regexp"
	"strings"
)

// MongoResponse - structure of the Mongo document sent by the db
type MongoResponse struct {
	MongoId    string  `json:"mongo-id" bson:"_id"`
	Name       string  `json:"name" bson:"_name"`
	ApiId      int     `json:"api-id" bson:"api_spec_id"`
	ApiVersion Version `json:"api-version"`
	OASVersion Version `json:"oas-version"`
	Commits    int     `json:"commits" bson:"commits"`
	Latest     bool    `json:"latest" bson:"latest"`
	Source     string  `json:"source"`

	Specification  bson.Raw `json:"-" bson:"api"`
	NameAlt        string   `json:"-" bson:"api_title"`
	ApiVersionAlt1 string   `json:"-" bson:"_version"`
	ApiVersionAlt2 string   `json:"-" bson:"api_version"`
	SourceAlt1 string `json:"-" bson:"_api_url"`
	SourceAlt2 string `json:"-" bson:"url"`
}

type Version struct {
	Raw        string `json:"raw"`
	Valid      bool   `json:"valid"`
	Major      string `json:"major"`
	Minor      string `json:"minor"`
	Patch      string `json:"patch"`
	Prerelease string `json:"prerelease"`
	Build      string `json:"build"`
}

// MongoResponseWithApi - structure containing both the mongo document and the embedding created by the backend
type MongoResponseWithApi struct {
	MongoResponse MongoResponse `json:"metadata"`
	Specification string        `json:"specification"`
}

// InitObject - function to fix the initiated object
func (b *MongoResponse) InitObject() {
	GetOasVersion(b)

	if strings.Compare(b.NameAlt, "") != 0 {
		b.Name = b.NameAlt
		b.ApiVersion = GetSemanticVersion(b.ApiVersionAlt2)
		b.Source = GetSource(b.SourceAlt2)
	} else {
		b.ApiVersion = GetSemanticVersion(b.ApiVersionAlt1)
		b.Source = GetSource(b.SourceAlt1)
	}
}

func GetOasVersion(specification *MongoResponse) {
	oasOpenapi := specification.Specification.Lookup("openapi").String()
	oasSwagger := specification.Specification.Lookup("swagger").String()
	oasOpenapi = strings.Trim(oasOpenapi, `\\"`)
	oasSwagger = strings.Trim(oasSwagger, `\\"`)

	if strings.Compare(oasOpenapi, "") != 0 {
		specification.OASVersion = GetSemanticVersion(oasOpenapi)
	} else {
		specification.OASVersion = GetSemanticVersion(oasSwagger)
	}
}

func GetSemanticVersion(version string) Version {
	regex := regexp.MustCompile("^(?P<major>0|[1-9]\\d*)\\.(?P<minor>0|[1-9]\\d*)\\.(?P<patch>0|[1-9]\\d*)(?:-(?P<prerelease>(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+(?P<build>[0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$")
	major := regex.SubexpIndex("major")
	minor := regex.SubexpIndex("minor")
	patch := regex.SubexpIndex("patch")
	prerelease := regex.SubexpIndex("prerelease")
	build := regex.SubexpIndex("build")

	var semanticVersion Version
	semanticVersion.Raw = strings.TrimLeft(version, "v")
	matches := regex.FindStringSubmatch(version)

	if matches != nil {
		semanticVersion.Valid = true
		semanticVersion.Major = matches[major]
		semanticVersion.Minor = matches[minor]
		semanticVersion.Patch = matches[patch]
		semanticVersion.Prerelease = matches[prerelease]
		semanticVersion.Build = matches[build]
	}

	return semanticVersion
}

func GetSource(url string) string {
	if strings.Contains(url, "github") {
		return "GitHub"
	} else if strings.Contains(url, "swagger") {
		return "SwaggerHub"
	}

	return ""
}
