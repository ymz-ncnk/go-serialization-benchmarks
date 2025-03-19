package protobuf_mus

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var GeneralFeatures = []serializer.Feature{
	serializer.Binary,
	serializer.Manual,
}

var Features = []serializer.Feature{
	serializer.Manual,
	serializer.Binary,
	serializer.Reuse,
	serializer.NotUnsafe,
	serializer.UnsafeStr,
	serializer.Unsafe,
	serializer.Varint,
	serializer.Raw,
}
