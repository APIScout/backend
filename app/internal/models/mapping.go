package models

var mapping = `{
  "mappings": {
    "properties": {
      "metadata": {
        "type": "nested",
        "properties": {
          "mongo-id": { "type": "text" },
          "length": { "type": "long" },
          "date": { "type": "date" },
          "api": {
            "type": "nested",
            "properties": {
              "name": { "type": "text" },
              "id": { "type": "text" },
              "commits": { "type": "long" },
              "latest": { "type": "boolean" },
              "source": { "type": "text" },
              "version": {
                "type": "nested",
                "properties": {
                  "raw": { "type": "version" },
                  "valid": { "type": "boolean" },
                  "major": { "type": "long" },
                  "minor": { "type": "long" },
                  "patch": { "type": "long" },
                  "prerelease": { "type": "text" },
                  "build": { "type": "text" }
                }
              }
			}
          },
          "metrics": {
            "type": "nested",
            "properties": {
              "security": {
                "type": "nested",
                "properties": {
                  "endpoints": { "type": "long" }
                }
              },
              "schema": {
                "type": "nested",
                "properties": {
                  "models": { "type": "long" },
                  "properties": { "type": "long" }
                }
              },
              "structure": {
                "type": "nested",
                "properties": {
                  "paths": { "type": "long" },
                  "operations": { "type": "long" },
                  "methods": { "type": "long" }
                }
              }
            }
          },
          "specification": {
            "type": "nested",
            "properties": {
              "type": { "type": "text" },
              "version": {
                "type": "nested",
                "properties": {
                  "raw": { "type": "version" },
                  "valid": { "type": "boolean" },
                  "major": { "type": "long" },
                  "minor": { "type": "long" },
                  "patch": { "type": "long" },
                  "prerelease": { "type": "text" },
                  "build": { "type": "text" }
                }
              }
            }
          }
        }
      },
      "embedding": {
        "type": "dense_vector",
        "dims": 512
      }
    }
  }
}`
