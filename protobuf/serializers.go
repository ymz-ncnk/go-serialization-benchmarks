package protobuf

import (
	"fmt"

	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const Protobuf = "protobuf"

var (
	SerializersRaw       = []serializer.Serializer[*DataRaw]{SerializerRaw{}}
	SerializersRawVarint = []serializer.Serializer[*DataRawVarint]{SerializerRawVarint{}}
)

func (d *DataRaw) EqualTo(ad *DataRaw) error {
	if d.Str != ad.Str {
		return fmt.Errorf("DataRaw.Str value %v != %v", ad.Str, d.Str)
	}
	if d.Bool != ad.Bool {
		return fmt.Errorf("DataRaw.Bool value %v != %v", ad.Bool, d.Bool)
	}
	if d.Int32 != ad.Int32 {
		return fmt.Errorf("DataRaw.Int32 value %v != %v", ad.Int32, d.Int32)
	}
	if d.Float64 != ad.Float64 {
		return fmt.Errorf("DataRaw.Float64 value %v != %v", ad.Float64, d.Float64)
	}
	if d.Time.Nanos != ad.Time.Nanos {
		return fmt.Errorf("DataRaw.Time value %v != %v", ad.Time, d.Time)
	}
	return nil
}

func (d *DataRawVarint) EqualTo(ad *DataRawVarint) error {
	if d.Str != ad.Str {
		return fmt.Errorf("DataRawVarint.Str value %v != %v", ad.Str, d.Str)
	}
	if d.Bool != ad.Bool {
		return fmt.Errorf("DataRawVarint.Bool value %v != %v", ad.Bool, d.Bool)
	}
	if d.Int32 != ad.Int32 {
		return fmt.Errorf("DataRawVarint.Int32 value %v != %v", ad.Int32, d.Int32)
	}
	if d.Float64 != ad.Float64 {
		return fmt.Errorf("DataRawVarint.Float64 value %v != %v", ad.Float64,
			d.Float64)
	}
	if d.Time.Nanos != ad.Time.Nanos {
		return fmt.Errorf("DataRawVarint.Time value %v != %v", ad.Time, d.Time)
	}
	return nil
}

func ToProtobufDataRaw(data serializer.Data) (d *DataRaw) {
	return &DataRaw{
		Str:     data.Str,
		Bool:    data.Bool,
		Int32:   data.Int32,
		Float64: data.Float64,
		Time:    timestamppb.New(data.Time),
	}
}

func ToProtobufDataRawVarint(data serializer.Data) (d *DataRawVarint) {
	return &DataRawVarint{
		Str:     data.Str,
		Bool:    data.Bool,
		Int32:   data.Int32,
		Float64: data.Float64,
		Time:    timestamppb.New(data.Time),
	}
}
