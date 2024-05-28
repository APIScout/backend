package operations

import (
	"strings"
	"testing"

	"backend/app/internal/retrieval"
)

func TestContainsOperatorString(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.source~="git"`})

	if filters == nil &&
		strings.Compare(filters[0].Lhs, "api.source") != 0 &&
		strings.Compare(filters[0].Operator, "~=") != 0 &&
		strings.Compare(filters[0].Rhs, `"git"`) != 0 &&
		err == nil {
		t.Fatal(filters)
	}
}

func TestContainsOperatorVersion(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.raw~=3.0`})

	if filters == nil &&
		strings.Compare(filters[0].Lhs, "api.source") != 0 &&
		strings.Compare(filters[0].Operator, "~=") != 0 &&
		strings.Compare(filters[0].Rhs, `"3.0"`) != 0 &&
		err == nil {
		t.Fatal(filters)
	}
}
