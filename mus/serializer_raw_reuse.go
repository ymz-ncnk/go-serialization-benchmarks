package mus

import (
	"time"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerRawReuse struct {
	bs []byte
}

func (s SerializerRawReuse) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.Raw, serializer.Reuse)
}

func (s SerializerRawReuse) Features() []serializer.Feature {
	return Features
}

func (s SerializerRawReuse) Marshal(data serializer.Data) (bs []byte, err error) {
	n := ord.MarshalString(data.Str, nil, s.bs)
	n += ord.MarshalBool(data.Bool, s.bs[n:])
	n += raw.MarshalInt32(data.Int32, s.bs[n:])
	n += raw.MarshalFloat64(data.Float64, s.bs[n:])
	n += raw.MarshalInt64(data.Time.UnixNano(), s.bs[n:])
	bs = s.bs[:n]
	return
}

func (s SerializerRawReuse) Unmarshal(bs []byte) (data serializer.Data, err error) {
	var (
		n    int
		n1   int
		nano int64
	)
	data.Str, n, err = ord.UnmarshalString(nil, bs)
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
