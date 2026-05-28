package main

import (
	"fmt"
)

func converter(unit string, temp float64) (float64, string) {
	switch unit {
	case "C", "c":
		return temp*9.0/5.0 + 32, "F"
	case "F", "f":
		return (temp - 32) * 5.0 / 9.0, "C"
	case "K", "k":
		return temp - 273.15, "C"
	default:
		fmt.Println("Ошибка: используйте C, F или K")
		return 0, ""
	}
}

func main() {
	var unit string
	var temp float64

	fmt.Print("Введите температуру: ")
	fmt.Scanf("%f %s", &temp, &unit)

	result, targetUnit := converter(unit, temp)

	if targetUnit != "" {
		fmt.Printf("%.2f %s\n", result, targetUnit)
	}
}
