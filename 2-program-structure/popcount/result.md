goos: darwin
goarch: arm64
pkg: example.com/popcount
BenchmarkPopCount1-12           1000000000               0.2599 ns/op          0 B/op          0 allocs/op
BenchmarkPopCount2-12           393921139                2.991 ns/op           0 B/op          0 allocs/op
BenchmarkPopCount3-12           16620114                71.79 ns/op            0 B/op          0 allocs/op
PASS
ok      example.com/popcount    3.770s
