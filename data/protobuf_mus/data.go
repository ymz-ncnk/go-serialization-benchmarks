package general

import "github.com/mus-format/ext-protobuf-go"

type Data struct {
	Str     string
	Bool    bool
	Int32   int32
	Float64 float64
	Time    ext.Timestamp
}
