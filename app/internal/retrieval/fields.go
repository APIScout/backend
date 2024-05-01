package retrieval

import (
	"errors"
	"slices"

	"backend/app/internal/models"
)

func FilterFields(response models.SpecificationWithApi, fields []string) (models.SpecificationWithApi, error) {
	var filteredResponse models.SpecificationWithApi

	for _, field := range fields {
		if slices.Contains(models.PossibleFilters, field) {
			switch field {
			case "metadata":
				filteredResponse.MongoDocument = response.MongoDocument
			case "specification":
				filteredResponse.Specification = response.Specification
			}
		} else {
			return models.SpecificationWithApi{}, errors.New("the given field does not exist")
		}
	}

	return filteredResponse, nil
}
