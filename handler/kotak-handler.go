package handler

// Kotak struct type
type Kotak struct {
	Width, Height float64
}

// NewKotakHandler handler
func NewKotakHandler(width, height float64) Kotak {
	return Kotak{width, height}
}

// Luas is receiver method for Kotak struct, implementing geometry interface, value receiver
func (r Kotak) Luas() float64 { // method on value
	return r.Width * r.Height
}

// Keliling is receiver method for Kotak struct, implementing geometry interface, value receiver
func (r Kotak) Keliling() float64 { // method on value
	return 2*r.Width + 2*r.Height
}
