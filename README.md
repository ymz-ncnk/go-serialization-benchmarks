# go-serialization-benchmarks
In this benchmarks:
- All serializers use the same data, it is generated once and then used by 
  everyone.
- Each serializer is described with a set of features.
- One serializer can have several benchmark results. For example, mus can have 
  `mus` and `mus+unsafe` results. The last one indicates that the results were 
  obtained with the `unsafe` feature enabled.
- Unmarshalled data are compared to the original data.

# List of Features
Each feature describes a serializer:
- `reflect` - it uses reflection.
- `codegen` - it uses code generation.
- `manual` - there are only serialization primitives, so you have to use them 
  manually.
- `text` - it has text serialization format.
- `binary` -  it has binary serialization format.
- `varint` - it supports varint encoding.
- `int` - it supports `int` type.

Features that must be in the result name when used:
- `reuse` -  it supports buffer reuse.
- `unsafe` - it supports unsafe code.
- `unsafestr` - it supports only unsafe string serialization.
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
already exist. Then you have implement 
[serializer.Serializer\[serializer.Data\]](serializer/serializer.go) interface 
(if you use own `Data` make shure it implements `EqualTo(data Data) error` 
method, an example can be found in [serializer.Data](serializer/data.go)). Then 
you have to define
```go
var Serializers = []serializer.Serializer[serializer.Data]{Serializer{}}
```
variable. Note that it can contain several serializers that produce different 
results.

Doing this you can use:
- [serializer.NewResultName(...)](serializer/result_name.go) - which creates a 
  correct result name.
- [serializer.BufSize](serializer/serializer.go) - defines the recommended 
  buffer size for reuse.

Also, if you want to run benchmarks from your own project, there is the
[benchser.BenchmarkSerializer(...)](benchser/benchser.go) function.
  
# Benchmarks
|            NAME             | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|-----------------------------|------------------|-------|--------|------|-----------|
| mus+reuse+unsafe            |         17610296 | 58.02 |     58 |    0 |         0 |
| benc+reuse+unsafestr        |         15161148 | 69.75 |     60 |    0 |         0 |
| mus+unsafe                  |         14892512 | 81.61 |     58 |   64 |         1 |
| beebop200sc+notunsafe+reuse |         13270548 | 84.52 |     61 |   48 |         1 |
| benc+unsafestr              |         12957506 | 91.53 |     60 |   64 |         1 |
| benc+reuse                  |         11744248 |  97.1 |     60 |   48 |         1 |
| beebop200sc+notunsafe       |          9812706 | 112.3 |     61 |  112 |         2 |
| mus+notunsafe               |         12524446 | 112.4 |     58 |  112 |         2 |
| benc                        |         10632272 | 115.8 |     64 |  112 |         2 |
| mus+reuse                   |          9704544 | 116.4 |     59 |   48 |         1 |
| mus                         |          8126935 | 150.6 |     59 |  112 |         2 |
| protobuf                    |          2063260 | 543.8 |     72 |  271 |         4 |
| json                        |           414622 |  2705 |    150 |  600 |         9 |
| gob                         |            67324 | 17705 |    159 | 9406 |       233 |

, where `iterations count`, `ns/op`, `B/op`, `allocs/op` are standard 
`go test -bench=.` results and `B/size` - determines how many bytes were used on 
average by the serializer to encode `Data`.
    
# Features
- benc: `binary`, `manual`, `reuse`, `unsafestr`
- beebop200sc+notunsafe: `binary`, `codegen`, `notunsafe`, `reuse`
- gob: `binary`
- json: `reflect`, `text`
- mus: `binary`, `manual`, `reuse`, `unsafe`, `varint`
- protobuf: `binary`, `codegen`, `varint`
  