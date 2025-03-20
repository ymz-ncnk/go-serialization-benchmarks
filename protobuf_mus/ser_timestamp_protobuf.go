package protobuf_mus

import (
	"github.com/mus-format/mus-go/varint"
	data_protobuf_mus "github.com/ymz-ncnk/go-serialization-benchmarks/data/protobuf_mus"
)

var TimestampProtobuf = timestampProtobuf{}

type timestampProtobuf struct{}

func (s timestampProtobuf) Marshal(tm data_protobuf_mus.Timestamp, bs []byte) (n int) {
	size := s.size(tm)
	n += varint.PositiveInt.Marshal(size, bs)
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

func (s timestampProtobuf) Unmarshal(bs []byte) (tm data_protobuf_mus.Timestamp, n int, err error) {
	size, n, err := varint.PositiveInt.Unmarshal(bs)
	if err != nil {
		return
	}
	var (
		tag uint64
		n1  int
	)
	for n < size {
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

func (s timestampProtobuf) Size(tm data_protobuf_mus.Timestamp) (size int) {
	size = s.size(tm)
	return size + varint.PositiveInt.Size(size)
}

func (s timestampProtobuf) size(tm data_protobuf_mus.Timestamp) (size int) {
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
