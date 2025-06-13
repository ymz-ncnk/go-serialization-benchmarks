package protobuf_mus

import (
	"fmt"

	ext "github.com/mus-format/ext-protobuf-go"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
)

type SerializerUnsafeReuse struct {
	bs []byte
}

func (s SerializerUnsafeReuse) Marshal(data Data) (bs []byte,
	err error) {
	var n int
	if data.Str != "" {
		n += varint.Uint64.Marshal(strFieldTag, s.bs[n:])
		n += ext.StringUnsafe.Marshal(data.Str, s.bs[n:])
	}
	if data.Bool {
		n += varint.Uint64.Marshal(boolFieldTag, s.bs[n:])
		n += unsafe.Bool.Marshal(data.Bool, s.bs[n:])
	}
	if data.Int32 != 0 {
		n += varint.Uint64.Marshal(int32FieldTag, s.bs[n:])
		n += unsafe.Int32.Marshal(data.Int32, s.bs[n:])
	}
	if data.Float64 != 0 {
		n += varint.Uint64.Marshal(float64FieldTag, s.bs[n:])
		n += unsafe.Float64.Marshal(data.Float64, s.bs[n:])
	}
	if data.Time.Seconds != 0 || data.Time.Nanos != 0 {
		n += varint.Uint64.Marshal(timeFieldTag, s.bs[n:])
		n += ext.TimestampProtobuf.Marshal(data.Time, s.bs[n:])
	}
	bs = s.bs[:n]
	return
}

func (s SerializerUnsafeReuse) Unmarshal(bs []byte) (data Data,
	err error) {
	var (
		n, n1 int
		l     = len(bs)
		tag   uint64
	)
	for n < l {
		tag, n1, err = varint.Uint64.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		switch tag {
		case strFieldTag:
			data.Str, n1, err = ext.StringUnsafe.Unmarshal(bs[n:])
		case boolFieldTag:
			data.Bool, n1, err = unsafe.Bool.Unmarshal(bs[n:])
		case int32FieldTag:
			data.Int32, n1, err = unsafe.Int32.Unmarshal(bs[n:])
		case float64FieldTag:
			data.Float64, n1, err = unsafe.Float64.Unmarshal(bs[n:])
		case timeFieldTag:
			data.Time, n1, err = ext.TimestampProtobuf.Unmarshal(bs[n:])
		default:
			err = fmt.Errorf("unexpected tag %v", tag)
		}
		n += n1
		if err != nil {
			return
		}
	}
	return
}

func (s SerializerUnsafeReuse) Name() benchser.ResultName {
	return benchser.NewResultName(Protobuf, s.features()...)
}

func (s SerializerUnsafeReuse) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s SerializerUnsafeReuse) features() []benchser.Feature {
	return []benchser.Feature{benchser.Unsafe, benchser.Reuse}
}
