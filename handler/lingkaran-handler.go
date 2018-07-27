package handler

import "math"

type lingkaran struct {
	radius float64
}

// NewLingkaranHandler handler
func NewLingkaranHandler(radius float64) lingkaran {
	return lingkaran{radius}
}

// Luas is receiver method for Lingkaran struct, implementing geometry interface, pointer receiver
func (c *lingkaran) Luas() float64 { // method on pointer
	return math.Pi * c.radius * c.radius
}

// Keliling is receiver method for Lingkaran struct, implementing geometry interface, pointer receiver
func (c *lingkaran) Keliling() float64 { // method on pointer
	return 2 * math.Pi * c.radius
}
