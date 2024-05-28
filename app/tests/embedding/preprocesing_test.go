package embedding

import (
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
