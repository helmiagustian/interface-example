package handler

// Kotak struct type, implementing geometry interface
type Kotak struct {
	width, height float64
}

// NewKotakHandler handler
func NewKotakHandler(width, height float64) Kotak {
	return Kotak{width, height}
}

// Luas is receiver method for Kotak struct
func (r Kotak) Luas() float64 {
	return r.width * r.height
}

// Keliling is receiver method for Kotak struct
func (r Kotak) Keliling() float64 {
	return 2*r.width + 2*r.height
}
