// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-05-04 19:43:32.019436285 +0800 CST m=+0.056436945

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "title": "kindleM API",
        "contact": {},
        "license": {},
        "version": "0.0.1"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/chapters": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get manga chapter list",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "manga_url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.MangaDetail"
                        }
                    }
                }
            }
        },
        "/download": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "download chapter list",
                "parameters": [
                    {
                        "description": " ",
                        "name": "download_list",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ChapterRowList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ChapterList"
                        }
                    }
                }
            }
        },
        "/search/{query}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "search manga",
                "operationId": "search-manga",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "query",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Manga"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Chapter": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "manga_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "model.ChapterList": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "created": {
                        "type": "string"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "link": {
                        "type": "string"
                    },
                    "manga_id": {
                        "type": "integer"
                    },
                    "status": {
                        "type": "boolean"
                    },
                    "title": {
                        "type": "string"
                    },
                    "total": {
                        "type": "integer"
                    },
                    "updated": {
                        "type": "string"
                    }
                }
            }
        },
        "model.ChapterRow": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ChapterRowList": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "link": {
                        "type": "string"
                    },
                    "title": {
                        "type": "string"
                    }
                }
            }
        },
        "model.Manga": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "author": {
                    "type": "string"
                },
                "cover": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "intro": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "source": {
                    "type": "integer"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "model.MangaDetail": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "author": {
                    "type": "string"
                },
                "chapters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Chapter"
                    }
                },
                "cover": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "intro": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "source": {
                    "type": "integer"
                },
                "updated": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
