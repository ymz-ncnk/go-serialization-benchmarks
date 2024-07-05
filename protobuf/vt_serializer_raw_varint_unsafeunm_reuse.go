package protobuf

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

type VTSerializerRawVarintUnsafeUnmReuse struct {
	bs []byte
}

func (s VTSerializerRawVarintUnsafeUnmReuse) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw, serializer.Varint,
		serializer.UnsafeUnm, serializer.Reuse)
}

func (s VTSerializerRawVarintUnsafeUnmReuse) Features() []serializer.Feature {
	return Features
}

func (s VTSerializerRawVarintUnsafeUnmReuse) Marshal(data *DataRawVarint) (
	bs []byte, err error) {
	n, err := data.MarshalToSizedBufferVT(s.bs)
	if err != nil {
		return
	}
	bs = s.bs[len(s.bs)-n:]
	return
}

func (s VTSerializerRawVarintUnsafeUnmReuse) Unmarshal(bs []byte) (
	data *DataRawVarint, err error) {
	data = &DataRawVarint{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}
