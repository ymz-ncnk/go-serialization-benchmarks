package mus

import (
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/common"
)

type SerializerNotUnsafe struct{}

func (s SerializerNotUnsafe) Marshal(data common.Data) (bs []byte, err error) {
	n := ord.String.Size(data.Str)
	n += unsafe.Bool.Size(data.Bool)
	n += unsafe.Int32.Size(data.Int32)
	n += unsafe.Float64.Size(data.Float64)
	n += unsafe.TimeUnixNanoUTC.Size(data.Time)
	bs = make([]byte, n)
	n = ord.String.Marshal(data.Str, bs)
	n += unsafe.Bool.Marshal(data.Bool, bs[n:])
	n += unsafe.Int32.Marshal(data.Int32, bs[n:])
	n += unsafe.Float64.Marshal(data.Float64, bs[n:])
	unsafe.TimeUnixNanoUTC.Marshal(data.Time, bs[n:])
	return
}

func (s SerializerNotUnsafe) Unmarshal(bs []byte) (data common.Data, err error) {
	var n, n1 int
	data.Str, n, err = ord.String.Unmarshal(bs)
	if err != nil {
		return
	}
	data.Bool, n1, err = unsafe.Bool.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Int32, n1, err = unsafe.Int32.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Float64, n1, err = unsafe.Float64.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Time, _, err = unsafe.TimeUnixNanoUTC.Unmarshal(bs[n:])
	return
}

func (s SerializerNotUnsafe) Name() benchser.ResultName {
	return benchser.NewResultName(MUS, s.features()...)
}

func (s SerializerNotUnsafe) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s SerializerNotUnsafe) features() []benchser.Feature {
	return []benchser.Feature{benchser.NotUnsafe}
}
