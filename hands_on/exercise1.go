package hands_on

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	Length, Width float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}

func (s Square) Area() float64 {
	return s.Length * s.Width
}

func Info(s Shape) {
	fmt.Println(`Area of the shape is:`, s.Area())
}
