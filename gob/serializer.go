package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type Serializer struct{}

func (s Serializer) Name() serializer.ResultName {
	return serializer.NewResultName("gob")
}

func (s Serializer) Features() []serializer.Feature {
	return Features
}

func (s Serializer) Marshal(data serializer.Data) (bs []byte, err error) {
	var buf bytes.Buffer
	err = gob.NewEncoder(&buf).Encode(data)
	return buf.Bytes(), err
}

func (s Serializer) Unmarshal(bs []byte) (data serializer.Data, err error) {
	err = gob.NewDecoder(bytes.NewReader(bs)).Decode(&data)
	return
}

func NewSerializer() Serializer {
	gob.Register(serializer.Data{})
	return Serializer{}
}
