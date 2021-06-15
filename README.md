# Testing with Go
Explore the Go language by writing tests

### Writing a test is just like writing a function, with a few rules
* It needs to be in a file with a name like `xxx_test.go`
* The test function must start with the word `Test`
* The test function takes one argument only `t *testing.T`
* In order to use the `*testing.T` type, you need to `import "testing"`

For now, it's enough to know that our `t` of type `*testing.T` is our "hook" into the testing framework so you can do things like `t.Fail()` when we want to fail.

### Go doc

Another quality of life feature of Go is the documentation. You can launch the docs locally by running `godoc -http :8000`. If you go to `localhost:8000/pkg` you will see all the packages installed on your system.

### The TDD process and why the steps are important

* Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an   easy to understand description of the failure

* Writing the smallest amount of code to make it pass so we know we have working software

* Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with

### Benchmarking

Writing benchmarks in Go is another first-class feature of the language and it is very similar to writing tests.

```
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```

The `testing.B` gives you access to the cryptically named `b.N`.
When the benchmark code is executed, it runs `b.N` times and measures how long it takes.
The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.
To run the benchmarks do `go test -bench=.`

```
goos: linux
goarch: amd64
pkg: github.com/petrostrak/Testing-with-Go/03-Iteration
cpu: AMD FX(tm)-6300 Six-Core Processor             
BenchmarkRepeat-6        4993839               240.9 ns/op
PASS
ok      github.com/petrostrak/Testing-with-Go/03-Iteration      1.458s
```
What `240.9 ns/op` means is our function takes on average 240 nanoseconds to run (on my computer).
NOTE by default Benchmarks are run sequentially.

### Coverage tool
Go's built-in testing toolkit features a coverage tool. Whilst striving for 100% coverage should not be your end goal, the coverage tool can help identify areas of your code not covered by tests. If you have been strict with TDD, it's quite likely you'll have close to 100% coverage anyway.
Try running `go test -cover` you should see:
```
PASS
coverage: 100.0% of statements
```
### reflect.DeepEqual
Go does not let you use equality operators with slices. You could write a function to iterate over each got and want slice and check their values but for convenience sake, we can use `reflect.DeepEqual` which is useful for seeing if any two variables are the same.
```
if !reflect.DeepEqual(got, want) {
	t.Errorf("got %v, want %v", got, want)
}
```

It's important to note that `reflect.DeepEqual` is not "type safe" - the code will compile even if you did something a bit silly. So while using `reflect.DeepEqual` is a convenient way of comparing slices (and other things) you must be careful when using it.



