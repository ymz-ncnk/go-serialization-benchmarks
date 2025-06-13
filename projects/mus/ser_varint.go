package mus

import (
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/common"
)

type SerializerRawVarint struct{}

func (s SerializerRawVarint) Marshal(data common.Data) (bs []byte, err error) {
	n := ord.String.Size(data.Str)
	n += ord.Bool.Size(data.Bool)
	n += varint.Int32.Size(data.Int32)
	n += raw.Float64.Size(data.Float64)
	n += raw.TimeUnixNanoUTC.Size(data.Time)
	bs = make([]byte, n)
	n = ord.String.Marshal(data.Str, bs)
	n += ord.Bool.Marshal(data.Bool, bs[n:])
	n += varint.Int32.Marshal(data.Int32, bs[n:])
	n += raw.Float64.Marshal(data.Float64, bs[n:])
	raw.TimeUnixNanoUTC.Marshal(data.Time, bs[n:])
	return
}

func (s SerializerRawVarint) Unmarshal(bs []byte) (data common.Data, err error) {
	var n, n1 int
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
	data.Time, _, err = raw.TimeUnixNanoUTC.Unmarshal(bs[n:])
	return
}

func (s SerializerRawVarint) Name() benchser.ResultName {
	return benchser.NewResultName(MUS, s.features()...)
}

func (s SerializerRawVarint) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s SerializerRawVarint) features() []benchser.Feature {
	return []benchser.Feature{benchser.Varint}
}
