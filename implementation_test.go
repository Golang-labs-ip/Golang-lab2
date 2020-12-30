package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"5 4 2 - 3 * +",
			"+ 5 * - 4 2 3"},

		{"1 2a + 3 - 4",
			"Incorrect statement..."},

		{"124.63 2412 + 386 23 3 ^ * -",
			"- + 124.63 2412 * 386 ^ 23 3"},

		{"5 32 + 3 96 * - 234 - 8 24 ^ +",
			"+ - - + 5 32 * 3 96 234 ^ 8 24"},

		{"1 2 + 3 4 5 ^ * - 6 - 7 8 * + 9 - 10 + 11 4 * -",
			"- + - + - - + 1 2 * 3 ^ 4 5 6 * 7 8 9 10 * 11 4"},

		{"10 22 ^ 3 ^",
			"^ ^ 10 22 3"},

		{"1 2",
			"dont enought args..."},

		{"10 9 - 3 +",
			"+ - 10 9 3"},
	}

	for _, test := range tests {
		assert.Equal(t, PostfixToPrefix(test.input), test.expected)
	}
}
func ExamplePostfixToPrefix() {
	res := PostfixToPrefix("1 2 + 3 -")
	fmt.Println(res)
}
