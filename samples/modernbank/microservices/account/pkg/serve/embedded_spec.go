// Code generated by go-swagger; DO NOT EDIT.

package serve

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the Account Microservice, responsible for managing accounts and their balances in Modern Bank.",
    "title": "Account",
    "version": "1.0.0"
  },
  "host": "account",
  "basePath": "/v1",
  "paths": {
    "/accounts": {
      "get": {
        "description": "Lists all accounts for a given customer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Lists all accounts for a given customer",
        "operationId": "listAccounts",
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/account"
              }
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "404": {
            "description": "Owner not found",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      },
      "post": {
        "description": "Creates a new account for a given customer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Create a new account for a customer",
        "operationId": "createAccount",
        "parameters": [
          {
            "enum": [
              "cash",
              "saving"
            ],
            "type": "string",
            "name": "type",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/account"
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Owner of the account",
          "name": "owner",
          "in": "query",
          "required": true
        },
        {
          "type": "string",
          "name": "x-request-id",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-flags",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-parentspanid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-sampled",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-spanId",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-traceid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "b3",
          "in": "header"
        }
      ]
    },
    "/accounts/{number}/balance/{delta}": {
      "put": {
        "description": "Change the balance of an account",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Change the balance of an account",
        "operationId": "changeBalance",
        "parameters": [
          {
            "type": "number",
            "format": "int64",
            "description": "Account number",
            "name": "number",
            "in": "path",
            "required": true
          },
          {
            "type": "number",
            "format": "currency",
            "description": "The amount to change the balance by (+/-)",
            "name": "delta",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "404": {
            "description": "Account not found",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "x-request-id",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-flags",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-parentspanid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-sampled",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-spanId",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-traceid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "b3",
          "in": "header"
        }
      ]
    },
    "/health": {
      "post": {
        "description": "returns 200",
        "tags": [
          "health"
        ],
        "summary": "returns 200 to prove the service is alive",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "x-request-id",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-flags",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-parentspanid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-sampled",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-spanId",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-traceid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "b3",
          "in": "header"
        }
      ]
    }
  },
  "definitions": {
    "account": {
      "type": "object",
      "required": [
        "number",
        "balance",
        "owner",
        "type"
      ],
      "properties": {
        "balance": {
          "type": "number",
          "format": "currency"
        },
        "number": {
          "type": "integer",
          "format": "int64"
        },
        "owner": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "enum": [
            "cash",
            "saving"
          ]
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about Accounts",
      "name": "account"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the Account Microservice, responsible for managing accounts and their balances in Modern Bank.",
    "title": "Account",
    "version": "1.0.0"
  },
  "host": "account",
  "basePath": "/v1",
  "paths": {
    "/accounts": {
      "get": {
        "description": "Lists all accounts for a given customer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Lists all accounts for a given customer",
        "operationId": "listAccounts",
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/account"
              }
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "404": {
            "description": "Owner not found",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      },
      "post": {
        "description": "Creates a new account for a given customer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Create a new account for a customer",
        "operationId": "createAccount",
        "parameters": [
          {
            "enum": [
              "cash",
              "saving"
            ],
            "type": "string",
            "name": "type",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/account"
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Owner of the account",
          "name": "owner",
          "in": "query",
          "required": true
        },
        {
          "type": "string",
          "name": "x-request-id",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-flags",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-parentspanid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-sampled",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-spanId",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-traceid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "b3",
          "in": "header"
        }
      ]
    },
    "/accounts/{number}/balance/{delta}": {
      "put": {
        "description": "Change the balance of an account",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Change the balance of an account",
        "operationId": "changeBalance",
        "parameters": [
          {
            "type": "number",
            "format": "int64",
            "description": "Account number",
            "name": "number",
            "in": "path",
            "required": true
          },
          {
            "type": "number",
            "format": "currency",
            "description": "The amount to change the balance by (+/-)",
            "name": "delta",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "404": {
            "description": "Account not found",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "x-request-id",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-flags",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-parentspanid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-sampled",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-spanId",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-traceid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "b3",
          "in": "header"
        }
      ]
    },
    "/health": {
      "post": {
        "description": "returns 200",
        "tags": [
          "health"
        ],
        "summary": "returns 200 to prove the service is alive",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "x-request-id",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-flags",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-parentspanid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-sampled",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-spanId",
          "in": "header"
        },
        {
          "type": "string",
          "name": "x-b3-traceid",
          "in": "header"
        },
        {
          "type": "string",
          "name": "b3",
          "in": "header"
        }
      ]
    }
  },
  "definitions": {
    "account": {
      "type": "object",
      "required": [
        "number",
        "balance",
        "owner",
        "type"
      ],
      "properties": {
        "balance": {
          "type": "number",
          "format": "currency"
        },
        "number": {
          "type": "integer",
          "format": "int64"
        },
        "owner": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "enum": [
            "cash",
            "saving"
          ]
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about Accounts",
      "name": "account"
    }
  ]
}`))
}
