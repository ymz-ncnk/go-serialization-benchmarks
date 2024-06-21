package bebop200sc

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

const Beebop200sc = "beebop200sc"

type Serializer struct{}

func (s Serializer) Name() serializer.ResultName {
	return serializer.NewResultName(Beebop200sc)
}

func (s Serializer) Features() []serializer.Feature {
	return Features
}
func (s Serializer) Marshal(data Data) (bs []byte, err error) {
	bs = data.MarshalBebop()
	return
}

func (s Serializer) Unmarshal(bs []byte) (data Data, err error) {
	err = data.UnmarshalBebop(bs)
	return
}
