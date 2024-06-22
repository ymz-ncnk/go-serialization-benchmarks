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
- `raw` - it supports raw encoding.
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
    
## Fastest Safe
|    NAME     | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|-------------|------------------|-------|--------|------|-----------|
| mus         |         14043190 | 79.82 |     58 |   48 |         1 |
| beebop200sc |         13137957 |  86.1 |     61 |   48 |         1 |
| benc        |         11942839 | 95.74 |     60 |   48 |         1 |
| protobuf    |          2225994 | 499.1 |     72 |  271 |         4 |
| json        |           408625 |  2743 |    150 |  600 |         9 |
| gob         |            67514 | 17837 |    159 | 9408 |       233 |
  
## Fastest Unsafe
| NAME | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|------|------------------|-------|--------|------|-----------|
| mus  |         18238142 | 57.49 |     58 |    0 |         0 |
| benc |         15114085 | 69.05 |     60 |    0 |         0 |
  
## All
|            NAME             | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|-----------------------------|------------------|-------|--------|------|-----------|
| mus+reuse+unsafe            |         18238142 | 57.49 |     58 |    0 |         0 |
| benc+raw+reuse+unsafestr    |         15114085 | 69.05 |     60 |    0 |         0 |
| mus+notunsafe+reuse         |         14043190 | 79.82 |     58 |   48 |         1 |
| mus+unsafe                  |         14791004 | 82.39 |     58 |   64 |         1 |
| beebop200sc+notunsafe+reuse |         13137957 |  86.1 |     61 |   48 |         1 |
| benc+raw+unsafestr          |         12219870 | 91.02 |     60 |   64 |         1 |
| mus+raw+reuse               |         12096433 | 92.55 |     58 |   48 |         1 |
| benc+raw+reuse              |         11942839 | 95.74 |     60 |   48 |         1 |
| mus+notunsafe               |         12510670 | 112.1 |     58 |  112 |         2 |
| mus+raw+reuse+varint        |          9862287 |   113 |     59 |   48 |         1 |
| beebop200sc+notunsafe       |          9998587 | 113.6 |     61 |  112 |         2 |
| benc+raw                    |         10507029 | 122.9 |     64 |  112 |         2 |
| mus+raw                     |          9550530 |   127 |     58 |  112 |         2 |
| mus+raw+varint              |          8192626 | 153.9 |     59 |  112 |         2 |
| protobuf                    |          2225994 | 499.1 |     72 |  271 |         4 |
| json                        |           408625 |  2743 |    150 |  600 |         9 |
| gob                         |            67514 | 17837 |    159 | 9408 |       233 |

, where `iterations count`, `ns/op`, `B/op`, `allocs/op` are standard 
`go test -bench=.` results and `B/size` - determines how many bytes were used on 
average by the serializer to encode `Data`.  
  
# Features
- benc: `binary`, `manual`, `raw`, `reuse`, `unsafestr`
- beebop200sc: `binary`, `codegen`, `notunsafe`, `reuse`
- gob: `binary`
- json: `reflect`, `text`
- mus: `binary`, `manual`, `raw`, `reuse`, `unsafe`, `varint`
- protobuf: `binary`, `codegen`, `varint`
  