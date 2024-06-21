# go-serialization-benchmarks
In this benchmarks, all serializers use the same data. It is generated once and 
then used by everyone. Also each serializer is described with a set of features.
One serializer can have several benchmark results. For example, serializer mus
can have `mus` and `mus+unsafe` results. The last one indicates that the results
were obtained with the `unsafe` feature enabled.

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
- `fixbuf` - if a fixed buffer is used.

This list is not fixed and can be expanded.

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
already exist.Then you have implement 
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
|         NAME         | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|----------------------|------------------|-------|--------|------|-----------|
| mus+reuse+unsafe     |         17677365 | 57.03 |     58 |    0 |         0 |
| benc+reuse+unsafestr |         15267850 | 68.38 |     60 |    0 |         0 |
| mus+unsafe           |         15048591 | 81.52 |     58 |   64 |         1 |
| beebop200sc+reuse    |         13366515 | 84.58 |     61 |   48 |         1 |
| benc+unsafestr       |         12774222 | 92.93 |     60 |   64 |         1 |
| benc+reuse           |         12132196 | 93.81 |     60 |   48 |         1 |
| mus+reuse            |          9877123 | 112.9 |     59 |   48 |         1 |
| beebop200sc          |          9649923 | 114.9 |     61 |  112 |         2 |
| benc                 |         11313277 | 121.5 |     64 |  112 |         2 |
| mus                  |          7978897 | 150.2 |     59 |  112 |         2 |
| protobuf             |          2159122 | 514.7 |     72 |  271 |         4 |
| json                 |           417366 |  2712 |    150 |  600 |         9 |
| gob                  |            67326 | 17724 |    159 | 9407 |       233 |

, where `iterations count`, `ns/op`, `B/op`, `allocs/op` are standard 
`go test -bench=.` results and `B/size` - determines how many bytes were used on 
average by the serializer to encode `Data`.
    
# Features
- benc: `binary`, `manual`, `reuse`, `unsafestr`
- beebop200sc: `binary`, `codegen`, `reuse`
- gob: `binary`
- json: `reflect`, `text`
- mus: `binary`, `manual`, `reuse`, `unsafe`, `varint`
- protobuf: `binary`, `codegen`, `varint`
  