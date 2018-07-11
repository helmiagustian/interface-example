package handler

import "math"

// Lingkaran struct type, implementing geometry interface
type Lingkaran struct {
	radius float64
}

// NewLingkaranHandler handler
func NewLingkaranHandler(radius float64) Lingkaran {
	return Lingkaran{radius}
}

// Luas is receiver method for Lingkaran struct
func (c *Lingkaran) Luas() float64 {
	return math.Pi * c.radius * c.radius
}

// Keliling is receiver method for Lingkaran struct
func (c *Lingkaran) Keliling() float64 {
	return 2 * math.Pi * c.radius
}
