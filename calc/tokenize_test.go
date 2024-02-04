package calc_test

import (
	"strings"
	"testing"

	"github.com/Frank-Mayer/calc/calc"
)

func TestTokenize(t *testing.T) {
	t.Parallel()
	testTokenize(t, "1+42", []string{"1", "+", "42"})
	testTokenize(t, "1-42", []string{"1", "-", "42"})
	testTokenize(t, " 1 - 42 ", []string{"1", "-", "42"})
	testTokenize(t, "1*42", []string{"1", "*", "42"})
	testTokenize(t, "1/42", []string{"1", "/", "42"})
	testTokenize(t, "-42", []string{"-", "42"})
	testTokenize(t, "2*(4+2)", []string{"2", "*", "(", "4", "+", "2", ")"})
	testTokenize(t, "2*(4+2*3)", []string{"2", "*", "(", "4", "+", "2", "*", "3", ")"})
	testTokenize(t, "2 * (4 + 2 * 3)", []string{"2", "*", "(", "4", "+", "2", "*", "3", ")"})
	testTokenize(t, "2 * (4 + 2 * (3 + 4))", []string{"2", "*", "(", "4", "+", "2", "*", "(", "3", "+", "4", ")", ")"})
}

func testTokenize(t *testing.T, input string, expected []string) {
	t.Run(input, func(t *testing.T) {
		tokens := calc.Tokenize(input)
		compareTokens(t, tokens, expected)
	})
}

func compareTokens(t *testing.T, tokens, expected []string) {
	if len(tokens) != len(expected) {
		t.Errorf("Expected [%s], got [%s]", strings.Join(expected, ", "), strings.Join(tokens, ", "))
	}
	for i, token := range tokens {
		if token != expected[i] {
			t.Errorf("Expected [%s], got [%s]", strings.Join(expected, ", "), strings.Join(tokens, ", "))
		}
	}
}
