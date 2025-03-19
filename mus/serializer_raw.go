package mus

import (
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerRaw struct{}

func (s SerializerRaw) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.Raw)
}

func (s SerializerRaw) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Raw)
}

func (s SerializerRaw) Marshal(data general.Data) (bs []byte, err error) {
	n := ord.String.Size(data.Str)
	n += ord.Bool.Size(data.Bool)
	n += raw.Int32.Size(data.Int32)
	n += raw.Float64.Size(data.Float64)
	n += raw.TimeUnixNanoUTC.Size(data.Time)
	bs = make([]byte, n)
	n = ord.String.Marshal(data.Str, bs)
	n += ord.Bool.Marshal(data.Bool, bs[n:])
	n += raw.Int32.Marshal(data.Int32, bs[n:])
	n += raw.Float64.Marshal(data.Float64, bs[n:])
	raw.TimeUnixNanoUTC.Marshal(data.Time, bs[n:])
	return
}

func (s SerializerRaw) Unmarshal(bs []byte) (data general.Data, err error) {
	var (
		n  int
		n1 int
	)
	data.Str, n, err = ord.String.Unmarshal(bs)
	if err != nil {
		return
	}
	data.Bool, n1, err = ord.Bool.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Int32, n1, err = raw.Int32.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Float64, n1, err = raw.Float64.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Time, n, err = raw.TimeUnixNanoUTC.Unmarshal(bs[n:])
	n += n1
	return
}
