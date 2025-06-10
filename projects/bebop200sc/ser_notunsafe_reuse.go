package bebop200sc

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data "github.com/ymz-ncnk/go-serialization-benchmarks/data/bebop"
)

type SerializerNotUnsafeReuse struct {
	bs []byte
}

func (s SerializerNotUnsafeReuse) Marshal(d data.Data) (bs []byte, err error) {
	n := d.MarshalBebopTo(s.bs)
	bs = s.bs[:n]
	return
}

func (s SerializerNotUnsafeReuse) Unmarshal(bs []byte) (d data.Data, err error) {
	err = d.UnmarshalBebop(bs)
	return
}

func (s SerializerNotUnsafeReuse) Name() benchser.ResultName {
	return benchser.NewResultName(Bebop200sc, s.features()...)
}

func (s SerializerNotUnsafeReuse) Features() []benchser.Feature {
	return append(GeneralFeatures, s.features()...)
}

func (s SerializerNotUnsafeReuse) features() []benchser.Feature {
	return []benchser.Feature{benchser.NotUnsafe, benchser.Reuse}
}
