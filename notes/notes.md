# Go Fundamentals

## Writing Tests
Tests operate like a function with the following rules
- It needs to be in a file with a name similar to ``xxx_test.go``
- The test function must start with the word `Test`
- The test function takes one argument only `t *testing.T`
- To use the `*testing.T` type, you need to `import "testing"`

## Go doc
It is possible to launch Go docs locally by running `godoc -http:8000` in the terminal. Then going to link `localhost:8000/pkg` to see the packages installed on the system. 

If `godoc` is not installed, it can be installed manually using `go install golang.org/x/tools/cmd/godoc@latest`.


## Benchmarking
Benchmarks operate within the testing file and have the following rules
- The benchmark function must start with the word `Benchmark`
- The benchmark function takes one argument only `t *testing.B`
- To run benchmarks, enter in the command line `go test -bench=.`

## Test coverage
To check that your tests cover all functions run the following in the terminal `go test -cover`

## Pointers
In Go, pointers are automatically dereferenced. However, it is best to 
