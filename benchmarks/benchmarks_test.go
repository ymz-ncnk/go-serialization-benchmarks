package benchmarks

import (
	"runtime"
	"testing"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

	"github.com/ymz-ncnk/go-serialization-benchmarks/bebop200sc"
	"github.com/ymz-ncnk/go-serialization-benchmarks/protobuf"
)

const DataCount = 23000000

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
	benchmarkBebop200sc(wantFeatures, data, b)
}

func benchmarkGeneralDataSerializers(wantFeatures []serializer.Feature,
	data []serializer.Data, b *testing.B) {
	generalDataSerializers := GeneralDataSerializers()
	for i := 0; i < len(generalDataSerializers); i++ {
		benchser.BenchmarkSerializer(generalDataSerializers[i], wantFeatures, data,
			b)
	}
}

func benchmarkProtobuf(wantFeatures []serializer.Feature,
	data []serializer.Data, b *testing.B) {
	protobufDataRaw := toCustomData(data, protobuf.ToProtobufDataRaw)
	for i := 0; i < len(protobuf.SerializersRaw); i++ {
		benchser.BenchmarkSerializer(protobuf.SerializersRaw[i], wantFeatures,
			protobufDataRaw, b)
	}
	runtime.GC()

	protobufDataRawVarint := toCustomData(data, protobuf.ToProtobufDataRawVarint)
	for i := 0; i < len(protobuf.SerializersRaw); i++ {
		benchser.BenchmarkSerializer(protobuf.SerializersRawVarint[i], wantFeatures,
			protobufDataRawVarint, b)
	}
	runtime.GC()
}

func benchmarkBebop200sc(wantFeatures []serializer.Feature,
	data []serializer.Data, b *testing.B) {
	bebop200scData := toCustomData(data, bebop200sc.ToBebop200scData)
	for i := 0; i < len(bebop200sc.Serializers); i++ {
		benchser.BenchmarkSerializer(bebop200sc.Serializers[i], wantFeatures,
			bebop200scData, b)
	}
	runtime.GC()
}

func toCustomData[T any](data []serializer.Data,
	fn func(data serializer.Data) T) (d []T) {
	l := len(data)
	d = make([]T, l)
	for i := 0; i < l; i++ {
		d[i] = fn(data[i])
	}
	return
}
