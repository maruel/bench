# bench

A collection of low level Go and CPU benchmarks to compare behavior
across OSes and CPU architecture. This package has no dependency.

Usage:

```
go get -d -u github.com/maruel/bench
go test -cpu 1 -bench=. github.com/maruel/bench -benchtime=100ms
```
