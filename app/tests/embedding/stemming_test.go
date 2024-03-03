package embedding

import (
	"strings"
	"testing"

	"backend/app/internal/embedding"
)


func TestStemmingEmptyString(t *testing.T) {
	fragments := []string{""}
	res := embedding.Stemming(fragments)

	if res[0] != "" {
		t.Fatal(res)
	}
}

func TestStemmingEmptyArray(t *testing.T)  {
	var fragments []string
	res := embedding.Stemming(fragments)

	if len(res) != 0 {
		t.Fatal(res)
	}
}

func TestStemmingCorrectlyFormattedQuery(t *testing.T) {
	fragments := []string{"this is a test query"}
	res := embedding.Stemming(fragments)

	if len(res) != 1 || strings.Compare(res[0], "this is a test queri") != 0 {
		t.Fatal(res)
	}
}

func TestStemmingCorrectlyFormattedQueries(t *testing.T) {
	fragments := []string{"this is a test query", "this is another different test query"}
	res := embedding.Stemming(fragments)

	if len(res) != 2 ||
		strings.Compare(res[0], "this is a test queri") != 0 ||
		strings.Compare(res[1], "this is anoth differ test queri") != 0 {

		t.Fatal(res)
	}
}

func TestStemmingIncorrectlyFormattedQueries(t *testing.T) {
	fragments := []string{"**incorrectly** !formatted   test ..query..   "}
	res := embedding.Stemming(fragments)

	if len(fragments) != 1 || strings.Compare(res[0], "incorrect format test queri") != 0 {
		t.Fatal(res)
	}
}

