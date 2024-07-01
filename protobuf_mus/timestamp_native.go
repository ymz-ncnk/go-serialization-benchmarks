package protobuf_mus

import (
	"github.com/mus-format/mus-go/varint"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MarshalTimestampNative(tm *timestamppb.Timestamp, bs []byte) (n int) {
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

func UnmarshalTimestampNative(bs []byte) (tm *timestamppb.Timestamp, n int,
	err error) {
	var (
		n1  int
		l   = len(bs)
		tag uint64
	)
	tm = &timestamppb.Timestamp{}
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

func SizeTimestampNative(tm *timestamppb.Timestamp) (size int) {
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

func EqualTimestampNative(t1, t2 *timestamppb.Timestamp) bool {
	return t1.Seconds == t2.Seconds && t1.Nanos == t2.Nanos
}
