package protobuf_mus

import (
	"fmt"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerMUSRawReuse struct {
	bs []byte
}

func (s SerializerMUSRawReuse) Name() serializer.ResultName {
	return serializer.NewResultName(Protobuf, serializer.Raw, serializer.Reuse)
}

func (s SerializerMUSRawReuse) Features() []serializer.Feature {
	return Features
}

func (s SerializerMUSRawReuse) Marshal(data serializer.Data) (bs []byte,
	err error) {
	var (
		n         int
		timestamp = NewTimestamp(data.Time)
	)
	if data.Str != "" {
		n += varint.MarshalUint64(strFieldTag, s.bs[n:])
		n += ord.MarshalString(data.Str, lenM,
			s.bs[n:])
	}
	if data.Bool {
		n += varint.MarshalUint64(boolFieldTag, s.bs[n:])
		n += ord.MarshalBool(data.Bool, s.bs[n:])
	}
	if data.Int32 != 0 {
		n += varint.MarshalUint64(int32FieldTag, s.bs[n:])
		n += raw.MarshalInt32(data.Int32, s.bs[n:])
	}
	if data.Float64 != 0 {
		n += varint.MarshalUint64(float64FieldTag, s.bs[n:])
		n += raw.MarshalFloat64(data.Float64, s.bs[n:])
	}
	if timestamp.Seconds != 0 || timestamp.Nanos != 0 {
		n += varint.MarshalUint64(timeFieldTag, s.bs[n:])
		n += varint.MarshalPositiveInt(SizeTimestamp(timestamp), s.bs[n:])
		n += MarshalTimestamp(timestamp, s.bs[n:])
	}
	bs = s.bs[:n]
	return
}

func (s SerializerMUSRawReuse) Unmarshal(bs []byte) (data serializer.Data,
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
