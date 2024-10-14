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
        "/actors": {
            "get": {
                "description": "Get all actors",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actor"
                ],
                "summary": "Get all actors",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Actor"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create an actor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actor"
                ],
                "summary": "Create an actor",
                "parameters": [
                    {
                        "description": "Actor payload",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ActorRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Actor"
                        }
                    }
                }
            }
        },
        "/actors/{id}": {
            "get": {
                "description": "Get an actor",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actor"
                ],
                "summary": "Get an actor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "actorId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Actor"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an actor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actor"
                ],
                "summary": "Update an actor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "actorId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Actor"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an actor with the given ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actor"
                ],
                "summary": "Delete an actor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "actorId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Actor deleted successfully"
                    }
                }
            }
        },
        "/films/{id}": {
            "get": {
                "description": "Get a film with the given ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Film"
                ],
                "summary": "Get a film",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "filmId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Film"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a film with the given ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Film"
                ],
                "summary": "Delete a film",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "filmId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Film deleted successfully"
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Actor": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "lastUpdate": {
                    "type": "string"
                }
            }
        },
        "entity.Film": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "languageId": {
                    "type": "integer"
                },
                "lastUpdate": {
                    "type": "string"
                },
                "length": {
                    "type": "integer"
                },
                "originalLanguageId": {
                    "type": "integer"
                },
                "rating": {
                    "type": "string"
                },
                "releaseYear": {
                    "type": "integer"
                },
                "rentalDuration": {
                    "type": "integer"
                },
                "rentalRate": {
                    "type": "number"
                },
                "replacementCost": {
                    "type": "number"
                },
                "special_features": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ActorRequest": {
            "type": "object",
            "required": [
                "firstName",
                "lastName"
            ],
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
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
	Title:            "API for Advanced Web",
	Description:      "API for Advanced Web",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
