package retrieval

import (
	"log"
	"strings"

	"backend/app/internal/models"
)

func ParseDSLRequest(dslString string) {
	filters := createFilters(strings.Split(dslString, " "))

	for _, filter := range filters {
		log.Print(filter)
	}
}

func createFilters(filtersRaw []string) []models.Filter {
	var filters []models.Filter

	for _, filterRaw := range filtersRaw {
		for _, operator := range models.Operators {
			if strings.Contains(filterRaw, operator) {
				sides := strings.Split(filterRaw, operator)

				if _, in := models.TypesMap[sides[0]]; in {
					if strings.Compare(operator, "<>") == 0 {
						limits := strings.Split(strings.Trim(sides[1], "[()]"), ",")
						bracketL := models.BracketsMap[string(sides[1][0])]
						bracketR := models.BracketsMap[string(sides[1][len(sides[1])-1])]

						if len(limits) != 2 {
							filters = append(filters, models.Filter{
								Lhs:      sides[0],
								Operator: operator,
								Rhs:      sides[1],
							})

							break
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

						filterL.Validate()
						filterR.Validate()
						filters = append(filters, filterL)
						filters = append(filters, filterR)
					} else {
						filter := models.Filter{
							Lhs:      sides[0],
							Operator: operator,
							Rhs:      sides[1],
						}

						filter.Validate()
						filters = append(filters, filter)
					}
				} else {
					filters = append(filters, models.Filter{
						Lhs:      sides[0],
						Operator: operator,
						Rhs:      sides[1],
					})
				}

				break
			}
		}
	}

	return filters
}
