package dsl

import (
	"testing"

	"backend/app/internal/retrieval"
)

func TestTesting (t *testing.T) {
	// Add range operator to version
	filters, _ := retrieval.CreateFilters([]string{`api.version.raw<>[2.0.0, 5.0.0]`, `api.source=="github"`})
	retrieval.CreateEsFilter(filters)
}
