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
        "contact": {
            "name": "Dimas Restu Hidayanto",
            "url": "https://github.com/dimaskiddo",
            "email": "drh.dimasrestu@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/whatsapp": {
            "get": {
                "description": "Get The Server Status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Root"
                ],
                "summary": "Show The Status of The Server",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/auth": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get Authentication Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Root"
                ],
                "summary": "Generate Authentication Token",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/login": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get QR Code for WhatsApp Multi-Device Login",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json",
                    "text/html"
                ],
                "tags": [
                    "WhatsApp Authentication"
                ],
                "summary": "Generate QR Code for WhatsApp Multi-Device Login",
                "parameters": [
                    {
                        "enum": [
                            "html",
                            "json"
                        ],
                        "type": "string",
                        "default": "html",
                        "description": "Change Output Format in HTML or JSON",
                        "name": "output",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/logout": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Make Device Logout from WhatsApp Multi-Device",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Authentication"
                ],
                "summary": "Logout Device from WhatsApp Multi-Device",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/send/audio": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Audio Message to Spesific Phone Number",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Send Audio Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination Phone Number",
                        "name": "msisdn",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Audio File",
                        "name": "audio",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/send/contact": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Contact Message to Spesific Phone Number",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Send Contact Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination Phone Number",
                        "name": "msisdn",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contact Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contact Phone",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/send/document": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Document Message to Spesific Phone Number",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Send Document Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination Phone Number",
                        "name": "msisdn",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Document File",
                        "name": "document",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/send/image": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Image Message to Spesific Phone Number",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Send Image Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination Phone Number",
                        "name": "msisdn",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Caption Image Message",
                        "name": "caption",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image File",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "Is View Once",
                        "name": "viewonce",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/send/link": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Link Message to Spesific Phone Number",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Send Link Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination Phone Number",
                        "name": "msisdn",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Link Caption",
                        "name": "caption",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Link URL",
                        "name": "url",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/send/location": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Location Message to Spesific Phone Number",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Send Location Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination Phone Number",
                        "name": "msisdn",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Location Latitude",
                        "name": "latitude",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Location Longitude",
                        "name": "longitude",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/send/sticker": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Sticker Message to Spesific Phone Number",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Send Sticker Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination Phone Number",
                        "name": "msisdn",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Sticker File",
                        "name": "sticker",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/send/text": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Text Message to Spesific Phone Number",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Send Text Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination Phone Number",
                        "name": "msisdn",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Text Message",
                        "name": "message",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/send/video": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Video Message to Spesific Phone Number",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Send Video Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination Phone Number",
                        "name": "msisdn",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Caption Video Message",
                        "name": "caption",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Video File",
                        "name": "video",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "Is View Once",
                        "name": "viewonce",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        },
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.x",
	Host:             "127.0.0.1:3000",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Go WhatsApp Multi-Device REST API",
	Description:      "This is WhatsApp Multi-Device Implementation in Go REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
