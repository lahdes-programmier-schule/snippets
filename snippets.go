package snippets

import (
	"fmt"
	"os"
	"strconv"
)

func GetUserInputString(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

func GetUserInputNumber(prompt string) int {
	input := GetUserInputString(prompt)
	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return int(i)
}

func GetUserInputNumberInt32(prompt string) int32 {
	input := GetUserInputString(prompt)
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return int32(i)
}

func GetUserInputNumberInt64(prompt string) int64 {
	input := GetUserInputString(prompt)
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}

func GetUserInputNumberUint32(prompt string) uint32 {
	input := GetUserInputString(prompt)
	i, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return uint32(i)
}

func GetUserInputNumberUint64(prompt string) uint64 {
	input := GetUserInputString(prompt)
	i, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}
