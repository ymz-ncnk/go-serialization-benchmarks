package mus

import (
	"time"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerNotUnsafeReuse struct {
	bs []byte
}

func (s SerializerNotUnsafeReuse) Name() serializer.ResultName {
	return serializer.NewResultName(MUS, serializer.NotUnsafe, serializer.Reuse)
}

func (s SerializerNotUnsafeReuse) Features() []serializer.Feature {
	return Features
}

func (s SerializerNotUnsafeReuse) Marshal(data serializer.Data) (bs []byte, err error) {
	n := ord.MarshalString(data.Str, nil, s.bs)
	n += unsafe.MarshalBool(data.Bool, s.bs[n:])
	n += unsafe.MarshalInt32(data.Int32, s.bs[n:])
	n += unsafe.MarshalFloat64(data.Float64, s.bs[n:])
	n += unsafe.MarshalInt64(data.Time.UnixNano(), s.bs[n:])
	bs = s.bs[:n]
	return
}

func (s SerializerNotUnsafeReuse) Unmarshal(bs []byte) (data serializer.Data, err error) {
	var (
		n    int
		n1   int
		nano int64
	)
	data.Str, n, err = ord.UnmarshalString(nil, bs)
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
