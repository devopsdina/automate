package api

func init() {
	Swagger.Add("cfgmgmt", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/cfgmgmt/cfgmgmt.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v0/cfgmgmt/errors": {
      "get": {
        "summary": "GetErrors",
        "description": "Returns a list of the most common errors reported for infra nodes' most recent Chef Infra Client runs.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetErrors",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.Errors"
            }
          }
        },
        "parameters": [
          {
            "name": "size",
            "description": "The number of results to return.\nIf set to zero, the default size of 10 will be used. Set to a negative\nvalue for unlimited results.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "filter",
            "description": "Filters in the request select the nodes from which the errors are\ncollected. The same filters may be specified for this request as for other\nNodes requests, with the exception of 'status' which is not valid for this\nrequest.",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/node_metadata_counts": {
      "get": {
        "summary": "GetNodeMetadataCounts",
        "description": "For each type of field requested this returns distinct values the amount of each. For example, \nif the 'platform' field is requested 'windows' 10, 'redhat' 5, and 'ubuntu' 8 could be returned. \nThe number next to each represents the number of nodes with that type of platform.\n\nExample:\nrequest\n` + "`" + `` + "`" + `` + "`" + `\ncfgmgmt/node_metadata_counts?type=platform\u0026type=status\n` + "`" + `` + "`" + `` + "`" + `\nresponse\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"types\": [\n{\n\"values\": [\n{\n\"value\": \"mac_os_x 10.11.5\",\n\"count\": 28\n},\n{\n\"value\": \"linux 8.9\",\n\"count\": 1\n},\n{\n\"value\": \"macos 8.9\",\n\"count\": 1\n},\n{\n\"value\": \"windows 8.9\",\n\"count\": 1\n}\n],\n\"type\": \"platform\"\n},\n{\n\"value\": [\n{\n\"value\": \"missing\",\n\"count\": 29\n},\n{\n\"value\": \"failure\",\n\"count\": 2\n}\n],\n\"type\": \"status\"\n}\n]\n}\n` + "`" + `` + "`" + `` + "`" + `\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetNodeMetadataCounts",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.NodeMetadataCounts"
            }
          }
        },
        "parameters": [
          {
            "name": "type",
            "description": "Types of node fields to collect value counts for.",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "filter",
            "description": "Filters to apply to the counts returned.",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "start",
            "description": "Earliest most recent check-in node information to return.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "end",
            "description": "Latest most recent check-in node information to return.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/node_runs_daily_status_time_series": {
      "get": {
        "summary": "GetNodeRunsDailyStatusTimeSeries",
        "description": "Provides the status of runs for each 24-hour duration. For multiple runs in one 24-hour duration, \nthe most recent failed run will be returned. If there are no failed runs the most recent successful \nrun will be returned. If no runs are found in the 24-hour duration, the status will be \"missing\" \nand no run information will be returned.\n\nExample:\nrequest\n` + "`" + `` + "`" + `` + "`" + `\ncfgmgmt/node_runs_daily_status_time_series?node_id=507bd518-5c18-4c2d-a445-60fe7dde9961\u0026days_ago=3\n` + "`" + `` + "`" + `` + "`" + `\nresponse\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"durations\": [\n{\n\"start\": \"2020-04-25T19:00:00Z\",\n\"end\": \"2020-04-26T18:59:59Z\",\n\"status\": \"missing\",\n\"run_id\": \"\"\n},\n{\n\"start\": \"2020-04-26T19:00:00Z\",\n\"end\": \"2020-04-27T18:59:59Z\",\n\"status\": \"missing\",\n\"run_id\": \"\"\n},\n{\n\"start\": \"2020-04-27T19:00:00Z\",\n\"end\": \"2020-04-28T18:59:59Z\",\n\"status\": \"failure\",\n\"run_id\": \"b7904f41-68b5-44ec-9da6-cf2481ff8600\"\n}\n]\n}\n` + "`" + `` + "`" + `` + "`" + `\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetNodeRunsDailyStatusTimeSeries",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.NodeRunsDailyStatusTimeSeries"
            }
          }
        },
        "parameters": [
          {
            "name": "node_id",
            "description": "Node ID of the runs.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "days_ago",
            "description": "Number of past days.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/nodes": {
      "get": {
        "summary": "GetNodes",
        "description": "Returns a list of infra nodes that have checked in to Automate.\nAdding a filter makes a list of all nodes that meet the filter criteria.\nFilters for the same field are ORd together, while filters across different fields are ANDed together.\nSupports pagination, filtering (with wildcard support), and sorting.\nLimited to 10k results.\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\ncfgmgmt/nodes?pagination.page=1\u0026pagination.size=100\u0026sorting.field=name\u0026sorting.order=ASC\u0026filter=name:mySO*\u0026filter=platform:ubun*\n` + "`" + `` + "`" + `` + "`" + `\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetNodes",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "type": "array",
              "items": {
                "type": "object"
              }
            }
          }
        },
        "parameters": [
          {
            "name": "filter",
            "description": "Filters to apply to the request for nodes list.",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "pagination.page",
            "description": "Page number of the results to return.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "pagination.size",
            "description": "Amount of results to include per page.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "sorting.field",
            "description": "Field to sort the list results on.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "sorting.order",
            "description": "Order the results should be returned in.",
            "in": "query",
            "required": false,
            "type": "string",
            "enum": [
              "ASC",
              "DESC"
            ],
            "default": "ASC"
          },
          {
            "name": "start",
            "description": "Earliest most recent check-in node information to return.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "end",
            "description": "Latest most recent check-in node information to return.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/nodes/{node_id}/attribute": {
      "get": {
        "summary": "GetAttributes",
        "description": "Returns the latest reported attributes for the provided node ID.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:get\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetAttributes",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.NodeAttribute"
            }
          }
        },
        "parameters": [
          {
            "name": "node_id",
            "description": "Chef guid for the requested node.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/nodes/{node_id}/runs": {
      "get": {
        "summary": "GetRuns",
        "description": "Returns a list of run metadata (id, start and end time, and status) for the provided node ID.\nSupports pagination.\nAccepts a ` + "`" + `start` + "`" + ` parameter to denote start date for the list and a filter of type ` + "`" + `status` + "`" + `.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetRuns",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "type": "array",
              "items": {
                "type": "object"
              }
            }
          }
        },
        "parameters": [
          {
            "name": "node_id",
            "description": "Chef guid for the node.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "filter",
            "description": "Filters to apply to the request for runs list.",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "pagination.page",
            "description": "Page number of the results to return.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "pagination.size",
            "description": "Amount of results to include per page.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "start",
            "description": "Earliest (in history) run information to return for the runs list.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "end",
            "description": "Latest (in history) run information to return for the runs list.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/nodes/{node_id}/runs/{run_id}": {
      "get": {
        "summary": "GetNodeRun",
        "description": "Returns the infra run report for the provided node ID and run ID.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:get\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetNodeRun",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.Run"
            }
          }
        },
        "parameters": [
          {
            "name": "node_id",
            "description": "Chef guid for the requested node.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "run_id",
            "description": "Run id for the node.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "end_time",
            "description": "End time on the node's run.",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "date-time"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/organizations": {
      "get": {
        "summary": "GetOrganizations",
        "description": "Returns a list of all organizations associated with nodes that have checked in to Automate.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetOrganizations",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "type": "array",
              "items": {
                "type": "object"
              }
            }
          }
        },
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/policy_revision/{revision_id}": {
      "get": {
        "summary": "GetPolicyCookbooks",
        "description": "Returns Policy Names with a list of cookbook names and associated policy identifiers based on a policy revision ID.\nPolicy revision ids are sent with an infra run report and identifies which instance of a policy the node used for this run.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetPolicyCookbooks",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.PolicyCookbooks"
            }
          }
        },
        "parameters": [
          {
            "name": "revision_id",
            "description": "Revision id for the policy.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/source_fqdns": {
      "get": {
        "summary": "GetSourceFqdns",
        "description": "Returns a list of all Chef Infra Servers associated with nodes that have checked in to Automate.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetSourceFqdns",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "type": "array",
              "items": {
                "type": "object"
              }
            }
          }
        },
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/stats/checkin_counts_timeseries": {
      "get": {
        "summary": "GetCheckinCountsTimeSeries",
        "description": "Returns a daily time series of unique node check-ins for the number of days requested.\nIf days ago value is empty, api will return the default 1 day ago results.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetCheckInCountsTimeSeries",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.CheckInCountsTimeSeries"
            }
          }
        },
        "parameters": [
          {
            "name": "filter",
            "description": "List of filters to be applied to the time series.",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "days_ago",
            "description": "Number of past days to create the time series.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/stats/missing_node_duration_counts": {
      "get": {
        "summary": "GetMissingNodeDurationCounts",
        "description": "Returns a count of missing nodes for the provided durations.\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\ncfgmgmt/stats/missing_node_duration_counts?durations=3d\u0026durations=1w\u0026durations=2w\u0026durations=1M\u0026durations=3M\n` + "`" + `` + "`" + `` + "`" + `\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetMissingNodeDurationCounts",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.MissingNodeDurationCounts"
            }
          }
        },
        "parameters": [
          {
            "name": "durations",
            "description": "A valid duration is any number zero or greater with one of these characters 'h', 'd', 'w', or 'M'. \n'h' is hours\n'd' is days\n'w' is weeks\n'M' is months\nExamples '12h', '3d', and '1M'.\nWill contain one or many.",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/stats/node_counts": {
      "get": {
        "summary": "GetNodesCounts",
        "description": "Returns totals for failed, success, missing, and overall total infra nodes that have reported into Automate.\nSupports filtering.\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\ncfgmgmt/stats/node_counts?filter=name:mySO*\u0026filter=platform:ubun*\n` + "`" + `` + "`" + `` + "`" + `\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetNodesCounts",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.NodesCounts"
            }
          }
        },
        "parameters": [
          {
            "name": "filter",
            "description": "List of filters to be applied to the node count results.",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "start",
            "description": "Earliest node check-in.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "end",
            "description": "Latest node check-in.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/stats/run_counts": {
      "get": {
        "summary": "GetRunsCounts",
        "description": "Returns totals for failed and successful runs given a ` + "`" + `node_id` + "`" + `.\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\ncfgmgmt/stats/run_counts?node_id=821fff07-abc9-4160-96b1-83d68ae5cfdd\u0026start=2019-11-02\n` + "`" + `` + "`" + `` + "`" + `\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetRunsCounts",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.RunsCounts"
            }
          }
        },
        "parameters": [
          {
            "name": "filter",
            "description": "List of filters to be applied to the run count results.",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "start",
            "description": "Earliest (in history) run information to return for the run counts.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "end",
            "description": "Latest (in history) run information to return for the run counts.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "node_id",
            "description": "Node id associated with the run.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/suggestions": {
      "get": {
        "summary": "GetSuggestions",
        "description": "Returns possible filter values given a valid ` + "`" + `type` + "`" + ` parameter. All values returned until two or more\ncharacters are provided for the ` + "`" + `text` + "`" + ` parameter.\nSupports wildcard (* and ?).\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\ncfgmgmt/suggestions?type=environment\u0026text=_d\n` + "`" + `` + "`" + `` + "`" + `\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ninfra:nodes:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetSuggestions",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "type": "array",
              "items": {
                "type": "object"
              }
            }
          }
        },
        "parameters": [
          {
            "name": "type",
            "description": "Field for which suggestions are being returned.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "text",
            "description": "Text to search on for the type value.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "filter",
            "description": "Filters to be applied to the results.",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          }
        ],
        "tags": [
          "ConfigMgmt"
        ]
      }
    },
    "/api/v0/cfgmgmt/version": {
      "get": {
        "operationId": "GetVersion",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.common.version.VersionInfo"
            }
          }
        },
        "tags": [
          "hidden"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.cfgmgmt.response.CheckInCounts": {
      "type": "object",
      "properties": {
        "start": {
          "type": "string"
        },
        "end": {
          "type": "string"
        },
        "count": {
          "type": "integer",
          "format": "int32"
        },
        "total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.CheckInCountsTimeSeries": {
      "type": "object",
      "properties": {
        "counts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.CheckInCounts"
          },
          "title": "List of daily checkin counts"
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.ChefError": {
      "type": "object",
      "properties": {
        "class": {
          "type": "string",
          "description": "Class for the error."
        },
        "message": {
          "type": "string",
          "description": "Error message for the failed run."
        },
        "backtrace": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Stacktrace for the failure."
        },
        "description": {
          "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.Description",
          "description": "Description for the error."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.CookbookLock": {
      "type": "object",
      "properties": {
        "cookbook": {
          "type": "string",
          "description": "Cookbook name."
        },
        "policy_identifier": {
          "type": "string",
          "description": "Policy identifier for the cookbook lock."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.CountedDuration": {
      "type": "object",
      "properties": {
        "duration": {
          "type": "string",
          "title": "Duration of the count. Example '3d'"
        },
        "count": {
          "type": "integer",
          "format": "int32",
          "title": "The number of nodes for this duration"
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.Deprecation": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "description": "Message for the deprecation."
        },
        "url": {
          "type": "string",
          "description": "Url reference for the deprecation."
        },
        "location": {
          "type": "string",
          "description": "Location of the deprecated code."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.Description": {
      "type": "object",
      "properties": {
        "title": {
          "type": "string",
          "description": "Title for the error description."
        },
        "sections": {
          "type": "array",
          "items": {
            "type": "object"
          },
          "description": "More information about the error."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.ErrorCount": {
      "type": "object",
      "properties": {
        "count": {
          "type": "integer",
          "format": "int32"
        },
        "type": {
          "type": "string"
        },
        "error_message": {
          "type": "string"
        }
      },
      "title": "ErrorCount gives the number of occurrences (count) of the error specified by\nthe type and message among the nodes included by the request parameters"
    },
    "chef.automate.api.cfgmgmt.response.Errors": {
      "type": "object",
      "properties": {
        "errors": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.ErrorCount"
          }
        }
      },
      "description": "Errors contains a list of the most common Chef Infra error type/message\ncombinations among nodes in the active project as filtered according to the\nrequest."
    },
    "chef.automate.api.cfgmgmt.response.ExpandedRunList": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Id of the run list collection."
        },
        "run_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.RunList"
          },
          "description": "Intentionally blank."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.MissingNodeDurationCounts": {
      "type": "object",
      "properties": {
        "counted_durations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.CountedDuration"
          },
          "title": "List of counted durations"
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.NodeAttribute": {
      "type": "object",
      "properties": {
        "node_id": {
          "type": "string",
          "description": "The chef_guid associated with the node."
        },
        "name": {
          "type": "string",
          "description": "Name of the node."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Run list for the node."
        },
        "chef_environment": {
          "type": "string",
          "description": "The environment for the node."
        },
        "normal": {
          "type": "string",
          "description": "Stringified json of the normal attributes for the node."
        },
        "default": {
          "type": "string",
          "description": "Stringified json of the default attributes for the node."
        },
        "override": {
          "type": "string",
          "description": "Stringified json of the override attributes for the node."
        },
        "automatic": {
          "type": "string",
          "description": "Stringified json of the automatic attributes for the node."
        },
        "normal_value_count": {
          "type": "integer",
          "format": "int32",
          "description": "Count of normal attributes on the node."
        },
        "default_value_count": {
          "type": "integer",
          "format": "int32",
          "description": "Count of default attributes on the node."
        },
        "override_value_count": {
          "type": "integer",
          "format": "int32",
          "description": "Count of override attributes on the node."
        },
        "all_value_count": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of attributes on the node."
        },
        "automatic_value_count": {
          "type": "integer",
          "format": "int32",
          "description": "Count of automatic attributes on the node."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.NodeMetadataCounts": {
      "type": "object",
      "properties": {
        "types": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.TypeCount"
          },
          "title": "Field Types for a node with value counts"
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.NodeRunsDailyStatusTimeSeries": {
      "type": "object",
      "properties": {
        "durations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.RunDurationStatus"
          },
          "title": "runs status of a 24-hour duration"
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.NodesCounts": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of nodes that have reported in to Automate."
        },
        "success": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of nodes that have reported in to Automate whose last run was successful."
        },
        "failure": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of nodes that have reported in to Automate whose last run was failed."
        },
        "missing": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of nodes that have been labeled as 'missing' as determined by node lifecycle settings."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.PolicyCookbooks": {
      "type": "object",
      "properties": {
        "policy_name": {
          "type": "string",
          "description": "Name of the policy."
        },
        "cookbook_locks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.CookbookLock"
          },
          "description": "Intentionally blank."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.Resource": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "description": "Resource type."
        },
        "name": {
          "type": "string",
          "description": "Name for the resource."
        },
        "id": {
          "type": "string",
          "description": "Id of the resource."
        },
        "duration": {
          "type": "string",
          "description": "Duration of the resource processing."
        },
        "delta": {
          "type": "string",
          "description": "Change diff for the resource (if it was changed during the run)."
        },
        "cookbook_name": {
          "type": "string",
          "description": "Cookbook name associated with the resource."
        },
        "cookbook_version": {
          "type": "string",
          "description": "Version of the cookbook associated with the resource."
        },
        "status": {
          "type": "string",
          "description": "Status of the resource (e.g. 'up-to-date')."
        },
        "recipe_name": {
          "type": "string",
          "description": "Name of the recipe associated with the resource."
        },
        "result": {
          "type": "string",
          "description": "String result of the resource."
        },
        "conditional": {
          "type": "string",
          "description": "Conditional rule associated with the resource."
        },
        "ignore_failure": {
          "type": "boolean",
          "format": "boolean",
          "description": "Boolean that denotes whether or not the resource failure should be ignored."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.Run": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Id of the infra node run."
        },
        "node_id": {
          "type": "string",
          "description": "The chef_guid associated with the node."
        },
        "node_name": {
          "type": "string",
          "description": "Name of the node."
        },
        "organization": {
          "type": "string",
          "description": "The organization the node is associated with."
        },
        "start_time": {
          "type": "string",
          "format": "date-time",
          "description": "Start time of the infra run."
        },
        "end_time": {
          "type": "string",
          "format": "date-time",
          "description": "End time of the infra run."
        },
        "source": {
          "type": "string",
          "description": "Source of the node run (e.g. chef-solo)."
        },
        "status": {
          "type": "string",
          "description": "Status of the infra node run."
        },
        "total_resource_count": {
          "type": "integer",
          "format": "int32",
          "description": "Resource count reported on the infra node run."
        },
        "updated_resource_count": {
          "type": "integer",
          "format": "int32",
          "description": "Count of resources updated in the infra node run."
        },
        "chef_version": {
          "type": "string",
          "description": "Chef-client version on the node."
        },
        "uptime_seconds": {
          "type": "integer",
          "format": "int32",
          "description": "Count in seconds that the node has been active."
        },
        "environment": {
          "type": "string",
          "description": "The environment for the node."
        },
        "fqdn": {
          "type": "string",
          "description": "FQDN of the node."
        },
        "source_fqdn": {
          "type": "string",
          "description": "Chef server associated with the node."
        },
        "ipaddress": {
          "type": "string",
          "description": "IP Address for the node."
        },
        "resources": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.Resource"
          },
          "description": "Intentionally blank."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Run list for the node."
        },
        "deprecations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.Deprecation"
          },
          "description": "Intentionally blank."
        },
        "error": {
          "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.ChefError",
          "description": "Chef Error information, available on failed runs."
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Unused field."
        },
        "resource_names": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of resource names for the node."
        },
        "recipes": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of recipes the node calls."
        },
        "chef_tags": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of tags associated with the node."
        },
        "cookbooks": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of cookbooks associated with the node."
        },
        "platform": {
          "type": "string",
          "description": "Full platform string for the node (family + version)."
        },
        "platform_family": {
          "type": "string",
          "description": "Platform family for the node."
        },
        "platform_version": {
          "type": "string",
          "description": "Platform version for the node."
        },
        "roles": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of roles associated with the node."
        },
        "policy_name": {
          "type": "string",
          "description": "Policy name associated with the node."
        },
        "policy_group": {
          "type": "string",
          "description": "Policy group associated with the node."
        },
        "policy_revision": {
          "type": "string",
          "description": "Policy revision associated with the node."
        },
        "expanded_run_list": {
          "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.ExpandedRunList",
          "description": "Expanded run list for the node."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects the node belongs to."
        },
        "versioned_cookbooks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.VersionedCookbook"
          },
          "description": "List of versioned cookbooks associated with the node."
        },
        "ip6address": {
          "type": "string",
          "description": "IP 6 Address for the node."
        },
        "timezone": {
          "type": "string",
          "title": "timezone of the node"
        },
        "domain": {
          "type": "string"
        },
        "hostname": {
          "type": "string"
        },
        "memory_total": {
          "type": "string"
        },
        "macaddress": {
          "type": "string"
        },
        "dmi_system_serial_number": {
          "type": "string"
        },
        "dmi_system_manufacturer": {
          "type": "string"
        },
        "virtualization_role": {
          "type": "string"
        },
        "virtualization_system": {
          "type": "string"
        },
        "kernel_version": {
          "type": "string"
        },
        "kernel_release": {
          "type": "string"
        },
        "cloud_provider": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.RunDurationStatus": {
      "type": "object",
      "properties": {
        "start": {
          "type": "string",
          "title": "Start of the duration (RFC3339)"
        },
        "end": {
          "type": "string",
          "title": "End of the duration (RFC3339)"
        },
        "status": {
          "type": "string",
          "title": "Prominent Run's status"
        },
        "run_id": {
          "type": "string",
          "title": "Prominent Run's ID"
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.RunList": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "description": "Type of run list item (e.g. 'recipe')."
        },
        "name": {
          "type": "string",
          "description": "Name of run list item."
        },
        "version": {
          "type": "string",
          "description": "Version of run list item."
        },
        "skipped": {
          "type": "boolean",
          "format": "boolean",
          "description": "Boolean denoting whether or not the run list item was skipped."
        },
        "children": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.RunList"
          },
          "description": "Intentionally blank."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.RunsCounts": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of run reports that have landed in Automate for the node."
        },
        "success": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of successful run reports that have landed in Automate for the node."
        },
        "failure": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of failed run reports that have landed in Automate for the node."
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.TypeCount": {
      "type": "object",
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cfgmgmt.response.ValueCount"
          },
          "title": "Values of the field type with a count for each"
        },
        "type": {
          "type": "string",
          "title": "The field type counted"
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.ValueCount": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string",
          "title": "The value counted"
        },
        "count": {
          "type": "integer",
          "format": "int32",
          "title": "The number of this distinct value"
        }
      }
    },
    "chef.automate.api.cfgmgmt.response.VersionedCookbook": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the cookbook."
        },
        "version": {
          "type": "string",
          "description": "Version of the cookbook."
        }
      }
    },
    "chef.automate.api.common.ExportData": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string",
          "format": "byte",
          "description": "Exported reports in JSON or CSV."
        }
      }
    },
    "chef.automate.api.common.query.Pagination": {
      "type": "object",
      "properties": {
        "page": {
          "type": "integer",
          "format": "int32",
          "description": "Page number of the results to return."
        },
        "size": {
          "type": "integer",
          "format": "int32",
          "description": "Amount of results to include per page."
        }
      }
    },
    "chef.automate.api.common.query.SortOrder": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC"
    },
    "chef.automate.api.common.query.Sorting": {
      "type": "object",
      "properties": {
        "field": {
          "type": "string",
          "description": "Field to sort the list results on."
        },
        "order": {
          "$ref": "#/definitions/chef.automate.api.common.query.SortOrder",
          "description": "Order the results should be returned in."
        }
      }
    },
    "chef.automate.api.common.version.VersionInfo": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "sha": {
          "type": "string"
        },
        "built": {
          "type": "string"
        }
      }
    },
    "google.protobuf.Any": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string",
          "description": "A URL/resource name that uniquely identifies the type of the serialized\nprotocol buffer message. This string must contain at least\none \"/\" character. The last segment of the URL's path must represent\nthe fully qualified name of the type (as in\n` + "`" + `path/google.protobuf.Duration` + "`" + `). The name should be in a canonical form\n(e.g., leading \".\" is not accepted).\n\nIn practice, teams usually precompile into the binary all types that they\nexpect it to use in the context of Any. However, for URLs which use the\nscheme ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, or no scheme, one can optionally set up a type\nserver that maps type URLs to message definitions as follows:\n\n* If no scheme is provided, ` + "`" + `https` + "`" + ` is assumed.\n* An HTTP GET on the URL must yield a [google.protobuf.Type][]\n  value in binary format, or produce an error.\n* Applications are allowed to cache lookup results based on the\n  URL, or have them precompiled into a binary to avoid any\n  lookup. Therefore, binary compatibility needs to be preserved\n  on changes to types. (Use versioned type names to manage\n  breaking changes.)\n\nNote: this functionality is not currently available in the official\nprotobuf release, and it is not used for type URLs beginning with\ntype.googleapis.com.\n\nSchemes other than ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` (or the empty scheme) might be\nused with implementation specific semantics."
        },
        "value": {
          "type": "string",
          "format": "byte",
          "description": "Must be a valid serialized protocol buffer of the above specified type."
        }
      },
      "description": "` + "`" + `Any` + "`" + ` contains an arbitrary serialized protocol buffer message along with a\nURL that describes the type of the serialized message.\n\nProtobuf library provides support to pack/unpack Any values in the form\nof utility functions or additional generated methods of the Any type.\n\nExample 1: Pack and unpack a message in C++.\n\n    Foo foo = ...;\n    Any any;\n    any.PackFrom(foo);\n    ...\n    if (any.UnpackTo(\u0026foo)) {\n      ...\n    }\n\nExample 2: Pack and unpack a message in Java.\n\n    Foo foo = ...;\n    Any any = Any.pack(foo);\n    ...\n    if (any.is(Foo.class)) {\n      foo = any.unpack(Foo.class);\n    }\n\n Example 3: Pack and unpack a message in Python.\n\n    foo = Foo(...)\n    any = Any()\n    any.Pack(foo)\n    ...\n    if any.Is(Foo.DESCRIPTOR):\n      any.Unpack(foo)\n      ...\n\n Example 4: Pack and unpack a message in Go\n\n     foo := \u0026pb.Foo{...}\n     any, err := ptypes.MarshalAny(foo)\n     ...\n     foo := \u0026pb.Foo{}\n     if err := ptypes.UnmarshalAny(any, foo); err != nil {\n       ...\n     }\n\nThe pack methods provided by protobuf library will by default use\n'type.googleapis.com/full.type.name' as the type URL and the unpack\nmethods only use the fully qualified type name after the last '/'\nin the type URL, for example \"foo.bar.com/x/y.z\" will yield type\nname \"y.z\".\n\n\nJSON\n====\nThe JSON representation of an ` + "`" + `Any` + "`" + ` value uses the regular\nrepresentation of the deserialized, embedded message, with an\nadditional field ` + "`" + `@type` + "`" + ` which contains the type URL. Example:\n\n    package google.profile;\n    message Person {\n      string first_name = 1;\n      string last_name = 2;\n    }\n\n    {\n      \"@type\": \"type.googleapis.com/google.profile.Person\",\n      \"firstName\": \u003cstring\u003e,\n      \"lastName\": \u003cstring\u003e\n    }\n\nIf the embedded message type is well-known and has a custom JSON\nrepresentation, that representation will be embedded adding a field\n` + "`" + `value` + "`" + ` which holds the custom JSON in addition to the ` + "`" + `@type` + "`" + `\nfield. Example (for message [google.protobuf.Duration][]):\n\n    {\n      \"@type\": \"type.googleapis.com/google.protobuf.Duration\",\n      \"value\": \"1.212s\"\n    }"
    },
    "google.protobuf.NullValue": {
      "type": "string",
      "enum": [
        "NULL_VALUE"
      ],
      "default": "NULL_VALUE",
      "description": "` + "`" + `NullValue` + "`" + ` is a singleton enumeration to represent the null value for the\n` + "`" + `Value` + "`" + ` type union.\n\n The JSON representation for ` + "`" + `NullValue` + "`" + ` is JSON ` + "`" + `null` + "`" + `.\n\n - NULL_VALUE: Null value."
    },
    "grpc.gateway.runtime.StreamError": {
      "type": "object",
      "properties": {
        "grpc_code": {
          "type": "integer",
          "format": "int32"
        },
        "http_code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "http_status": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/google.protobuf.Any"
          }
        }
      }
    }
  },
  "x-stream-definitions": {
    "chef.automate.api.common.ExportData": {
      "type": "object",
      "properties": {
        "result": {
          "$ref": "#/definitions/chef.automate.api.common.ExportData"
        },
        "error": {
          "$ref": "#/definitions/grpc.gateway.runtime.StreamError"
        }
      },
      "title": "Stream result of chef.automate.api.common.ExportData"
    }
  }
}
`)
}
