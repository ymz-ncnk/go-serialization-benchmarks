package gob

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var Serializers = []serializer.Serializer[serializer.Data]{NewSerializer()}
