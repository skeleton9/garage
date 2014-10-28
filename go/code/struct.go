package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

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
