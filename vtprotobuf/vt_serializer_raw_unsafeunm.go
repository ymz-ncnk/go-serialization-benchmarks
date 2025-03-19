package vtprotobuf

import (
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type VTSerializerRawUnsafeUnm struct{}

func (s VTSerializerRawUnsafeUnm) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw,
		serializer.UnsafeUnm)
}

func (s VTSerializerRawUnsafeUnm) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Raw, serializer.UnsafeUnm)
}

func (s VTSerializerRawUnsafeUnm) Marshal(data *data_proto.DataRaw) (bs []byte,
	err error) {
	bs = make([]byte, data.SizeVT())
	_, err = data.MarshalToSizedBufferVT(bs)
	return
}

func (s VTSerializerRawUnsafeUnm) Unmarshal(bs []byte) (data *data_proto.DataRaw,
	err error) {
	data = &data_proto.DataRaw{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}
