go test -bench BenchmarkSerializers -count=10 -timeout=30m > ./results/benchmarks.txt
benchstat ./results/benchmarks.txt > ./results/benchstat.txt

go generate