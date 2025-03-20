package general

type Data struct {
	Str     string
	Bool    bool
	Int32   int32
	Float64 float64
	Time    Timestamp
}

type Timestamp struct {
	Seconds int64
	Nanos   int32
}
