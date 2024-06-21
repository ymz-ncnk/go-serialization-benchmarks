package mus

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var Serializers = []serializer.Serializer[serializer.Data]{
	SerializerRaw{},
	SerializerRawVarint{},
	SerializerRawVarintReuse{make([]byte, serializer.BufSize)},
	SerializerUnsafe{},
	SerializerUnsafeReuse{make([]byte, serializer.BufSize)},
	SerializerNotUnsafe{},
}
