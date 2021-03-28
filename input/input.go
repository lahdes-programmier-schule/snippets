package input

import (
	"fmt"
)

func GetUserInputString(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

