package vtprotobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const VTProtobuf = "vtprotobuf"

var (
	SerializersRaw = []serializer.Serializer[*data_proto.DataRaw]{
		VTSerializerRaw{},
		VTSerializerRawReuse{bs: make([]byte, serializer.BufSize)},
		VTSerializerRawUnsafeUnm{},
		VTSerializerRawUnsafeUnmReuse{bs: make([]byte, serializer.BufSize)},
	}
	SerializersVarint = []serializer.Serializer[*data_proto.DataRawVarint]{
		VTSerializerVarint{},
		VTSerializerVarintReuse{bs: make([]byte, serializer.BufSize)},
		VTSerializerVarintUnsafeUnm{},
		VTSerializerVarintUnsafeUnmReuse{bs: make([]byte, serializer.BufSize)},
	}
)

func ToProtobufDataRaw(data general.Data) (d *data_proto.DataRaw) {
	return &data_proto.DataRaw{
		Str:     data.Str,
		Bool:    data.Bool,
		Int32:   data.Int32,
		Float64: data.Float64,
		Time:    timestamppb.New(data.Time),
	}
}

func ToProtobufDataRawVarint(data general.Data) (d *data_proto.DataRawVarint) {
	return &data_proto.DataRawVarint{
		Str:     data.Str,
		Bool:    data.Bool,
		Int32:   data.Int32,
		Float64: data.Float64,
		Time:    timestamppb.New(data.Time),
	}
}
