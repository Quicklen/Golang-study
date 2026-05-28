package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.Intn(100) + 1
	attempts := 0
	var number int

	fmt.Println("Я загадал число от 1 до 100. Угадывай!")

	for {
		fmt.Print("Твоё число: ")
		fmt.Scanln(&number)
		attempts++

		if number > random {
			fmt.Println("Меньше!")
		} else if number < random {
			fmt.Println("Больше!")
		} else {
			fmt.Printf("Угадал за %d попыток! 🎉\n", attempts)
			break
		}
	}
}
