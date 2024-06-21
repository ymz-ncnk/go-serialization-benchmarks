package bebop200sc

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

type SerializerNotUnsafeReuse struct {
	bs []byte
}

func (s SerializerNotUnsafeReuse) Name() serializer.ResultName {
	return serializer.NewResultName(Beebop200sc, serializer.NotUnsafe,
		serializer.Reuse)
}

func (s SerializerNotUnsafeReuse) Features() []serializer.Feature {
	return Features
}
func (s SerializerNotUnsafeReuse) Marshal(data Data) (bs []byte, err error) {
	n := data.MarshalBebopTo(s.bs)
	bs = s.bs[:n]
	return
}

func (s SerializerNotUnsafeReuse) Unmarshal(bs []byte) (data Data, err error) {
	err = data.UnmarshalBebop(bs)
	return
}
