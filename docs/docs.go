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
        "/auth/forgot-password": {
            "post": {
                "description": "Receives the user email and if the email is valid, send a verification via email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Request new password via e-mail",
                "parameters": [
                    {
                        "description": "Reset password info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authentication.ForgotPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No content"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
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
        "/auth/update-password": {
            "post": {
                "description": "Checks the token on the request and updates the password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Update password after email verification",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User verification token",
                        "name": "t",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Update password info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authentication.UpdatePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No content"
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
                    "Course"
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
                                                "$ref": "#/definitions/courses.ResponseContent"
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
                    "Course"
                ],
                "summary": "Create a course",
                "parameters": [
                    {
                        "description": "Request payload for creating a new course",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/courses.RequestContent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/courses.ResponseContent"
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
        "/courses/{id}": {
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
                    "Course"
                ],
                "summary": "Get course by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the course to be retrieved",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/courses.ResponseContent"
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
                    "Course"
                ],
                "summary": "Update course by ID",
                "parameters": [
                    {
                        "description": "Request payload for updating course information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/courses.RequestContent"
                        }
                    },
                    {
                        "type": "string",
                        "description": "ID of course to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/courses.ResponseContent"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Course not found"
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
                    "Course"
                ],
                "summary": "Delete a course by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the course to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "OK"
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
                                                "$ref": "#/definitions/users.ResponseContent"
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
                            "$ref": "#/definitions/users.RequestContent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.ResponseContent"
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
        "/users/me": {
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
                "summary": "Get authenticated user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.ResponseContent"
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
                            "$ref": "#/definitions/users.ResponseContent"
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
                            "$ref": "#/definitions/users.RequestContentUpdate"
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
                            "$ref": "#/definitions/users.ResponseContent"
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
                    "204": {
                        "description": "No content"
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
        "authentication.ForgotPasswordRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
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
        "authentication.UpdatePasswordRequest": {
            "type": "object",
            "properties": {
                "password": {
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
        "courses.ContentRequest": {
            "type": "object",
            "properties": {
                "URL": {
                    "type": "string"
                },
                "contentType": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "courses.ContentResponse": {
            "type": "object",
            "properties": {
                "URL": {
                    "type": "string"
                },
                "contentType": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "courses.RequestContent": {
            "type": "object",
            "properties": {
                "contentInHours": {
                    "type": "number"
                },
                "contents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/courses.ContentRequest"
                    }
                },
                "description": {
                    "type": "string"
                },
                "level": {
                    "type": "string"
                },
                "thumbnailURL": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "courses.ResponseContent": {
            "type": "object",
            "properties": {
                "contentInHours": {
                    "type": "number"
                },
                "contents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/courses.ContentResponse"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "level": {
                    "type": "string"
                },
                "thumbnailURL": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
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
        "users.RequestContent": {
            "type": "object",
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "users.RequestContentUpdate": {
            "type": "object",
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "users.ResponseContent": {
            "type": "object",
            "properties": {
                "birthDate": {
                    "type": "string"
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
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
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
            "description": "Insert the token withou \"Bearer\" prefix.",
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
