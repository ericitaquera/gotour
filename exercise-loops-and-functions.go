package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	margin := 0.000000000001

	for {
		diff := math.Abs(math.Pow(z, 2) - x)
		if diff < margin {
			return z
		}
		z -= (z*z - x) / (2 * z)
	}

	return z

}

func main() {
	fmt.Println("minha:", Sqrt(3))
	fmt.Println("math", math.Sqrt(3))
	//fmt.Println(Sqrt(1))
	//fmt.Println(Sqrt(3))
	//fmt.Println(Sqrt(4))
}
