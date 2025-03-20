package vtprotobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
)

type VTSerializerVarintUnsafeUnmReuse struct {
	bs []byte
}

func (s VTSerializerVarintUnsafeUnmReuse) Marshal(data *data_proto.DataRawVarint) (
	bs []byte, err error) {
	n, err := data.MarshalToSizedBufferVT(s.bs)
	if err != nil {
		return
	}
	bs = s.bs[len(s.bs)-n:]
	return
}

func (s VTSerializerVarintUnsafeUnmReuse) Unmarshal(bs []byte) (
	data *data_proto.DataRawVarint, err error) {
	data = &data_proto.DataRawVarint{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}

func (s VTSerializerVarintUnsafeUnmReuse) Name() benchser.ResultName {
	return benchser.NewResultName(VTProtobuf, s.features()...)
}

func (s VTSerializerVarintUnsafeUnmReuse) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s VTSerializerVarintUnsafeUnmReuse) features() []benchser.Feature {
	return []benchser.Feature{benchser.Varint, benchser.UnsafeUnm,
		benchser.Reuse}
}
