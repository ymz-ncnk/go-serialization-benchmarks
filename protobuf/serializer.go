package protobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
	"google.golang.org/protobuf/proto"
)

const Protobuf = "protobuf"

type Serializer struct{}

func (s Serializer) Name() serializer.ResultName {
	return serializer.NewResultName(Protobuf)
}

func (s Serializer) Features() []serializer.Feature {
	return Features
}

func (s Serializer) Marshal(data *Data) (bs []byte, err error) {
	return proto.Marshal(data)
}

func (s Serializer) Unmarshal(bs []byte) (data *Data, err error) {
	data = &Data{}
	err = proto.Unmarshal(bs, data)
	return
}
