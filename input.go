package input

import (
	"fmt"
)

func getUserInputString(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

