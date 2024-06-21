package mus

import (
	"time"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerRaw struct{}

func (s SerializerRaw) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.Raw)
}

func (s SerializerRaw) Features() []serializer.Feature {
	return Features
}

func (s SerializerRaw) Marshal(data serializer.Data) (bs []byte, err error) {
	nano := data.Time.UnixNano()
	n := ord.SizeString(data.Str)
	n += ord.SizeBool(data.Bool)
	n += raw.SizeInt32(data.Int32)
	n += raw.SizeFloat64(data.Float64)
	n += raw.SizeInt64(nano)
	bs = make([]byte, n)
	n = ord.MarshalString(data.Str, bs)
	n += ord.MarshalBool(data.Bool, bs[n:])
	n += raw.MarshalInt32(data.Int32, bs[n:])
	n += raw.MarshalFloat64(data.Float64, bs[n:])
	raw.MarshalInt64(nano, bs[n:])
	return
}

func (s SerializerRaw) Unmarshal(bs []byte) (data serializer.Data, err error) {
	var (
		n    int
		n1   int
		nano int64
	)
	data.Str, n, err = ord.UnmarshalString(bs)
	if err != nil {
		return
	}
	data.Bool, n1, err = ord.UnmarshalBool(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Int32, n1, err = raw.UnmarshalInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Float64, n1, err = raw.UnmarshalFloat64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	nano, _, err = raw.UnmarshalInt64(bs[n:])
	if err != nil {
		return
	}
	data.Time = time.Unix(0, nano)
	return
}
