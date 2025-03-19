package vtprotobuf

import (
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type VTSerializerVarintReuse struct {
	bs []byte
}

func (s VTSerializerVarintReuse) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Varint,
		serializer.Reuse)
}

func (s VTSerializerVarintReuse) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Varint, serializer.Reuse)
}

func (s VTSerializerVarintReuse) Marshal(data *data_proto.DataRawVarint) (
	bs []byte, err error) {
	n, err := data.MarshalToSizedBufferVT(s.bs)
	if err != nil {
		return
	}
	bs = s.bs[len(s.bs)-n:]
	return
}

func (s VTSerializerVarintReuse) Unmarshal(bs []byte) (
	data *data_proto.DataRawVarint, err error) {
	data = &data_proto.DataRawVarint{}
	err = data.UnmarshalVT(bs)
	return
}
