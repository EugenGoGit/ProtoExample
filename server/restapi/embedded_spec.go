package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Key.proto",
    "version": "version not set"
  },
  "paths": {
    "/v1/keys": {
      "get": {
        "tags": [
          "KeyService"
        ],
        "summary": "Get key list (with explicit page number)",
        "operationId": "KeyService_GetPage",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "name": "page_number",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "page_size",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/exampleKeyPageResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        }
      },
      "post": {
        "tags": [
          "KeyService"
        ],
        "summary": "Adding a new key",
        "operationId": "KeyService_AddKey",
        "parameters": [
          {
            "description": "The key resource to create.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/exampleV1RfidDto"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/exampleV1RfidDto"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        }
      }
    },
    "/v1/keys-continued-list": {
      "get": {
        "tags": [
          "KeyService"
        ],
        "summary": "Get continued key list by token (another GetPage variant)",
        "operationId": "KeyService_GetPageContinued",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "name": "page_size",
            "in": "query"
          },
          {
            "type": "string",
            "name": "page_token",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/exampleListKeysResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        }
      }
    },
    "/v1/keys/{id}": {
      "get": {
        "tags": [
          "KeyService"
        ],
        "summary": "Get key by ID",
        "operationId": "KeyService_GetKey",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/exampleV1RfidDto"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "KeyService"
        ],
        "summary": "Delete key by ID",
        "operationId": "KeyService_DeleteKey",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/exampleV1RfidDto"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "exampleAddingKeyStatus": {
      "type": "object",
      "title": "Represents the status of key adding",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int32"
        },
        "status": {
          "type": "boolean"
        }
      }
    },
    "exampleKeyPageResponse": {
      "type": "object",
      "title": "Response for GetPage",
      "properties": {
        "count": {
          "type": "integer",
          "format": "int32"
        },
        "is_first": {
          "type": "boolean"
        },
        "is_last": {
          "type": "boolean"
        },
        "keys": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/exampleV1RfidDto"
          }
        },
        "page_number": {
          "type": "integer",
          "format": "int32"
        },
        "page_size": {
          "type": "integer",
          "format": "int32"
        },
        "total_count": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "exampleListKeysResponse": {
      "type": "object",
      "title": "Response for GetPageContinued",
      "properties": {
        "keys": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/exampleV1RfidDto"
          }
        },
        "next_page_token": {
          "type": "string"
        }
      }
    },
    "exampleV1RfidDto": {
      "type": "object",
      "title": "Represents the key",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int32"
        },
        "key": {
          "type": "string"
        }
      }
    },
    "protobufAny": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "runtimeError": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufAny"
          }
        },
        "error": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "runtimeStreamError": {
      "type": "object",
      "properties": {
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufAny"
          }
        },
        "grpc_code": {
          "type": "integer",
          "format": "int32"
        },
        "http_code": {
          "type": "integer",
          "format": "int32"
        },
        "http_status": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}