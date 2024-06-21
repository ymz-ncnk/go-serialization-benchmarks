package serializer

import "sort"

const ResultNameSeparator = "+"

func NewResultName(resultName string, features ...Feature) ResultName {
	sort.Slice(features, func(i, j int) bool {
		return features[i] < features[j]
	})
	for i := 0; i < len(features); i++ {
		resultName += ResultNameSeparator
		resultName += string(features[i])
	}
	return ResultName(resultName)
}

type ResultName string
