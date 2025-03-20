package json

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
)

var GeneralFeatures = []benchser.Feature{
	benchser.Reflect,
	benchser.Text,
	benchser.Int,
}

var Serializers = []benchser.Serializer[general.Data]{Serializer{}}
