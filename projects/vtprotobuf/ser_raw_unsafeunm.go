package vtprotobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
)

type VTSerializerRawUnsafeUnm struct{}

func (s VTSerializerRawUnsafeUnm) Marshal(data *data_proto.DataRaw) (bs []byte,
	err error) {
	bs = make([]byte, data.SizeVT())
	_, err = data.MarshalToSizedBufferVT(bs)
	return
}

func (s VTSerializerRawUnsafeUnm) Unmarshal(bs []byte) (data *data_proto.DataRaw,
	err error) {
	data = &data_proto.DataRaw{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}

func (s VTSerializerRawUnsafeUnm) Name() benchser.ResultName {
	return benchser.NewResultName(VTProtobuf, s.features()...)
}

func (s VTSerializerRawUnsafeUnm) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s VTSerializerRawUnsafeUnm) features() []benchser.Feature {
	return []benchser.Feature{benchser.Raw, benchser.UnsafeUnm}
}
