package handler

// Kotak struct type
type Kotak struct {
	width, height float64
}

// NewKotakHandler handler
func NewKotakHandler(width, height float64) Kotak {
	return Kotak{width, height}
}

// Luas is receiver method for Kotak struct, implementing geometry interface
func (r Kotak) Luas() float64 { // method on value
	return r.width * r.height
}

// Keliling is receiver method for Kotak struct, implementing geometry interface
func (r Kotak) Keliling() float64 { // method on value
	return 2*r.width + 2*r.height
}
