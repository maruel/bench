# bench

A collection of low level Go and CPU benchmarks to compare behavior
across OSes and CPU architecture. This package has no dependency.

## Usage

```
go get -d -u github.com/maruel/bench
go test -cpu 1 -bench=. github.com/maruel/bench -benchtime=100ms
```

## Inspection

It is useful to inspect the assembly, to compare micro benchmarks between each
others or across platforms.

```
go test -c
go tool objdump -S -s BenchmarkDivision ./bench.test
```
