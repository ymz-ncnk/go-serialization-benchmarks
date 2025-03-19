package main

import (
	"strings"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchmarks"
)

func MakeFeaturesList() (featuresList []string, err error) {
	type FeatureMap map[string]string
	m := map[string]FeatureMap{}
	d, err := benchmarks.AllSerializers()
	if err != nil {
		return
	}
	for i := range d {
		var name string
		name, err = d[i].Name().SerializerName()
		if err != nil {
			return
		}
		if _, pst := m[name]; !pst {
			m[name] = make(FeatureMap)
		}
		for j := range d[i].Features() {
			f := d[i].Features()
			m[name][string(f[j])] = ""
		}
	}

	for serName, features := range m {
		var (
			slist = serName + ": "
			arr   = []string{}
		)
		for f := range features {
			arr = append(arr, f)
		}
		slist += strings.Join(arr, ",")
		featuresList = append(featuresList, slist)
	}
	return
}

// func featuresToStrs(features []serializer.Feature) (strs []string) {
// 	l := len(features)
// 	strs = make([]string, l)
// 	for i := 0; i < l; i++ {
// 		strs[i] = "`" + string(features[i]) + "`"
// 	}
// 	sort.Slice(strs, func(i, j int) bool {
// 		return strs[i] < strs[j]
// 	})
// 	return
// }
