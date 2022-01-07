// Code generated by go-swagger; DO NOT EDIT.

package restapi

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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Dell",
    "version": "1.0"
  },
  "host": "localhost:8081",
  "paths": {
    "/user": {
      "post": {
        "description": "Create a new user.",
        "tags": [
          "user"
        ],
        "summary": "Create New User",
        "operationId": "post-user",
        "parameters": [
          {
            "x-examples": {},
            "description": "Post the necessary fields for the API to create a new user.",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User Created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Missing Required Information",
            "schema": {
              "type": "null"
            }
          },
          "409": {
            "description": "Email Already Taken",
            "schema": {
              "type": "null"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "description": "Retrieve the information of the user with the matching user ID.",
        "tags": [
          "user"
        ],
        "summary": "Get User Info by User ID",
        "operationId": "get-users-userId",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "User Found",
            "schema": {
              "$ref": "#/definitions/User"
            },
            "examples": {
              "Get User Alice Smith": {
                "dateOfBirth": "1997-10-31",
                "email": "alice.smith@gmail.com",
                "emailVerified": true,
                "firstName": "Alice",
                "id": 142,
                "lastName": "Smith",
                "signUpDate": "2019-08-24"
              }
            }
          },
          "404": {
            "description": "User Not Found",
            "schema": {
              "type": "null"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "title": "User",
      "required": [
        "firstName",
        "lastName",
        "email",
        "dateOfBirth",
        "emailVerified"
      ],
      "properties": {
        "createDate": {
          "description": "The date that the user was created.",
          "type": "string",
          "format": "date"
        },
        "dateOfBirth": {
          "type": "string",
          "format": "date",
          "example": "1997-10-31"
        },
        "email": {
          "type": "string",
          "format": "email"
        },
        "emailVerified": {
          "description": "Set to true if the user's email has been verified.",
          "type": "boolean"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "description": "Unique identifier for the given user.",
          "type": "integer"
        },
        "lastName": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "name": "user"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Dell",
    "version": "1.0"
  },
  "host": "localhost:8081",
  "paths": {
    "/user": {
      "post": {
        "description": "Create a new user.",
        "tags": [
          "user"
        ],
        "summary": "Create New User",
        "operationId": "post-user",
        "parameters": [
          {
            "x-examples": {},
            "description": "Post the necessary fields for the API to create a new user.",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User Created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Missing Required Information",
            "schema": {
              "type": "null"
            }
          },
          "409": {
            "description": "Email Already Taken",
            "schema": {
              "type": "null"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "description": "Retrieve the information of the user with the matching user ID.",
        "tags": [
          "user"
        ],
        "summary": "Get User Info by User ID",
        "operationId": "get-users-userId",
        "parameters": [
          {
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "User Found",
            "schema": {
              "$ref": "#/definitions/User"
            },
            "examples": {
              "Get User Alice Smith": {
                "dateOfBirth": "1997-10-31",
                "email": "alice.smith@gmail.com",
                "emailVerified": true,
                "firstName": "Alice",
                "id": 142,
                "lastName": "Smith",
                "signUpDate": "2019-08-24"
              }
            }
          },
          "404": {
            "description": "User Not Found",
            "schema": {
              "type": "null"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "title": "User",
      "required": [
        "firstName",
        "lastName",
        "email",
        "dateOfBirth",
        "emailVerified"
      ],
      "properties": {
        "createDate": {
          "description": "The date that the user was created.",
          "type": "string",
          "format": "date"
        },
        "dateOfBirth": {
          "type": "string",
          "format": "date",
          "example": "1997-10-31"
        },
        "email": {
          "type": "string",
          "format": "email"
        },
        "emailVerified": {
          "description": "Set to true if the user's email has been verified.",
          "type": "boolean"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "description": "Unique identifier for the given user.",
          "type": "integer"
        },
        "lastName": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "name": "user"
    }
  ]
}`))
}
