package vtprotobuf

import (
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type VTSerializerRaw struct{}

func (s VTSerializerRaw) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw)
}

func (s VTSerializerRaw) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Raw)
}

func (s VTSerializerRaw) Marshal(data *data_proto.DataRaw) (bs []byte, err error) {
	bs = make([]byte, data.SizeVT())
	_, err = data.MarshalToSizedBufferVT(bs)
	return
}

func (s VTSerializerRaw) Unmarshal(bs []byte) (data *data_proto.DataRaw, err error) {
	data = &data_proto.DataRaw{}
	err = data.UnmarshalVT(bs)
	return
}
