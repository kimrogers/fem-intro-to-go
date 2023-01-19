package main

import "fmt"

func main() {
	var sentence = "Learning Go"

	for index, value := range sentence {
		if index%2 == 1 {
			fmt.Println(string(value))
		}
	}
}
