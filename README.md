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
| mus          |         15015302 | 74.93 |     58 |   48 |         1 |
| bebop200sc   |         14287324 | 79.89 |     61 |   48 |         1 |
| benc         |         12620012 | 90.82 |     60 |   48 |         1 |
| protobuf_mus |          6715797 |   163 |     69 |   48 |         1 |
| protobuf     |          2666491 | 464.2 |     70 |  271 |         4 |
| json         |           438956 |  2659 |    150 |  600 |         9 |
| gob          |            75608 | 16408 |    159 | 9407 |       233 |
  
## Fastest Unsafe
|     NAME     | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|--------------|------------------|-------|--------|------|-----------|
| mus          |         19392787 |    53 |     58 |    0 |         0 |
| benc         |         16042617 | 64.95 |     60 |    0 |         0 |
| protobuf_mus |          8486268 | 127.5 |     69 |    0 |         0 |
  
## All
|               NAME               | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|----------------------------------|------------------|-------|--------|------|-----------|
| mus+reuse+unsafe                 |         19392787 |    53 |     58 |    0 |         0 |
| benc+raw+reuse+unsafestr         |         16042617 | 64.95 |     60 |    0 |         0 |
| mus+notunsafe+reuse              |         15015302 | 74.93 |     58 |   48 |         1 |
| mus+unsafe                       |         15711453 | 76.87 |     58 |   64 |         1 |
| bebop200sc+notunsafe+reuse       |         14287324 | 79.89 |     61 |   48 |         1 |
| benc+raw+unsafestr               |         13961784 | 85.31 |     60 |   64 |         1 |
| mus+raw+reuse                    |         13010559 | 86.94 |     58 |   48 |         1 |
| benc+raw+reuse                   |         12620012 | 90.82 |     60 |   48 |         1 |
| mus+notunsafe                    |         13135551 | 106.5 |     58 |  112 |         2 |
| bebop200sc+notunsafe             |         10095086 | 107.9 |     61 |  112 |         2 |
| mus+raw                          |         11000094 | 111.4 |     58 |  112 |         2 |
| mus+raw+reuse+varint             |          9998512 | 112.8 |     59 |   48 |         1 |
| benc+raw                         |         11287188 | 114.4 |     64 |  112 |         2 |
| protobuf_mus+reuse+unsafe        |          8486268 | 127.5 |     69 |    0 |         0 |
| mus+raw+varint                   |          9191444 | 135.7 |     59 |  112 |         2 |
| protobuf_mus+raw+reuse           |          6715797 |   163 |     69 |   48 |         1 |
| protobuf_mus+raw+reuse+varint    |          6351121 | 173.5 |     70 |   48 |         1 |
| protobuf_mus+native+reuse+unsafe |          7304798 | 178.1 |     69 |  144 |         2 |
| protobuf_mus+unsafe              |          6621054 | 186.3 |     69 |   79 |         1 |
| protobuf_mus+raw                 |          6077431 | 212.9 |     69 |  127 |         2 |
| protobuf_mus+raw+varint          |          5650884 | 230.6 |     70 |  127 |         2 |
| protobuf+raw+varint              |          2666491 | 464.2 |     70 |  271 |         4 |
| protobuf+raw                     |          2361918 | 472.1 |     69 |  271 |         4 |
| json                             |           438956 |  2659 |    150 |  600 |         9 |
| gob                              |            75608 | 16408 |    159 | 9407 |       233 |

, where `iterations count`, `ns/op`, `B/op`, `allocs/op` are standard 
`go test -bench=.` output and `B/size` - determines how many bytes were used on 
average by the serializer to encode `Data`.  
  
# Features
- benc: `binary`, `manual`, `raw`, `reuse`, `unsafestr`
- bebop200sc: `binary`, `codegen`, `notunsafe`, `reuse`
- gob: `binary`, `int`
- json: `int`, `reflect`, `text`
- mus: `binary`, `int`, `manual`, `notunsafe`, `raw`, `reuse`, `unsafe`, `unsafestr`, `varint`
- protobuf: `binary`, `codegen`, `raw`, `varint`
  
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
- `native` - when one format is implemented by a serializer of another format 
  and native data is used.

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