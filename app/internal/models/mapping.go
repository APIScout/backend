package models

var mapping = `{
  "mappings": {
    "properties": {
      "metadata": {
        "type": "nested",
        "properties": {
          "mongo-id": { "type": "text" },
          "api": {
            "type": "nested",
            "properties": {
              "name": { "type": "text" },
              "id": { "type": "text" },
              "commits": { "type": "short" },
              "latest": { "type": "boolean" },
              "source": { "type": "text" },
              "version": {
                "type": "nested",
                "properties": {
                  "raw": { "type": "version" },
                  "valid": { "type": "boolean" },
                  "major": { "type": "short" },
                  "minor": { "type": "short" },
                  "patch": { "type": "short" },
                  "prerelease": { "type": "text" },
                  "build": { "type": "text" }
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
                  "major": { "type": "short" },
                  "minor": { "type": "short" },
                  "patch": { "type": "short" },
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
