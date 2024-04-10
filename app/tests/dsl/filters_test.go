package dsl

import (
	"testing"

	"backend/app/internal/retrieval"
)

func TestNonExistentFilter(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{"api.version.something>3"})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}

func TestNonExistentOperator(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{"api.version.raw!!3"})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}

func TestWrongFormatOnlyFilter(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{"api.version.raw"})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}

func TestWrongFormatOnlyOperator(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{">="})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}

func TestWrongFormatOnlyRightHandSide(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{"5.0.0"})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}

func TestWrongFormatFilterAndOperator(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{"api.version.raw=="})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}

func TestWrongFormatOperatorAndRightHandSide(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{"==3.0.0"})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}

func TestWrongRightHandSideVersion(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{"api.version.raw==3"})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}

func TestWrongRightHandSideString(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{"api.source==20"})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}

func TestWrongRightHandSideBoolean(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{"api.latest==ciao"})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}

func TestWrongRightHandSideInteger(t *testing.T) {
	filters, err := retrieval.CreateFilters([]string{"api.version.minor==true"})

	if filters != nil && err == nil {
		t.Fatal(filters)
	}
}
