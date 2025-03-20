package vtprotobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
)

type VTSerializerRawReuse struct {
	bs []byte
}

func (s VTSerializerRawReuse) Marshal(data *data_proto.DataRaw) (bs []byte,
	err error) {
	n, err := data.MarshalToSizedBufferVT(s.bs)
	if err != nil {
		return
	}
	bs = s.bs[len(s.bs)-n:]
	return
}

func (s VTSerializerRawReuse) Unmarshal(bs []byte) (data *data_proto.DataRaw,
	err error) {
	data = &data_proto.DataRaw{}
	err = data.UnmarshalVT(bs)
	return
}

func (s VTSerializerRawReuse) Name() benchser.ResultName {
	return benchser.NewResultName(VTProtobuf, s.features()...)
}

func (s VTSerializerRawReuse) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s VTSerializerRawReuse) features() []benchser.Feature {
	return []benchser.Feature{benchser.Raw, benchser.Reuse}
}
