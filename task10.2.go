package main

import (
	"fmt"
)

func main() {

	fmt.Println("Проблемы округления процентов")

	fmt.Println("Введите сумму вклада:")
	var summ float64
	fmt.Scan(&summ)
	summ = summ * 100

	fmt.Println("Введите ежемесячный процент капитализации:")
	var percent float64
	fmt.Scan(&percent)
	percent /= 100

	fmt.Println("Введите на сколько лет вклад:")
	var month int
	fmt.Scan(&month)
	month *= 12

	remainder := 0.0
	prevSumm := 0.0
	fmt.Println("--------------------")

	for i := 0; i < month; i++ {
		summ += summ * percent
		prevSumm = summ
		summ = float64(int(summ))
		remainder += prevSumm - summ
	}

	fmt.Println("Сумма вкладчика", summ/100)
	fmt.Println("Сумма банка", remainder/100)
}
