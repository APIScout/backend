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
