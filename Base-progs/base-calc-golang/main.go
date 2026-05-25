package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func main() {
	var result, first, second int

	var UserInput string

	fmt.Scanln(&UserInput)
	if UserInput != "+" && UserInput != "-" && UserInput != "*" && UserInput != "/" {
		return
	}
	fmt.Scanln(&first)
	fmt.Scanln(&second)

	if UserInput == "+" {
		result = add(first, second)
	} else if UserInput == "-" {
		result = sub(first, second)
	} else if UserInput == "*" {
		result = mul(first, second)
	} else if UserInput == "/" {
		result = div(first, second)
	} else {
		fmt.Println("Fuck you dump")
		return
	}

	fmt.Printf("%d", result)
}
