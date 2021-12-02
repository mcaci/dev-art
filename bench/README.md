# Introduction to benchmarks in Go

Benchmarks in Go are in many ways similar to unit tests but have key differencies and serve a different purpose. As they are not as known as unit tests in Go, this article aims to give an introductory look at Go's benchmarks: how to create, how to run them, how to read the results and a few pointers to some advanced topics in creating benchmark tests in Go.

Benchmarks in Go are functions that test the performance of Go code, they are included the `testing` package of the standard Go library and thus available without any dependecies to external libraries.

When executing a benchmark, you are provided with some information about the execution time and, if requested, the memory footprint of the code under benchmark.

```shell
$go test -benchmem -run ^$ -bench ^(Benchmark1Sort)$
goos: linux
goarch: amd64
Benchmark1Sort-12          10000        105705 ns/op        8224 B/op          2 allocs/op
PASS
ok      _/home/mcaci/code/github.com/mcaci/dev-art/go-bench 1.083s
```

## Creating a benchmark in Go

To create a benchmark, you need to import in your go file the `testing` package and create benchmark functions in a similar way that test functions are created.

For example, when defining unit tests, we write a function in the form of `func TestAny(t *testing)` at the beginning, instead, when we define benchmarks we will create a __`func BenchmarkAny(b *testing.B)`__.

A notable difference in Go's benchmarks with respect to unit tests is the __for loop from 0 to `b.N`__. In fact a benchmark is run multiple times in order to make sure that enough data is collected to improve the accuracy of the performance measurement of the code under benchmark.

The field `b.N` is not a fixed value but it is dynamically adapted to make sure that the benchmark function is run for at least 1 second.

Here to show is a comparison between a benchmark and a test function:

```go
func Benchmark1Sort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sort.Ints(generateSlice(1000))
    }
}
```

```go
func Test1Sort(t *testing.T) {
    slice := generateSlice(1000)
    if len(slice) != 1000 {
        t.Errorf("unexpected slice size: %d", len(slice))
    }
}
```

## Running benchmarks

The starting point for running Go's benchmarks is the `go test` command and here we will see what we need to make sure we're not just running the unit tests.

### Basic usage

```shell
go test -bench .
```

By itself, `go test` runs only unit tests, so we need to add the flag `-bench` to instruct go test to run also the benchmarks.

Specifically, this command runs all unit tests and benchmarks in the current package as denoted by the __.__ added as argument for the `-bench` flag.

The "__.__" value is actually a regular expression which can describe what benchmarks will be executed. For example `go test -bench ^Benchmark1Sort$` will run the benchmark named _Benchmark1Sort_.

Same as when running unit tests, you can add the `-v` flag for _verbose_, which will show more details on the benchmarks executed as well as any printed output, logs, fmt.Prints and so on, or _add a path_ (like "./...") to look for benchmarks on a specific package (or all packages and subpackages).

```shell
go test -bench . -v
go test -bench . ./...
```

### Running benchmarks only

To filter out all the unit tests from `go test`'s execution the `-run ^$` flag should be used.

```shell
go test -run ^$ -bench .
```

The flag `-run` by itself is used to specify which unit tests should be run. Its argument is a regular expression. When we use `^$` as argument  we are effectively filtering out all tests, which means only the benchmarks present in the current package will be executed.

### Running multiple times

Simply add the `-count` parameter to run your benchmark as many times as the specified count: the outcome of all the executions will be shown in the output.

```shell
$ go test -bench ^Benchmark1Sort$ -run ^$ -count 4
goos: linux
goarch: amd64
Benchmark1Sort-12          10207            134834 ns/op
Benchmark1Sort-12           7554            175572 ns/op
Benchmark1Sort-12           7904            148960 ns/op
Benchmark1Sort-12           8568            147594 ns/op
PASS
ok      _/home/mcaci/code/github.com/mcaci/dev-art/go-bench     7.339s
```

This flag is useful when sampling the outcomes of multiple runs to make statistical analisys on the benchmark data.

## Reading benchmark results

Let's take again the following example and run it with `go test -bench` to examine its output.

```go
func Benchmark1Sort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sort.Ints(generateSlice(1000))
    }
}
```

### With execution time

For the first analysis we run the benchmark with `go test -bench ^Benchmark1Sort$ -run ^$`

```shell
$ go test -bench ^Benchmark1Sort$ -run ^$
goos: linux
goarch: amd64
Benchmark1Sort-12           9252            110547 ns/op
PASS
ok      _/home/mcaci/code/github.com/mcaci/dev-art/go-bench     1.053s
```

The output shown is present in any benchmark execution and it shows:

- The information about the __enviroment__ where Go is run, which is also obtained by running `go env GOOS GOARCH` (case sensitive)
  - In our example they are __goos: linux__ and __goarch: amd64__.
- The __benchmark row__ composed of:
  - The __name of the benchmark run__, _Benchmark1Sort-12_, that is itself composed of the function name, _Benchmark1Sort_, followed by the number of CPUs used for the benchmark run, _12_.
  - The __number of times__ the loop has been executed, _9252_.
  - The __average runtime__, expressed in nanoseconds per operation, of the tested function, `sort.Ints(generateSlice(1000))`, which is in this case _110547 ns/op_.
- The information about the benchmark overall status, the package(s) under benchmark and the total time for the execution.

Quick note on the number of CPUs: this parameter can be specified by using the `-cpu` flag; the benchmark will be run multiple times once per CPU defined in the flag.

```shell
$ go test -bench ^Benchmark1Sort$ -run ^$ -cpu 1,2,4
goos: linux
goarch: amd64
Benchmark1Sort              9280            113086 ns/op
Benchmark1Sort-2            9379            117156 ns/op
Benchmark1Sort-4            8637            118818 ns/op
PASS
ok      _/home/mcaci/code/github.com/mcaci/dev-art/go-bench     3.234s
```

If this flag is omitted, a default value is taken from the environment variable _GOMAXPROCS_ and the number of CPUs is not printed in the output when it's equal to 1.

### With execution time and memory

To add the information about memory footprint in the output you can add the `-benchmem` flag as follows.

```shell
$ go test -bench ^Benchmark1Sort$ -run ^$ -benchmem

goos: linux
goarch: amd64
Benchmark1Sort-12          10327            116903 ns/op            8224 B/op          2 allocs/op
PASS
ok      _/home/mcaci/code/github.com/mcaci/dev-art/go-bench     2.128s
```

Two new columns have been added in the output of the benchmark row:

- the __number of bytes__ required by the operation under benchmark, _8224 B/op_
- the __number of allocations__ done by the operation under benchmark, _2 allocs/op_

## Writing more complex benchmarks

Here are some examples of how to write more complex benchmarks.

### StartTimer/StopTimer/ResetTimer

When there is the need to do some setup before actually measuring the time spent to execute code to benchmark, the usage of `StartTimer`, `StopTimer` and `ResetTimer` helps to isolate the bits of code that actually need to be taken into account by the benchmark tools.

Let's take the previous snippet, isolate the creation of the slice from the sorting operation and just measure the execution of the latter.

To do so we can write:

```go
func Benchmark2aSort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        b.StopTimer()
        s := generateSlice(1000)
        b.StartTimer()
        sort.Ints(s)
    }
}
```

By using `b.StopTimer()` we signal that from this point on the execution is not going to be part of the benchmark until `b.StartTimer()` is invoked, which means that in each loop, we only consider the data collected during the execution of `sort.Ints(s)` for the benchmark.

If we want to prepare the slice at the beginning and make it an invariant for the benchmark we can write instead:

```go
func Benchmark2bSort(b *testing.B) {
    s := generateSlice(1000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        sort.Ints(s)
    }
}
```

By using `b.ResetTimer()` we discard all data collected so far and start anew the collection of data for the benchmark, effectively ignoring the execution time of the `generateSlice` call in the overall results.

### Benchmark test cases and subbenchmarks

Like tests, also benchmarks can benefit from the structure of test cases and execution loop to create subbenchmarks.

Let's see an example:

```go
func Benchmark3Sort(b *testing.B) {
    benchData := map[string]struct {
        size int
    }{
        "with size 1000":    {size: 1000},
        "with size 10000":   {size: 10000},
        "with size 100000":  {size: 100000},
        "with size 1000000": {size: 1000000},
    }
    b.ResetTimer()
    for benchName, data := range benchData {
        b.StopTimer()
        s := generateSlice(data.size)
        b.StartTimer()
        b.Run(benchName, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                sort.Ints(s)
            }
        })
    }
}
```

In this example we make use of a `map[string]struct{...}` to define our benchmark cases and data as we would do for complex tests with test cases and we invoke the `b.Run(name string, f func(*testing.B))` to create _subbenchmarks_ that execute separately our benchmark tests.

```shell
$ go test -bench ^Benchmark3Sort$ -run ^$
goos: linux
goarch: amd64
Benchmark3Sort/with_size_1000000-12                   10         130396565 ns/op
Benchmark3Sort/with_size_1000-12                   23210             58078 ns/op
Benchmark3Sort/with_size_10000-12                   1300            865703 ns/op
Benchmark3Sort/with_size_100000-12                   118           8718656 ns/op
PASS
ok      _/home/mcaci/code/github.com/mcaci/dev-art/go-bench     6.670s
```

Notice that the name of the benchmark cases are appended to the _benchmark name_ as part of the output of the benchmark operation as __benchmark_name/benchmark_case_name-number-of-cpus__.

## Parting thoughts

There is still a long way to describe how benchmarks work in Go and to get deeper knowledge in how to write them effectively. One of the main topics that would need its own article would be the benchmark of concurrent code in Go with the usage of the `b.RunParallel` calls, however I hope this article helps in giving the basics of benchmarks in Go and some grounds to explore the functionalities and tools that have not mentioned here.

You can find me up on twitter @nikiforos_frees or here on dev.to @mcaci and I'm looking forward to hearing your questions or comments.

__This was Michele, thanks for reading!__

## References

- [Go's testing package](https://golang.org/pkg/testing/) and [go cmd testing flags](https://golang.org/cmd/go/#hdr-Testing_flags) from the Go team
- Dave Cheney's [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [Justforfunc](https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6)'s article about benchmarks by Francesc Campoy
- Source of benchmark examples inside [instana.com's Practical Golang benchmarks](https://www.instana.com/blog/practical-golang-benchmarks/)
- Talk on advanced benchmarks by Daniel Marti given at the DotGo19 conference: [Optimizing go code without a blindfold](https://www.dotconferences.com/2019/03/daniel-marti-optimizing-go-code-without-a-blindfold)
