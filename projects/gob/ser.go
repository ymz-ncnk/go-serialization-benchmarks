package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/common"
)

func NewSerializer() Serializer {
	gob.Register(common.Data{})
	return Serializer{}
}

type Serializer struct{}

func (s Serializer) Marshal(data common.Data) (bs []byte, err error) {
	var buf bytes.Buffer
	err = gob.NewEncoder(&buf).Encode(data)
	return buf.Bytes(), err
}

func (s Serializer) Unmarshal(bs []byte) (data common.Data, err error) {
	err = gob.NewDecoder(bytes.NewReader(bs)).Decode(&data)
	return
}

func (s Serializer) Name() benchser.ResultName {
	return benchser.NewResultName(Gob)
}

func (s Serializer) Features() []benchser.Feature {
	return GeneralFeatures
}
