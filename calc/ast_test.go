package calc_test

import (
	"testing"

	"github.com/tsukinoko-kun/calc/calc"
)

func TestAst(t *testing.T) {
	t.Parallel()
	testAst(t, "1+42", 1.0+42.0)
	testAst(t, "1-42", 1.0-42.0)
	testAst(t, " 1 - 42 ", 1.0-42.0)
	testAst(t, "1*42", 1.0*42.0)
	testAst(t, "1/42", 1.0/42.0)
	testAst(t, "-42", -42.0)
	testAst(t, "2*(4+2)", 2.0*(4.0+2.0))
	testAst(t, "2*(4+2*3)", 2.0*(4.0+2.0*3.0))
	testAst(t, "2 * (4 + 2 * 3)", 2.0*(4.0+2.0*3.0))
	testAst(t, "2 * (4 + 2 * (3 + 4))", 2.0*(4.0+2.0*(3.0+4.0)))
	testAst(t, "1*(2.3+4.5)", 1.0*(2.3+4.5))
	testAst(t, "6/2*(2+1)", 9.0)
	testAst(t, "6/2(2+1)", 9.0)
	testAst(t, "6/(2(2+1))", 1.0)

	testAstFail(t, "1+")
	testAstFail(t, "1+*2")
	testAstFail(t, "42/0")
}

func testAst(t *testing.T, input string, expected float64) {
	t.Run(input, func(t *testing.T) {
		root, err := calc.Ast(calc.Tokenize(input))
		if err != nil {
			t.Errorf("Error parsing %s: %s", input, err)
			return
		}
		result, err := root.Eval()
		if err != nil {
			t.Errorf("Error evaluating %s: %s", input, err)
			return
		}
		if result != expected {
			t.Errorf("%s = %f, expected %f", input, result, expected)
			return
		}
	})
}

func testAstFail(t *testing.T, input string) {
	t.Run(input, func(t *testing.T) {
		root, err := calc.Ast(calc.Tokenize(input))
		if err != nil {
			return
		}
		result, err := root.Eval()
		if err != nil {
			return
		}
		t.Errorf("Expected error parsing %s, got %f", input, result)
	})
}
