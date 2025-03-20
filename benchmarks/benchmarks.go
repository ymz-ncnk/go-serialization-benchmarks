package benchmarks

import (
	"errors"
	"flag"
	"strings"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benc"
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/gob"
	"github.com/ymz-ncnk/go-serialization-benchmarks/json"
	"github.com/ymz-ncnk/go-serialization-benchmarks/mus"
	"github.com/ymz-ncnk/go-serialization-benchmarks/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/vtprotobuf"
)

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

func GeneralDataSerializers() (
	serializers []benchser.Serializer[general.Data]) {
	serializers = []benchser.Serializer[general.Data]{}
	serializers = append(serializers, json.Serializers...)
	serializers = append(serializers, gob.Serializers...)
	serializers = append(serializers, mus.Serializers...)
	serializers = append(serializers, benc.Serializers...)
	return
}

func parseFeatures() (features []benchser.Feature, err error) {
	var (
		fs = flag.NewFlagSet("my", flag.ContinueOnError)
		f  = fs.String("f", "", "a list of features, separeted by ','")
	)
	if err = fs.Parse(flag.Args()); err != nil {
		return
	}
	if len(*f) < 1 {
		return
	}
	var (
		strs = strings.Split(*f, ",")
		l    = len(strs)
	)
	features = make([]benchser.Feature, l)
	for i := range l {
		features[i] = benchser.Feature(strs[i])
	}
	return
}
