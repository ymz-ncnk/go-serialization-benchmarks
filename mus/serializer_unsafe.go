package mus

import (
	"github.com/mus-format/mus-go/unsafe"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerUnsafe struct{}

func (s SerializerUnsafe) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.Unsafe)
}

func (s SerializerUnsafe) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Unsafe)
}

func (s SerializerUnsafe) Marshal(data general.Data) (bs []byte, err error) {
	n := unsafe.String.Size(data.Str)
	n += unsafe.Bool.Size(data.Bool)
	n += unsafe.Int32.Size(data.Int32)
	n += unsafe.Float64.Size(data.Float64)
	n += unsafe.TimeUnixNanoUTC.Size(data.Time)
	bs = make([]byte, n)
	n = unsafe.String.Marshal(data.Str, bs)
	n += unsafe.Bool.Marshal(data.Bool, bs[n:])
	n += unsafe.Int32.Marshal(data.Int32, bs[n:])
	n += unsafe.Float64.Marshal(data.Float64, bs[n:])
	unsafe.TimeUnixNanoUTC.Marshal(data.Time, bs[n:])
	return
}

func (s SerializerUnsafe) Unmarshal(bs []byte) (data general.Data, err error) {
	var (
		n  int
		n1 int
	)
	data.Str, n, err = unsafe.String.Unmarshal(bs)
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
	data.Time, n, err = unsafe.TimeUnixNanoUTC.Unmarshal(bs[n:])
	n += n1
	return
}
