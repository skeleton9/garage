# Go Variables

### Variable Names in mixed caps

Go is not using underscore in names. So for multiword names, use `MixedCaps` or
`mixedCaps`.

### Names in packages

- Only names start with upper case is visible outside a package.
- Once imported, package name becomes an accessor for its contents.
  `import "bytes"`, and use `bytes.Buffer`.

### Variable declaration

```go
var x int
var x, y int
var p *int
var a [3]int
f func(a, b int) int {}

var (
  Name string = "bob"
  Age int = 12
  )
```

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

### Constants

```go
const Pi = 3.14
```
Constants can be character, string, boolean, or numeric values.

You can not use *short assignment* for constants.

### Basic Types

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
