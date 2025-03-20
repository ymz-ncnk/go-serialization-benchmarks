package benchser

import (
	"errors"
	"regexp"
	"sort"
)

const ResultNameSeparator = "+"

var SerializerNameRegexp = regexp.MustCompile(`(.+?)(?:\+|$)`)

func NewResultName(serializerName string, features ...Feature) ResultName {
	sort.Slice(features, func(i, j int) bool {
		return features[i] < features[j]
	})
	resultName := serializerName
	for i := range features {
		resultName += ResultNameSeparator
		resultName += string(features[i])
	}
	return ResultName(resultName)
}

type ResultName string

func (r ResultName) SerializerName() (name string, err error) {
	strs := SerializerNameRegexp.FindStringSubmatch(string(r))
	if len(strs) != 2 {
		err = errors.New("SerializerNameRegexp failed to find a value")
		return
	}
	name = strs[1]
	return
}
