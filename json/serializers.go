package json

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

var Serializers = []serializer.Serializer[general.Data]{Serializer{}}
