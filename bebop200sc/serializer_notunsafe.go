package bebop200sc

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

type SerializerNotUnsafe struct{}

func (s SerializerNotUnsafe) Name() serializer.ResultName {
	return serializer.NewResultName(Bebop200sc, serializer.NotUnsafe)
}

func (s SerializerNotUnsafe) Features() []serializer.Feature {
	return Features
}
func (s SerializerNotUnsafe) Marshal(data Data) (bs []byte, err error) {
	bs = data.MarshalBebop()
	return
}

func (s SerializerNotUnsafe) Unmarshal(bs []byte) (data Data, err error) {
	err = data.UnmarshalBebop(bs)
	return
}
