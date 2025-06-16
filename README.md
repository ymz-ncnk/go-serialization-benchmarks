# go-serialization-benchmarks
In this benchmarks:
- All serializers use the same data. It is generated once and then reused by all 
  of them.
- Each serializer is described with a set of features.
- A single serializer can have multiple benchmark results. For example, MUS has 
  results labeled `mus+raw` and `mus+unsafe`. The former means the raw encoding 
  was used during benchmarking, while the latter refers to the use of unsafe 
  code.
- Unmarshalled data is compared to the original data.
  
## Benchmark Results
Benchmark results are available in [results/benchmarks.txt](results/benchmarks.txt), 
and the corresponding `benchstat` output is in [results/benchstat.txt](results/benchstat.txt).
  
## Fastest Safe
|     NAME     |  NS/OP   | B/SIZE |  B/OP   | ALLOCS/OP |
|--------------|----------|--------|---------|-----------|
| mus          |   102.90 |  58.00 |   48.00 |         1 |
| bebop200sc   |   108.80 |  61.00 |   48.00 |         1 |
| benc         |   111.10 |  58.00 |   48.00 |         1 |
| protobuf_mus |   163.70 |  69.00 |   48.00 |         1 |
| vtprotobuf   |   192.50 |  69.00 |  192.00 |         3 |
| protobuf     |   531.70 |  69.00 |  271.00 |         4 |
| json         |  2779.00 | 150.00 |  600.00 |         9 |
| gob          | 17050.00 | 159.00 | 9493.50 |       195 |

## Fastest Unsafe
|     NAME     | NS/OP  | B/SIZE |  B/OP  | ALLOCS/OP |
|--------------|--------|--------|--------|-----------|
| mus          |  77.21 |  58.00 |   0.00 |         0 |
| benc         |  85.29 |  58.00 |   0.00 |         0 |
| protobuf_mus | 136.40 |  69.00 |   0.00 |         0 |
| vtprotobuf   | 161.40 |  69.00 | 144.00 |         2 |

, where `ns/op`, `B/op`, `allocs/op` are standard `go test -bench=.` output and 
`B/size` - determines how many bytes were used on average by the serializer to 
encode data.
  
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

# Recomendation
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
  