package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Frank-Mayer/calc/calc"
)

func main() {
	inputList := os.Args[1:]
	root, err := calc.Ast(calc.Tokenize(strings.Join(inputList, "")))
	if err != nil {
		fmt.Println(err)
		return
	}
	if v, err := root.Eval(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
