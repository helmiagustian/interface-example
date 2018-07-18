package main

import (
	"fmt"

	"github.com/helmiagustian/interface-example/handler"
)

// geometry interface
type geometry interface {
	Luas() float64
	Keliling() float64
}

// ukuran function
func ukuran(g geometry) {
	fmt.Printf("Luasnya adalah %v\n", g.Luas())
	fmt.Printf("Kelilingnya adalah %v\n", g.Keliling())
}

func main() {

	fmt.Println("Kotak 1")
	k1 := handler.NewKotakHandler(3, 4)
	ukuran(k1)

	fmt.Println("Kotak 2")
	k2 := handler.Kotak{
		Height: 6,
		Width:  5,
	}
	ukuran(k2)

	fmt.Println("Lingkaran")
	l := handler.NewLingkaranHandler(5)
	ukuran(&l)

	fmt.Println("Segitiga 1")
	var st1 handler.SegiTiga
	st1 = handler.SegiTiga{5}
	ukuran(st1)

	// example not implementing
	fmt.Println("Segitiga 2")
	st2 := handler.SegiTiga{10}
	fmt.Printf("Luasnya adalah %v\n", st2.Luas())
	fmt.Printf("Kelilingnya adalah %v\n", st2.Keliling())
}
