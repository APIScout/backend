package embedding

import (
	"strings"
	"testing"

	"backend/app/internal/embedding"
)

func TestExtractEmptyString(t *testing.T) {
	fragments := []string{""}
	res := embedding.ExtractTags(fragments)

	if res[0] != "" {
		t.Fatal(res)
	}
}

func TestExtractEmptyArray(t *testing.T) {
	var fragments []string
	res := embedding.ExtractTags(fragments)

	if len(res) != 0 {
		t.Fatal(res)
	}
}

func TestExtractDocument(t *testing.T) {
	fragments := []string{"{\n  \"openapi\": \"3.0.0\",\n  \"info\": {\n    \"title\": \"test title\",\n    \"description\": \"test description\",\n    \"version\": \"0.0.1\",\n    \"summary\": 'test \"summary\"'\n  }\n}"}
	res := embedding.ExtractTags(fragments)

	if len(res) != 1 || strings.Compare(res[0], "test title test description test \"summary\"") != 0 {
		t.Fatal(res)
	}
}

func TestExtractDocuments(t *testing.T) {
	fragments := []string{"{\n  \"openapi\": \"3.0.0\",\n  \"info\": {\n    \"title\": \"test title\",\n    \"description\": \"test description\",\n    \"version\": \"0.0.1\",\n    \"summary\": 'test \"summary\"'\n  }\n}", "\"version\": \"0.0.1\""}
	res := embedding.ExtractTags(fragments)

	if len(res) != 2 ||
		strings.Compare(res[0], "test title test description test \"summary\"") != 0 ||
		strings.Compare(res[1], "") != 0 {

		t.Fatal(res)
	}
}
