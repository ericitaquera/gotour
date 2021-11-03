package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([]uint8, dx, dx)
	b := make([][]uint8, dy, dy)
	printSlice("a", a)
	for i := range b {
		println(i)
		for j, v := range a {
			v = uint8((i + j) / 2)
			println(j, v)
			a[i] = v
		}
		b[i] = a
	}
	fmt.Println(a, b)

	return b
}

func main() {
	pic.Show(Pic)
	//println ("teste")
	//Pic(10,10)
}

func printSlice(s string, x []uint8) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
