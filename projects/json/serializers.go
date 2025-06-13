package json

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/common"
)

var GeneralFeatures = []benchser.Feature{
	benchser.Reflect,
	benchser.Text,
	benchser.Int,
}

var Serializers = []benchser.Serializer[common.Data]{Serializer{}}
