package protobuf

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

type VTSerializerRawVarintUnsafeUnm struct{}

func (s VTSerializerRawVarintUnsafeUnm) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw, serializer.Varint,
		serializer.UnsafeUnm)
}

func (s VTSerializerRawVarintUnsafeUnm) Features() []serializer.Feature {
	return Features
}

func (s VTSerializerRawVarintUnsafeUnm) Marshal(data *DataRawVarint) (bs []byte,
	err error) {
	bs = make([]byte, data.SizeVT())
	_, err = data.MarshalToSizedBufferVT(bs)
	return
}

func (s VTSerializerRawVarintUnsafeUnm) Unmarshal(bs []byte) (
	data *DataRawVarint, err error) {
	data = &DataRawVarint{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}
