//go:generate go run ./gen/readme/
package benchmarks

import (
	"flag"
	"runtime"
	"strings"
	"testing"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/common"
	data_protobuf "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/benc"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/gob"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/json"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/mus"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/protobuf_mus"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/vtprotobuf"

	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/bebop200sc"
	"github.com/ymz-ncnk/go-serialization-benchmarks/projects/protobuf"
)

const DataCount = 20000000

func BenchmarkSerializers(b *testing.B) {
	wantFeatures, err := parseFeatures()
	if err != nil {
		b.Fatal(err)
	}
	data, err := benchser.CommonData(DataCount)
	if err != nil {
		b.Fatal(err)
	}
	benchmarkCommonDataSerializers(wantFeatures, data, b)
	benchmarkProtobuf(wantFeatures, data, b)
	benchmarkProtobufMUS(wantFeatures, data, b)
	benchmarkVTProtobuf(wantFeatures, data, b)
	benchmarkBebop200sc(wantFeatures, data, b)
}

func benchmarkCommonDataSerializers(wantFeatures []benchser.Feature,
	data []common.Data, b *testing.B) {
	s := commonDataSerializers()
	for i := range s {
		benchser.BenchmarkSerializer(s[i], wantFeatures, data, b)
	}
}

func benchmarkProtobuf(wantFeatures []benchser.Feature,
	data []common.Data, b *testing.B) {
	var (
		dr = toCustomData(data, data_protobuf.ToProtobufDataRaw)
		sr = protobuf.SerializersRaw
	)
	for i := range sr {
		benchser.BenchmarkSerializer(sr[i], wantFeatures, dr, b)
	}
	runtime.GC()
	var (
		dv = toCustomData(data, data_protobuf.ToProtobufDataVarint)
		sv = protobuf.SerializersVarint
	)
	for i := range sv {
		benchser.BenchmarkSerializer(sv[i], wantFeatures, dv, b)
	}
	runtime.GC()
}

func benchmarkVTProtobuf(wantFeatures []benchser.Feature, data []common.Data,
	b *testing.B) {
	var (
		dr = toCustomData(data, data_protobuf.ToProtobufDataRaw)
		sr = vtprotobuf.SerializersRaw
	)
	for i := range sr {
		benchser.BenchmarkSerializer(sr[i], wantFeatures, dr, b)
	}
	runtime.GC()
	var (
		dv = toCustomData(data, data_protobuf.ToProtobufDataVarint)
		sv = vtprotobuf.SerializersVarint
	)
	for i := range sv {
		benchser.BenchmarkSerializer(sv[i], wantFeatures, dv, b)
	}
	runtime.GC()
}

func benchmarkProtobufMUS(wantFeatures []benchser.Feature,
	data []common.Data, b *testing.B) {
	var (
		d = toCustomData(data, protobuf_mus.ToProtobufMUSData)
		s = protobuf_mus.Serializers
	)
	for i := range s {
		benchser.BenchmarkSerializer(s[i], wantFeatures, d, b)
	}
	runtime.GC()
	var (
		pd = toCustomData(data, data_protobuf.ToProtobufDataRaw)
		ns = protobuf_mus.SerializersNative
	)
	for i := range ns {
		benchser.BenchmarkSerializer(ns[i], wantFeatures, pd, b)
	}
	runtime.GC()
}

func benchmarkBebop200sc(wantFeatures []benchser.Feature, data []common.Data,
	b *testing.B) {
	var (
		d = toCustomData(data, bebop200sc.ToBebop200scData)
		s = bebop200sc.Serializers
	)
	for i := range s {
		benchser.BenchmarkSerializer(s[i], wantFeatures, d, b)
	}
	runtime.GC()
}

func commonDataSerializers() (
	serializers []benchser.Serializer[common.Data]) {
	serializers = []benchser.Serializer[common.Data]{}
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

func toCustomData[T any](data []common.Data,
	fn func(data common.Data) T) (d []T) {
	l := len(data)
	d = make([]T, l)
	for i := range l {
		d[i] = fn(data[i])
	}
	return
}
