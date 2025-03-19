package benchmarks

import (
	"runtime"
	"testing"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
	"github.com/ymz-ncnk/go-serialization-benchmarks/vtprotobuf"

	"github.com/ymz-ncnk/go-serialization-benchmarks/bebop200sc"
	"github.com/ymz-ncnk/go-serialization-benchmarks/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/protobuf_mus"
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
	benchmarkVTProtobuf(wantFeatures, data, b)
	benchmarkBebop200sc(wantFeatures, data, b)
}

func benchmarkGeneralDataSerializers(wantFeatures []serializer.Feature,
	data []general.Data, b *testing.B) {
	generalDataSerializers := GeneralDataSerializers()
	for i := range generalDataSerializers {
		benchser.BenchmarkSerializer(generalDataSerializers[i], wantFeatures, data,
			b)
	}
}

func benchmarkProtobuf(wantFeatures []serializer.Feature,
	data []general.Data, b *testing.B) {
	protobufDataRaw := toCustomData(data, protobuf.ToProtobufDataRaw)
	for i := range protobuf.SerializersRaw {
		benchser.BenchmarkSerializer(protobuf.SerializersRaw[i], wantFeatures,
			protobufDataRaw, b)
	}
	runtime.GC()

	for i := range protobuf_mus.SerializersNative {
		benchser.BenchmarkSerializer(protobuf_mus.SerializersNative[i],
			wantFeatures, protobufDataRaw, b)
	}
	runtime.GC()

	protobufDataRawVarint := toCustomData(data, protobuf.ToProtobufDataRawVarint)
	for i := range protobuf.SerializersVarint {
		benchser.BenchmarkSerializer(protobuf.SerializersVarint[i], wantFeatures,
			protobufDataRawVarint, b)
	}
	runtime.GC()
}

func benchmarkVTProtobuf(wantFeatures []serializer.Feature,
	data []general.Data, b *testing.B) {
	protobufDataRaw := toCustomData(data, vtprotobuf.ToProtobufDataRaw)
	for i := range vtprotobuf.SerializersRaw {
		benchser.BenchmarkSerializer(vtprotobuf.SerializersRaw[i], wantFeatures,
			protobufDataRaw, b)
	}
	runtime.GC()

	for i := range protobuf_mus.SerializersNative {
		benchser.BenchmarkSerializer(protobuf_mus.SerializersNative[i],
			wantFeatures, protobufDataRaw, b)
	}
	runtime.GC()

	protobufDataRawVarint := toCustomData(data, vtprotobuf.ToProtobufDataRawVarint)
	for i := range vtprotobuf.SerializersVarint {
		benchser.BenchmarkSerializer(vtprotobuf.SerializersVarint[i], wantFeatures,
			protobufDataRawVarint, b)
	}
	runtime.GC()
}

func benchmarkBebop200sc(wantFeatures []serializer.Feature,
	data []general.Data, b *testing.B) {
	bebop200scData := toCustomData(data, bebop200sc.ToBebop200scData)
	for i := range bebop200sc.Serializers {
		benchser.BenchmarkSerializer(bebop200sc.Serializers[i], wantFeatures,
			bebop200scData, b)
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
