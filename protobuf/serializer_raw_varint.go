package protobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
	"google.golang.org/protobuf/proto"
)

type SerializerRawVarint struct{}

func (s SerializerRawVarint) Name() serializer.ResultName {
	return serializer.NewResultName(Protobuf, serializer.Raw, serializer.Varint)
}

func (s SerializerRawVarint) Features() []serializer.Feature {
	return Features
}

func (s SerializerRawVarint) Marshal(data *DataRawVarint) (bs []byte, err error) {
	return proto.Marshal(data)
}

func (s SerializerRawVarint) Unmarshal(bs []byte) (data *DataRawVarint, err error) {
	data = &DataRawVarint{}
	err = proto.Unmarshal(bs, data)
	return
}
