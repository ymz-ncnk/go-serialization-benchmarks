package protobuf

import (
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
	"google.golang.org/protobuf/proto"
)

type SerializerRaw struct{}

func (s SerializerRaw) Name() serializer.ResultName {
	return serializer.NewResultName(Protobuf, serializer.Raw)
}

func (s SerializerRaw) Features() []serializer.Feature {
	return Features
}

func (s SerializerRaw) Marshal(data *data_proto.DataRaw) (bs []byte, err error) {
	return proto.Marshal(data)
}

func (s SerializerRaw) Unmarshal(bs []byte) (data *data_proto.DataRaw, err error) {
	data = &data_proto.DataRaw{}
	err = proto.Unmarshal(bs, data)
	return
}
