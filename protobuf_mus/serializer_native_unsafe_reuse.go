package protobuf_mus

import (
	"fmt"

	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerNativeUnsafeReuse struct {
	bs []byte
}

func (s SerializerNativeUnsafeReuse) Name() serializer.ResultName {
	return serializer.NewResultName(Protobuf, serializer.Native,
		serializer.Unsafe, serializer.Reuse)
}

func (s SerializerNativeUnsafeReuse) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Native, serializer.Unsafe,
		serializer.Reuse)
}

func (s SerializerNativeUnsafeReuse) Marshal(data *data_proto.DataRaw) (bs []byte,
	err error) {
	var n int
	if data.Str != "" {
		n += varint.Uint64.Marshal(strFieldTag, s.bs[n:])
		n += unsafe.String.Marshal(data.Str, s.bs[n:])
	}
	if data.Bool {
		n += varint.Uint64.Marshal(boolFieldTag, s.bs[n:])
		n += unsafe.Bool.Marshal(data.Bool, s.bs[n:])
	}
	if data.Int32 != 0 {
		n += varint.Uint64.Marshal(int32FieldTag, s.bs[n:])
		n += unsafe.Int32.Marshal(data.Int32, s.bs[n:])
	}
	if data.Float64 != 0 {
		n += varint.Uint64.Marshal(float64FieldTag, s.bs[n:])
		n += unsafe.Float64.Marshal(data.Float64, s.bs[n:])
	}
	if data.Time != nil {
		n += varint.Uint64.Marshal(timeFieldTag, s.bs[n:])
		n += TimestampNativeProtobuf.Marshal(data.Time, s.bs[n:])
	}
	bs = s.bs[:n]
	return
}

func (s SerializerNativeUnsafeReuse) Unmarshal(bs []byte) (
	data *data_proto.DataRaw, err error) {
	var (
		n, n1 int
		l     = len(bs)
		tag   uint64
	)
	data = &data_proto.DataRaw{}
	for n < l {
		tag, n1, err = varint.Uint64.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		switch tag {
		case strFieldTag:
			data.Str, n1, err = unsafe.String.Unmarshal(bs[n:])
		case boolFieldTag:
			data.Bool, n1, err = unsafe.Bool.Unmarshal(bs[n:])
		case int32FieldTag:
			data.Int32, n1, err = unsafe.Int32.Unmarshal(bs[n:])
		case float64FieldTag:
			data.Float64, n1, err = unsafe.Float64.Unmarshal(bs[n:])
		case timeFieldTag:
			data.Time, n1, err = TimestampNativeProtobuf.Unmarshal(bs[n:])
		default:
			err = fmt.Errorf("unexpected tag %v", tag)
		}
		n += n1
		if err != nil {
			return
		}
	}
	return
}
