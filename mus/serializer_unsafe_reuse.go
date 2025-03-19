package mus

import (
	"github.com/mus-format/mus-go/unsafe"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerUnsafeReuse struct {
	bs []byte
}

func (s SerializerUnsafeReuse) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.Unsafe, serializer.Reuse)
}

func (s SerializerUnsafeReuse) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Unsafe, serializer.Reuse)
}

func (s SerializerUnsafeReuse) Marshal(data general.Data) (bs []byte, err error) {
	n := unsafe.String.Marshal(data.Str, s.bs)
	n += unsafe.Bool.Marshal(data.Bool, s.bs[n:])
	n += unsafe.Int32.Marshal(data.Int32, s.bs[n:])
	n += unsafe.Float64.Marshal(data.Float64, s.bs[n:])
	n += unsafe.TimeUnixNanoUTC.Marshal(data.Time, s.bs[n:])
	bs = s.bs[:n]
	return
}

func (s SerializerUnsafeReuse) Unmarshal(bs []byte) (data general.Data, err error) {
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
	data.Time, n1, err = unsafe.TimeUnixNanoUTC.Unmarshal(bs[n:])
	n += n1
	return
}
