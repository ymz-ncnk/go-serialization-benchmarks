package protobuf_mus

import (
	"fmt"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerMUSVarint struct{}

func (s SerializerMUSVarint) Name() serializer.ResultName {
	return serializer.NewResultName(Protobuf, serializer.Varint)
}

func (s SerializerMUSVarint) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Reuse)
}

func (s SerializerMUSVarint) Marshal(data general.Data) (bs []byte,
	err error) {
	var (
		n         int
		timestamp = NewTimestamp(data.Time)
	)
	if data.Str != "" {
		n += varint.Uint64.Size(strFieldTag)
		n += ord.String.Size(data.Str)
	}
	if data.Bool {
		n += varint.Uint64.Size(boolFieldTag)
		n += ord.Bool.Size(data.Bool)
	}
	if data.Int32 != 0 {
		n += varint.Uint64.Size(int32FieldTag)
		n += varint.Int32.Size(data.Int32)
	}
	if data.Float64 != 0 {
		n += varint.Uint64.Size(float64FieldTag)
		n += raw.Float64.Size(data.Float64)
	}
	if timestamp.Seconds != 0 || timestamp.Nanos != 0 {
		n += varint.Uint64.Size(timeFieldTag)
		n += TimestampProtobuf.Size(timestamp)
	}

	bs = make([]byte, n)
	n = 0

	if data.Str != "" {
		n += varint.Uint64.Marshal(strFieldTag, bs[n:])
		n += ord.String.Marshal(data.Str, bs[n:])
	}
	if data.Bool {
		n += varint.Uint64.Marshal(boolFieldTag, bs[n:])
		n += ord.Bool.Marshal(data.Bool, bs[n:])
	}
	if data.Int32 != 0 {
		n += varint.Uint64.Marshal(int32FieldTag, bs[n:])
		n += varint.Int32.Marshal(data.Int32, bs[n:])
	}
	if data.Float64 != 0 {
		n += varint.Uint64.Marshal(float64FieldTag, bs[n:])
		n += raw.Float64.Marshal(data.Float64, bs[n:])
	}
	if timestamp.Seconds != 0 || timestamp.Nanos != 0 {
		n += varint.Uint64.Marshal(timeFieldTag, bs[n:])
		n += TimestampProtobuf.Marshal(timestamp, bs[n:])
	}
	return
}

func (s SerializerMUSVarint) Unmarshal(bs []byte) (data general.Data,
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
			data.Str, n1, err = ord.String.Unmarshal(bs[n:])
		case boolFieldTag:
			data.Bool, n1, err = ord.Bool.Unmarshal(bs[n:])
		case int32FieldTag:
			data.Int32, n1, err = varint.Int32.Unmarshal(bs[n:])
		case float64FieldTag:
			data.Float64, n1, err = raw.Float64.Unmarshal(bs[n:])
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
