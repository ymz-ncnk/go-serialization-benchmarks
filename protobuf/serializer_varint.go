package protobuf

import (
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
	"google.golang.org/protobuf/proto"
)

type SerializerVarint struct{}

func (s SerializerVarint) Name() serializer.ResultName {
	return serializer.NewResultName(Protobuf, serializer.Varint)
}

func (s SerializerVarint) Features() []serializer.Feature {
	return Features
}

func (s SerializerVarint) Marshal(data *data_proto.DataRawVarint) (bs []byte,
	err error) {
	return proto.Marshal(data)
}

func (s SerializerVarint) Unmarshal(bs []byte) (
	data *data_proto.DataRawVarint, err error) {
	data = &data_proto.DataRawVarint{}
	err = proto.Unmarshal(bs, data)
	return
}
