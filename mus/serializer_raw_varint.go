package mus

import (
	"time"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

const MUS = "mus"

type SerializerRawVarint struct{}

func (s SerializerRawVarint) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.Raw, serializer.Varint)
}

func (s SerializerRawVarint) Features() []serializer.Feature {
	return Features
}

func (s SerializerRawVarint) Marshal(data serializer.Data) (bs []byte, err error) {
	nano := data.Time.UnixNano()
	n := ord.SizeString(data.Str)
	n += ord.SizeBool(data.Bool)
	n += varint.SizeInt32(data.Int32)
	n += varint.SizeFloat64(data.Float64)
	n += raw.SizeInt64(nano)
	bs = make([]byte, n)
	n = ord.MarshalString(data.Str, bs)
	n += ord.MarshalBool(data.Bool, bs[n:])
	n += varint.MarshalInt32(data.Int32, bs[n:])
	n += varint.MarshalFloat64(data.Float64, bs[n:])
	raw.MarshalInt64(nano, bs[n:])
	return
}

func (s SerializerRawVarint) Unmarshal(bs []byte) (data serializer.Data, err error) {
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
	data.Int32, n1, err = varint.UnmarshalInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Float64, n1, err = varint.UnmarshalFloat64(bs[n:])
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
