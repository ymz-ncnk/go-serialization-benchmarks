package protobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type VTSerializerRawUnsafeUnmReuse struct {
	bs []byte
}

func (s VTSerializerRawUnsafeUnmReuse) Name() serializer.ResultName {
	return serializer.NewResultName(VTProtobuf, serializer.Raw,
		serializer.UnsafeUnm, serializer.Reuse)
}

func (s VTSerializerRawUnsafeUnmReuse) Features() []serializer.Feature {
	return Features
}

func (s VTSerializerRawUnsafeUnmReuse) Marshal(data *DataRaw) (bs []byte,
	err error) {
	n, err := data.MarshalToSizedBufferVT(s.bs)
	if err != nil {
		return
	}
	bs = s.bs[len(s.bs)-n:]
	return
}

func (s VTSerializerRawUnsafeUnmReuse) Unmarshal(bs []byte) (data *DataRaw,
	err error) {
	data = &DataRaw{}
	err = data.UnmarshalVTUnsafe(bs)
	return
}
