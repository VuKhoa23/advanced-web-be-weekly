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
                            "$ref": "#/definitions/model.HttpResponse-array_entity_Actor"
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
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-entity_Actor"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
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
                        "example": 1,
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
                            "$ref": "#/definitions/model.HttpResponse-entity_Actor"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
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
                        "example": 1,
                        "description": "actorId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Actor payload",
                        "name": "request",
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
                            "$ref": "#/definitions/model.HttpResponse-entity_Actor"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
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
                        "example": 1,
                        "description": "actorId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Actor deleted successfully"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
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
                        "example": 1,
                        "description": "actorId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Actor payload",
                        "name": "request",
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
                        "example": 1,
                        "description": "actorId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Actor deleted successfully"
                    }
                }
            }
        },
        "/films": {
            "get": {
                "description": "Get all films",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Film"
                ],
                "summary": "Get all films",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-array_entity_Film"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a film",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Film"
                ],
                "summary": "Update a film",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "filmId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Film payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FilmRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-entity_Film"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a film",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Film"
                ],
                "summary": "Create a film",
                "parameters": [
                    {
                        "description": "Film payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FilmRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-entity_Film"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
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
                        "example": 1,
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
                            "$ref": "#/definitions/model.HttpResponse-entity_Film"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
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
                        "example": 1,
                        "description": "filmId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Film deleted successfully"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.HttpResponse-any"
                        }
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
                "specialFeatures": {
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
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1
                }
            }
        },
        "model.FilmRequest": {
            "type": "object",
            "required": [
                "description",
                "languageId",
                "length",
                "originalLanguageId",
                "rating",
                "releaseYear",
                "rentalDuration",
                "rentalRate",
                "replacementCost",
                "specialFeatures",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 10
                },
                "languageId": {
                    "type": "integer",
                    "minimum": 0
                },
                "length": {
                    "type": "integer",
                    "minimum": 1
                },
                "originalLanguageId": {
                    "type": "integer"
                },
                "rating": {
                    "type": "string"
                },
                "releaseYear": {
                    "type": "integer",
                    "minimum": 0
                },
                "rentalDuration": {
                    "type": "integer",
                    "minimum": 1
                },
                "rentalRate": {
                    "type": "number",
                    "minimum": 0
                },
                "replacementCost": {
                    "type": "number",
                    "minimum": 0
                },
                "specialFeatures": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255
                }
            }
        },
        "model.HttpResponse-any": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.HttpResponse-array_entity_Actor": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Actor"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.HttpResponse-array_entity_Film": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Film"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.HttpResponse-entity_Actor": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/entity.Actor"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.HttpResponse-entity_Film": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/entity.Film"
                },
                "message": {
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
