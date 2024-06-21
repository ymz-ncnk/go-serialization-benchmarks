package main

import (
	"sort"
	"strings"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchmarks"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

func MakeFeaturesList() (featuresList []string) {
	serializers, err := benchmarks.FirstOneSerializers()
	if err != nil {
		return
	}
	featuresList = []string{}
	for i := 0; i < len(serializers); i++ {
		s := serializers[i]
		featuresList = append(featuresList,
			string(s.Name())+": "+strings.Join(featuresToStrs(s.Features()), ", "))
	}
	sort.Slice(featuresList, func(i, j int) bool {
		return featuresList[i][0] < featuresList[j][0]
	})
	return
}

func featuresToStrs(features []serializer.Feature) (strs []string) {
	l := len(features)
	strs = make([]string, l)
	for i := 0; i < l; i++ {
		strs[i] = "`" + string(features[i]) + "`"
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	return
}
