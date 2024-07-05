package protobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type VTSerializerRawVarintReuse struct {
	bs []byte
}

func (s VTSerializerRawVarintReuse) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw, serializer.Varint,
		serializer.Reuse)
}

func (s VTSerializerRawVarintReuse) Features() []serializer.Feature {
	return Features
}

func (s VTSerializerRawVarintReuse) Marshal(data *DataRawVarint) (bs []byte,
	err error) {
	n, err := data.MarshalToSizedBufferVT(s.bs)
	if err != nil {
		return
	}
	bs = s.bs[len(s.bs)-n:]
	return
}

func (s VTSerializerRawVarintReuse) Unmarshal(bs []byte) (data *DataRawVarint, err error) {
	data = &DataRawVarint{}
	err = data.UnmarshalVT(bs)
	return
}
