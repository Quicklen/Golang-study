package main

import "fmt"

func main() {
	var number int
	var sum int
	var count int
	var min, max int
	first := true

	fmt.Println("Вводи числа (0 для завершения):")

	for {
		fmt.Scan(&number)

		if number == 0 {
			break
		}

		if first {
			min = number
			max = number
			first = false
		}

		if number > max {
			max = number
		}
		if number < min {
			min = number
		}

		sum += number
		count++
	}

	if count == 0 {
		fmt.Println("Не введено ни одного числа")
		return
	}

	mean := float64(sum) / float64(count)

	fmt.Printf("Количество: %d\n", count)
	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Min: %d\n", min)
	fmt.Printf("Max: %d\n", max)
	fmt.Printf("Среднее арифметическое: %.2f\n", mean)
}
