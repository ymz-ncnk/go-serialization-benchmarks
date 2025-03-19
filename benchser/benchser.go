package benchser

import (
	"slices"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/ymz-ncnk/go-serialization-benchmarks/data/general"
	"github.com/ymz-ncnk/go-serialization-benchmarks/serializer"
)

type Data[T any] interface {
	EqualTo(T) error
}

func BenchmarkSerializer[T Data[T]](s serializer.Serializer[T],
	wantFeatures []serializer.Feature,
	data []T,
	b *testing.B,
) {
	l := len(wantFeatures)
	if l == 0 || (l > 0 && hasFeatures(s, wantFeatures)) {
		b.Run(string(s.Name()), func(b *testing.B) {
			if b.N > len(data) {
				b.Fatal("too little data, make the benchmarks.DataCount constant larger")
			}
			b.ResetTimer()
			doMarshalUnmarshalCheck(s, data, b)
			b.StopTimer()
			b.ReportAllocs()
			ReportBSizeMetric(s, data, b)
		})
	}
}

func ReportBSizeMetric[T any](s serializer.Serializer[T], d []T, b *testing.B) {
	var (
		err      error
		bs       []byte
		dataSize int
	)
	for i := 0; i < b.N; i++ {
		bs, err = s.Marshal(d[i])
		if err != nil {
			b.Fatal(err)
		}
		dataSize += len(bs)
	}
	b.ReportMetric(float64(dataSize/b.N), "B/size")
}

func GenerateData(count int) (d []general.Data, err error) {
	d = make([]general.Data, count)
	for i := range d {
		d[i] = general.Data{
			Str:     gofakeit.UUID(),
			Bool:    gofakeit.Bool(),
			Int32:   gofakeit.Int32(),
			Float64: gofakeit.Float64(),
			Time:    gofakeit.Date(),
		}
	}
	return
}

func doMarshalUnmarshalCheck[T Data[T]](s serializer.Serializer[T], d []T,
	b *testing.B) {
	var (
		err  error
		bs   []byte
		data T
	)
	for i := 0; i < b.N; i++ {
		bs, err = s.Marshal(d[i])
		if err != nil {
			b.Fatal(err)
		}
		data, err = s.Unmarshal(bs)
		if err != nil {
			b.Fatal(err)
		}
		if err := d[i].EqualTo(data); err != nil {
			b.Fatal(err)
		}
	}
}

func hasFeatures(s serializer.SerializerDesc, features []serializer.Feature) bool {
	for i := range features {
		if !slices.Contains(s.Features(), features[i]) {
			return false
		}
	}
	return true
}
