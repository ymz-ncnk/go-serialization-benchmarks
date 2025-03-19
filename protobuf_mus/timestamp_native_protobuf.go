package protobuf_mus

import (
	"github.com/mus-format/mus-go/varint"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	secondsFieldTag = protowire.EncodeTag(1, protowire.VarintType)
	nanosFieldTag   = protowire.EncodeTag(2, protowire.VarintType)
)

var TimestampNativeProtobuf = timestampNativeProtobuf{}

// timestampNativeProtobuf implements the mus.Serializer interface for timestamppb.Timestamp.
type timestampNativeProtobuf struct{}

func (s timestampNativeProtobuf) Marshal(tm *timestamppb.Timestamp, bs []byte) (n int) {
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

func (timestampNativeProtobuf) Unmarshal(bs []byte) (tm *timestamppb.Timestamp, n int,
	err error) {
	n, err = varint.PositiveInt.Skip(bs)
	if err != nil {
		return
	}
	var (
		n1  int
		l   = len(bs)
		tag uint64
	)
	tm = &timestamppb.Timestamp{}
	for {
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
		if n == l {
			return
		}
	}
}

func (s timestampNativeProtobuf) Size(tm *timestamppb.Timestamp) (size int) {
	size = s.size(tm)
	return size + varint.PositiveInt.Size(size)
}

func (s timestampNativeProtobuf) size(tm *timestamppb.Timestamp) (size int) {
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

// type timestampProtobuf struct{}

// func (s timestampProtobuf) Marshal(tm *timestamppb.Timestamp, bs []byte) (n int) {
// 	size := s.Size(tm)
// 	if size > 0 {
// 		n += varint.PositiveInt.Marshal(size, bs)
// 		if tm.Seconds != 0 {
// 			n += varint.Uint64.Marshal(secondsFieldTag, bs[n:])
// 			n += varint.PositiveInt64.Marshal(tm.Seconds, bs[n:])
// 		}
// 		if tm.Nanos != 0 {
// 			n += varint.Uint64.Marshal(nanosFieldTag, bs[n:])
// 			n += varint.PositiveInt32.Marshal(tm.Nanos, bs[n:])
// 		}
// 	}
// 	return
// }

// func (s timestampProtobuf) Unmarshal(bs []byte) (tm *timestamppb.Timestamp, n int,
// 	err error) {
// 	n, err = varint.PositiveInt.Skip(bs)
// 	if err != nil {
// 		return
// 	}
// 	var (
// 		n1  int
// 		l   = len(bs)
// 		tag uint64
// 	)
// 	tm = &timestamppb.Timestamp{}
// 	for n < l {
// 		tag, n1, err = varint.Uint64.Unmarshal(bs[n:])
// 		n += n1
// 		if err != nil {
// 			return
// 		}
// 		switch tag {
// 		case secondsFieldTag:
// 			tm.Seconds, n1, err = varint.PositiveInt64.Unmarshal(bs[n:])
// 		case nanosFieldTag:
// 			tm.Nanos, n1, err = varint.PositiveInt32.Unmarshal(bs[n:])
// 		}
// 		n += n1
// 		if err != nil {
// 			return
// 		}
// 	}
// 	return
// }

// func (s timestampProtobuf) Size(tm *timestamppb.Timestamp) (size int) {
// 	if tm.Seconds != 0 {
// 		size += varint.Uint64.Size(secondsFieldTag)
// 		size += varint.PositiveInt64.Size(tm.Seconds)
// 	}
// 	if tm.Nanos != 0 {
// 		size += varint.Uint64.Size(nanosFieldTag)
// 		size += varint.PositiveInt32.Size(tm.Nanos)
// 	}
// 	return
// }

func EqualNativeTimestamp(t1, t2 *timestamppb.Timestamp) bool {
	return t1.Seconds == t2.Seconds && t1.Nanos == t2.Nanos
}
