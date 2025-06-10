package protobuf_mus

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	data_protobuf_mus "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf_mus"
	"google.golang.org/protobuf/encoding/protowire"
)

const Protobuf = "protobuf_mus"

var GeneralFeatures = []benchser.Feature{
	benchser.Binary,
	benchser.Manual,
}

var (
	strFieldTag     = protowire.EncodeTag(1, protowire.BytesType)
	boolFieldTag    = protowire.EncodeTag(2, protowire.VarintType)
	int32FieldTag   = protowire.EncodeTag(3, protowire.Fixed32Type)
	float64FieldTag = protowire.EncodeTag(4, protowire.Fixed64Type)
	timeFieldTag    = protowire.EncodeTag(5, protowire.BytesType)
)

var Serializers = []benchser.Serializer[data_protobuf_mus.Data]{
	SerializerVarint{},
	SerializerVarintReuse{bs: make([]byte, benchser.BufSize)},
	SerializerRaw{},
	SerializerRawReuse{bs: make([]byte, benchser.BufSize)},
	SerializerUnsafe{},
	SerializerUnsafeReuse{bs: make([]byte, benchser.BufSize)},
}

var SerializersNative = []benchser.Serializer[*data_proto.DataRaw]{
	SerializerNativeUnsafeReuse{bs: make([]byte, benchser.BufSize)},
}
