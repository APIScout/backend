package embedding

import (
	"strings"
	"testing"

	"backend/app/internal/embedding"
)

func TestPreprocessEmptyString(t *testing.T) {
	fragments := []string{""}
	res := embedding.PreprocessFragment(fragments, true)

	if res[0] != "" {
		t.Fatal(res)
	}
}

func TestPreprocessEmptyArray(t *testing.T) {
	var fragments []string
	res := embedding.PreprocessFragment(fragments, true)

	if len(res) != 0 {
		t.Fatal(res)
	}
}

func TestPreprocessQuery(t *testing.T) {
	fragments := []string{"this is a test query"}
	res := embedding.PreprocessFragment(fragments, true)

	if len(res) != 1 || strings.Compare(res[0], "test queri") != 0 {
		t.Fatal(res)
	}
}

func TestPreprocessQueries(t *testing.T) {
	fragments := []string{"this is a test query", "this is another different test query"}
	res := embedding.PreprocessFragment(fragments, true)

	if len(res) != 2 ||
		strings.Compare(res[0], "test queri") != 0 ||
		strings.Compare(res[1], "differ test queri") != 0 {

		t.Fatal(res)
	}
}

func TestPreprocessDocument(t *testing.T) {
	fragments := []string{"{\n  \"openapi\": \"3.0.0\",\n  \"info\": {\n    \"title\": \"test title\",\n    \"description\": \"test description\",\n    \"version\": \"0.0.1\",\n    \"summary\": 'test \"summary\"'\n  }\n}"}
	res := embedding.PreprocessFragment(fragments, false)

	if len(res) != 1 || strings.Compare(res[0], "test titl test descript test summari") != 0 {
		t.Fatal(res)
	}
}

func TestPreprocessDocuments(t *testing.T) {
	fragments := []string{"{\n  \"openapi\": \"3.0.0\",\n  \"info\": {\n    \"title\": \"test title\",\n    \"description\": \"test description\",\n    \"version\": \"0.0.1\",\n    \"summary\": 'test \"summary\"'\n  }\n}", "\"version\": \"0.0.1\""}
	res := embedding.PreprocessFragment(fragments, false)

	if len(res) != 2 ||
		strings.Compare(res[0], "test titl test descript test summari") != 0 ||
		strings.Compare(res[1], "") != 0 {

		t.Fatal(res)
	}
}
