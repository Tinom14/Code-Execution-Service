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
        "/login": {
            "post": {
                "description": "Проверяет логин и пароль, возвращает токен",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Аутентификация пользователя",
                "parameters": [
                    {
                        "description": "Данные пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UserHandlerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный вход",
                        "schema": {
                            "$ref": "#/definitions/types.SessionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/protected-endpoint": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Позволяет доступ только авторизованным пользователям",
                "tags": [
                    "auth"
                ],
                "summary": "Проверка токена",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Токен авторизации, формат: Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Создаёт нового пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "register"
                ],
                "summary": "Регистрация пользователя",
                "parameters": [
                    {
                        "description": "Данные пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UserHandlerRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/result/{task_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Возвращает результат выполнения задачи по ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Получить результат задачи",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID задачи",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Токен авторизации, формат: Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetTaskResultHandlerResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/status/{task_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Возвращает текущий статус задачи по ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Получить статус задачи",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID задачи",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Токен авторизации, формат: Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetTaskStatusHandlerResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Запускает новую задачу и возвращает ее ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Создать задачу",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Токен авторизации, формат: Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.PostTaskHandlerResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.GetTaskResultHandlerResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                }
            }
        },
        "types.GetTaskStatusHandlerResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "types.PostTaskHandlerResponse": {
            "type": "object",
            "properties": {
                "task_id": {
                    "type": "string"
                }
            }
        },
        "types.SessionResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "types.UserHandlerRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
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
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Task API",
	Description:      "API для создания задач и получения их статуса/результата.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
