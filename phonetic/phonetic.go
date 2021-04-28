package phonetic

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// Mapping from ints 0-9 to strings
type PhoneticDigits [10]string

var Digits PhoneticDigits = [10]string{
	"Zero",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
}

// PhoneticizeInt will take an integer and return a string without spaces
// representing the phonetic equivalent of that integer, as well as an error.
func PhoneticizeInt(n int) (string, error) {
	if n < 0 {
		return "", errors.New(fmt.Sprintf("Cannot phoneticize %d. Must be positive.\n", n))
	}

	// Lenght of integer.
	length := int(math.Log10(float64(n))) + 1

	output := make([]string, length)

	for i := length - 1; i >= 0; i-- {
		// Get last/least significant digit.
		d := n % 10

		str, _ := phoneticizeDigit(d)

		output[i] = str
		n /= 10
	}

	return strings.Join(output, ""), nil
}

// Look up digit in Digits array.
func phoneticizeDigit(n int) (string, error) {
	if n < 0 || n > 9 {
		return "", errors.New(fmt.Sprintf("Cannot phoneticize digit %d, must be in range 0-9.\n", n))
	}
	return Digits[n], nil
}
