package handler

import "math"

// Lingkaran struct type
type Lingkaran struct {
	radius float64
}

// NewLingkaranHandler handler
func NewLingkaranHandler(radius float64) Lingkaran {
	return Lingkaran{radius}
}

// Luas is receiver method for Lingkaran struct, implementing geometry interface
func (c *Lingkaran) Luas() float64 { // method on pointer
	return math.Pi * c.radius * c.radius
}

// Keliling is receiver method for Lingkaran struct, implementing geometry interface
func (c *Lingkaran) Keliling() float64 { // method on pointer
	return 2 * math.Pi * c.radius
}
