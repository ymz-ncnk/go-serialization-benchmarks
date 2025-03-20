package benc

import (
	"time"

	bstd "github.com/deneonet/benc/std"
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
)

type SerializerRawUnsafeStrReuse struct {
	bs []byte
}

func (s SerializerRawUnsafeStrReuse) Marshal(data general.Data) (bs []byte,
	err error) {
	n := bstd.MarshalUnsafeString(0, s.bs, data.Str)
	n = bstd.MarshalBool(n, s.bs, data.Bool)
	n = bstd.MarshalInt32(n, s.bs, data.Int32)
	n = bstd.MarshalFloat64(n, s.bs, data.Float64)
	n = bstd.MarshalInt64(n, s.bs, data.Time.UnixNano())
	bs = s.bs[:n]
	return
}

func (s SerializerRawUnsafeStrReuse) Unmarshal(bs []byte) (data general.Data,
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

func (s SerializerRawUnsafeStrReuse) Name() benchser.ResultName {
	return benchser.NewResultName(Benc, s.features()...)
}

func (s SerializerRawUnsafeStrReuse) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s SerializerRawUnsafeStrReuse) features() []benchser.Feature {
	return []benchser.Feature{benchser.Raw, benchser.UnsafeStr,
		benchser.Reuse}
}
