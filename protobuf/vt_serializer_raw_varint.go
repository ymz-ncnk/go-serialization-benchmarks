package protobuf

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

type VTSerializerRawVarint struct{}

func (s VTSerializerRawVarint) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw, serializer.Varint)
}

func (s VTSerializerRawVarint) Features() []serializer.Feature {
	return Features
}

func (s VTSerializerRawVarint) Marshal(data *DataRawVarint) (bs []byte, err error) {
	bs = make([]byte, data.SizeVT())
	_, err = data.MarshalToSizedBufferVT(bs)
	return
}

func (s VTSerializerRawVarint) Unmarshal(bs []byte) (data *DataRawVarint, err error) {
	data = &DataRawVarint{}
	err = data.UnmarshalVT(bs)
	return
}
