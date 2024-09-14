<img src="https://img.shields.io/badge/go%20version-min%201.21-green" alt="Go version"/>

# go-topmap

Top items map (string of int), can be used to maintain a list of counters and get the top N items.

## Installation

```go
go get -u github.com/dmachard/go-topmap
```

## Usage example

Example to use this module

```go
// init topmap with 5 max items
tmap := NewTopMap(5)

// add items to the map
tmap.Record("oranges", 2)
tmap.Record("apples", 10)

// incremente items
tmap.Inc("oranges")


// get items
items = tmap.Get()
fmt.Println(items)
[{apples 10} {oranges 3}]
```

## Testing

```bash
=== RUN   TestTopMapRecord
--- PASS: TestTopMapRecord (0.00s)
=== RUN   TestTopMapInc
--- PASS: TestTopMapInc (0.00s)
=== RUN   TestTopMapGetSorted
--- PASS: TestTopMapGetSorted (0.00s)
=== RUN   TestTopMapRecordMax
--- PASS: TestTopMapRecordMax (0.00s)
PASS
ok      dmachard/go-topmap      0.002s
```

## Benchmark

```bash
goos: linux
goarch: amd64
pkg: dmachard/go-topmap
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkTopMapRecord/TopMapSize:_10_-4                  4745701               247.1 ns/op
BenchmarkTopMapRecord/TopMapSize:_100_-4                  604754              1761 ns/op
BenchmarkTopMapRecord/TopMapSize:_1000_-4                  56764             21874 ns/op
BenchmarkTopMapRecord/TopMapSize:_10000_-4                 13870      

BenchmarkTopMapRecord/TopMapSize:_10_-4                  4745701               247.1 ns/op
BenchmarkTopMapRecord/TopMapSize:_100_-4                  604754              1761 ns/op
BenchmarkTopMapRecord/TopMapSize:_1000_-4                  56764             21874 ns/op
BenchmarkTopMapRecord/TopMapSize:_10000_-4                 13870            113485 ns/op

BenchmarkTopMapGet/TopMapSize:_10_-4                     1012234              1264 ns/op
BenchmarkTopMapGet/TopMapSize:_100_-4                      75252             15170 ns/op
BenchmarkTopMapGet/TopMapSize:_1000_-4                      5414            199167 ns/op
BenchmarkTopMapGet/TopMapSize:_10000_-4                      429           2792158 ns/op
PASS
ok      dmachard/go-topmap      14.896s
```