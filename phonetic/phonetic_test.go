package phonetic_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mattConn/kg_mattconn_2020/phonetic"
)

func TestPhoneticizeInt(t *testing.T) {
	tests := []struct {
		input []int
		want  string
	}{
		{
			input: []int{3, 25, 209},
			want:  "Three,TwoFive,TwoZeroNine",
		},
		{
			input: []int{10, 300, 5},
			want:  "OneZero,ThreeZeroZero,Five",
		},
	}

	for _, test := range tests {
		output := make([]string, len(test.input))
		i := 0
		for _, n := range test.input {
			val, err := phonetic.PhoneticizeInt(n)
			if err != nil {
				t.Error(err)
				break
			}

			output[i] = val
			i++
		}

		outputStr := strings.Join(output, ",")
		if outputStr != test.want {
			t.Errorf("Incorrect phoneticization.\nInput: %s\nWanted: %s\nGot: %s\n", fmt.Sprint(test.input), test.want, outputStr)
		}
	}
}
