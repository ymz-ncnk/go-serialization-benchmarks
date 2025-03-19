package vtprotobuf

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var GeneralFeatures = []serializer.Feature{
	serializer.Binary,
	serializer.Codegen,
}

var Features = []serializer.Feature{
	serializer.Binary,
	serializer.Codegen,
	serializer.Raw,
	serializer.Varint,
	serializer.UnsafeUnm,
	serializer.Reuse,
}
