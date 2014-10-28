Go Functions
===

A GO function looks like:
```
func func_name(params) return_values {
  code in side
}
```

#### Variadic Functions

The final parameter in a function signature may have a type prefixed with `...`.
A function with such a parameter is called *variadic* and may be invoked with
zero or more arguments for that parameter.

```go
// fmt.Print is defined as below
func Print(a ...interface{}) (n int, err error)

// call it like:
fmt.Print("hello", "jack", "age", 23)
// "hello jack age 23"
```

#### Name the return value

```go
func test(x int) (r1 int, r2 int) {
  r1 = x*10
  r2 = x*100
  return   // r1, r2 will be returned
}
```

#### Return multiple values

Go is also capable of returning multiple values from a function:

```go
func test(x int) (int, int) {
  return x, x*10
}

func main() {
  a, b := test(1) // 1, 10
}
```

Multiple values are often used to return an error value along with the result
`(x, err := f())`, or a boolean to indicate success `(x, ok := f())`.

#### Closure

It is possible to create functions inside of a function. Local function can
access other local variables.

```go
package main
import "fmt"

func main() {
	a := 5

	inc := func() {
		a += 1
	}

	inc() // a == 6
	inc() // a == 7
	fmt.Println(a) //7
}
```

#### Defer

Go has a special statement called defer which schedules a function call to be
run after the function completes.

```go
package main

import "fmt"

func test() {
	fmt.Println("I am deferred")
}

func main() {
	defer test()
	fmt.Println("I will be first printed")
}
```

Will produce:

```
I will be first printed
I am deferred
```

`defer` is often used when resources need to be freed in some way. For example
when we open a file we need to make sure to close it later. With defer:

```go
f, _ := os.Open(filename)
defer f.Close()
```

This has 3 advantages:

1. it keeps our Close call near our Open call so its easier to understand
2. if our function had multiple return statements (perhaps one in an if and
   one in an else) Close will happen before both of them and
3. deferred functions are run even if a run-time panic occurs.

#### Panic & Recover

- Panic

The `panic` built-in function stops normal execution of the current
goroutine. When a function F calls panic, normal execution of F stops
immediately. Any functions whose execution was deferred by F are run in
the usual way, and then F returns to its caller.

```go
func panic(v interface{})
```
- Recover

```go
// The recover built-in function allows a program to manage behavior of a
// panicking goroutine. Executing a call to recover inside a deferred
// function (but not any function called by it) stops the panicking sequence
// by restoring normal execution and retrieves the error value passed to the
// call of panic. If recover is called outside the deferred function it will
// not stop a panicking sequence. In this case, or when the goroutine is not
// panicking, or if the argument supplied to panic was nil, recover returns
// nil. Thus the return value from recover reports whether the goroutine is
// panicking.
func recover() interface{}
```

We can use `recover` in a deferred function to get panic info.

```go
package main

import "fmt"

func main() {
    defer func() {
        str := recover()
        fmt.Println(str) // PANIC
    }()
    panic("PANIC")
}
```
