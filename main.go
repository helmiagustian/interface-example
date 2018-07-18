package main

import (
	"fmt"
	"math"

	"github.com/helmiagustian/interface-example/handler"
)

// geometry interface
type geometry interface {
	Luas() float64
	Keliling() float64
}

// SegiTiga struct type
type SegiTiga struct {
	sisi float64
}

// ukuran function
func ukuran(g geometry) {
	fmt.Printf("Luasnya adalah %v\n", g.Luas())
	fmt.Printf("Kelilingnya adalah %v\n", g.Keliling())
}

func main() {

	fmt.Println("Kotak")
	r := handler.NewKotakHandler(3, 4)
	ukuran(r)

	fmt.Println("Lingkaran")
	c := handler.NewLingkaranHandler(5)
	ukuran(&c)

	fmt.Println("Segitiga")
	st := SegiTiga{5}
	ukuran(st)
}

// Luas is receiver method for SegiTiga struct, implementing geometry interface
func (r SegiTiga) Luas() float64 { // method on value
	return (0.5 * r.sisi) * math.Sqrt((r.sisi*r.sisi)-(0.5*r.sisi*0.5*r.sisi))
}

// Keliling is receiver method for SegiTiga struct, implementing geometry interface
func (r SegiTiga) Keliling() float64 { // method on value
	return r.sisi * 3
}
