package benc

import (
	"time"

	"github.com/deneonet/benc/bstd"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type SerializerRawReuse struct {
	bs []byte
}

func (s SerializerRawReuse) Name() serializer.ResultName {
	return serializer.NewResultName(Benc, serializer.Raw, serializer.Reuse)
}

func (s SerializerRawReuse) Features() []serializer.Feature {
	return Features
}

func (s SerializerRawReuse) Marshal(data general.Data) (bs []byte, err error) {
	var n int
	n, err = bstd.MarshalString(n, s.bs, data.Str)
	if err != nil {
		return
	}
	n = bstd.MarshalBool(n, s.bs, data.Bool)
	n = bstd.MarshalInt32(n, s.bs, data.Int32)
	n = bstd.MarshalFloat64(n, s.bs, data.Float64)
	n = bstd.MarshalInt64(n, s.bs, data.Time.UnixNano())
	bs = s.bs[:n]
	return
}

func (s SerializerRawReuse) Unmarshal(bs []byte) (data general.Data, err error) {
	var (
		n   int
		n64 int64
	)
	n, data.Str, err = bstd.UnmarshalString(n, bs)
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
