package protobuf

import (
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	data_proto "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf"

	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const Protobuf = "protobuf"

var (
	SerializersRaw = []serializer.Serializer[*data_proto.DataRaw]{
		SerializerRaw{},
	}
	SerializersVarint = []serializer.Serializer[*data_proto.DataRawVarint]{
		SerializerVarint{},
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
