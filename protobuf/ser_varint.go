package protobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"google.golang.org/protobuf/proto"
)

type SerializerVarint struct{}

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

func (s SerializerVarint) Name() benchser.ResultName {
	return benchser.NewResultName(Protobuf, s.features()...)
}

func (s SerializerVarint) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s SerializerVarint) features() []benchser.Feature {
	return []benchser.Feature{benchser.Varint}
}
