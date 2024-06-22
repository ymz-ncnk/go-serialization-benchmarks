package main

import (
	"sort"
	"strings"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchmarks"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

func MakeFeaturesList() (featuresList []string, err error) {
	descs, err := benchmarks.FirstOneSerializerDescs()
	if err != nil {
		return
	}
	featuresList = []string{}
	var (
		desc           serializer.SerializerDesc
		serializerName string
	)
	for i := 0; i < len(descs); i++ {
		desc = descs[i]
		serializerName, err = desc.Name().SerializerName()
		if err != nil {
			return
		}
		featuresList = append(featuresList,
			serializerName+": "+strings.Join(featuresToStrs(desc.Features()), ", "))
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
