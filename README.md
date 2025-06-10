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
| mus          |         12573285 | 90.82 |     58 |   48 |         1 |
| bebop200sc   |         12233728 | 93.62 |     61 |   48 |         1 |
| benc         |         11444528 | 98.88 |     58 |   48 |         1 |
| protobuf_mus |          7712743 | 139.7 |     69 |   48 |         1 |
| vtprotobuf   |          8138298 | 154.7 |     69 |  192 |         3 |
| protobuf     |          2539180 | 449.5 |     69 |  271 |         4 |
| json         |           458604 |  2552 |    150 |  600 |         9 |
| gob          |            87592 | 14654 |    159 | 9495 |       195 |
  
## Fastest Unsafe
|     NAME     | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|--------------|------------------|-------|--------|------|-----------|
| mus          |         14824605 | 69.77 |     58 |    0 |         0 |
| benc         |         13746213 | 75.89 |     58 |    0 |         0 |
| protobuf_mus |          9582236 | 108.5 |     69 |    0 |         0 |
| vtprotobuf   |          8504149 | 138.7 |     69 |  144 |         2 |
  
## All
|               NAME                | ITERATIONS COUNT | NS/OP | B/SIZE | B/OP | ALLOCS/OP |
|-----------------------------------|------------------|-------|--------|------|-----------|
| mus+reuse+unsafe                  |         14824605 | 69.77 |     58 |    0 |         0 |
| benc+raw+reuse+unsafestr          |         13746213 | 75.89 |     58 |    0 |         0 |
| mus+unsafe                        |         12485073 | 89.89 |     58 |   64 |         1 |
| mus+notunsafe+reuse               |         12573285 | 90.82 |     58 |   48 |         1 |
| bebop200sc+notunsafe+reuse        |         12233728 | 93.62 |     61 |   48 |         1 |
| benc+raw+unsafestr                |         11988807 | 98.27 |     58 |   64 |         1 |
| benc+raw+reuse                    |         11444528 | 98.88 |     58 |   48 |         1 |
| mus+raw+reuse                     |         11204305 | 101.6 |     58 |   48 |         1 |
| protobuf_mus+reuse+unsafe         |          9582236 | 108.5 |     69 |    0 |         0 |
| mus+reuse+varint                  |          9528118 | 111.3 |     58 |   48 |         1 |
| mus+notunsafe                     |         11284736 | 119.8 |     58 |  112 |         2 |
| bebop200sc+notunsafe              |          9227256 | 122.4 |     61 |  112 |         2 |
| mus+raw                           |         10191066 | 126.2 |     58 |  112 |         2 |
| benc+raw                          |          9358150 | 134.5 |     62 |  112 |         2 |
| vtprotobuf+raw+reuse+unsafeunm    |          8504149 | 138.7 |     69 |  144 |         2 |
| protobuf_mus+raw+reuse            |          7712743 | 139.7 |     69 |   48 |         1 |
| vtprotobuf+reuse+unsafeunm+varint |          7770396 | 146.3 |     70 |  144 |         2 |
| mus+varint                        |          8550482 | 147.3 |     58 |  112 |         2 |
| protobuf_mus+unsafe               |          8148639 | 150.6 |     69 |   79 |         1 |
| vtprotobuf+raw+reuse              |          8138298 | 154.7 |     69 |  192 |         3 |
| protobuf_mus+reuse+varint         |          6931992 | 157.3 |     70 |   48 |         1 |
| protobuf_mus+native+reuse+unsafe  |          6947217 | 170.2 |     69 |  144 |         2 |
| vtprotobuf+reuse+varint           |          7619318 | 170.5 |     70 |  192 |         3 |
| vtprotobuf+raw+unsafeunm          |          7944594 | 181.4 |     69 |  223 |         3 |
| protobuf_mus+raw                  |          6999267 | 182.6 |     69 |  127 |         2 |
| protobuf_mus+reuse                |          5981740 | 196.4 |     70 |  127 |         2 |
| vtprotobuf+unsafeunm+varint       |          7398261 | 196.9 |     70 |  223 |         3 |
| vtprotobuf+raw                    |          7372633 |   204 |     69 |  271 |         4 |
| vtprotobuf+varint                 |          6893404 | 218.9 |     70 |  271 |         4 |
| protobuf+raw                      |          2539180 | 449.5 |     69 |  271 |         4 |
| protobuf+varint                   |          2740602 | 456.6 |     70 |  271 |         4 |
| json                              |           458604 |  2552 |    150 |  600 |         9 |
| gob                               |            87592 | 14654 |    159 | 9495 |       195 |

, where `iterations count`, `ns/op`, `B/op`, `allocs/op` are standard 
`go test -bench=.` output and `B/size` - determines how many bytes were used on 
average by the serializer to encode `Data`.  
  
# Features
- json: reflect,text,int
- gob: binary,int
- mus: manual,unsafestr,raw,reuse,notunsafe,binary,codegen,varint,unsafe
- benc: binary,codegen,manual,raw,reuse,unsafestr
- protobuf: binary,codegen,raw,varint
- vtprotobuf: raw,reuse,unsafeunm,varint,binary,codegen
  
# List of Features
Each feature describes a property of a serializer:
- `reflect` – uses reflection.
- `codegen` – uses code generation.
- `manual` – only provides serialization primitives, requiring manual usage.
- `text` – uses a text-based serialization format.
- `binary` – uses a binary serialization format.
- `varint` – supports varint encoding.
- `raw` – supports raw encoding.
- `int` – supports the int type.
- `native` – uses native data structures.

Features that must appear in the result name when used:
- `reuse` – supports buffer reuse.
- `unsafe` – uses unsafe code.
- `unsafestr` – uses unsafe code only for string serialization.
- `unsafeunm` – uses unsafe code only for unmarshalling.
- `notunsafe` – uses unsafe code for all types except string, and copies data 
during unmarshalling.
- `fixbuf` - uses a fixed buffer.

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
```bash
go test -bench=.
```
To filter serializers - for example, by `binary` and `manual` types:
```bash
go test -bench=. -- -f binary,manual
```
To run benchmarks for one particular case just name it, for example:
```bash
go test -bench=BenchmarkSerializers/mus
```
Or to see the results obtained using the `reuse` feature:
```bash
go test -bench=/.+reuse
```

# Generate README.md
```bash
go generate
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