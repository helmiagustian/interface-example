package handler

type Kotak struct {
	width, height float64
}

// NewKotakHandler handler
func NewKotakHandler(width, height float64) Kotak {
	return Kotak{width, height}
}

func (r Kotak) Luas() float64 {
	return r.width * r.height
}

func (r Kotak) Keliling() float64 {
	return 2*r.width + 2*r.height
}
