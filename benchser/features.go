package benchser

const (
	Manual    Feature = "manual"
	Codegen   Feature = "codegen"
	Reflect   Feature = "reflect"
	Text      Feature = "text"
	Binary    Feature = "binary"
	Reuse     Feature = "reuse"
	Unsafe    Feature = "unsafe"
	UnsafeStr Feature = "unsafestr"
	NotUnsafe Feature = "notunsafe"
	Varint    Feature = "varint"
	Raw       Feature = "raw"
	Int       Feature = "int"
	FixBuf    Feature = "fixbuf"
	Native    Feature = "native"
	UnsafeUnm Feature = "unsafeunm"
)

type Feature string
