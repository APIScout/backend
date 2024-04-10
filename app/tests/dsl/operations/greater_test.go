package operations

import (
	"strings"
	"testing"

	"backend/app/internal/retrieval"
)

func TestGreaterOperatorInteger(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.major>4`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[0].Operator, ">") != 0 ||
		strings.Compare(filters[0].Rhs, "4") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestGreaterOperatorVersion(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.raw>4.0.0`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[0].Operator, ">") != 0 ||
		strings.Compare(filters[0].Rhs, "4.0.0") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestGreaterEqualOperatorInteger(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.major>=4`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[0].Operator, ">=") != 0 ||
		strings.Compare(filters[0].Rhs, "4") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestGreaterEqualOperatorVersion(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.raw>=4.0.0`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[0].Operator, ">=") != 0 ||
		strings.Compare(filters[0].Rhs, "4.0.0") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}
