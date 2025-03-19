package vtprotobuf

import (
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type VTSerializerVarintUnsafeUnmReuse struct {
	bs []byte
}

func (s VTSerializerVarintUnsafeUnmReuse) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Varint,
		serializer.UnsafeUnm, serializer.Reuse)
}

func (s VTSerializerVarintUnsafeUnmReuse) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Varint, serializer.UnsafeUnm,
		serializer.Reuse)
}

func (s VTSerializerVarintUnsafeUnmReuse) Marshal(data *data_proto.DataRawVarint) (
	bs []byte, err error) {
	n, err := data.MarshalToSizedBufferVT(s.bs)
	if err != nil {
		return
	}
	bs = s.bs[len(s.bs)-n:]
	return
}

func (s VTSerializerVarintUnsafeUnmReuse) Unmarshal(bs []byte) (
	data *data_proto.DataRawVarint, err error) {
	data = &data_proto.DataRawVarint{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}
