package vtprotobuf

import (
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type VTSerializerVarintUnsafeUnm struct{}

func (s VTSerializerVarintUnsafeUnm) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Varint,
		serializer.UnsafeUnm)
}

func (s VTSerializerVarintUnsafeUnm) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Varint, serializer.UnsafeUnm)
}

func (s VTSerializerVarintUnsafeUnm) Marshal(data *data_proto.DataRawVarint) (
	bs []byte, err error) {
	bs = make([]byte, data.SizeVT())
	_, err = data.MarshalToSizedBufferVT(bs)
	return
}

func (s VTSerializerVarintUnsafeUnm) Unmarshal(bs []byte) (
	data *data_proto.DataRawVarint, err error) {
	data = &data_proto.DataRawVarint{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}
