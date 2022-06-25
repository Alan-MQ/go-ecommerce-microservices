// Package openapi GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package openapi

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mehdi Hadeli",
            "url": "https://github.com/mehdihadeli"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "get": {
                "description": "Get all products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all product",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/getting_products.GetProductsResponseDto"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new product item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Create product",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "CreateProductRequestDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateProductRequestDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateProductResponseDto"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Get product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/getting_product_by_id.GetProductByIdResponseDto"
                        }
                    }
                }
            },
            "put": {
                "description": "Update existing product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update product",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "UpdateProductRequestDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/updating_product.UpdateProductRequestDto"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Delete existing product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ProductDto": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "productId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "dtos.CreateProductRequestDto": {
            "type": "object",
            "required": [
                "description",
                "name",
                "price"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 5000,
                    "minLength": 0
                },
                "name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 0
                },
                "price": {
                    "type": "number",
                    "minimum": 0
                }
            }
        },
        "dtos.CreateProductResponseDto": {
            "type": "object",
            "properties": {
                "productId": {
                    "type": "string"
                }
            }
        },
        "getting_product_by_id.GetProductByIdResponseDto": {
            "type": "object",
            "properties": {
                "product": {
                    "$ref": "#/definitions/dto.ProductDto"
                }
            }
        },
        "getting_products.GetProductsResponseDto": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ProductDto"
                    }
                }
            }
        },
        "updating_product.UpdateProductRequestDto": {
            "type": "object",
            "required": [
                "description",
                "name",
                "price",
                "productId"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 5000,
                    "minLength": 0
                },
                "name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 0
                },
                "price": {
                    "type": "number",
                    "minimum": 0
                },
                "productId": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 0
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
