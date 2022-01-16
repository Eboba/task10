package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Разложение ex в ряд Тейлора")

	fmt.Println("Введите значение:")
	var x float64
	fmt.Scan(&x)

	fmt.Println("Введите точность вычисления значения функции:")
	var epsilon float64
	fmt.Scan(&epsilon)
	epsilon = 1 / math.Pow(10, epsilon)

	result := 0.0
	prevResult := 0.0
	fact := 1
	n := 1

	for {

		result += math.Pow(x, float64(n)) / float64(fact)

		if n > 0 {
			fact *= n
		}

		if math.Abs(result-prevResult) < epsilon {
			fmt.Println(result)
			break
		}
		n++
		prevResult = result
	}
}
