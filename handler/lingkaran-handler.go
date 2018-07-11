package handler

import "math"

type Lingkaran struct {
	radius float64
}

// NewLingkaranHandler handler
func NewLingkaranHandler(radius float64) Lingkaran {
	return Lingkaran{radius}
}

func (c *Lingkaran) Luas() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Lingkaran) Keliling() float64 {
	return 2 * math.Pi * c.radius
}
