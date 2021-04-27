package main

import (
	"fmt"
	"os"
)

// 1. get input
// 2. check if input exists
// 3. check if each string is a valid integer
// 4. slice each string, getting phonetic equivalent, store in array
// 5. print

func main() {
	input := os.Args[1:]
	if len(input) == 0 {
		fmt.Println("Missing arguments.")
	}
}
