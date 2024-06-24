package json

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var Features = []serializer.Feature{
	serializer.Reflect,
	serializer.Text,
	serializer.Int,
}
