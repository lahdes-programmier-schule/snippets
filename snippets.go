package snippets

import (
	"fmt"
	"strconv"
)

func GetUserInputString(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

func GetUserInputNumber(prompt string) int32 {
    input := GetUserInputString(prompt)
    i, err := strconv.Atoi(s)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return i
}
