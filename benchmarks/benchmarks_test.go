package benchmarks

import (
	"testing"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"

	"github.com/ymz-ncnk/go-serialization-benchmarks/bebop200sc"
	"github.com/ymz-ncnk/go-serialization-benchmarks/protobuf"
)

const DataCount = 23000000

// go test -bench=. -- -f textser
func BenchmarkSerializers(b *testing.B) {
	wantFeatures, err := parseFeatures()
	if err != nil {
		b.Fatal(err)
	}
	data, err := benchser.GenerateData(DataCount)
	if err != nil {
		b.Fatal(err)
	}

	generalDataSerializers := GeneralDataSerializers()
	for i := 0; i < len(generalDataSerializers); i++ {
		benchser.BenchmarkSerializer(generalDataSerializers[i], wantFeatures, data,
			b)
	}

	protobufData := protobuf.ToProtobufData(data)
	for i := 0; i < len(protobuf.Serializers); i++ {
		benchser.BenchmarkSerializer(protobuf.Serializers[i], wantFeatures,
			protobufData, b)
	}

	bebop200scData := bebop200sc.ToBebop200scData(data)
	for i := 0; i < len(bebop200sc.Serializers); i++ {
		benchser.BenchmarkSerializer(bebop200sc.Serializers[i], wantFeatures,
			bebop200scData, b)
	}
}
