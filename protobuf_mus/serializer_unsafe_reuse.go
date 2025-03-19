package protobuf_mus

import (
	"fmt"

	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerMUSUnsafeReuse struct {
	bs []byte
}

func (s SerializerMUSUnsafeReuse) Name() serializer.ResultName {
	return serializer.NewResultName(Protobuf, serializer.Unsafe, serializer.Reuse)
}

func (s SerializerMUSUnsafeReuse) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Unsafe, serializer.Reuse)
}

func (s SerializerMUSUnsafeReuse) Marshal(data general.Data) (bs []byte,
	err error) {
	var (
		n         int
		timestamp = NewTimestamp(data.Time)
	)
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
	if timestamp.Seconds != 0 || timestamp.Nanos != 0 {
		n += varint.Uint64.Marshal(timeFieldTag, s.bs[n:])
		n += TimestampProtobuf.Marshal(timestamp, s.bs[n:])
	}
	bs = s.bs[:n]
	return
}

func (s SerializerMUSUnsafeReuse) Unmarshal(bs []byte) (data general.Data,
	err error) {
	var (
		n, n1 int
		l     = len(bs)
		tag   uint64
	)
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
			var timestamp Timestamp
			timestamp, n1, err = TimestampProtobuf.Unmarshal(bs[n:])
			if err != nil {
				return
			}
			data.Time = timestamp.ToTime()
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
