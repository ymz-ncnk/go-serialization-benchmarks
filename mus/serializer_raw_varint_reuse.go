package mus

import (
	"time"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerRawVarintReuse struct {
	bs []byte
}

func (s SerializerRawVarintReuse) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.Raw, serializer.Varint,
		serializer.Reuse)
}

func (s SerializerRawVarintReuse) Features() []serializer.Feature {
	return Features
}

func (s SerializerRawVarintReuse) Marshal(data serializer.Data) (bs []byte, err error) {
	n := ord.MarshalString(data.Str, s.bs)
	n += ord.MarshalBool(data.Bool, s.bs[n:])
	n += varint.MarshalInt32(data.Int32, s.bs[n:])
	n += varint.MarshalFloat64(data.Float64, s.bs[n:])
	n += raw.MarshalInt64(data.Time.UnixNano(), s.bs[n:])
	bs = s.bs[:n]
	return
}

func (s SerializerRawVarintReuse) Unmarshal(bs []byte) (data serializer.Data, err error) {
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
	nano, _, err = unsafe.UnmarshalInt64(bs[n:])
	if err != nil {
		return
	}
	data.Time = time.Unix(0, nano)
	return
}
