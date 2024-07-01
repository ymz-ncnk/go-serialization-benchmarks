package bebop200sc

import (
	"fmt"
	"time"

	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

const Bebop200sc = "bebop200sc"

var Serializers = []serializer.Serializer[Data]{
	SerializerNotUnsafe{},
	SerializerNotUnsafeReuse{bs: make([]byte, serializer.BufSize)},
}

func (d Data) EqualTo(ad Data) error {
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
	t := time.Unix(0, (d.Time.UnixNano()/100)*100).UTC()
	if !t.Equal(ad.Time) {
		return fmt.Errorf("Data.Time value %v != %v", ad.Time, d.Time)
	}
	return nil
}

func ToBebop200scData(data serializer.Data) (d Data) {
	return Data{
		Str:     data.Str,
		Bool:    data.Bool,
		Int32:   data.Int32,
		Float64: data.Float64,
		Time:    data.Time,
	}
}
