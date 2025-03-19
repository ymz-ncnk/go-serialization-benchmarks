package vtprotobuf

import (
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type VTSerializerRawReuse struct {
	bs []byte
}

func (s VTSerializerRawReuse) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw, serializer.Reuse)
}

func (s VTSerializerRawReuse) Features() []serializer.Feature {
	return append(GeneralFeatures, serializer.Raw, serializer.Reuse)
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
