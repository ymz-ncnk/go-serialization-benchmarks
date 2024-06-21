package json

import (
	"encoding/json"

	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type Serializer struct{}

func (s Serializer) Name() serializer.ResultName {
	return serializer.NewResultName("json")
}

func (s Serializer) Features() (features []serializer.Feature) {
	return Features
}

func (s Serializer) Marshal(data serializer.Data) (bs []byte, err error) {
	return json.Marshal(data)
}

func (s Serializer) Unmarshal(bs []byte) (data serializer.Data, err error) {
	err = json.Unmarshal(bs, &data)
	return
}
