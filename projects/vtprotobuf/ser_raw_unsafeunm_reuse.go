package vtprotobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
)

type VTSerializerRawUnsafeUnmReuse struct {
	bs []byte
}

func (s VTSerializerRawUnsafeUnmReuse) Marshal(data *data_proto.DataRaw) (
	bs []byte, err error) {
	n, err := data.MarshalToSizedBufferVT(s.bs)
	if err != nil {
		return
	}
	bs = s.bs[len(s.bs)-n:]
	return
}

func (s VTSerializerRawUnsafeUnmReuse) Unmarshal(bs []byte) (
	data *data_proto.DataRaw, err error) {
	data = &data_proto.DataRaw{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}

func (s VTSerializerRawUnsafeUnmReuse) Name() benchser.ResultName {
	return benchser.NewResultName(VTProtobuf, s.features()...)
}

func (s VTSerializerRawUnsafeUnmReuse) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s VTSerializerRawUnsafeUnmReuse) features() []benchser.Feature {
	return []benchser.Feature{benchser.Raw, benchser.UnsafeUnm,
		benchser.Reuse}
}
