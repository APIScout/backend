package embedding

import (
	"os"
	"testing"

	"backend/app/internal/embedding"
)

func TestPipelineEmptyString(t *testing.T) {
	if os.Getenv("MODELS_HOST") == "" {
		t.Skip("Skipping testing in CI environment")
	}

	fragments := []string{""}
	res, _, _ := embedding.PerformPipeline(fragments, true)

	if res == nil {
		t.Fatal(res)
	}
}

func TestPipelineEmptyArray(t *testing.T) {
	if os.Getenv("MODELS_HOST") == "" {
		t.Skip("Skipping testing in CI environment")
	}

	var fragments []string
	res, _, _ := embedding.PerformPipeline(fragments, true)

	if res != nil {
		t.Fatal(res)
	}
}

func TestPipelineQuery(t *testing.T) {
	if os.Getenv("MODELS_HOST") == "" {
		t.Skip("Skipping testing in CI environment")
	}

	fragments := []string{"this is a test query"}
	res, _, _ := embedding.PerformPipeline(fragments, true)

	if len(res.Predictions) != 1 || len(res.Predictions[0]) != 512 {
		t.Fatal(res)
	}
}

func TestPipelineQueries(t *testing.T) {
	if os.Getenv("MODELS_HOST") == "" {
		t.Skip("Skipping testing in CI environment")
	}

	fragments := []string{"this is a test query", "this is another different test query"}
	res, _, _ := embedding.PerformPipeline(fragments, true)

	if len(res.Predictions) != 2 || len(res.Predictions[0]) != 512 || len(res.Predictions[1]) != 512 {
		t.Fatal(res)
	}
}

func TestPipelineQueries1(t *testing.T) {
	if os.Getenv("MODELS_HOST") == "" {
		t.Skip("Skipping testing in CI environment")
	}

	fragments := []string{"this is a https://twitter.com/_geodatasource/profile_image?size=original query", "this is another http://www.google.com"}
	res, _, _ := embedding.PerformPipeline(fragments, true)

	if len(res.Predictions) != 2 || len(res.Predictions[0]) != 512 || len(res.Predictions[1]) != 512 {
		t.Fatal(res)
	}
}

func TestPipelineDocument(t *testing.T) {
	if os.Getenv("MODELS_HOST") == "" {
		t.Skip("Skipping testing in CI environment")
	}

	fragments := []string{"{\n  \"openapi\": \"3.0.0\",\n  \"info\": {\n    \"title\": \"test title\",\n    \"description\": \"test description\",\n    \"version\": \"0.0.1\",\n    \"summary\": 'test \"summary\"'\n  }\n}"}
	res, _, _ := embedding.PerformPipeline(fragments, false)

	if len(res.Predictions) != 1 || len(res.Predictions[0]) != 512 {
		t.Fatal(res)
	}
}

func TestPipelineDocuments(t *testing.T) {
	if os.Getenv("MODELS_HOST") == "" {
		t.Skip("Skipping testing in CI environment")
	}

	fragments := []string{"{\n  \"openapi\": \"3.0.0\",\n  \"info\": {\n    \"title\": \"test title\",\n    \"description\": \"test description\",\n    \"version\": \"0.0.1\",\n    \"summary\": 'test \"summary\"'\n  }\n}", "\"version\": \"0.0.1\""}
	res, _, _ := embedding.PerformPipeline(fragments, false)

	if len(res.Predictions) != 2 || len(res.Predictions[0]) != 512 || len(res.Predictions[1]) != 512 {
		t.Fatal(res)
	}
}
