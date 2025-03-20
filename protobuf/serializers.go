package protobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
)

const Protobuf = "protobuf"

var GeneralFeatures = []benchser.Feature{
	benchser.Binary,
	benchser.Codegen,
}

var (
	SerializersRaw    = []benchser.Serializer[*data_proto.DataRaw]{SerializerRaw{}}
	SerializersVarint = []benchser.Serializer[*data_proto.DataRawVarint]{SerializerVarint{}}
)
