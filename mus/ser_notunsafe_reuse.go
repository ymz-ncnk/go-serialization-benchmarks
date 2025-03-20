package mus

import (
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
)

type SerializerNotUnsafeReuse struct {
	bs []byte
}

func (s SerializerNotUnsafeReuse) Marshal(data general.Data) (bs []byte, err error) {
	n := ord.String.Marshal(data.Str, s.bs)
	n += unsafe.Bool.Marshal(data.Bool, s.bs[n:])
	n += unsafe.Int32.Marshal(data.Int32, s.bs[n:])
	n += unsafe.Float64.Marshal(data.Float64, s.bs[n:])
	n += unsafe.TimeUnixNanoUTC.Marshal(data.Time, s.bs[n:])
	bs = s.bs[:n]
	return
}

func (s SerializerNotUnsafeReuse) Unmarshal(bs []byte) (data general.Data, err error) {
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

func (s SerializerNotUnsafeReuse) Name() benchser.ResultName {
	return benchser.NewResultName(MUS, s.features()...)
}

func (s SerializerNotUnsafeReuse) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s SerializerNotUnsafeReuse) features() []benchser.Feature {
	return []benchser.Feature{benchser.NotUnsafe, benchser.Reuse}
}
