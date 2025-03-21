// Code generated by bebopc-go; DO NOT EDIT.

package bebop

import (
	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
	"io"
	"time"
)

var _ bebop.Record = &Data{}

type Data struct {
	Str string
	Bool bool
	Int32 int32
	Float64 float64
	Time time.Time
}

func (bbp Data) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.Str)))
	copy(buf[at+4:at+4+len(bbp.Str)], []byte(bbp.Str))
	at += 4 + len(bbp.Str)
	iohelp.WriteBoolBytes(buf[at:], bbp.Bool)
	at += 1
	iohelp.WriteInt32Bytes(buf[at:], bbp.Int32)
	at += 4
	iohelp.WriteFloat64Bytes(buf[at:], bbp.Float64)
	at += 8
	if (bbp.Time).IsZero() {
		iohelp.WriteInt64Bytes(buf[at:], 0)
	} else {
		iohelp.WriteInt64Bytes(buf[at:], ((bbp.Time).UnixNano() / 100))
	}
	at += 8
	return at
}

func (bbp *Data) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	bbp.Str, err = iohelp.ReadStringBytes(buf[at:])
	if err != nil {
		return err
	}
	at += 4 + len(bbp.Str)
	if len(buf[at:]) < 1 {
		return io.ErrUnexpectedEOF
	}
	bbp.Bool = iohelp.ReadBoolBytes(buf[at:])
	at += 1
	if len(buf[at:]) < 4 {
		return io.ErrUnexpectedEOF
	}
	bbp.Int32 = iohelp.ReadInt32Bytes(buf[at:])
	at += 4
	if len(buf[at:]) < 8 {
		return io.ErrUnexpectedEOF
	}
	bbp.Float64 = iohelp.ReadFloat64Bytes(buf[at:])
	at += 8
	if len(buf[at:]) < 8 {
		return io.ErrUnexpectedEOF
	}
	bbp.Time = iohelp.ReadDateBytes(buf[at:])
	at += 8
	return nil
}

func (bbp Data) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(len(bbp.Str)))
	w.Write([]byte(bbp.Str))
	iohelp.WriteBool(w, bbp.Bool)
	iohelp.WriteInt32(w, bbp.Int32)
	iohelp.WriteFloat64(w, bbp.Float64)
	if (bbp.Time).IsZero() {
		iohelp.WriteInt64(w, 0)
	} else {
		iohelp.WriteInt64(w, ((bbp.Time).UnixNano() / 100))
	}
	return w.Err
}

func (bbp *Data) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.Str = iohelp.ReadString(r)
	bbp.Bool = iohelp.ReadBool(r)
	bbp.Int32 = iohelp.ReadInt32(r)
	bbp.Float64 = iohelp.ReadFloat64(r)
	bbp.Time = iohelp.ReadDate(r)
	return r.Err
}

func (bbp Data) Size() int {
	bodyLen := 0
	bodyLen += 4 + len(bbp.Str)
	bodyLen += 1
	bodyLen += 4
	bodyLen += 8
	bodyLen += 8
	return bodyLen
}

func (bbp Data) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeData(r *iohelp.ErrorReader) (Data, error) {
	v := Data{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeDataFromBytes(buf []byte) (Data, error) {
	v := Data{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

