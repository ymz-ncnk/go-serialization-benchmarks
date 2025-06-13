package benc

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/common"
)

const Benc = "benc"

var GeneralFeatures = []benchser.Feature{
	benchser.Binary,
	benchser.Codegen,
	benchser.Manual,
}

var Serializers = []benchser.Serializer[common.Data]{
	SerializerRaw{},
	SerializerRawReuse{make([]byte, benchser.BufSize)},
	SerializerRawUnsafeStr{},
	SerializerRawUnsafeStrReuse{make([]byte, benchser.BufSize)},
}
