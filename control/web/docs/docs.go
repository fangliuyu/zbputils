// Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/getBotList": {
            "get": {
                "description": "获取机器人qq号",
                "responses": {}
            }
        },
        "/api/getFriendList": {
            "get": {
                "description": "获取好友列表",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 123456,
                        "description": "机器人qq号",
                        "name": "selfId",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/getGroupList": {
            "get": {
                "description": "获取群列表",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 123456,
                        "description": "机器人qq号",
                        "name": "selfId",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/getPermCode": {
            "get": {
                "description": "授权码",
                "responses": {}
            }
        },
        "/api/getRequestList": {
            "get": {
                "description": "获取所有的请求",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 123456,
                        "description": "机器人qq号",
                        "name": "selfId",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/getUserInfo": {
            "get": {
                "description": "获得用户信息",
                "responses": {}
            }
        },
        "/api/handleRequest": {
            "post": {
                "description": "处理一个请求",
                "parameters": [
                    {
                        "type": "string",
                        "default": "abc",
                        "description": "事件id",
                        "name": "flag",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "abc",
                        "description": "原因",
                        "name": "reason",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "default": true,
                        "description": "是否同意",
                        "name": "approve",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/api/login": {
            "post": {
                "description": "前端登录",
                "parameters": [
                    {
                        "type": "string",
                        "default": "xiaoguofan",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "123456",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/logout": {
            "get": {
                "description": "登出",
                "responses": {}
            }
        },
        "/api/manage/getAllPlugin": {
            "get": {
                "description": "获取所有插件的状态",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "群号",
                        "name": "groupId",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/manage/getPlugin": {
            "get": {
                "description": "获取某个插件的状态",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "群号",
                        "name": "groupId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "antibuse",
                        "description": "插件名",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/manage/updateAllPluginStatus": {
            "post": {
                "description": "更改某群所有插件状态",
                "parameters": [
                    {
                        "description": "修改插件状态入参",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.AllPluginStatusParams"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/manage/updatePluginStatus": {
            "post": {
                "description": "更改某一个插件状态",
                "parameters": [
                    {
                        "description": "修改插件状态入参",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.PluginStatusParams"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/manage/updateResponseStatus": {
            "post": {
                "description": "更改某一个群响应",
                "parameters": [
                    {
                        "description": "修改群响应入参",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.ResponseStatusParams"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/sendMsg": {
            "post": {
                "description": "前端调用发送信息",
                "parameters": [
                    {
                        "description": "发消息参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.SendMsgParams"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "types.AllPluginStatusParams": {
            "type": "object",
            "properties": {
                "groupId": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "types.PluginStatusParams": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "groupId": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "types.ResponseStatusParams": {
            "type": "object",
            "properties": {
                "groupId": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "types.SendMsgParams": {
            "type": "object",
            "properties": {
                "gidList": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "message": {
                    "type": "string"
                },
                "selfId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "zbp api",
	Description:      "zbp restful api document",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
