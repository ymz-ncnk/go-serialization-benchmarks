package protobuf_mus

import (
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/protobuf"
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

var (
	lenM mus.MarshallerFn[int] = func(t int, bs []byte) (n int) {
		return varint.MarshalPositiveInt32(int32(t), bs)
	}
	lenU mus.UnmarshallerFn[int] = func(bs []byte) (t int, n int, err error) {
		t32, n, err := varint.UnmarshalPositiveInt32(bs)
		t = int(t32)
		return
	}
	lenS mus.SizerFn[int] = func(t int) (size int) {
		return varint.SizePositiveInt32(int32(t))
	}
)

var Serializers = []serializer.Serializer[serializer.Data]{
	SerializerMUSRawVarint{},
	SerializerMUSRawVarintReuse{bs: make([]byte, serializer.BufSize)},
	SerializerMUSRaw{},
	SerializerMUSRawReuse{bs: make([]byte, serializer.BufSize)},
	SerializerMUSUnsafe{},
	SerializerMUSUnsafeReuse{bs: make([]byte, serializer.BufSize)},
}

var SerializersNative = []serializer.Serializer[*protobuf.DataRaw]{
	SerializerNativeUnsafeReuse{bs: make([]byte, serializer.BufSize)},
}
