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

### The TDD process and why the steps are important

* Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an   easy to understand description of the failure

* Writing the smallest amount of code to make it pass so we know we have working software

* Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with