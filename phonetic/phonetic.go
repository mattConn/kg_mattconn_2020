package phonetic

// mapping from ints 0-9 to strings
type PhoneticDigits []string

func NewDigits() PhoneticDigits {
	return []string{
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
}

func (p PhoneticDigits) PhoneticizeInt(n int) (string, error) {
	return "", nil
}
func (p PhoneticDigits) PhoneticizeDigit(n int) (string, error) {
	return "", nil
}
