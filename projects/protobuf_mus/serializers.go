package protobuf_mus

import (
	"fmt"

	"github.com/mus-format/ext-protobuf-go"
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/common"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
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

var Serializers = []benchser.Serializer[Data]{
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

func (d Data) EqualTo(ad Data) (err error) {
	if d.Str != ad.Str {
		return fmt.Errorf("Data.Str value %v != %v", ad.Str, d.Str)
	}
	if d.Bool != ad.Bool {
		return fmt.Errorf("Data.Bool value %v != %v", ad.Bool, d.Bool)
	}
	if d.Int32 != ad.Int32 {
		return fmt.Errorf("Data.Int32 value %v != %v", ad.Int32, d.Int32)
	}
	if d.Float64 != ad.Float64 {
		return fmt.Errorf("Data.Float64 value %v != %v", ad.Float64, d.Float64)
	}
	if d.Time.Seconds != ad.Time.Seconds && d.Time.Nanos != ad.Time.Nanos {
		return fmt.Errorf("Data.Time value %v != %v", ad.Time, d.Time)
	}
	return nil
}

func ToProtobufMUSData(data common.Data) (d Data) {
	return Data{
		Str:     data.Str,
		Bool:    data.Bool,
		Int32:   data.Int32,
		Float64: data.Float64,
		Time: ext.Timestamp{Seconds: data.Time.Unix(),
			Nanos: int32(data.Time.Nanosecond())},
	}
}
