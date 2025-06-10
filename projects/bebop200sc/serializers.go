package bebop200sc

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data "github.com/ymz-ncnk/go-serialization-benchmarks/data/bebop"
)

const Bebop200sc = "bebop200sc"

var GeneralFeatures = []benchser.Feature{
	benchser.Codegen,
	benchser.Binary,
}

var Serializers = []benchser.Serializer[data.Data]{
	SerializerNotUnsafe{},
	SerializerNotUnsafeReuse{bs: make([]byte, benchser.BufSize)},
}
