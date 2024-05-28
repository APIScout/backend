package operations

import (
	"strings"
	"testing"

	"backend/app/internal/retrieval"
)

func TestRangeOperatorSquaresInteger(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.major<>[1,5]`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[0].Operator, ">=") != 0 ||
		strings.Compare(filters[0].Rhs, "1") != 0 ||
		strings.Compare(filters[1].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[1].Operator, "<=") != 0 ||
		strings.Compare(filters[1].Rhs, "5") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestRangeOperatorSquareRoundInteger(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.major<>[1,5)`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[0].Operator, ">=") != 0 ||
		strings.Compare(filters[0].Rhs, "1") != 0 ||
		strings.Compare(filters[1].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[1].Operator, "<") != 0 ||
		strings.Compare(filters[1].Rhs, "5") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestRangeOperatorRoundSquareInteger(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.major<>(1,5]`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[0].Operator, ">") != 0 ||
		strings.Compare(filters[0].Rhs, "1") != 0 ||
		strings.Compare(filters[1].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[1].Operator, "<=") != 0 ||
		strings.Compare(filters[1].Rhs, "5") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestRangeOperatorRoundsInteger(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.major<>(1,5)`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[0].Operator, ">") != 0 ||
		strings.Compare(filters[0].Rhs, "1") != 0 ||
		strings.Compare(filters[1].Lhs, "api.version.major") != 0 ||
		strings.Compare(filters[1].Operator, "<") != 0 ||
		strings.Compare(filters[1].Rhs, "5") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestRangeOperatorSquaresVersion(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.raw<>[1.0.0,3.0.0]`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[0].Operator, ">=") != 0 ||
		strings.Compare(filters[0].Rhs, "1.0.0") != 0 ||
		strings.Compare(filters[1].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[1].Operator, "<=") != 0 ||
		strings.Compare(filters[1].Rhs, "3.0.0") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestRangeOperatorSquareRoundVersion(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.raw<>[1.0.0,3.0.0)`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[0].Operator, ">=") != 0 ||
		strings.Compare(filters[0].Rhs, "1.0.0") != 0 ||
		strings.Compare(filters[1].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[1].Operator, "<") != 0 ||
		strings.Compare(filters[1].Rhs, "3.0.0") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestRangeOperatorRoundSquareVersion(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.raw<>(1.0.0,3.0.0]`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[0].Operator, ">") != 0 ||
		strings.Compare(filters[0].Rhs, "1.0.0") != 0 ||
		strings.Compare(filters[1].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[1].Operator, "<=") != 0 ||
		strings.Compare(filters[1].Rhs, "3.0.0") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}

func TestRangeOperatorRoundsVersion(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{`api.version.raw<>(1.0.0,3.0.0)`})

	if filters == nil ||
		strings.Compare(filters[0].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[0].Operator, ">") != 0 ||
		strings.Compare(filters[0].Rhs, "1.0.0") != 0 ||
		strings.Compare(filters[1].Lhs, "api.version.raw") != 0 ||
		strings.Compare(filters[1].Operator, "<") != 0 ||
		strings.Compare(filters[1].Rhs, "3.0.0") != 0 ||
		err != nil {
		t.Fatal(filters)
	}
}
