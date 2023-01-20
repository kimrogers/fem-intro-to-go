package utils

import (
	"fmt"
	"strings"
)

func printNum(num int) {
	fmt.Println("Current Number:", num)
}

// Add adds multiple numbers together, returns the total
func Add(nums ...int) int {
	total := 0
	for _, value := range nums {
		printNum(value)
		total += value
	}
	return total
}

func MakeExcited(sentence string) string {
	return strings.ToUpper(sentence) + "!"
}
