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
| mus          |         14082382 | 80.55 |     58 |   48 |         1 |
| bebop200sc   |         13028316 | 85.52 |     61 |   48 |         1 |
| benc         |         11802933 | 94.81 |     60 |   48 |         1 |
| protobuf_mus |          6688456 |   163 |     69 |   48 |         1 |
| vtprotobuf   |          6871676 | 191.3 |     69 |  192 |         3 |
| protobuf     |          2460066 | 503.1 |     70 |  271 |         4 |
| json         |           421777 |  2697 |    150 |  600 |         9 |
| gob          |            66589 | 17880 |    159 | 9407 |       233 |
  
## Fastest Unsafe
|     NAME     | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|--------------|------------------|-------|--------|------|-----------|
| mus          |         18103663 |    57 |     58 |    0 |         0 |
| benc         |         14829669 | 69.76 |     60 |    0 |         0 |
| protobuf_mus |          8152317 | 129.4 |     69 |    0 |         0 |
| vtprotobuf   |          7640108 | 165.6 |     69 |  144 |         2 |
  
## All
|                 NAME                  | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|---------------------------------------|------------------|-------|--------|------|-----------|
| mus+reuse+unsafe                      |         18103663 |    57 |     58 |    0 |         0 |
| benc+raw+reuse+unsafestr              |         14829669 | 69.76 |     60 |    0 |         0 |
| mus+notunsafe+reuse                   |         14082382 | 80.55 |     58 |   48 |         1 |
| mus+unsafe                            |         14908098 | 81.72 |     58 |   64 |         1 |
| bebop200sc+notunsafe+reuse            |         13028316 | 85.52 |     61 |   48 |         1 |
| benc+raw+unsafestr                    |         12691201 | 92.63 |     60 |   64 |         1 |
| mus+raw+reuse                         |         12016142 | 92.98 |     58 |   48 |         1 |
| benc+raw+reuse                        |         11802933 | 94.81 |     60 |   48 |         1 |
| bebop200sc+notunsafe                  |         10276396 |   108 |     61 |  112 |         2 |
| mus+raw+reuse+varint                  |          9901872 | 112.2 |     59 |   48 |         1 |
| mus+notunsafe                         |         12403351 | 113.3 |     58 |  112 |         2 |
| mus+raw                               |          9719802 |   122 |     58 |  112 |         2 |
| benc+raw                              |         10587306 | 122.1 |     64 |  112 |         2 |
| protobuf_mus+reuse+unsafe             |          8152317 | 129.4 |     69 |    0 |         0 |
| mus+raw+varint                        |          8684614 | 144.8 |     59 |  112 |         2 |
| protobuf_mus+raw+reuse                |          6688456 |   163 |     69 |   48 |         1 |
| vtprotobuf+raw+reuse+unsafeunm        |          7640108 | 165.6 |     69 |  144 |         2 |
| vtprotobuf+raw+reuse+unsafeunm+varint |          7736064 | 174.1 |     70 |  144 |         2 |
| protobuf_mus+raw+reuse+varint         |          5891563 |   186 |     70 |   48 |         1 |
| protobuf_mus+unsafe                   |          6586830 | 187.4 |     69 |   79 |         1 |
| vtprotobuf+raw+reuse                  |          6871676 | 191.3 |     69 |  192 |         3 |
| protobuf_mus+native+reuse+unsafe      |          6588710 | 194.1 |     69 |  144 |         2 |
| vtprotobuf+raw+reuse+varint           |          6774465 | 198.3 |     70 |  192 |         3 |
| vtprotobuf+raw+unsafeunm              |          6554413 | 212.2 |     69 |  223 |         3 |
| vtprotobuf+raw+unsafeunm+varint       |          6402992 | 218.7 |     70 |  223 |         3 |
| vtprotobuf+raw                        |          6292698 | 227.9 |     69 |  271 |         4 |
| protobuf_mus+raw                      |          5407700 | 231.8 |     69 |  127 |         2 |
| protobuf_mus+raw+varint               |          5300852 | 243.2 |     70 |  127 |         2 |
| vtprotobuf+raw+varint                 |          5882449 | 246.4 |     70 |  271 |         4 |
| protobuf+raw+varint                   |          2460066 | 503.1 |     70 |  271 |         4 |
| protobuf+raw                          |          2213274 | 504.4 |     69 |  271 |         4 |
| json                                  |           421777 |  2697 |    150 |  600 |         9 |
| gob                                   |            66589 | 17880 |    159 | 9407 |       233 |

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