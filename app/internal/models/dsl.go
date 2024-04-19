package models

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"time"
)

var Operators = []string{"==", "!=", "~=", "<>", ">=", "<=", ">", "<"}
var NegOperators = []string{"!="}

var OperatorsMap = map[string][]string{
	"==": {"bool", "str", "int", "version", "date"},
	"!=": {"bool", "str", "int", "version", "date"},
	"~=": {"str", "version"},
	">=": {"int", "version", "date"},
	"<=": {"int", "version", "date"},
	">":  {"int", "version", "date"},
	"<":  {"int", "version", "date"},
}

var TypesMap = map[string]string{
	"date": "date",

	"api.version.raw":        "version",
	"api.version.valid":      "bool",
	"api.version.major":      "int",
	"api.version.minor":      "int",
	"api.version.patch":      "int",
	"api.version.prerelease": "str",
	"api.version.build":      "str",
	"api.name":               "str",
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

var OperatorToEsMap = map[string]string{
	"==": "term",
	"!=": "term",
	"~=": "regexp",
	">=": "gte",
	"<=": "lte",
	">":  "gt",
	"<":  "lt",
}

var BracketsMap = map[string]string{
	"[": ">=",
	"(": ">",
	"]": "<=",
	")": "<",
}

type Filter struct {
	Lhs      string
	Operator string
	Rhs      string
}

func (filter *Filter) Validate() error {
	acceptedType := TypesMap[filter.Lhs]
	operatorTypes := OperatorsMap[filter.Operator]

	if slices.Contains(operatorTypes, acceptedType) {
		switch acceptedType {
		case "str":
			if strings.Compare(string(filter.Rhs[0]), `"`) == 0 &&
				strings.Compare(string(filter.Rhs[len(filter.Rhs)-1]), `"`) == 0 &&
				strings.Compare(filter.Rhs, "") != 0 {

				filter.Rhs = strings.Trim(filter.Rhs, `"`)
				return nil
			}
		case "version":
			if filter.Operator == "~=" || GetSemanticVersion(filter.Rhs).Valid {
				return nil
			}
		case "bool":
			if strings.Compare(filter.Rhs, "true") == 0 || strings.Compare(filter.Rhs, "false") == 0 {
				return nil
			}
		case "int":
			if _, err := strconv.ParseInt(filter.Rhs, 10, 64); err == nil {
				return nil
			}
		case "date":
			if res, err := time.Parse("02/01/2006", filter.Rhs); err == nil {
				filter.Rhs = res.Format(time.RFC3339)
				return nil
			}
		}
	}

	return errors.New("the type of the right hand side does not match the ones of the operator")
}

func (filter *Filter) ToEsFilter() (string, bool) {
	var esFilter strings.Builder
	esOperator := OperatorToEsMap[filter.Operator]
	positive := !slices.Contains(NegOperators, filter.Operator)

	switch esOperator {
	case "term":
		esFilter.WriteString(`"term": {"metadata.` + filter.Lhs + `": "` + filter.Rhs + `"}`)
	case "regexp":
		esFilter.WriteString(`"regexp": {"metadata.` + filter.Lhs + `": ".*` + filter.Rhs + `.*"}`)
	case "gte", "lte", "gt", "lt":
		rhs := filter.Rhs

		if TypesMap[filter.Lhs] == "version" || TypesMap[filter.Lhs] == "date" {
			rhs = `"` + rhs + `"`
		}

		esFilter.WriteString(`"range": {"metadata.` + filter.Lhs + `": {"` + esOperator + `": ` + rhs + `}}`)
	}

	return esFilter.String(), positive
}
