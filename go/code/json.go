package main

import "fmt"

type Rectangle struct {
	height, width int
}

func (r Rectangle) area() int {
	return r.height * r.width
}

func main() {
	r := Rectangle{3, 4}
	fmt.Println(r.area())
}
