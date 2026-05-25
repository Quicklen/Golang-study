package main

import (
	"fmt"
	"unicode"
)

func checkpass(str string) bool {
	hasupper := false
	hasdigit := false
	hassymbol := false
	var lenght int = len(str)
	if lenght < 8 {
		return false
	}
	for _, char := range str {
		if unicode.IsUpper(char) {
			hasupper = true
		}
		if unicode.IsDigit(char) {
			hasdigit = true
		}
		if unicode.IsLower(char) {
			hassymbol = true
		}
	}

	if hasupper && hasdigit && hassymbol {
		return true
	} else {
		return false
	}

}

func main() {
	var str string = "kaEina288"

	if checkpass(str) == false {
		fmt.Println("Pass bad")
	} else {
		fmt.Println("Pass good")
	}
}
