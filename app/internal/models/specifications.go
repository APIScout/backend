package models

type SpecificationsRequest struct {
	Specifications []Specification
}

type Specification map[string]interface{}
