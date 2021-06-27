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

### Table driven tests
[Table driven tests](https://github.com/golang/go/wiki/TableDrivenTests) are useful when you want to build a list of test cases that can be tested in the same manner.
```
func TestArea(t *testing.T) {

    areaTests := []struct {
        shape Shape
        want  float64
    }{
        {Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},
    }

    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }

}
```
The only new syntax here is creating an "anonymous struct", `areaTests`. We are declaring a slice of structs by using `[]struct`. Then we fill the slice with cases. We then iterate over them just like we do any other slice, using the struct fields to run our tests.

### http.HandlerFunc

>The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
```
type HandlerFunc func(ResponseWriter, *Request)
```

From the documentation, we see that type `HandlerFunc` has already implemented the `ServeHTTP` method. By type casting our `PlayerServer` function with it, we have now implemented the required `Handler`.

### http.ListenAndServe(":8080"...)
`ListenAndServe` takes a port to listen on a `Handler`. If there is a problem the web server will return an error, an example of that might be the port already being listened to.

### JSON Decoding and Encoding
```
var got []Player
err := json.NewDecoder(response.Body).Decode(&got)
```
To parse JSON into our data model we create a `Decoder` from `encoding/json` package and then call its `Decode` method. To create a `Decoder` it needs an `io.Reader` to read from which in our case is our response spy's `Body`.
`Decode` takes the address of the thing we are trying to decode into which is why we declare an empty slice of `Player` the line before.
Parsing JSON can fail so `Decode` can return an `error`. There's no point continuing the test if that fails so we check for the error and stop the test with `t.Fatalf` if it happens. Notice that we print the response body along with the error as it's important for someone running the test to see what string cannot be parsed.

```
leagueTable := []Player{
    {"Petros", 20},
}

json.NewEncoder(w).Encode(leagueTable)
```
Notice the lovely symmetry in the standard library.
* To create an Encoder you need an `io.Writer` which is what http.ResponseWriter implements.
* To create a Decoder you need an io.Reader which the Body field of our response spy implements.












