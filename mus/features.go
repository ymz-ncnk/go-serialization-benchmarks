package mus

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var GeneralFeatures = []serializer.Feature{
	serializer.Binary,
	serializer.Codegen,
	serializer.Manual,
	serializer.UnsafeStr,
}
