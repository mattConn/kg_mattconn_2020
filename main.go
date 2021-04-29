package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mattConn/kg_mattconn_2020/phonetic"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Missing arguments.")
		return
	}

	output := make([]string, len(args))
	i := 0
	for _, input := range args {
		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Invalid argument: %s\n", input)
			return
		}

		output[i], err = phonetic.PhoneticizeInt(n)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		i++
	}

	fmt.Println(strings.Join(output, ","))
}
