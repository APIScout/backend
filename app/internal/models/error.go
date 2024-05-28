package models

// HTTPError - structure of the error response sent by the backend
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}
