package main

// BenchmarkTable is a collection of BenchmarkResult, representing a full table.
type BenchmarkTable []BenchmarkResult

// BenchmarkResult holds the processed data for a single benchmark case.
type BenchmarkResult struct {
	Name     string  // Short name like "mus", "benc", etc.
	NsOp     float64 // Time in nanoseconds per operation.
	BSize    float64 // Size in bytes.
	BOp      float64 // Bytes allocated per operation.
	AllocsOp float64 // Allocations per operation.
}
