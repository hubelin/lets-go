package shapes

import "math"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	height float64
	width  float64
}

func (r *Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
