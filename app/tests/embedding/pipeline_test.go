package embedding

import (
	"testing"

	"backend/app/internal/embedding"
)

func TestPipelineEmptyString(t *testing.T) {
	fragments := []string{""}
	res := embedding.PerformPipeline(fragments, true)

	if len(res.Predictions) != 1 || len(res.Predictions[0]) != 512 {
		t.Fatal(res)
	}
}

func TestPipelineEmptyArray(t *testing.T) {
	var fragments []string
	res := embedding.PerformPipeline(fragments, true)

	if len(res.Predictions) != 0 {
		t.Fatal(res)
	}
}

func TestPipelineQuery(t *testing.T) {
	fragments := []string{"this is a test query"}
	res := embedding.PerformPipeline(fragments, true)

	if len(res.Predictions) != 1 || len(res.Predictions[0]) != 512 {
		t.Fatal(res)
	}
}

func TestPipelineQueries(t *testing.T) {
	fragments := []string{"this is a test query", "this is another different test query"}
	res := embedding.PerformPipeline(fragments, true)

	if len(res.Predictions) != 2 || len(res.Predictions[0]) != 512 || len(res.Predictions[1]) != 512 {
		t.Fatal(res)
	}
}

func TestPipelineDocument(t *testing.T) {
	fragments := []string{"{\n  \"openapi\": \"3.0.0\",\n  \"info\": {\n    \"title\": \"test title\",\n    \"description\": \"test description\",\n    \"version\": \"0.0.1\",\n    \"summary\": 'test \"summary\"'\n  }\n}"}
	res := embedding.PerformPipeline(fragments, false)

	if len(res.Predictions) != 1 || len(res.Predictions[0]) != 512 {
		t.Fatal(res)
	}
}

func TestPipelineDocuments(t *testing.T) {
	fragments := []string{"{\n  \"openapi\": \"3.0.0\",\n  \"info\": {\n    \"title\": \"test title\",\n    \"description\": \"test description\",\n    \"version\": \"0.0.1\",\n    \"summary\": 'test \"summary\"'\n  }\n}", "\"version\": \"0.0.1\""}
	res := embedding.PerformPipeline(fragments, false)

	if len(res.Predictions) != 2 || len(res.Predictions[0]) != 512 || len(res.Predictions[1]) != 512 {
		t.Fatal(res)
	}
}
