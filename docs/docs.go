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
        "/weather": {
            "get": {
                "description": "Get weather infos by given city name in query",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "weather"
                ],
                "summary": "Show weather info",
                "parameters": [
                    {
                        "type": "string",
                        "format": "city",
                        "description": "Weather search by city name",
                        "name": "city",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.StoreData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.StoreData": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "last_updated": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "localtime": {
                    "type": "string"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "temp_c": {
                    "type": "number"
                },
                "temp_f": {
                    "type": "number"
                }
            }
        },
        "utils.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        }
    }
}`

// SwaggerInfo
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7070",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Weather App",
	Description:      "This service is a weather web server that utilizes an external API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
