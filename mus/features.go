package mus

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var Features = []serializer.Feature{
	serializer.Manual,
	serializer.Binary,
	serializer.Reuse,
	serializer.NotUnsafe,
	serializer.UnsafeStr,
	serializer.Unsafe,
	serializer.Varint,
	serializer.Raw,
	serializer.Int,
}
