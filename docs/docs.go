// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://cybersafe.academy.com/support/terms",
        "contact": {
            "name": "CyberSafe support team",
            "url": "http://cybersafe.academy.com/support/contact",
            "email": "support@cybersafe.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/license/mit/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Authenticates an user and generates an access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User login information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authentication.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authentication.TokenContent"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/auth/logoff": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "Language": []
                    }
                ],
                "description": "Logs off an user",
                "tags": [
                    "Authentication"
                ],
                "summary": "User logoff",
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "Language": []
                    }
                ],
                "description": "Refreshes the token for an authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User login refresh",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authentication.TokenContent"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/courses": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "Language": []
                    }
                ],
                "tags": [
                    "Courses"
                ],
                "summary": "List courses with paginated response",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit of elements per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/pagination.PaginationData"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "$ref": "#/definitions/course.ResponseContent"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "default": {
                        "description": "Standard error example object",
                        "schema": {
                            "$ref": "#/definitions/components.Response"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "Language": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "List users with paginated response",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit of elements per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/pagination.PaginationData"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "$ref": "#/definitions/user.ResponseContent"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "default": {
                        "description": "Standard error example object",
                        "schema": {
                            "$ref": "#/definitions/components.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "Language": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "Request payload for creating a new user",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RequestContent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/pagination.PaginationData"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "$ref": "#/definitions/user.ResponseContent"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "default": {
                        "description": "Standard error example object",
                        "schema": {
                            "$ref": "#/definitions/components.Response"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "Language": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the user to be retrieved",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.ResponseContent"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "default": {
                        "description": "Standard error example object",
                        "schema": {
                            "$ref": "#/definitions/components.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "Language": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user by ID",
                "parameters": [
                    {
                        "description": "Request payload for updating user information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RequestContent"
                        }
                    },
                    {
                        "type": "string",
                        "description": "ID of user to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.ResponseContent"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "User not found"
                    },
                    "default": {
                        "description": "Standard error example object",
                        "schema": {
                            "$ref": "#/definitions/components.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "Language": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete a user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the user to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/pagination.PaginationData"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "$ref": "#/definitions/user.ResponseContent"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "default": {
                        "description": "Standard error example object",
                        "schema": {
                            "$ref": "#/definitions/components.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authentication.LoginRequest": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "authentication.TokenContent": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "expiresIn": {
                    "type": "number"
                },
                "tokenType": {
                    "type": "string"
                }
            }
        },
        "components.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "description": {
                    "type": "string",
                    "example": "Bad Request"
                },
                "error_details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/components.ErrorDetail"
                    }
                },
                "id": {
                    "type": "string",
                    "example": "c77fa521-99b1-4c54-9a8d-4b6902912eb0"
                }
            }
        },
        "components.ErrorDetail": {
            "type": "object",
            "properties": {
                "attribute": {
                    "type": "string",
                    "example": "field name with error or key for help messages"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "explanatory messages about the attribute error"
                    ]
                }
            }
        },
        "components.Response": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/components.Error"
                }
            }
        },
        "course.ContentFields": {
            "type": "object",
            "properties": {
                "PDFURL": {
                    "type": "string",
                    "example": "https://pdf.com"
                },
                "contentType": {
                    "type": "string",
                    "example": "youtube"
                },
                "id": {
                    "type": "string"
                },
                "imageURL": {
                    "type": "string",
                    "example": "https://image.com"
                },
                "youtubeURL": {
                    "type": "string",
                    "example": "https://www.youtube.com/watch?v=mvV7tzRm8Pk"
                }
            }
        },
        "course.ResponseContent": {
            "type": "object",
            "properties": {
                "contentInHours": {
                    "type": "number",
                    "example": 24.5
                },
                "contents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/course.ContentFields"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string",
                    "example": "Course Description"
                },
                "id": {
                    "type": "string"
                },
                "level": {
                    "type": "string",
                    "example": "advanced"
                },
                "name": {
                    "type": "string",
                    "example": "Example Course"
                },
                "thumbnailURL": {
                    "type": "string",
                    "example": "https://image.com"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "pagination.PaginationData": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "user.RequestContent": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "cpf": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.ResponseContent": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "cpf": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token. e.g: Bearer eyJhbGciO....",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "CyberSafe Academy API",
	Description:      "This REST API contains all services for the CyberSafe plataform.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
