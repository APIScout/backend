package retrieval

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func CreateKnnQuery(embedding []float32, filters string, pageSize int, page int, k int) string {
	var query strings.Builder

	query.WriteString(`{"from": ` + strconv.Itoa((page - 1) * pageSize) + `, "size": ` + strconv.Itoa(pageSize))

	if embedding != nil {
		query.WriteString(`, "knn": {"field": "embedding", "query_vector": [`)

		for ind, el := range embedding {
			value := reflect.ValueOf(el).Float()
			query.WriteString(fmt.Sprintf(`%f`, value))

			if ind != len(embedding)-1 {
				query.WriteString(", ")
			}
		}

		query.WriteString(`], "k": ` + strconv.Itoa(k) + `, "num_candidates": 10000, "filter": {`)
		query.WriteString(filters)
		query.WriteString(`}}}`)
	} else {
		query.WriteString(`, "query": {`)
		query.WriteString(filters)
		query.WriteString("}}")
	}

	return query.String()
}
