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
        "/api/v1/search/{key}": {
            "get": {
                "description": "关键词搜索新闻, 按照时间排序",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "summary": "关键词搜索",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键字",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.NewsModel"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.NewsModel": {
            "type": "object",
            "properties": {
                "abstract": {
                    "type": "string"
                },
                "author": {
                    "type": "string"
                },
                "data_source": {
                    "description": "required",
                    "type": "string"
                },
                "img_url": {
                    "type": "string"
                },
                "is_hot": {
                    "type": "boolean"
                },
                "list_url": {
                    "description": "required",
                    "type": "string"
                },
                "news_url": {
                    "description": "required",
                    "type": "string"
                },
                "page_url": {
                    "description": "required",
                    "type": "string"
                },
                "publish_time": {
                    "type": "string"
                },
                "rank": {
                    "type": "integer"
                },
                "title": {
                    "description": "required",
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
