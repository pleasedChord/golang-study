package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circleb struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return r.Width * r.Height
}

func (c Circleb) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circleb) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	rectangle := Rectangle{5, 3}
	fmt.Println("rectangle.area=", rectangle.Area(), "rectangle.Perimeter=", rectangle.Perimeter())

	circleb := Circleb{6}
	fmt.Println("circleb.area=", circleb.Area(), "circleb.Perimeter=", circleb.Perimeter())
}
