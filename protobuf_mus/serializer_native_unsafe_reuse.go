package protobuf_mus

import (
	"fmt"

	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/protobuf"
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
	return Features
}

func (s SerializerNativeUnsafeReuse) Marshal(data *protobuf.DataRaw) (bs []byte,
	err error) {
	var n int
	if data.Str != "" {
		n += varint.MarshalUint64(strFieldTag, s.bs[n:])
		n += unsafe.MarshalString(data.Str, lenM, s.bs[n:])
	}
	if data.Bool {
		n += varint.MarshalUint64(boolFieldTag, s.bs[n:])
		n += unsafe.MarshalBool(data.Bool, s.bs[n:])
	}
	if data.Int32 != 0 {
		n += varint.MarshalUint64(int32FieldTag, s.bs[n:])
		n += unsafe.MarshalInt32(data.Int32, s.bs[n:])
	}
	if data.Float64 != 0 {
		n += varint.MarshalUint64(float64FieldTag, s.bs[n:])
		n += unsafe.MarshalFloat64(data.Float64, s.bs[n:])
	}
	if data.Time != nil && (data.Time.Seconds != 0 || data.Time.Nanos != 0) {
		n += varint.MarshalUint64(timeFieldTag, s.bs[n:])
		n += varint.MarshalPositiveInt(SizeTimestampNative(data.Time), s.bs[n:])
		n += MarshalTimestampNative(data.Time, s.bs[n:])
	}
	bs = s.bs[:n]
	return
}

func (s SerializerNativeUnsafeReuse) Unmarshal(bs []byte) (
	data *protobuf.DataRaw, err error) {
	var (
		n, n1 int
		l     = len(bs)
		tag   uint64
	)
	data = &protobuf.DataRaw{}
	for n < l {
		tag, n1, err = varint.UnmarshalUint64(bs[n:])
		n += n1
		if err != nil {
			return
		}
		switch tag {
		case strFieldTag:
			data.Str, n1, err = unsafe.UnmarshalString(lenU,
				bs[n:])
		case boolFieldTag:
			data.Bool, n1, err = unsafe.UnmarshalBool(bs[n:])
		case int32FieldTag:
			data.Int32, n1, err = unsafe.UnmarshalInt32(bs[n:])
		case float64FieldTag:
			data.Float64, n1, err = unsafe.UnmarshalFloat64(bs[n:])
		case timeFieldTag:
			_, n1, err = varint.UnmarshalPositiveInt(bs[n:])
			n += n1
			if err != nil {
				return
			}
			data.Time, n1, err = UnmarshalTimestampNative(bs[n:])
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
