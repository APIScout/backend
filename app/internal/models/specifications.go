package models

// SpecificationsRequest - structure of the request to be sent to the backend whenever new specifications are added
type SpecificationsRequest struct {
	Specifications []SpecificationBackend `json:"specifications"`
}

// SpecificationBackend - type of the single specification
type SpecificationBackend map[string]interface{}
