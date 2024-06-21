package benc

import (
	"time"

	"github.com/deneonet/benc"
	"github.com/deneonet/benc/bstd"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerRawUnsafeStr struct{}

func (s SerializerRawUnsafeStr) Name() serializer.ResultName {
	return serializer.NewResultName(Benc, serializer.Raw, serializer.UnsafeStr)
}

func (s SerializerRawUnsafeStr) Features() []serializer.Feature {
	return Features
}

func (s SerializerRawUnsafeStr) Marshal(data serializer.Data) (bs []byte,
	err error) {
	n, err := bstd.SizeString(data.Str)
	if err != nil {
		return
	}
	n += bstd.SizeBool()
	n += bstd.SizeInt64()
	n += bstd.SizeFloat64()
	n += bstd.SizeInt64()
	n, bs = benc.Marshal(n)

	n, err = bstd.MarshalUnsafeString(n, bs, data.Str)
	if err != nil {
		return
	}
	n = bstd.MarshalBool(n, bs, data.Bool)
	n = bstd.MarshalInt32(n, bs, data.Int32)
	n = bstd.MarshalFloat64(n, bs, data.Float64)
	n = bstd.MarshalInt64(n, bs, data.Time.UnixNano())
	bs = bs[:n]
	return
}

func (s SerializerRawUnsafeStr) Unmarshal(bs []byte) (data serializer.Data,
	err error) {
	var (
		n   int
		n64 int64
	)
	n, data.Str, err = bstd.UnmarshalUnsafeString(n, bs)
	if err != nil {
		return
	}
	n, data.Bool, err = bstd.UnmarshalBool(n, bs)
	if err != nil {
		return
	}
	n, data.Int32, err = bstd.UnmarshalInt32(n, bs)
	if err != nil {
		return
	}
	n, data.Float64, err = bstd.UnmarshalFloat64(n, bs)
	if err != nil {
		return
	}
	_, n64, err = bstd.UnmarshalInt64(n, bs)
	if err != nil {
		return
	}
	data.Time = time.Unix(0, n64)
	return
}
