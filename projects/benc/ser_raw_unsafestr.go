package benc

import (
	"time"

	bstd "github.com/deneonet/benc/std"
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/common"
)

type SerializerRawUnsafeStr struct{}

func (s SerializerRawUnsafeStr) Marshal(data common.Data) (bs []byte,
	err error) {
	size := bstd.SizeString(data.Str)
	size += bstd.SizeBool()
	size += bstd.SizeInt64()
	size += bstd.SizeFloat64()
	size += bstd.SizeInt64()
	bs = make([]byte, size)
	n := bstd.MarshalUnsafeString(0, bs, data.Str)
	n = bstd.MarshalBool(n, bs, data.Bool)
	n = bstd.MarshalInt32(n, bs, data.Int32)
	n = bstd.MarshalFloat64(n, bs, data.Float64)
	n = bstd.MarshalInt64(n, bs, data.Time.UnixNano())
	bs = bs[:n]
	return
}

func (s SerializerRawUnsafeStr) Unmarshal(bs []byte) (data common.Data,
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

func (s SerializerRawUnsafeStr) Name() benchser.ResultName {
	return benchser.NewResultName(Benc, s.features()...)
}

func (s SerializerRawUnsafeStr) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s SerializerRawUnsafeStr) features() []benchser.Feature {
	return []benchser.Feature{benchser.Raw, benchser.UnsafeStr}
}
