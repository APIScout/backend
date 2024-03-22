package models

// SpecificationsRequest - structure of the request to be sent to the backend whenever new specifications are added
type SpecificationsRequest struct {
	Specifications []Specification `json:"specifications"`
}

// Specification - type of the single specification
type Specification map[string]interface{}
