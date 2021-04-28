package phonetic

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// mapping from ints 0-9 to strings
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

func PhoneticizeInt(n int) (string, error) {
	if n < 0 {
		return "", errors.New(fmt.Sprintf("Cannot phoneticize %d. Must be positive.\n", n))
	}
	length := int(math.Log10(float64(n))) + 1
	output := make([]string, length)
	i := length - 1

	for i >= 0 {
		d := n % 10
		str, err := phoneticizeDigit(d)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Cannot phoneticize integer %d. %s", n, err.Error()))
		}
		output[i] = str
		n /= 10
		i--
	}

	return strings.Join(output, ""), nil
}

func phoneticizeDigit(n int) (string, error) {
	if n < 0 || n > 9 {
		return "", errors.New(fmt.Sprintf("Cannot phoneticize digit %d, must be in range 0-9.\n", n))
	}
	return Digits[n], nil
}
