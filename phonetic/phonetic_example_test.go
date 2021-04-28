package phonetic_test

import (
	"fmt"
	"strings"

	"github.com/mattConn/kg_mattconn_2020/phonetic"
)

func ExamplePhoneticizeInt() {
	examples := []struct {
		name  string
		input []int
	}{
		{
			name:  "Simple",
			input: []int{3, 25, 209},
		},
		{
			name:  "Negative",
			input: []int{-10, 300, 5},
		},
		{
			name:  "Empty",
			input: []int{},
		},
	}

	for _, example := range examples {
		fmt.Println(example.name)
		fmt.Println(fmt.Sprint(example.input))

		output := make([]string, len(example.input))
		i := 0
		for _, n := range example.input {
			str, err := phonetic.PhoneticizeInt(n)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			output[i] = str
			i++
		}
		fmt.Println(strings.Join(output, ","))
		fmt.Println()

	}

	// Output:
	// Simple
	// [3 25 209]
	// Three,TwoFive,TwoZeroNine

	// Negative
	// [-10 300 5]
	// Cannot phoneticize -10. Must be positive.

	// Empty
	// []
	//

}
