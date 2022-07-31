package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	a := float64(2)
	ans := float64(0)

	for {
		a = a - (a*a-x)/(2*a)
		if math.Abs(ans-a) < 1e-15 {
			break
		}

		ans = a
	}

	return ans
}

func main() {
	fmt.Println("Square root using Sqrt: ", Sqrt(1711))
	fmt.Println("Square root using math.Sqrt: ", math.Sqrt(1711))
}
