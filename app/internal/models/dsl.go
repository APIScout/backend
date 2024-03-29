package models

import (
	"slices"
	"strconv"
	"strings"
)

var Operators = []string{"==", "!=", "~=", "<>", ">=", "<=", ">", "<"}

var OperatorsMap = map[string][]string{
	"==": {"bool", "str", "int", "version"},
	"!=": {"bool", "str", "int", "version"},
	"~=": {"str", "version"},
	">=": {"int", "version"},
	"<=": {"int", "version"},
	">":  {"int", "version"},
	"<":  {"int", "version"},
}

var TypesMap = map[string]string{
	"api.version.raw":        "version",
	"api.version.valid":      "bool",
	"api.version.major":      "int",
	"api.version.minor":      "int",
	"api.version.patch":      "int",
	"api.version.prerelease": "str",
	"api.version.build":      "str",
	"api.commits":            "int",
	"api.latest":             "bool",
	"api.source":             "str",

	"specification.version.raw":        "version",
	"specification.version.valid":      "bool",
	"specification.version.major":      "int",
	"specification.version.minor":      "int",
	"specification.version.patch":      "int",
	"specification.version.prerelease": "str",
	"specification.version.build":      "str",
	"specification.type":               "str",
}

var BracketsMap = map[string]string{
	"[": ">=",
	"(": ">",
	"]": "<=",
	")": "<",
}

type Filter struct {
	Valid    bool
	Lhs      string
	Operator string
	Rhs      string
}

func (filter *Filter) Validate() {
	acceptedType := TypesMap[filter.Lhs]
	operatorTypes := OperatorsMap[filter.Operator]

	if slices.Contains(operatorTypes, acceptedType) {
		switch acceptedType {
		case "str":
			if strings.Compare(filter.Rhs, "") != 0 {
				filter.Rhs = strings.Trim(filter.Rhs, `"`)
				filter.Valid = true
			}
		case "version":
			if GetSemanticVersion(filter.Rhs).Valid {
				filter.Valid = true
			}
		case "bool":
			if strings.Compare(filter.Rhs, "true") == 0 || strings.Compare(filter.Rhs, "false") == 0 {
				filter.Valid = true
			}
		case "int":
			if _, err := strconv.ParseInt(filter.Rhs, 10, 64); err == nil {
				filter.Valid = true
			}
		}
	}
}
