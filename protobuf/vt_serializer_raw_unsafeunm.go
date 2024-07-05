package protobuf

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

type VTSerializerRawUnsafeUnm struct{}

func (s VTSerializerRawUnsafeUnm) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw,
		serializer.UnsafeUnm)
}

func (s VTSerializerRawUnsafeUnm) Features() []serializer.Feature {
	return Features
}

func (s VTSerializerRawUnsafeUnm) Marshal(data *DataRaw) (bs []byte, err error) {
	bs = make([]byte, data.SizeVT())
	_, err = data.MarshalToSizedBufferVT(bs)
	return
}

func (s VTSerializerRawUnsafeUnm) Unmarshal(bs []byte) (data *DataRaw, err error) {
	data = &DataRaw{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}
