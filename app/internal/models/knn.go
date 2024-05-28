package models

// ParametersMap - The elements of the map are structured as follows: "queryKey": {default, min, max}
var ParametersMap = map[string][]int{
	"size": {10, 1, 120},
	"page": {1, 1, -1},
	"k":    {100, 1, 120},
}
