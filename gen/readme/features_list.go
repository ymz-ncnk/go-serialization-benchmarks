package main

import (
	"errors"
	"strings"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/benc"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/gob"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/json"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/mus"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/vtprotobuf"
)

func MakeFeaturesList() (featuresList []string, err error) {
	type FeatureMap map[string]string
	m := map[string]FeatureMap{}
	d, err := AllSerializers()
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

func AllSerializers() (serializers []benchser.SerializerDesc, err error) {
	serializers = []benchser.SerializerDesc{}

	// json
	if len(json.Serializers) == 0 {
		err = errors.New("json doesn't have any serializers")
	}
	for i := range json.Serializers {
		serializers = append(serializers, json.Serializers[i])
	}

	// gob
	if len(gob.Serializers) == 0 {
		err = errors.New("gob doesn't have any serializers")
	}
	for i := range gob.Serializers {
		serializers = append(serializers, gob.Serializers[i])
	}

	// mus
	if len(mus.Serializers) == 0 {
		err = errors.New("mus doesn't have any serializers")
	}
	for i := range mus.Serializers {
		serializers = append(serializers, mus.Serializers[i])
	}

	// benc
	if len(benc.Serializers) == 0 {
		err = errors.New("benc doesn't have any serializers")
	}
	for i := range benc.Serializers {
		serializers = append(serializers, benc.Serializers[i])
	}

	// protobuf
	if len(protobuf.SerializersRaw) == 0 {
		err = errors.New("protobuf doesn't have any serializers")
	}
	for i := range protobuf.SerializersRaw {
		serializers = append(serializers, protobuf.SerializersRaw[i])
	}
	if len(protobuf.SerializersVarint) == 0 {
		err = errors.New("protobuf doesn't have any serializers")
	}
	for i := range protobuf.SerializersVarint {
		serializers = append(serializers, protobuf.SerializersVarint[i])
	}

	// vtprotobuf
	if len(vtprotobuf.SerializersRaw) == 0 {
		err = errors.New("vtprotobuf doesn't have any serializers")
	}
	for i := range vtprotobuf.SerializersRaw {
		serializers = append(serializers, vtprotobuf.SerializersRaw[i])
	}
	if len(vtprotobuf.SerializersVarint) == 0 {
		err = errors.New("vtprotobuf doesn't have any serializers")
	}
	for i := range vtprotobuf.SerializersVarint {
		serializers = append(serializers, vtprotobuf.SerializersVarint[i])
	}
	return
}
