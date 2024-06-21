package benc

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var Features = []serializer.Feature{
	serializer.Manual,
	serializer.Binary,
	serializer.Reuse,
	serializer.UnsafeStr,
	serializer.Raw,
}
