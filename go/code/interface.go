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
