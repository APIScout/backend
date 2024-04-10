package models

import (
	"regexp"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// MongoResponse - structure of the Mongo document sent by the db
type MongoResponse struct {
	MongoId    string `bson:"_id"`
	Id         int    `bson:"api_spec_id"`
	Name       string `bson:"_name"`
	Commits    int    `bson:"commits"`
	Latest     bool   `bson:"latest"`
	OASType    string
	Source     string
	ApiVersion Version
	OASVersion Version

	SpecificationJson bson.Raw `json:"-" bson:"api"`
	NameAlt           string   `bson:"api_title"`
	ApiVersionAlt1    string   `bson:"_version"`
	ApiVersionAlt2    string   `bson:"api_version"`
	SourceAlt1        string   `bson:"_api_url"`
	SourceAlt2        string   `bson:"url"`
}

type MongoDocument struct {
	MongoId       string        `json:"mongo-id"`
	Api           Api           `json:"api"`
	Specification Specification `json:"specification"`
}

type Api struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Version Version `json:"version"`
	Commits int     `json:"commits"`
	Latest  bool    `json:"latest"`
	Source  string  `json:"source"`
}

type Specification struct {
	Version Version `json:"version"`
	Type    string  `json:"type"`
}

type Version struct {
	Raw        string `json:"raw"`
	Valid      bool   `json:"valid"`
	Major      int    `json:"major"`
	Minor      int    `json:"minor"`
	Patch      int    `json:"patch"`
	Prerelease string `json:"prerelease"`
	Build      string `json:"build"`
}

// MongoResponseWithApi - structure containing both the mongo document and the embedding created by the backend
type MongoResponseWithApi struct {
	MongoResponse MongoResponse `json:"metadata"`
	Specification string        `json:"specification"`
}

// InitObject - function to fix the initiated object
func (b *MongoResponse) InitObject() MongoDocument {
	GetOasVersion(b)

	if strings.Compare(b.NameAlt, "") != 0 {
		b.Name = b.NameAlt
		b.ApiVersion = GetSemanticVersion(b.ApiVersionAlt2)
		b.Source = GetSource(b.SourceAlt2)
	} else {
		b.ApiVersion = GetSemanticVersion(b.ApiVersionAlt1)
		b.Source = GetSource(b.SourceAlt1)
	}

	return MongoDocument{
		MongoId: b.MongoId,
		Api: Api{
			Id:      b.Id,
			Name:    b.Name,
			Version: b.ApiVersion,
			Commits: b.Commits,
			Latest:  b.Latest,
			Source:  b.Source,
		},
		Specification: Specification{
			Version: b.OASVersion,
			Type:    b.OASType,
		},
	}
}

func GetOasVersion(specification *MongoResponse) {
	oasOpenapi := specification.SpecificationJson.Lookup("openapi").String()
	oasSwagger := specification.SpecificationJson.Lookup("swagger").String()
	oasOpenapi = strings.Trim(oasOpenapi, `\\"`)
	oasSwagger = strings.Trim(oasSwagger, `\\"`)

	if strings.Compare(oasOpenapi, "") != 0 {
		specification.OASVersion = GetSemanticVersion(oasOpenapi)
		specification.OASType = "openapi"
	} else {
		specification.OASVersion = GetSemanticVersion(oasSwagger)
		specification.OASType = "swagger"
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
		major, _ := strconv.ParseInt(matches[major], 10, 32)
		minor, _ := strconv.ParseInt(matches[minor], 10, 32)
		patch, _ := strconv.ParseInt(matches[patch], 10, 32)

		semanticVersion.Valid = true
		semanticVersion.Major = int(major)
		semanticVersion.Minor = int(minor)
		semanticVersion.Patch = int(patch)
		semanticVersion.Prerelease = matches[prerelease]
		semanticVersion.Build = matches[build]
	}

	return semanticVersion
}

func GetSource(url string) string {
	if strings.Contains(url, "github") {
		return "github"
	} else if strings.Contains(url, "swagger") {
		return "swaggerhub"
	}

	return ""
}
