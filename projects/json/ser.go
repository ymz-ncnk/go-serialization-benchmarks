package json

import (
	"encoding/json"

	"github.com/ymz-ncnk/go-serialization-benchmarks/benchser"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
)

type Serializer struct{}

func (s Serializer) Marshal(data general.Data) (bs []byte, err error) {
	return json.Marshal(data)
}

func (s Serializer) Unmarshal(bs []byte) (data general.Data, err error) {
	err = json.Unmarshal(bs, &data)
	return
}

func (s Serializer) Name() benchser.ResultName {
	return benchser.NewResultName("json")
}

func (s Serializer) Features() (features []benchser.Feature) {
	return GeneralFeatures
}
