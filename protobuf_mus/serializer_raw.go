package protobuf_mus

import (
	"fmt"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerMUSRaw struct{}

func (s SerializerMUSRaw) Name() serializer.ResultName {
	return serializer.NewResultName(Protobuf, serializer.Raw)
}

func (s SerializerMUSRaw) Features() []serializer.Feature {
	return Features
}

func (s SerializerMUSRaw) Marshal(data serializer.Data) (bs []byte,
	err error) {
	var (
		n         int
		timestamp = NewTimestamp(data.Time)
	)
	if data.Str != "" {
		n += varint.SizeUint64(strFieldTag)
		n += ord.SizeString(data.Str, lenS)
	}
	if data.Bool {
		n += varint.SizeUint64(boolFieldTag)
		n += ord.SizeBool(data.Bool)
	}
	if data.Int32 != 0 {
		n += varint.SizeUint64(int32FieldTag)
		n += raw.SizeInt32(data.Int32)
	}
	if data.Float64 != 0 {
		n += varint.SizeUint64(float64FieldTag)
		n += raw.SizeFloat64(data.Float64)
	}
	if timestamp.Seconds != 0 || timestamp.Nanos != 0 {
		sizeTimestamp := SizeTimestamp(timestamp)
		n += varint.SizeUint64(timeFieldTag)
		n += varint.SizePositiveInt(sizeTimestamp)
		n += sizeTimestamp
	}

	bs = make([]byte, n)
	n = 0

	if data.Str != "" {
		n += varint.MarshalUint64(strFieldTag, bs[n:])
		n += ord.MarshalString(data.Str, lenM, bs[n:])
	}
	if data.Bool {
		n += varint.MarshalUint64(boolFieldTag, bs[n:])
		n += ord.MarshalBool(data.Bool, bs[n:])
	}
	if data.Int32 != 0 {
		n += varint.MarshalUint64(int32FieldTag, bs[n:])
		n += raw.MarshalInt32(data.Int32, bs[n:])
	}
	if data.Float64 != 0 {
		n += varint.MarshalUint64(float64FieldTag, bs[n:])
		n += raw.MarshalFloat64(data.Float64, bs[n:])
	}

	if timestamp.Seconds != 0 || timestamp.Nanos != 0 {
		n += varint.MarshalUint64(timeFieldTag, bs[n:])
		n += varint.MarshalPositiveInt(SizeTimestamp(timestamp), bs[n:])
		n += MarshalTimestamp(timestamp, bs[n:])
	}
	return
}

func (s SerializerMUSRaw) Unmarshal(bs []byte) (data serializer.Data,
	err error) {
	var (
		n, n1     int
		l         = len(bs)
		tag       uint64
		timestamp Timestamp
	)
	for n < l {
		tag, n1, err = varint.UnmarshalUint64(bs[n:])
		n += n1
		if err != nil {
			return
		}
		switch tag {
		case strFieldTag:
			data.Str, n1, err = ord.UnmarshalString(lenU, bs[n:])
		case boolFieldTag:
			data.Bool, n1, err = ord.UnmarshalBool(bs[n:])
		case int32FieldTag:
			data.Int32, n1, err = raw.UnmarshalInt32(bs[n:])
		case float64FieldTag:
			data.Float64, n1, err = raw.UnmarshalFloat64(bs[n:])
		case timeFieldTag:
			n1, err = varint.SkipPositiveInt(bs[n:])
			n += n1
			if err != nil {
				return
			}
			timestamp, n1, err = UnmarshalTimestamp(bs[n:])
		default:
			err = fmt.Errorf("unexpected tag %v", tag)
		}
		n += n1
		if err != nil {
			return
		}
	}
	data.Time = timestamp.ToTime()
	return
}
