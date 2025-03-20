package vtprotobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
)

type VTSerializerVarintUnsafeUnm struct{}

func (s VTSerializerVarintUnsafeUnm) Marshal(data *data_proto.DataRawVarint) (
	bs []byte, err error) {
	bs = make([]byte, data.SizeVT())
	_, err = data.MarshalToSizedBufferVT(bs)
	return
}

func (s VTSerializerVarintUnsafeUnm) Unmarshal(bs []byte) (
	data *data_proto.DataRawVarint, err error) {
	data = &data_proto.DataRawVarint{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}

func (s VTSerializerVarintUnsafeUnm) Name() benchser.ResultName {
	return benchser.NewResultName(VTProtobuf, s.features()...)
}

func (s VTSerializerVarintUnsafeUnm) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s VTSerializerVarintUnsafeUnm) features() []benchser.Feature {
	return []benchser.Feature{benchser.Varint, benchser.UnsafeUnm}
}
