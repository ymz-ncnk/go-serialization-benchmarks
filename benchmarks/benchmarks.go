package benchmarks

import (
	"errors"
	"flag"
	"strings"

	"github.com/ymz-ncnk/go-serialization-benchmarks/bebop200sc"
	"github.com/ymz-ncnk/go-serialization-benchmarks/benc"
	"github.com/ymz-ncnk/go-serialization-benchmarks/gob"
	"github.com/ymz-ncnk/go-serialization-benchmarks/json"
	"github.com/ymz-ncnk/go-serialization-benchmarks/mus"
	"github.com/ymz-ncnk/go-serialization-benchmarks/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

func FirstOneSerializerDescs() (serializers []serializer.SerializerDesc,
	err error) {
	if len(json.Serializers) == 0 {
		err = errors.New("json doesn't have any serializers")
	}
	if len(gob.Serializers) == 0 {
		err = errors.New("gob doesn't have any serializers")
	}
	if len(mus.Serializers) == 0 {
		err = errors.New("mus doesn't have any serializers")
	}
	if len(benc.Serializers) == 0 {
		err = errors.New("benc doesn't have any serializers")
	}
	serializers = []serializer.SerializerDesc{}
	serializers = append(serializers, json.Serializers[0])
	serializers = append(serializers, gob.Serializers[0])
	serializers = append(serializers, mus.Serializers[0])
	serializers = append(serializers, benc.Serializers[0])
	serializers = append(serializers, protobuf.Serializers[0])
	serializers = append(serializers, bebop200sc.Serializers[0])
	return
}

func GeneralDataSerializers() (
	serializers []serializer.Serializer[serializer.Data]) {
	serializers = []serializer.Serializer[serializer.Data]{}
	serializers = append(serializers, json.Serializers...)
	serializers = append(serializers, gob.Serializers...)
	serializers = append(serializers, mus.Serializers...)
	serializers = append(serializers, benc.Serializers...)
	return
}

func parseFeatures() (features []serializer.Feature, err error) {
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
	features = make([]serializer.Feature, l)
	for i := 0; i < l; i++ {
		features[i] = serializer.Feature(strs[i])
	}
	return
}
