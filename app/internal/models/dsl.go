package models

type DSL struct {
	API struct {
		Version string `json:"version"`
		Commits int    `json:"commits"`
		Latest  bool   `json:"latest"`
	} `json:"api"`
	Specification struct {
		Version string `json:"version"`
	} `json:"specification"`
	Metrics struct {
		Paths int `json:"metrics"`
	} `json:"metrics"`
}
