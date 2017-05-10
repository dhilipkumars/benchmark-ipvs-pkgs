# A Simple test to bench mark among two differnt IPVS libraries
1) libnetwork - from docker
2) seesaw - from google

```
$go test -v -bench=. -count=5 -cpu=1 -benchmem
BenchmarkSingleServiceLibNet         100          14072409 ns/op           47030 B/op        374 allocs/op
BenchmarkSingleServiceLibNet         100          15351626 ns/op           47030 B/op        374 allocs/op
BenchmarkSingleServiceLibNet         100          14986575 ns/op           47030 B/op        374 allocs/op
BenchmarkSingleServiceLibNet         100          14018903 ns/op           47030 B/op        374 allocs/op
BenchmarkSingleServiceLibNet         100          14917029 ns/op           47030 B/op        374 allocs/op
BenchmarkSingleServiceSeeSaw         100          15187018 ns/op           32090 B/op       1465 allocs/op
BenchmarkSingleServiceSeeSaw         100          14874869 ns/op           32088 B/op       1465 allocs/op
BenchmarkSingleServiceSeeSaw         100          14916282 ns/op           32089 B/op       1465 allocs/op
BenchmarkSingleServiceSeeSaw         100          13916363 ns/op           32093 B/op       1465 allocs/op
BenchmarkSingleServiceSeeSaw         100          14676425 ns/op           32089 B/op       1465 allocs/op
PASS
ok      github.com/dhilipkumars/benchmark-ipvs-pkgs     17.126s

```
