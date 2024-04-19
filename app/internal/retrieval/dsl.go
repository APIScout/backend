package retrieval

import (
	"errors"
	"slices"
	"strings"

	"backend/app/internal/models"
)

func ParseDSLRequest(dslString string) (*string, error) {
	filters, err := CreateFilters(strings.Split(dslString, " "))

	if err != nil {
		return nil, err
	}

	return CreateEsFilter(filters), nil
}

func CreateEsFilter(filters []models.Filter) *string {
	var must strings.Builder
	var mustNot strings.Builder
	var esFilter strings.Builder
	esFilter.WriteString(`"bool": {`)

	filterCommits := false

	for index, filter := range filters {
		pathArray := strings.Split(filter.Lhs, ".")
		path := strings.Join(slices.Delete(pathArray, len(pathArray)-1, len(pathArray)), ".")

		if len(path) != 0 {
			path = "." + path
		}

			parsed, positive := filter.ToEsFilter()
		query := `{"nested": {"path": "metadata` + path + `", "query": {` + parsed + "}}}"

		if index != len(filters)-1 {
			query += ","
		}

		if positive {
			must.WriteString(query)
		} else {
			mustNot.WriteString(query)
		}

		if strings.Compare(filter.Lhs, "api.commits") == 0 {
			filterCommits = true
		}
	}

	esFilter.WriteString(`"must": [` + must.String())

	if len(must.String()) != 0 {
		esFilter.WriteString(", ")
	}

	if !filterCommits {
		esFilter.WriteString(`{"nested": {"path": "metadata.api", "query": {"term": {"metadata.api.latest": true}}}}, `)
	}

	esFilter.WriteString(`{"nested": {"path": "metadata", "query": {"range": {"metadata.length": {"gte": 200}}}}}, `)
	esFilter.WriteString(`{"nested": {"path": "metadata.api", "query": {"regexp": {"metadata.api.name": ".+"}}}}], `)
	esFilter.WriteString(`"must_not": [` + mustNot.String() + `]`)
	esFilter.WriteString("}")
	res := esFilter.String()

	return &res
}

func CreateFilters(filtersRaw []string) ([]models.Filter, error) {
	var filters []models.Filter

	for _, filterRaw := range filtersRaw {
		for _, operator := range models.Operators {
			if strings.Contains(filterRaw, operator) {
				sides := strings.Split(filterRaw, operator)

				if _, in := models.TypesMap[sides[0]]; in {
					// Range operation is split into two operations.
					// e.g. api.commits<>[1,5] => api.commits>=1 api.commits<=5
					if strings.Compare(operator, "<>") == 0 {
						if _, in := models.BracketsMap[string(sides[1][0])]; !in {
							return nil, errors.New("in the range, you can only use [, ], ), or )")
						}

						if _, in := models.BracketsMap[string(sides[1][len(sides[1])-1])]; !in {
							return nil, errors.New("in the range, you can only use [, ], ), or )")
						}

						limits := strings.Split(strings.Trim(sides[1], "[()]"), ",")
						bracketL := models.BracketsMap[string(sides[1][0])]
						bracketR := models.BracketsMap[string(sides[1][len(sides[1])-1])]

						if len(limits) != 2 {
							return nil, errors.New("there are less than two elements in the range")
						}

						limitL := limits[0]
						limitR := limits[1]

						filterL := models.Filter{
							Lhs:      sides[0],
							Operator: bracketL,
							Rhs:      limitL,
						}

						filterR := models.Filter{
							Lhs:      sides[0],
							Operator: bracketR,
							Rhs:      limitR,
						}

						err := filterL.Validate()
						err = filterR.Validate()

						if err != nil {
							return nil, err
						}

						filters = append(filters, filterL)
						filters = append(filters, filterR)
					} else {
						if strings.Compare(sides[1], "") != 0 {
							filter := models.Filter{
								Lhs:      sides[0],
								Operator: operator,
								Rhs:      sides[1],
							}

							err := filter.Validate()

							if err != nil {
								return nil, err
							}

							filters = append(filters, filter)
						}
					}
				} else {
					return nil, errors.New("the given left hand side filter name does not exist")
				}

				break
			}
		}
	}

	return filters, nil
}
