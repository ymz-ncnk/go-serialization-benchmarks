package mus

import (
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerRawVarintReuse struct {
	bs []byte
}

func (s SerializerRawVarintReuse) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.Varint, serializer.Reuse)
}

func (s SerializerRawVarintReuse) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Varint, serializer.Reuse)
}

func (s SerializerRawVarintReuse) Marshal(data general.Data) (bs []byte, err error) {
	n := ord.String.Marshal(data.Str, s.bs)
	n += ord.Bool.Marshal(data.Bool, s.bs[n:])
	n += varint.Int32.Marshal(data.Int32, s.bs[n:])
	n += raw.Float64.Marshal(data.Float64, s.bs[n:])
	n += raw.TimeUnixNanoUTC.Marshal(data.Time, s.bs[n:])
	bs = s.bs[:n]
	return
}

func (s SerializerRawVarintReuse) Unmarshal(bs []byte) (data general.Data, err error) {
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
	data.Int32, n1, err = varint.Int32.Unmarshal(bs[n:])
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
