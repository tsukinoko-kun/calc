package calc

// tokenize takes an equation as a string and returns a slice of elements in the equation
// Example: "1+42" -> ["1", "+", "42"]
func Tokenize(input string) []string {
	tokens := []string{}
	currentToken := ""
	for _, char := range input {
		switch char {
		case ' ':
			continue
		case '+', '-', '*', '/', '(', ')':
			if currentToken != "" {
				tokens = append(tokens, currentToken)
			}
			tokens = append(tokens, string(char))
			currentToken = ""
		default:
			currentToken += string(char)
		}
	}
	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}
	return tokens
}
