package protobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type VTSerializerRawReuse struct {
	bs []byte
}

func (s VTSerializerRawReuse) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw, serializer.Reuse)
}

func (s VTSerializerRawReuse) Features() []serializer.Feature {
	return Features
}

func (s VTSerializerRawReuse) Marshal(data *DataRaw) (bs []byte, err error) {
	n, err := data.MarshalToSizedBufferVT(s.bs)
	if err != nil {
		return
	}
	bs = s.bs[len(s.bs)-n:]
	return
}

func (s VTSerializerRawReuse) Unmarshal(bs []byte) (data *DataRaw, err error) {
	data = &DataRaw{}
	err = data.UnmarshalVT(bs)
	return
}
