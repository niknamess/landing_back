// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/v1/form": {
            "get": {
                "description": "получение формы",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "action"
                ],
                "summary": "получить формы по user id",
                "parameters": [
                    {
                        "description": "Ввести UserID",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sober_driver_pkg_domain.GetFormsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sober_driver_pkg_domain.FormsResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/form/insert": {
            "post": {
                "description": "запись в форму инпуты",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "action"
                ],
                "summary": "записать в форму",
                "parameters": [
                    {
                        "description": "Форма для заполнения",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sober_driver_pkg_domain.FormToAction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sober_driver_pkg_domain.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/register": {
            "post": {
                "description": "регистрация с проверкой",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "регистрация пользователя",
                "parameters": [
                    {
                        "description": "Ввести данные",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sober_driver_pkg_domain.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sober_driver_pkg_domain.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/register/check": {
            "get": {
                "description": "проверка регистрации пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "проверка регистрации пользователя",
                "parameters": [
                    {
                        "description": "Ввести Email",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sober_driver_pkg_domain.RegisterCheckRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sober_driver_pkg_domain.RegisterResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "sober_driver_pkg_domain.Fio": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "sober_driver_pkg_domain.FormToAction": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "where_from": {
                    "type": "string"
                },
                "where_to": {
                    "type": "string"
                }
            }
        },
        "sober_driver_pkg_domain.FormsResponse": {
            "type": "object",
            "properties": {
                "forms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sober_driver_pkg_domain.FormToAction"
                    }
                }
            }
        },
        "sober_driver_pkg_domain.GetFormsRequest": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "sober_driver_pkg_domain.RegisterCheckRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "sober_driver_pkg_domain.RegisterRequest": {
            "type": "object",
            "properties": {
                "customer": {
                    "type": "boolean"
                },
                "driver": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string"
                },
                "fio": {
                    "$ref": "#/definitions/sober_driver_pkg_domain.Fio"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "sober_driver_pkg_domain.RegisterResponse": {
            "type": "object",
            "properties": {
                "status": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
