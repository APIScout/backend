// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/search": {
            "post": {
                "description": "Retrieve OpenAPI specifications matching the given query",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Search OpenAPI specifications",
                "parameters": [
                    {
                        "description": "Search query",
                        "name": "fragment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EmbeddingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/specification": {
            "post": {
                "description": "Insert new OpenAPI specifications in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "specification"
                ],
                "summary": "Insert OpenAPI specifications",
                "parameters": [
                    {
                        "description": "New Specifications",
                        "name": "specifications",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SpecificationsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/specification/{id}": {
            "get": {
                "description": "Retrieve a specific OpenAPI specification's content given a valid ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "specification"
                ],
                "summary": "Get OpenAPI specification",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Specification ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MongoResponseWithApi"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.EmbeddingRequest": {
            "type": "object",
            "properties": {
                "fragment": {
                    "type": "string"
                }
            }
        },
        "models.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "Bad Request"
                }
            }
        },
        "models.MongoResponse": {
            "type": "object",
            "properties": {
                "api_id": {
                    "type": "integer"
                },
                "api_version": {
                    "type": "string"
                },
                "commits_n": {
                    "type": "integer"
                },
                "is_latest": {
                    "type": "boolean"
                },
                "mongo_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "oas_version": {
                    "type": "string"
                }
            }
        },
        "models.MongoResponseWithApi": {
            "type": "object",
            "properties": {
                "metadata": {
                    "$ref": "#/definitions/models.MongoResponse"
                },
                "specification": {
                    "type": "string"
                }
            }
        },
        "models.Specification": {
            "type": "object",
            "additionalProperties": true
        },
        "models.SpecificationsRequest": {
            "type": "object",
            "properties": {
                "specifications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Specification"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "API Scout",
	Description:      "This is the backend for the API Scout platform.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
