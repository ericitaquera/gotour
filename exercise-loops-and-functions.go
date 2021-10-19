package main

import "fmt"

func Sqrt(x float64) float64 {
	i := 1
	root := 0.0
	for z := 1.0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Println(z)
		if z == root {
			return root
		}
		root = z
	}
	return root
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(4))
}
