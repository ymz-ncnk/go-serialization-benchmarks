package protobuf_mus

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
	"google.golang.org/protobuf/encoding/protowire"
)

const Protobuf = "protobuf_mus"

var (
	strFieldTag     = protowire.EncodeTag(1, protowire.BytesType)
	boolFieldTag    = protowire.EncodeTag(2, protowire.VarintType)
	int32FieldTag   = protowire.EncodeTag(3, protowire.Fixed32Type)
	float64FieldTag = protowire.EncodeTag(4, protowire.Fixed64Type)
	timeFieldTag    = protowire.EncodeTag(5, protowire.BytesType)
)

var Serializers = []serializer.Serializer[general.Data]{
	SerializerMUSVarint{},
	SerializerMUSVarintReuse{bs: make([]byte, serializer.BufSize)},
	SerializerMUSRaw{},
	SerializerMUSRawReuse{bs: make([]byte, serializer.BufSize)},
	SerializerMUSUnsafe{},
	SerializerMUSUnsafeReuse{bs: make([]byte, serializer.BufSize)},
}

var SerializersNative = []serializer.Serializer[*data_proto.DataRaw]{
	SerializerNativeUnsafeReuse{bs: make([]byte, serializer.BufSize)},
}
