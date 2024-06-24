package protobuf

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var Features = []serializer.Feature{
	serializer.Binary,
	serializer.Codegen,
	serializer.Raw,
	serializer.Varint,
}
