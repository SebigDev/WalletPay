// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/login": {
            "post": {
                "description": "Authenticates a person.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "autenticate a person",
                "parameters": [
                    {
                        "description": "Authenticate",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "description": "Get person by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.PersonResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "description": "Get all users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "summary": "get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.PersonResponse"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/users/onboard": {
            "post": {
                "description": "Create a new person.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "summary": "create a new person",
                "parameters": [
                    {
                        "description": "Create person",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePerson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/wallet/add": {
            "post": {
                "description": "Create a new wallet.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "create a new wallet",
                "parameters": [
                    {
                        "description": "Create wallet",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateWalletRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/wallet/deposit": {
            "post": {
                "description": "Deposit in a wallet.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "deposit in a wallet",
                "parameters": [
                    {
                        "description": "Deposit into a wallet",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DepositRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/wallet/withdraw": {
            "post": {
                "description": "Withdraw from a wallet.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Withdraw from a wallet",
                "parameters": [
                    {
                        "description": "Withdraws from a wallet",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.WithdrawRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreatePerson": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "emailAddress": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "houseNumber": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "postalCode": {
                    "type": "string"
                },
                "streetName": {
                    "type": "string"
                }
            }
        },
        "dto.CreateWalletRequest": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                }
            }
        },
        "dto.DepositRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "walletNo": {
                    "type": "string"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "properties": {
                "emailAddress": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.WithdrawRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "walletNo": {
                    "type": "string"
                }
            }
        },
        "responses.Amount": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "responses.PersonResponse": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "emailAddress": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "houseNumber": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "isVerified": {
                    "type": "boolean"
                },
                "lastName": {
                    "type": "string"
                },
                "postalCode": {
                    "type": "string"
                },
                "streetName": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                },
                "wallets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.WalletResponse"
                    }
                }
            }
        },
        "responses.WalletResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "$ref": "#/definitions/responses.Amount"
                },
                "number": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "GoApp Wallet API",
	Description:      "This is a GoApp project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
