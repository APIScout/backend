package dsl

import (
	"backend/app/internal/retrieval"
	"testing"
)

func TestRetrieval(t *testing.T) {
	retrieval.ParseDSLRequest("api.latest==true")
	retrieval.ParseDSLRequest("api.version.raw>=2.0.0")
	retrieval.ParseDSLRequest(`api.source=="github"`)
	retrieval.ParseDSLRequest("specification.version.raw<>[2.0.0,3.0.0)")
	retrieval.ParseDSLRequest("api.version.minor>3")
	retrieval.ParseDSLRequest("api.version.something>3")
	retrieval.ParseDSLRequest("api.version.raw<>[3.0.0")
}
