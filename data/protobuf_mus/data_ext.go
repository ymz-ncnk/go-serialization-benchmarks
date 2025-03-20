package general

import (
	"fmt"

	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
)

func (d Data) EqualTo(ad Data) (err error) {
	if d.Str != ad.Str {
		return fmt.Errorf("Data.Str value %v != %v", ad.Str, d.Str)
	}
	if d.Bool != ad.Bool {
		return fmt.Errorf("Data.Bool value %v != %v", ad.Bool, d.Bool)
	}
	if d.Int32 != ad.Int32 {
		return fmt.Errorf("Data.Int32 value %v != %v", ad.Int32, d.Int32)
	}
	if d.Float64 != ad.Float64 {
		return fmt.Errorf("Data.Float64 value %v != %v", ad.Float64, d.Float64)
	}
	if d.Time.Seconds != ad.Time.Seconds && d.Time.Nanos != ad.Time.Nanos {
		return fmt.Errorf("Data.Time value %v != %v", ad.Time, d.Time)
	}
	return nil
}

func ToProtobufMUSData(data general.Data) (d Data) {
	return Data{
		Str:     data.Str,
		Bool:    data.Bool,
		Int32:   data.Int32,
		Float64: data.Float64,
		Time: Timestamp{Seconds: data.Time.Unix(),
			Nanos: int32(data.Time.Nanosecond())},
	}
}
