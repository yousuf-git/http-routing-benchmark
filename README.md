# HTTP Routing Benchmark (Go)

This benchmark suite compares the overall performance of the following Go HTTP routers:

* [Shift](https://github.com/yousuf64/shift)
* [Gin](https://github.com/gin-gonic/gin)
* [Echo](https://github.com/labstack/echo)

Most of the routes and tests in this suite are inspired by the [julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) benchmark suite.

Benchmark suite can be run on [GitHub Actions](https://github.com/yousuf64/http-routing-benchmark/actions) as well.

## Results
Date: Mar 18, 2023

System specifications:
* 12th Gen Intel Core i7-1265U vPro (12 MB cache, 10 cores, up to 4.80 GHz Turbo)
* 32 GB (2x 16 GB), DDR4-3200
* Windows 10 Enterprise 22H2
* Go 1.19.4 (windows/amd64)

```
BenchmarkShift_CaseInsensitiveAll-12             1101076               949.6 ns/op             0 B/op          0 allocs/op
BenchmarkGin_CaseInsensitiveAll-12                773574              1636 ns/op               0 B/op          0 allocs/op
BenchmarkShift_GithubAll-12                        60122             21789 ns/op               0 B/op          0 allocs/op
BenchmarkGin_GithubAll-12                          27148             54872 ns/op            9911 B/op        154 allocs/op
BenchmarkEcho_GithubAll-12                         36133             29333 ns/op               0 B/op          0 allocs/op
BenchmarkShift_GPlusAll-12                       1847936               632.3 ns/op             0 B/op          0 allocs/op
BenchmarkGin_GPlusAll-12                         1421362               807.0 ns/op             0 B/op          0 allocs/op
BenchmarkEcho_GPlusAll-12                        1000000              1027 ns/op               0 B/op          0 allocs/op
BenchmarkShift_OverlappingRoutesAll-12            923389              1279 ns/op               0 B/op          0 allocs/op
BenchmarkEcho_OverlappingRoutesAll-12             285714              4038 ns/op               0 B/op          0 allocs/op
BenchmarkShift_ParseAll-12                       1225710               971.9 ns/op             0 B/op          0 allocs/op
BenchmarkGin_ParseAll-12                          762722              1475 ns/op               0 B/op          0 allocs/op
BenchmarkEcho_ParseAll-12                         748764              1731 ns/op               0 B/op          0 allocs/op
BenchmarkShift_RandomAll-12                       827026              1295 ns/op               0 B/op          0 allocs/op
BenchmarkGin_RandomAll-12                         231808              7166 ns/op            2201 B/op         34 allocs/op
BenchmarkEcho_RandomAll-12                        221920              5045 ns/op               0 B/op          0 allocs/op
BenchmarkShift_StaticAll-12                       416360              2784 ns/op               0 B/op          0 allocs/op
BenchmarkGin_StaticAll-12                         109315             12010 ns/op               0 B/op          0 allocs/op
BenchmarkEcho_StaticAll-12                         76387             14764 ns/op               0 B/op          0 allocs/op
```

* Column 1: Benchmark name
* Column 2: Number of iterations, higher means more confident result
* Column 3: Nanoseconds elapsed per operation (ns/op), lower is better
* Column 4: Number of bytes allocated on heap per operation (B/op), lower is better
* Column 5: Average allocations per operation (allocs/op), lower is better