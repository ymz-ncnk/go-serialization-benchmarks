package protobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"google.golang.org/protobuf/proto"
)

type SerializerRaw struct{}

func (s SerializerRaw) Marshal(data *data_proto.DataRaw) (bs []byte, err error) {
	return proto.Marshal(data)
}

func (s SerializerRaw) Unmarshal(bs []byte) (data *data_proto.DataRaw, err error) {
	data = &data_proto.DataRaw{}
	err = proto.Unmarshal(bs, data)
	return
}

func (s SerializerRaw) Name() benchser.ResultName {
	return benchser.NewResultName(Protobuf, s.features()...)
}

func (s SerializerRaw) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s SerializerRaw) features() []benchser.Feature {
	return []benchser.Feature{benchser.Raw}
}
