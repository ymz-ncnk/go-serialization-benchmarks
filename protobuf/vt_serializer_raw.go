package protobuf

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

type VTSerializerRaw struct{}

func (s VTSerializerRaw) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw)
}

func (s VTSerializerRaw) Features() []serializer.Feature {
	return Features
}

func (s VTSerializerRaw) Marshal(data *DataRaw) (bs []byte, err error) {
	bs = make([]byte, data.SizeVT())
	_, err = data.MarshalToSizedBufferVT(bs)
	return
}

func (s VTSerializerRaw) Unmarshal(bs []byte) (data *DataRaw, err error) {
	data = &DataRaw{}
	err = data.UnmarshalVT(bs)
	return
}
