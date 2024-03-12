package embedding

import (
	"strings"
	"testing"

	"backend/app/internal/embedding"
)

func TestStopwordsEmptyString(t *testing.T) {
	fragments := []string{""}
	res := embedding.StopWordRemoval(fragments)

	if res[0] != "" {
		t.Fatal(res)
	}
}

func TestStopwordsEmptyArray(t *testing.T) {
	var fragments []string
	res := embedding.StopWordRemoval(fragments)

	if len(res) != 0 {
		t.Fatal(res)
	}
}

func TestStopwordsCorrectlyFormattedQuery(t *testing.T) {
	fragments := []string{"this is a test query"}
	res := embedding.StopWordRemoval(fragments)

	if len(res) != 1 || strings.Compare(res[0], "test query") != 0 {
		t.Fatal(res)
	}
}

func TestStopwordsCorrectlyFormattedQueries(t *testing.T) {
	fragments := []string{"this is a test query", "this is another different test query"}
	res := embedding.StopWordRemoval(fragments)

	if len(res) != 2 ||
		strings.Compare(res[0], "test query") != 0 ||
		strings.Compare(res[1], "different test query") != 0 {

		t.Fatal(res)
	}
}

func TestStopwordsIncorrectlyFormattedQueries(t *testing.T) {
	fragments := []string{"**incorrectly** !formatted   test ..query..   "}
	res := embedding.StopWordRemoval(fragments)

	if len(fragments) != 1 || strings.Compare(res[0], "incorrectly formatted test query") != 0 {
		t.Fatal(res)
	}
}
