package vtprotobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
)

type VTSerializerRaw struct{}

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

func (s VTSerializerRaw) Name() benchser.ResultName {
	return benchser.NewResultName(VTProtobuf, s.features()...)
}

func (s VTSerializerRaw) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s VTSerializerRaw) features() []benchser.Feature {
	return []benchser.Feature{benchser.Raw}
}
