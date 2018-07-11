package main

import (
	"fmt"

	"github.com/helmiagustian/interface-example/handler"
)

type geometry interface {
	Luas() float64
	Keliling() float64
}

func ukuran(g geometry) {
	fmt.Println(g)
	fmt.Println(g.Luas())
	fmt.Println(g.Keliling())
}

func main() {
	r := handler.NewKotakHandler(3, 4)
	ukuran(r)

	fmt.Printf("\n")

	c := handler.NewLingkaranHandler(5)
	ukuran(&c)
}
