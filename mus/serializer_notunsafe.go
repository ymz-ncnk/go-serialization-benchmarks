package mus

import (
	"time"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerNotUnsafe struct{}

func (s SerializerNotUnsafe) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.NotUnsafe)
}

func (s SerializerNotUnsafe) Features() []serializer.Feature {
	return Features
}

func (s SerializerNotUnsafe) Marshal(data serializer.Data) (bs []byte, err error) {
	nano := data.Time.UnixNano()
	n := ord.SizeString(data.Str)
	n += unsafe.SizeBool(data.Bool)
	n += unsafe.SizeInt32(data.Int32)
	n += unsafe.SizeFloat64(data.Float64)
	n += unsafe.SizeInt64(nano)
	bs = make([]byte, n)
	n = ord.MarshalString(data.Str, bs)
	n += unsafe.MarshalBool(data.Bool, bs[n:])
	n += unsafe.MarshalInt32(data.Int32, bs[n:])
	n += unsafe.MarshalFloat64(data.Float64, bs[n:])
	unsafe.MarshalInt64(nano, bs[n:])
	return
}

func (s SerializerNotUnsafe) Unmarshal(bs []byte) (data serializer.Data, err error) {
	var (
		n    int
		n1   int
		nano int64
	)
	data.Str, n, err = ord.UnmarshalString(bs)
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
