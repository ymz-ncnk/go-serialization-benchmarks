package vtprotobuf

import (
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type VTSerializerVarint struct{}

func (s VTSerializerVarint) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Varint)
}

func (s VTSerializerVarint) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Varint)
}

func (s VTSerializerVarint) Marshal(data *data_proto.DataRawVarint) (
	bs []byte, err error) {
	bs = make([]byte, data.SizeVT())
	_, err = data.MarshalToSizedBufferVT(bs)
	return
}

func (s VTSerializerVarint) Unmarshal(bs []byte) (
	data *data_proto.DataRawVarint, err error) {
	data = &data_proto.DataRawVarint{}
	err = data.UnmarshalVT(bs)
	return
}
