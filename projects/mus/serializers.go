package mus

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
)

const MUS = "mus"

var GeneralFeatures = []benchser.Feature{
	benchser.Binary,
	benchser.Codegen,
	benchser.Manual,
	benchser.UnsafeStr,
}

var Serializers = []benchser.Serializer[general.Data]{
	SerializerRaw{},
	SerializerRawReuse{make([]byte, benchser.BufSize)},
	SerializerRawVarint{},
	SerializerRawVarintReuse{make([]byte, benchser.BufSize)},
	SerializerUnsafe{},
	SerializerUnsafeReuse{make([]byte, benchser.BufSize)},
	SerializerNotUnsafe{},
	SerializerNotUnsafeReuse{make([]byte, benchser.BufSize)},
}
