package bebop200sc

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var Features = []serializer.Feature{
	serializer.Codegen,
	serializer.Binary,
	serializer.NotUnsafe,
	serializer.Reuse,
}
