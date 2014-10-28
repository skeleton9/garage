Types in GO
===

Go is a **statically typed** programming language. This means that variables always
have a specific type and that type cannot change.

## Numbers

Generally, there are two kinds of numbers: integers and float point numbers.

#### Integers

Go's integer types are:

`uint8, uint16, uint32, uint64, int8, int16, int32 and int64`.

8, 16, 32, 64 tell us how many bits are used for the integer.

There are two alias, `byte` is same as `uint8`, and `rune` as `int32`.

There are 3 machine dependent integer types: `uint`, `int` and `uintptr`. Their
size depends on the type of the machine architecture.

#### Float Point Numbers

Some facts you should keep in mind:

- Float point numbers are inexact.
Do not compare two float point numbers to be equal.

- Float point number also have sizes, as `float32` and `float64`.

- `NaN` for "not a number".

- `complex64` and `complex128` for complex numbers.

#### Operators

+, -, *, /, %

## Strings
String literals can be created using double quotes `"Hello World"` or back ticks
``` `Hello World` ```.

Here are some basic operations for strings:

```go
package main

import "fmt"

func main() {
    fmt.Println(len("Hello World"))
    fmt.Println("Hello World"[1])
    fmt.Println("Hello " + "World")
}
```

## Booleans

A boolean value is a 1 bit integer type used to represent `true` and `false`.

Operators:

| && | and |
| || | or |
| !  | not |
