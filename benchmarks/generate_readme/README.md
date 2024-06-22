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