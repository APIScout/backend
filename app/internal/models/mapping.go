package models

var mapping = `{
	"mappings": {
		"properties": {
			"metadata": {
                "type": "nested",
                "properties": {
    				"mongo-id": { "type": "text" },
    				"name":     { "type": "text" },
    				"api-id":   { "type": "text" },
    				"commits":  { "type": "short" },
    				"latest":   { "type": "boolean" },
    				"source":   { "type": "text" },
    				"api-version": {
    					"type": "nested",
    					"properties": {
    						"raw":         { "type": "version" },
    						"valid":       { "type": "boolean" },
    						"major":       { "type": "text" },
    						"minor":       { "type": "text" },
    						"patch":       { "type": "text" },
    						"prerelease":  { "type": "text" },
    						"build":       { "type": "text" }
    					}
    				},
    				"oas-version": {
    					"type": "nested",
    					"properties": {
    						"raw":         { "type": "version" },
    						"valid":       { "type": "boolean" },
    						"major":       { "type": "text" },
    						"minor":       { "type": "text" },
							"patch":       { "type": "text" },
    						"prerelease":  { "type": "text" },
    						"build":       { "type": "text" }
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
