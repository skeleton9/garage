# Control Structures in Go

- `if` and `switch` accept an optional initialization statement like `for`.

## for

For is like C/Java, but the ( ) are gone and { } are required.

```go
package main
import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
```

It also take the work for `while`, there is no `while` in Go.

```go
sum : = 1
for sum < 10 {
  sum += sum
}

for {
  //endless loop
}
```


## if else

Similar to `for`, ( ) is gone and { } are required.

```go
if x < 0 {
  x = -x
}
```

You can put a short statement in the `if` expression, and the variable declared
in the short statement is only in scope util the end of the `if`.

```go
func pow(x, n, lim float64) float64{
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Println("%g >= %g\n", v, lim)
  }
  // v is not available here
  return lim
}
```

## switch case

Switch evaluate cases from top to bottom, and stops when a case succeeds.
You do not need to write break yourself. `default` is the last one to be
evaluated no matter where it is placed.

If no expression is provided for `switch`, it just like `switch true`.

```go
switch i {
  case 0, 1:
      doSomething()
  case 2:
      doSomething()
  default:
    fmt.Println(i)
}
```
- Type Switch [:question:](http://golang.org/doc/effective_go.html#type_switch)

```
switch t := t.(type) {
  case bool:
  ///
  case int:
  ///
  default:
  ///
}
```

## break, continue :question:

TBD
