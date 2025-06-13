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
To see the results obtained using the `reuse` feature:
```bash
go test -bench=/.+reuse
```

# Generate README.md
```bash
go generate
```

## Recomendation
When running benchmarks on a laptop, make sure that it is connected to a charger 
and the fan is at full speed.

# Contribution
First of all, you need to create a new package for your serializer if it doesn't
already exist. Then:
1. Implement [benchser.Serializer](benchser/serializer.go) interface.
2. If you use own `Data` make shure it implements `EqualTo(data Data) error` 
   method, also add `func ToYourData(data serializer.Data) (d Data)`
   function (an example can be found in [projects/bbebop200sc/serializers.go](projects/bebop200sc/serializers.go)).
3. Define
   ```go
   var Serializers = []benchser.Serializer[...]{...}
   ```
   variable. Note that it can contain several serializers that produce different
   results.
4. Create PR.

If you want to run benchmarks from your own project, there is the
[benchser.BenchmarkSerializer(...)](benchser/benchser.go) function.