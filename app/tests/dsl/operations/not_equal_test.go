package operations

import (
	"strings"
	"testing"

	"backend/app/internal/retrieval"
)

func TestNotEqualOperatorString(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.source!="github"`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.source") != 0 ||
		strings.Compare(filters[0].Operator, "!=") != 0 ||
		strings.Compare(filters[0].Rhs, "github") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestNotEqualOperatorBool(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.valid!=true`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.valid") != 0 ||
		strings.Compare(filters[0].Operator, "!=") != 0 ||
		strings.Compare(filters[0].Rhs, "true") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestNotEqualOperatorInt(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.major!=5`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[0].Operator, "!=") != 0 ||
		strings.Compare(filters[0].Rhs, "5") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestNotEqualOperatorVersion(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.raw!=3.0.0`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[0].Operator, "!=") != 0 ||
		strings.Compare(filters[0].Rhs, "3.0.0") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}
