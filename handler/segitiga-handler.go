package handler

import "math"

// SegiTiga struct type
type SegiTiga struct {
	Sisi float64
}

// Luas is receiver method for SegiTiga struct, implementing geometry interface
func (r SegiTiga) Luas() float64 { // method on value
	return (0.5 * r.Sisi) * math.Sqrt((r.Sisi*r.Sisi)-(0.5*r.Sisi*0.5*r.Sisi))
}

// Keliling is receiver method for SegiTiga struct, implementing geometry interface
func (r SegiTiga) Keliling() float64 { // method on value
	return r.Sisi * 3
}
