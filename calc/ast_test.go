package calc_test

import (
	"testing"

	"github.com/Frank-Mayer/calc/calc"
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
