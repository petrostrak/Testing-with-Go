# Testing with Go
Explore the Go language by writing tests

### Writing a test is just like writing a function, with a few rules
* It needs to be in a file with a name like `xxx_test.go`
* The test function must start with the word `Test`
* The test function takes one argument only `t *testing.T`
* In order to use the `*testing.T` type, you need to `import "testing"`

For now, it's enough to know that our `t` of type `*testing.T` is our "hook" into the testing framework so you can do things like `t.Fail()` when we want to fail.

### Go doc

Another quality of life feature of Go is the documentation. You can launch the docs locally by running `godoc -http :8000`. If you go to localhost:8000/pkg you will see all the packages installed on your system.
