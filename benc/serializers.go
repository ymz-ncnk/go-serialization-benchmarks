package benc

import "github.com/ymz-ncnk/go-serialization-benchmarks/serializer"

var Serializers = []serializer.Serializer[serializer.Data]{
	SerializerRaw{},
	SerializerRawReuse{make([]byte, serializer.BufSize)},
	SerializerRawUnsafeStr{},
	SerializerRawUnsafeStrReuse{make([]byte, serializer.BufSize)},
}
