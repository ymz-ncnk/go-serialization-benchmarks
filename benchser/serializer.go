package benchser

const BufSize = 1024

type Serializer[T any] interface {
	SerializerDesc
	Marshal(data T) (bs []byte, err error)
	Unmarshal(bs []byte) (data T, err error)
}

type SerializerDesc interface {
	Name() ResultName
	Features() []Feature
}
