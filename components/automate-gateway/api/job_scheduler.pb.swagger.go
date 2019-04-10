package api

func init() {
	Swagger.Add("job_scheduler", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/ingest/job_scheduler.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/retention/nodes/delete-nodes/config": {
      "post": {
        "operationId": "ConfigureDeleteNodesScheduler",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseConfigureDeleteNodesScheduler"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestSchedulerConfig"
            }
          }
        ],
        "tags": [
          "JobScheduler"
        ]
      }
    },
    "/retention/nodes/missing-nodes-deletion/config": {
      "post": {
        "operationId": "ConfigureMissingNodesForDeletionScheduler",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseConfigureMissingNodesForDeletionScheduler"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestSchedulerConfig"
            }
          }
        ],
        "tags": [
          "JobScheduler"
        ]
      }
    },
    "/retention/nodes/missing-nodes/config": {
      "post": {
        "operationId": "ConfigureNodesMissingScheduler",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseConfigureNodesMissingScheduler"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestSchedulerConfig"
            }
          }
        ],
        "tags": [
          "JobScheduler"
        ]
      }
    },
    "/retention/nodes/status": {
      "get": {
        "operationId": "GetStatusJobScheduler",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseJobSchedulerStatus"
            }
          }
        },
        "tags": [
          "JobScheduler"
        ]
      }
    }
  },
  "definitions": {
    "requestSchedulerConfig": {
      "type": "object",
      "properties": {
        "every": {
          "type": "string"
        },
        "threshold": {
          "type": "string"
        },
        "running": {
          "type": "boolean",
          "format": "boolean"
        }
      },
      "description": "SchedulerConfig\nThe job message to configure the Delete Node Job\nevery - It accepts '1h30m', '1m', '2h30m', ..."
    },
    "responseConfigureDeleteNodesScheduler": {
      "type": "object"
    },
    "responseConfigureMissingNodesForDeletionScheduler": {
      "type": "object"
    },
    "responseConfigureNodesMissingScheduler": {
      "type": "object"
    },
    "responseJob": {
      "type": "object",
      "properties": {
        "running": {
          "type": "boolean",
          "format": "boolean"
        },
        "name": {
          "type": "string"
        },
        "every": {
          "type": "string"
        },
        "last_run": {
          "type": "string"
        },
        "next_run": {
          "type": "string"
        },
        "last_elapsed": {
          "type": "string"
        },
        "started_on": {
          "type": "string"
        },
        "threshold": {
          "type": "string"
        }
      }
    },
    "responseJobSchedulerStatus": {
      "type": "object",
      "properties": {
        "running": {
          "type": "boolean",
          "format": "boolean"
        },
        "jobs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/responseJob"
          }
        }
      }
    }
  }
}
`)
}
