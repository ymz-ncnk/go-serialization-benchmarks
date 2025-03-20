package gob

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
)

const Gob = "gob"

var GeneralFeatures = []benchser.Feature{
	benchser.Binary,
	benchser.Int,
}

var Serializers = []benchser.Serializer[general.Data]{NewSerializer()}
