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