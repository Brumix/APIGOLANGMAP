// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/auth/login": {
            "post": {
                "description": "Autentica o utilizador e gera o token para os próximos acessos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Realizar autenticação",
                "parameters": [
                    {
                        "description": "Do login",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserOut"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Claims"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "Logout do utilizador invalidando o token atual",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Logout",
                "parameters": [
                    {
                        "description": "Logout",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserOut"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "bool"
                    },
                    "406": {
                        "description": "Cannot log out"
                    }
                }
            }
        },
        "/auth/refresh_token": {
            "put": {
                "description": "Atualiza o token de autenticação do utilizador invalidando o antigo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Atualiza token de autenticação",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Claims"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "406": {
                        "description": "Cannot invalidate old token"
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Registo do Utilizador",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Registar um utilizador",
                "parameters": [
                    {
                        "description": "Registar um utilizador",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserOut"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Claims"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
		"/alert/time": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Atualiza a periodicidade de alerta determinando o tempo máximo até dar uma pessoa como perdida",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Atualiza a periodicidade de alerta",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Update Alert",
                        "name": "Username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Alert"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Alert"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    }
                }
            },
        },
        "/follower": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Exibe a lista, sem todos os campos, de todos os followers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtem os Followers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Follower"
                            }
                        }
                    },
                    "404": {
                        "description": "Not found"
                    }
                }
            }
        },
        "/follower/assoc": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Associa um Follower a um User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Associa um Follower(User) a um User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Associate User as Follower",
                        "name": "follower",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Follower"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Follower"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Not found"
                    }
                }
            }
        },
        "/follower/deassoc": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Desassocia um Follower de um User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Desassocia um Follower(User) de um User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Deassociate Follower from User",
                        "name": "follower",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Follower"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Follower"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Not found"
                    }
                }
            }
        },
        "/position": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Exibe a lista da última localização do utilizador",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obter a última localização do utilizador",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Position"
                        }
                    },
                    "400": {
                        "description": "User Token Malformed"
                    },
                    "404": {
                        "description": "User Not found"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Cria uma localizacao de um utilizador em especifico",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Adicionar uma localizaçao",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Add Location",
                        "name": "evaluation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Position"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Position"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Not found"
                    }
                }
            }
        },
        "/position/history": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Exibe a lista de todas as localizações do utilizador",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtem todas as localizações do utilizador",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Position"
                            }
                        }
                    },
                    "400": {
                        "description": "User Token Malformed"
                    },
                    "404": {
                        "description": "User Not found"
                    }
                }
            }
        },
        "/position/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Exclui uma localização selecionada",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Exclui uma localização",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Position ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete succeeded!"
                    },
                    "404": {
                        "description": "None found!"
                    }
                }
            }
        },
        "/socket": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Inicia todos os recursos necessario para a criação de uma webSocket com o cliente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Iniciar conecção com a webSocket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Connection confirm"
                    },
                    "400": {
                        "description": "User Token Malformed"
                    },
                    "404": {
                        "description": "Connection failed"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Claims": {
            "type": "object",
            "properties": {
                "access_mode": {
                    "type": "integer"
                },
                "userid": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.Follower": {
            "type": "object",
            "properties": {
                "FollowerUserID": {
                    "type": "integer"
                },
                "UserID": {
                    "type": "integer"
                }
            }
        },
        "model.Position": {
            "type": "object",
            "required": [
                "Latitude",
                "Longitude"
            ],
            "properties": {
                "Latitude": {
                    "type": "number"
                },
                "Longitude": {
                    "type": "number"
                },
                "UserId": {
                    "type": "integer"
                }
            }
        },
		"model.Alert": {
            "type": "object",
            "properties": {
                "alertTime": {
                    "type": "integer"
                },
            }
        },
		"model.UserOut": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                },
				"password": {
                    "type": "string"
                },
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "access_mode": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "userFriends": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Follower"
                    }
                },
                "userPositions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Position"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
