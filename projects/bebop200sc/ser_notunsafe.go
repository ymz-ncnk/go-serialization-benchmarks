package bebop200sc

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
)

type SerializerNotUnsafe struct{}

func (s SerializerNotUnsafe) Marshal(data Data) (bs []byte, err error) {
	bs = data.MarshalBebop()
	return
}

func (s SerializerNotUnsafe) Unmarshal(bs []byte) (data Data, err error) {
	err = data.UnmarshalBebop(bs)
	return
}

func (s SerializerNotUnsafe) Name() benchser.ResultName {
	return benchser.NewResultName(Bebop200sc, s.features()...)
}

func (s SerializerNotUnsafe) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s SerializerNotUnsafe) features() []benchser.Feature {
	return []benchser.Feature{benchser.NotUnsafe}
}
