# go-serialization-benchmarks
In this benchmarks:
- All serializers use the same data, it is generated once and then used by 
  everyone.
- Each serializer is described with a set of features.
- One serializer can have several benchmark results. For example, MUS can have 
  `mus+raw` and `mus+unsafe` results. The first one indicates that the results
  were obtained with the `raw` feature enabled, the last one - that `unsafe` 
  feature was used.
- Unmarshalled data are compared to the original data.  
  
# Benchmarks  
## Fastest Safe
|     NAME     | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|--------------|------------------|-------|--------|------|-----------|
| mus          |         11795019 | 95.47 |     58 |   48 |         1 |
| bebop200sc   |         11282048 | 97.45 |     61 |   48 |         1 |
| benc         |         10793390 | 103.8 |     60 |   48 |         1 |
| vtprotobuf   |          7774521 | 169.4 |     69 |  192 |         3 |
| protobuf_mus |          6391887 | 170.4 |     69 |   48 |         1 |
| protobuf     |          2505753 | 452.1 |     69 |  271 |         4 |
| json         |           453517 |  2621 |    150 |  600 |         9 |
| gob          |            84352 | 15046 |    159 | 9493 |       195 |
  
## Fastest Unsafe
|     NAME     | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|--------------|------------------|-------|--------|------|-----------|
| mus          |         15055489 | 70.29 |     58 |    0 |         0 |
| benc         |         13512625 |  78.5 |     60 |    0 |         0 |
| protobuf_mus |          7528129 | 141.8 |     69 |    0 |         0 |
| vtprotobuf   |          7935742 |   143 |     70 |  144 |         2 |
  
## All
|                NAME                 | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|-------------------------------------|------------------|-------|--------|------|-----------|
| mus+reuse+unsafe                    |         15055489 | 70.29 |     58 |    0 |         0 |
| benc+raw+reuse+unsafestr            |         13512625 |  78.5 |     60 |    0 |         0 |
| mus+unsafe                          |         12787286 | 94.03 |     58 |   64 |         1 |
| mus+notunsafe+reuse                 |         11795019 | 95.47 |     58 |   48 |         1 |
| bebop200sc+notunsafe+reuse          |         11282048 | 97.45 |     61 |   48 |         1 |
| benc+raw+unsafestr                  |         12117286 | 97.97 |     60 |   64 |         1 |
| benc+raw+reuse                      |         10793390 | 103.8 |     60 |   48 |         1 |
| mus+raw+reuse                       |         10430548 |   109 |     58 |   48 |         1 |
| mus+notunsafe                       |         11046586 | 121.9 |     58 |  112 |         2 |
| mus+reuse+varint                    |          9054440 | 122.6 |     58 |   48 |         1 |
| mus+raw                             |         10465509 | 124.1 |     58 |  112 |         2 |
| bebop200sc+notunsafe                |          8968006 |   129 |     61 |  112 |         2 |
| benc+raw                            |          9251139 | 135.5 |     64 |  112 |         2 |
| protobuf_mus+reuse+unsafe           |          7528129 | 141.8 |     69 |    0 |         0 |
| vtprotobuf+reuse+unsafeunm+varint   |          7935742 |   143 |     70 |  144 |         2 |
| vtprotobuf+raw+reuse+unsafeunm      |          8473729 | 143.5 |     69 |  144 |         2 |
| mus+varint                          |          8519427 | 150.6 |     58 |  112 |         2 |
| protobuf_mus+native+reuse+unsafe    |          7905327 | 167.1 |     69 |  144 |         2 |
| vtprotobuf+raw+reuse                |          7774521 | 169.4 |     69 |  192 |         3 |
| protobuf_mus+raw+reuse              |          6391887 | 170.4 |     69 |   48 |         1 |
| protobuf_mus+native+reuse+unsafe#01 |          7853995 | 170.7 |     69 |  144 |         2 |
| vtprotobuf+reuse+varint             |          7155690 | 180.4 |     70 |  192 |         3 |
| protobuf_mus+reuse+varint           |          5802302 | 191.5 |     70 |   48 |         1 |
| protobuf_mus+unsafe                 |          6207885 | 195.8 |     69 |   79 |         1 |
| vtprotobuf+unsafeunm+varint         |          7210292 | 196.3 |     70 |  223 |         3 |
| vtprotobuf+raw+unsafeunm            |          7316026 | 197.1 |     69 |  223 |         3 |
| vtprotobuf+raw                      |          6619782 |   215 |     69 |  271 |         4 |
| vtprotobuf+varint                   |          6376152 | 232.2 |     70 |  271 |         4 |
| protobuf_mus+raw                    |          5035914 | 240.2 |     69 |  127 |         2 |
| protobuf_mus+varint                 |          4998715 | 253.9 |     70 |  127 |         2 |
| protobuf+raw                        |          2505753 | 452.1 |     69 |  271 |         4 |
| protobuf+varint                     |          2798874 | 454.9 |     70 |  271 |         4 |
| json                                |           453517 |  2621 |    150 |  600 |         9 |
| gob                                 |            84352 | 15046 |    159 | 9493 |       195 |

, where `iterations count`, `ns/op`, `B/op`, `allocs/op` are standard 
`go test -bench=.` output and `B/size` - determines how many bytes were used on 
average by the serializer to encode `Data`.  
  
# Features
- json: int,reflect,text
- gob: binary,int
- mus: binary,codegen,manual,varint,unsafestr,raw,reuse,unsafe,notunsafe
- benc: manual,binary,reuse,unsafestr,raw
- protobuf: binary,codegen,raw,varint
- vtprotobuf: binary,codegen,raw,reuse,unsafeunm,varint
  
# List of Features
Each feature describes a serializer:
- `reflect` - it uses reflection.
- `codegen` - it uses code generation.
- `manual` - there are only serialization primitives, so you have to use them 
  manually.
- `text` - it has text serialization format.
- `binary` -  it has binary serialization format.
- `varint` - it supports varint encoding.
- `raw` - it supports raw encoding.
- `int` - it supports `int` type.
- `native` - when a format is implemented with a set of primitives and native 
  data is used.

Features that must be in the result name when used:
- `reuse` -  it supports buffer reuse.
- `unsafe` - it supports unsafe code.
- `unsafestr` - it supports only unsafe string serialization.
- `unsafeunm` - it supports only unsafe unmarshalling.
- `notunsafe` - it uses the unsafe code for all types except `string` and copies
  data on unmarshal.
- `fixbuf` - if a fixed buffer is used.

This list can be expanded.

# Data
Randomly generated data has the following form:
```go
type Data struct {
  Str     string
  Bool    bool
  Int32   int32
  Float64 float64
  Time    time.Time
}
```
It does not have an `int` type because many serializers do not support it.

# Run Benchmarks
From the `benchmarks/` folder:
```bash
go test -bench=.
```
We can also filter the serializers by features, for example:
```bash
go test -bench=. -- -f binary,manual
```
In this case we will see the results of serializers that have both `binary`
and `manual` features.
To run benchmarks for one particular case just name it, for example:
```bash
go test -bench=BenchmarkSerializers/mus
```
or
```bash
go test -bench=/.+reuse
```
to see the results obtained using the `reuse` feature.

# Generate README.md
From the `benchmarks/` folder:
```bash
go generate ./...
```

## Recomendation
When creating `README.md` on a laptop, please make sure that it is connected to 
a charger and the fan is running at full speed.

# Contribution
First of all, you need to create a new package for your serializer if it doesn't
already exist. Then:
1. Implement [serializer.Serializer\[serializer.Data\]](serializer/serializer.go) 
   interface. Doing this you can use:
   - [serializer.NewResultName(...)](serializer/result_name.go) - which creates 
     a correct result name.
   - [serializer.BufSize](serializer/serializer.go) - defines the recommended 
     buffer size for reuse.
2. If you use own `Data` make shure it implements `EqualTo(data Data) error` 
   method, also add `func ToYourData(data serializer.Data) (d Data)`
   function (an example can be found in [bbebop200sc/serializers.go](bebop200sc/serializers.go)).
3. Define
   ```go
   var Serializers = []serializer.Serializer[serializer.Data]{Serializer{}}
   ```
   variable. Note that it can contain several serializers that produce different
   results.
4. Create PR.

If you want to run benchmarks from your own project, there is the
[benchser.BenchmarkSerializer(...)](benchser/benchser.go) function.