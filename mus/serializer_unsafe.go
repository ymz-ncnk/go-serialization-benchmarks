package mus

import (
	"time"

	"github.com/mus-format/mus-go/unsafe"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerUnsafe struct{}

func (s SerializerUnsafe) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.Unsafe)
}

func (s SerializerUnsafe) Features() []serializer.Feature {
	return Features
}

func (s SerializerUnsafe) Marshal(data serializer.Data) (bs []byte, err error) {
	n := unsafe.SizeString(data.Str)
	n += unsafe.SizeBool(data.Bool)
	n += unsafe.SizeInt32(data.Int32)
	n += unsafe.SizeFloat64(data.Float64)
	n += unsafe.SizeInt64(data.Time.UnixNano())
	bs = make([]byte, n)
	n = unsafe.MarshalString(data.Str, bs)
	n += unsafe.MarshalBool(data.Bool, bs[n:])
	n += unsafe.MarshalInt32(data.Int32, bs[n:])
	n += unsafe.MarshalFloat64(data.Float64, bs[n:])
	unsafe.MarshalInt64(data.Time.UnixNano(), bs[n:])
	return
}

func (s SerializerUnsafe) Unmarshal(bs []byte) (data serializer.Data, err error) {
	var (
		n    int
		n1   int
		nano int64
	)
	data.Str, n, err = unsafe.UnmarshalString(bs)
	if err != nil {
		return
	}
	data.Bool, n1, err = unsafe.UnmarshalBool(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Int32, n1, err = unsafe.UnmarshalInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	data.Float64, n1, err = unsafe.UnmarshalFloat64(bs[n:])
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
