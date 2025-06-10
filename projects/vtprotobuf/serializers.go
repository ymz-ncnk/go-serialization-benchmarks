package vtprotobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
)

const VTProtobuf = "vtprotobuf"

var GeneralFeatures = []benchser.Feature{
	benchser.Binary,
	benchser.Codegen,
}

var (
	SerializersRaw = []benchser.Serializer[*data_proto.DataRaw]{
		VTSerializerRaw{},
		VTSerializerRawReuse{bs: make([]byte, benchser.BufSize)},
		VTSerializerRawUnsafeUnm{},
		VTSerializerRawUnsafeUnmReuse{bs: make([]byte, benchser.BufSize)},
	}
	SerializersVarint = []benchser.Serializer[*data_proto.DataRawVarint]{
		VTSerializerVarint{},
		VTSerializerVarintReuse{bs: make([]byte, benchser.BufSize)},
		VTSerializerVarintUnsafeUnm{},
		VTSerializerVarintUnsafeUnmReuse{bs: make([]byte, benchser.BufSize)},
	}
)
