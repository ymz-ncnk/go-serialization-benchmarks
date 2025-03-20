package benchmarks

import (
	"runtime"
	"testing"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_bebop "github.com/ymz-ncnk/go-serialization-benchmarks/data/bebop"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	data_protobuf "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	data_protobuf_mus "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf_mus"
	"github.com/ymz-ncnk/go-serialization-benchmarks/protobuf_mus"
	"github.com/ymz-ncnk/go-serialization-benchmarks/vtprotobuf"

	"github.com/ymz-ncnk/go-serialization-benchmarks/bebop200sc"
	"github.com/ymz-ncnk/go-serialization-benchmarks/protobuf"
)

const DataCount = 20000000

func BenchmarkSerializers(b *testing.B) {
	wantFeatures, err := parseFeatures()
	if err != nil {
		b.Fatal(err)
	}
	data, err := benchser.GenerateData(DataCount)
	if err != nil {
		b.Fatal(err)
	}
	benchmarkGeneralDataSerializers(wantFeatures, data, b)
	benchmarkProtobuf(wantFeatures, data, b)
	benchmarkProtobufMUS(wantFeatures, data, b)
	benchmarkVTProtobuf(wantFeatures, data, b)
	benchmarkBebop200sc(wantFeatures, data, b)
}

func benchmarkGeneralDataSerializers(wantFeatures []benchser.Feature,
	data []general.Data, b *testing.B) {
	s := GeneralDataSerializers()
	for i := range s {
		benchser.BenchmarkSerializer(s[i], wantFeatures, data, b)
	}
}

func benchmarkProtobuf(wantFeatures []benchser.Feature,
	data []general.Data, b *testing.B) {
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

func benchmarkVTProtobuf(wantFeatures []benchser.Feature, data []general.Data,
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
	data []general.Data, b *testing.B) {
	var (
		d = toCustomData(data, data_protobuf_mus.ToProtobufMUSData)
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

func benchmarkBebop200sc(wantFeatures []benchser.Feature, data []general.Data,
	b *testing.B) {
	var (
		d = toCustomData(data, data_bebop.ToBebop200scData)
		s = bebop200sc.Serializers
	)
	for i := range s {
		benchser.BenchmarkSerializer(s[i], wantFeatures, d, b)
	}
	runtime.GC()
}

func toCustomData[T any](data []general.Data,
	fn func(data general.Data) T) (d []T) {
	l := len(data)
	d = make([]T, l)
	for i := range l {
		d[i] = fn(data[i])
	}
	return
}
