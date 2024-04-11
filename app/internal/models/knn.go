package models

// ParametersMap - The elements of the map are structured as follows: "queryKey": {default, min, max}
var ParametersMap = map[string][]int{
	"pageSize": {10, 1, 100},
	"page":     {1, 1, -1},
	"k":        {100, 1, 100},
}
