package models

type SpecificationsRequest struct {
	Specifications []Specification `json:"specifications"`
}

type Specification map[string]interface{}
