package protobuf_mus

import (
	"time"

	"github.com/mus-format/mus-go/varint"
	"google.golang.org/protobuf/encoding/protowire"
)

var (
	secondsFieldTag = protowire.EncodeTag(1, protowire.VarintType)
	nanosFieldTag   = protowire.EncodeTag(2, protowire.VarintType)
)

func NewTimestamp(tm time.Time) Timestamp {
	return Timestamp{Seconds: tm.Unix(), Nanos: int32(tm.Nanosecond())}
}

type Timestamp struct {
	Seconds int64
	Nanos   int32
}

func (t Timestamp) ToTime() time.Time {
	return time.Unix(int64(t.Seconds), int64(t.Nanos)).UTC()
}

func MarshalTimestamp(tm Timestamp, bs []byte) (n int) {
	if tm.Seconds != 0 {
		n += varint.MarshalUint64(secondsFieldTag, bs[n:])
		n += varint.MarshalPositiveInt64(tm.Seconds, bs[n:])
	}
	if tm.Nanos != 0 {
		n += varint.MarshalUint64(nanosFieldTag, bs[n:])
		n += varint.MarshalPositiveInt32(tm.Nanos, bs[n:])
	}
	return
}

func UnmarshalTimestamp(bs []byte) (tm Timestamp, n int,
	err error) {
	var (
		n1  int
		l   = len(bs)
		tag uint64
	)
	for n < l {
		tag, n1, err = varint.UnmarshalUint64(bs[n:])
		n += n1
		if err != nil {
			return
		}
		switch tag {
		case secondsFieldTag:
			tm.Seconds, n1, err = varint.UnmarshalPositiveInt64(bs[n:])
		case nanosFieldTag:
			tm.Nanos, n1, err = varint.UnmarshalPositiveInt32(bs[n:])
		}
		n += n1
		if err != nil {
			return
		}
	}
	return
}

func SizeTimestamp(tm Timestamp) (size int) {
	if tm.Seconds != 0 {
		size += varint.SizeUint64(secondsFieldTag)
		size += varint.SizePositiveInt64(tm.Seconds)
	}
	if tm.Nanos != 0 {
		size += varint.SizeUint64(nanosFieldTag)
		size += varint.SizePositiveInt32(tm.Nanos)
	}
	return
}

func EqualTimestamp(t1, t2 Timestamp) bool {
	return t1.Seconds == t2.Seconds && t1.Nanos == t2.Nanos
}
