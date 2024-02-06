package calc

import "fmt"

type TokenType uint8

const (
	TokenInvalid            TokenType = 0
	TokenNumber             TokenType = 1
	TokenOpeningParenthesis TokenType = 2
	TokenClosingParenthesis TokenType = 3
	TokenOperator           TokenType = 128
	TokenPlus               TokenType = TokenOperator + 1
	TokenMinus              TokenType = TokenOperator + 2
	TokenMultiply           TokenType = TokenOperator + 3
	TokenDivide             TokenType = TokenOperator + 4
)

type Token struct {
	Type TokenType
	Str  string
	Pos  int
}

func (t Token) String() string {
	return fmt.Sprintf("%d(%s)", t.Type, t.Str)
}

// tokenize takes an equation as a string and returns a slice of elements in the equation
// Example: "1+42" -> ["1", "+", "42"]
func Tokenize(input string) []Token {
	tokens := []Token{}
	currentToken := ""
	for i, char := range input {
		switch char {
		case ' ':
			continue
		case '+':
			if currentToken != "" {
				tokens = append(tokens, Token{Type: TokenNumber, Str: currentToken, Pos: i - len(currentToken)})
			}
			tokens = append(tokens, Token{Type: TokenPlus, Str: "+", Pos: i})
			currentToken = ""
		case '-':
			if currentToken != "" {
				tokens = append(tokens, Token{Type: TokenNumber, Str: currentToken, Pos: i - len(currentToken)})
			}
			tokens = append(tokens, Token{Type: TokenMinus, Str: "-", Pos: i})
			currentToken = ""
		case '*':
			if currentToken != "" {
				tokens = append(tokens, Token{Type: TokenNumber, Str: currentToken, Pos: i - len(currentToken)})
			}
			tokens = append(tokens, Token{Type: TokenMultiply, Str: "*", Pos: i})
			currentToken = ""
		case '/':
			if currentToken != "" {
				tokens = append(tokens, Token{Type: TokenNumber, Str: currentToken, Pos: i - len(currentToken)})
			}
			tokens = append(tokens, Token{Type: TokenDivide, Str: "/", Pos: i})
			currentToken = ""
		case '(':
			if currentToken != "" {
				tokens = append(tokens, Token{Type: TokenNumber, Str: currentToken, Pos: i - len(currentToken)})
			}
			tokens = append(tokens, Token{Type: TokenOpeningParenthesis, Str: "(", Pos: i})
			currentToken = ""
		case ')':
			if currentToken != "" {
				tokens = append(tokens, Token{Type: TokenNumber, Str: currentToken, Pos: i - len(currentToken)})
			}
			tokens = append(tokens, Token{Type: TokenClosingParenthesis, Str: ")", Pos: i})
			currentToken = ""
		default:
			currentToken += string(char)
		}
	}
	if currentToken != "" {
		tokens = append(tokens, Token{Type: TokenNumber, Str: currentToken, Pos: len(input) - len(currentToken)})
	}
	return tokens
}
