# Iteration

## Benchmark Test

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c03.Repeat("a")
	}
}
```

```bash
go test c03-for/repeat_test.go -bench=.
goos: darwin
goarch: arm64
BenchmarkRepeat-12      49776781                23.33 ns/op
PASS
ok      command-line-arguments  1.191s
```
