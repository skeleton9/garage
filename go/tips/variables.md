# Go Variables

### Variable declaration

Variable is a storage with a name and a type.

```go
var x int
var x, y int
var p *int
var a [3]int
f func(a, b int) int {}

// define multiple variables
var (
  Name string = "bob"
  Age int = 12
  )
```
As declare and init with a value is quite common, GO provide simple way to do it.
Use *Short assignment* like `a := 1` to declare a `var` with implicit type.
It is only available inside function. Outside a function, every construct
begins with a keyword (`var`, `func`, and so on)

- Examples

```go
package main

const Pi = 3.14
const Pi float32 = 3.1415

func main() {
  var x1, y1 = 3, 4 // ok, type can be implicit
  var x2, y2 int = 3, 5 // ok
  var a1, b1 := 3, 4 //wrong!!! When use :=, do not need var
  a2, b2 := 3, 4 //ok, short assignment
  a3, b3 = 3, 4 //wrong, if not with var, use :=
}
```

### Variable Names in mixed caps

Go is not using underscore in names. So for multiword names, use `MixedCaps` or
`mixedCaps`.

### Names in packages

- Only names start with upper case is visible outside a package.
- Once imported, package name becomes an accessor for its contents.
  `import "bytes"`, and use `bytes.Buffer`.



### Constants

```go
const Pi = 3.14
```
Constants can be character, string, boolean, or numeric values.

You can not use *short assignment* for constants.

### Note

The Go compiler won't allow you to create variables that you never use.
So you can use `_` for the temp variable that you won't use.

```go
var x = int[4]{1, 2, 3, 4}
var total int = 0
for _, value := range x {
    total += value
}
```
