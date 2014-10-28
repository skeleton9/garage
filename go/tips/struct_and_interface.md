Struct and Interface
===

## Struct

```go
package main

import (
  "fmt"
  "math"
)

type Circle struct {
  x, y, r float64
}

// Embed types
type MyCircle struct {
  Circle
  x float64
}

func (c *Circle) Area() float64 {
  return math.Pi * c.r * c.r
}

func (c *Circle) Dist() float64 {
  return math.Sqrt(c.x*c.x + c.y*c.y)
}

func main() {
  c1 := Circle{1, 2, 3}
  fmt.Println(c1.Area())
  fmt.Println(c1.Dist())

  var c2 Circle
  c2.x = 1
  c2.y = 2
  c2.r = 3
  fmt.Println(c2.Area())
  fmt.Println(c2.Dist())

  c3 := new(Circle)
  c3.x = 1
  c3.y = 2
  c3.r = 3
  fmt.Println(c3.Area())
  fmt.Println(c3.Dist())

  c4 := MyCircle{x: 10, Circle: Circle{1,2,3}}
  fmt.Println(c4.Area())
  fmt.Println(c4.Dist())
}
```

## Embed Types

Go supports relationships like A is a B, not A has a B by using an embedded
type. Also known as anonymous fields, embedded types look like this:

```go
type Person struct {
  Name string
  Age uint16
}

func (p *Person) Talk() {}

type Android struct {
  Person
  UUID String  // What if Android have a filed with the name `Name`???
              // Then, Name of Person's value can't be get directly!!!
              // see example above
}

a := new(Android)
a.Talk() // will call Person's Talk()
a.Person.Talk() //same as above
```

## Interface

See example:

```go
package main

import (
  "fmt"
  "math"
)

type Shape interface {
  area() float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

type Circle struct {
  x, y, r float64
}

func (r *Rectangle) area() float64 {
  return math.Abs((r.x2 - r.x1) * (r.y1 * r.y2))
}

func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

func totalArea(shapes ...Shape) float64 {
  var total float64 = 0.0
  for _, item := range shapes {
    total += item.area()
  }
  return total
}

type MultipleShapes struct {
  shapes []Shape
}

func (ms *MultipleShapes) totalArea() float64 {
  var total float64 = 0.0
  for _, item := range ms.shapes {
    total += item.area()
  }
  return total
}

type Shapes []Shape

func (sps Shapes) totalArea() float64 {
  var total float64 = 0.0
  for _, item := range sps {
    total += item.area()
  }
  return total

}

func main() {
  c := Circle{1, 2, 3}
  r := Rectangle{1, 3, 2, 5}
  fmt.Println(totalArea(&c, &r))

  ms := MultipleShapes{[]Shape{&c, &r}}
  fmt.Println(ms.totalArea())

  shapes := Shapes{&c, &r}
  fmt.Println(shapes.totalArea())
}

```
