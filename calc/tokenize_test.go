package calc_test

import (
	"testing"

	"github.com/Frank-Mayer/calc/calc"
)

func TestTokenize(t *testing.T) {
	t.Parallel()
	testTokenize(t, "1+42", []calc.TokenType{calc.TokenNumber, calc.TokenPlus, calc.TokenNumber})
	testTokenize(t, "1-42", []calc.TokenType{calc.TokenNumber, calc.TokenMinus, calc.TokenNumber})
	testTokenize(t, " 1 - 42 ", []calc.TokenType{calc.TokenNumber, calc.TokenMinus, calc.TokenNumber})
	testTokenize(t, "1*42", []calc.TokenType{calc.TokenNumber, calc.TokenMultiply, calc.TokenNumber})
	testTokenize(t, "1/42", []calc.TokenType{calc.TokenNumber, calc.TokenDivide, calc.TokenNumber})
	testTokenize(t, "-42", []calc.TokenType{calc.TokenMinus, calc.TokenNumber})
	testTokenize(t, "2*(4+2)", []calc.TokenType{calc.TokenNumber, calc.TokenMultiply, calc.TokenOpeningParenthesis, calc.TokenNumber, calc.TokenPlus, calc.TokenNumber, calc.TokenClosingParenthesis})
	testTokenize(t, "2*(4+2*3)", []calc.TokenType{calc.TokenNumber, calc.TokenMultiply, calc.TokenOpeningParenthesis, calc.TokenNumber, calc.TokenPlus, calc.TokenNumber, calc.TokenMultiply, calc.TokenNumber, calc.TokenClosingParenthesis})
	testTokenize(t, "2 * (4 + 2 * 3)", []calc.TokenType{calc.TokenNumber, calc.TokenMultiply, calc.TokenOpeningParenthesis, calc.TokenNumber, calc.TokenPlus, calc.TokenNumber, calc.TokenMultiply, calc.TokenNumber, calc.TokenClosingParenthesis})
	testTokenize(t, "2 * (4 + 2 * (3 + 4))", []calc.TokenType{calc.TokenNumber, calc.TokenMultiply, calc.TokenOpeningParenthesis, calc.TokenNumber, calc.TokenPlus, calc.TokenNumber, calc.TokenMultiply, calc.TokenOpeningParenthesis, calc.TokenNumber, calc.TokenPlus, calc.TokenNumber, calc.TokenClosingParenthesis, calc.TokenClosingParenthesis})
}

func testTokenize(t *testing.T, input string, expected []calc.TokenType) {
	t.Run(input, func(t *testing.T) {
		tokens := calc.Tokenize(input)
		compareTokens(t, tokens, expected)
	})
}

func compareTokens(t *testing.T, tokens []calc.Token, expected []calc.TokenType) {
	if len(tokens) != len(expected) {
		t.Errorf("Expected %d tokens, got %d", len(expected), len(tokens))
	}
	for i, token := range tokens {
		if token.Type&expected[i] == 0 {
			t.Errorf("Expected token %d to be of type %d, got %d", i, expected[i], token.Type)
		}
	}
}
