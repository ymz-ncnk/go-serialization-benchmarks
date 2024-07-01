package gob

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

const Gob = "gob"

var Serializers = []serializer.Serializer[serializer.Data]{NewSerializer()}
