// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Jeisa Raja",
            "url": "https://github.com/jeisaRaja"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/games/": {
            "post": {
                "description": "This endpoint allows the user to create a new game in the system by providing necessary game details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Create a new game entry",
                "parameters": [
                    {
                        "description": "Game details",
                        "name": "game",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.GameCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Game created successfully",
                        "schema": {
                            "$ref": "#/definitions/data.Game"
                        }
                    }
                }
            }
        },
        "/games/{id}": {
            "get": {
                "description": "This endpoint retrieves the details of a specific game by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Show details of a specific game",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Game ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Game details fetched successfully",
                        "schema": {
                            "$ref": "#/definitions/main.GameResponse"
                        }
                    },
                    "404": {
                        "description": "Game not found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "data.Game": {
            "type": "object",
            "properties": {
                "developer": {
                    "type": "string"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/data.Genres"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "platform": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/data.Platform"
                    }
                },
                "price": {
                    "type": "integer"
                },
                "publisher": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "rating_count": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "data.Genres": {
            "type": "string",
            "enum": [
                "action",
                "adventure",
                "action-adventure",
                "puzzle",
                "role-playing",
                "simulation",
                "strategy",
                "sports",
                "mmo",
                "platformer"
            ],
            "x-enum-varnames": [
                "Action",
                "Adventure",
                "ActionAdventure",
                "Puzzle",
                "RolePlaying",
                "Simulation",
                "Strategy",
                "Sports",
                "MMO",
                "Platformer"
            ]
        },
        "data.Platform": {
            "type": "string",
            "enum": [
                "playstation",
                "xbox",
                "nintento switch",
                "pc",
                "ios",
                "android"
            ],
            "x-enum-varnames": [
                "Playstation",
                "Xbox",
                "NintentoSwitch",
                "PC",
                "IOS",
                "Android"
            ]
        },
        "main.GameCreateRequest": {
            "type": "object",
            "required": [
                "developer",
                "genres",
                "platform",
                "publisher",
                "title",
                "year"
            ],
            "properties": {
                "developer": {
                    "type": "string"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/data.Genres"
                    }
                },
                "platform": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/data.Platform"
                    }
                },
                "price": {
                    "type": "integer",
                    "minimum": 0
                },
                "publisher": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                },
                "year": {
                    "type": "integer",
                    "minimum": 1900
                }
            }
        },
        "main.GameResponse": {
            "type": "object",
            "properties": {
                "games": {
                    "$ref": "#/definitions/data.Game"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:8000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "GameVault API",
	Description:      "sample desc",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
