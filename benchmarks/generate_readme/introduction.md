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