package protobuf_mus

import (
	"time"

	"github.com/mus-format/mus-go/varint"
)

// var (
// 	secondsFieldTag = protowire.EncodeTag(1, protowire.VarintType)
// 	nanosFieldTag   = protowire.EncodeTag(2, protowire.VarintType)
// )

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

var TimestampProtobuf = timestampProtobuf{}

type timestampProtobuf struct{}

func (s timestampProtobuf) Marshal(tm Timestamp, bs []byte) (n int) {
	size := s.size(tm)
	n += varint.PositiveInt.Marshal(size, bs[n:])
	if tm.Seconds != 0 {
		n += varint.Uint64.Marshal(secondsFieldTag, bs[n:])
		n += varint.PositiveInt64.Marshal(tm.Seconds, bs[n:])
	}
	if tm.Nanos != 0 {
		n += varint.Uint64.Marshal(nanosFieldTag, bs[n:])
		n += varint.PositiveInt32.Marshal(tm.Nanos, bs[n:])
	}
	return
}

func (s timestampProtobuf) Unmarshal(bs []byte) (tm Timestamp, n int, err error) {
	n, err = varint.PositiveInt.Skip(bs)
	if err != nil {
		return
	}
	var (
		n1  int
		l   = len(bs)
		tag uint64
	)
	for n < l {
		tag, n1, err = varint.Uint64.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		switch tag {
		case secondsFieldTag:
			tm.Seconds, n1, err = varint.PositiveInt64.Unmarshal(bs[n:])
		case nanosFieldTag:
			tm.Nanos, n1, err = varint.PositiveInt32.Unmarshal(bs[n:])
		}
		n += n1
		if err != nil {
			return
		}
	}
	return
}

func (s timestampProtobuf) Size(tm Timestamp) (size int) {
	size = s.size(tm)
	return size + varint.PositiveInt.Size(size)
}

func (s timestampProtobuf) size(tm Timestamp) (size int) {
	if tm.Seconds != 0 {
		size += varint.Uint64.Size(secondsFieldTag)
		size += varint.PositiveInt64.Size(tm.Seconds)
	}
	if tm.Nanos != 0 {
		size += varint.Uint64.Size(nanosFieldTag)
		size += varint.PositiveInt32.Size(tm.Nanos)
	}
	return
}

func EqualTimestamp(t1, t2 Timestamp) bool {
	return t1.Seconds == t2.Seconds && t1.Nanos == t2.Nanos
}
