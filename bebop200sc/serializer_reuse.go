package bebop200sc

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

type SerializerReuse struct {
	bs []byte
}

func (s SerializerReuse) Name() serializer.ResultName {
	return serializer.NewResultName(Beebop200sc, serializer.Reuse)
}

func (s SerializerReuse) Features() []serializer.Feature {
	return Features
}
func (s SerializerReuse) Marshal(data Data) (bs []byte, err error) {
	n := data.MarshalBebopTo(s.bs)
	bs = s.bs[:n]
	return
}

func (s SerializerReuse) Unmarshal(bs []byte) (data Data, err error) {
	err = data.UnmarshalBebop(bs)
	return
}
