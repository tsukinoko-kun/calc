package calc

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// Ast takes a slice of tokens and returns an abstract syntax tree
func Ast(tokens []string) (AstNode, error) {
	nodes := make([]AstNode, len(tokens))
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token {
		case "+":
			nodes[i] = &plusNode{}
		case "-":
			nodes[i] = &minusNode{}
		case "*":
			nodes[i] = &timesNode{}
		case "/":
			nodes[i] = &divideNode{}
		case "(":
			// find the matching closing bracket
			bracketCount := 1
			for j := i + 1; j < len(tokens); j++ {
				if tokens[j] == "(" {
					bracketCount++
				} else if tokens[j] == ")" {
					bracketCount--
					if bracketCount == 0 {
						subTree, err := Ast(tokens[i+1 : j])
						if err != nil {
							return nil, err
						}
						nodes[i] = &bracketNode{subTree}
						i = j
						break
					}
				}
			}
			if bracketCount != 0 {
				return nil, fmt.Errorf("Unmatched (")
			}
		case ")":
			return nil, fmt.Errorf("Unexpected )")
		default:
			number, err := ParseNumber(token)
			if err != nil {
				return nil, err
			}
			nodes[i] = &numberNode{number}
		}
	}

	l := len(nodes)

	// handle multiplication and division
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == nil {
			continue
		}
		switch node := nodes[i].(type) {
		case *timesNode:
			if i == 0 {
				return nil, fmt.Errorf("Unexpected * at beginning of input")
			} else if i == l-1 {
				return nil, fmt.Errorf("Unexpected end of input after *")
			} else {
				if prev, err := prevNode(nodes, i); err != nil {
					return nil, errors.Wrap(err, "Failed to find previous node from *")
				} else {
					node.left = prev
				}
				if next, err := nextNode(nodes, i); err != nil {
					return nil, errors.Wrap(err, "Failed to find next node from *")
				} else {
					node.right = next
				}
			}
		case *divideNode:
			if i == 0 {
				return nil, fmt.Errorf("Unexpected / at beginning of input")
			} else if i == l-1 {
				return nil, fmt.Errorf("Unexpected end of input after /")
			} else {
				if prev, err := prevNode(nodes, i); err != nil {
					return nil, errors.Wrap(err, "Failed to find previous node from /")
				} else {
					node.left = prev
				}
				if next, err := nextNode(nodes, i); err != nil {
					return nil, errors.Wrap(err, "Failed to find next node from /")
				} else {
					node.right = next
				}
			}
		}
	}

	// handle addition and subtraction
	for i := 0; i < l; i++ {
		if nodes[i] == nil {
			continue
		}
		switch node := nodes[i].(type) {
		case *plusNode:
			if i == l-1 {
				return nil, fmt.Errorf("Unexpected end of input after +")
			} else {
				if i == 0 {
					node.left = &numberNode{0}
				} else {
					if prev, err := prevNode(nodes, i); err != nil {
						node.left = &numberNode{0}
					} else {
						node.left = prev
					}
				}
				if next, err := nextNode(nodes, i); err != nil {
					return nil, errors.Wrap(err, "Failed to find next node from +")
				} else {
					node.right = next
				}
			}
		case *minusNode:
			if i == l-1 {
				return nil, fmt.Errorf("Failed to find next node from -")
			} else {
				if i == 0 {
					node.left = &numberNode{0}
				} else {
					if prev, err := prevNode(nodes, i); err != nil {
						node.left = &numberNode{0}
					} else {
						node.left = prev
					}
				}
				if next, err := nextNode(nodes, i); err != nil {
					return nil, errors.Wrap(err, "Failed to find next node from -")
				} else {
					node.right = next
				}
			}
		}
	}

	// find the root node
	var root AstNode
	c := 0
	for i := 0; i < l; i++ {
		if nodes[i] != nil {
			root = nodes[i]
			c++
		}
	}
	switch c {
	case 0:
		return nil, fmt.Errorf("No nodes found")
	case 1:
		return root, nil
	default:
		return nil, fmt.Errorf("More than one root node found")
	}
}

// prevNode tries to find the previous node in the slice of nodes that is not nil
func prevNode(nodes []AstNode, i int) (AstNode, error) {
	for j := i - 1; j >= 0; j-- {
		if nodes[j] != nil {
			node := nodes[j]
			nodes[j] = nil
			return node, nil
		}
	}
	return nil, fmt.Errorf("Unexpected beginning of input")
}

// nextNode tries to find the next node in the slice of nodes that is not nil
func nextNode(nodes []AstNode, i int) (AstNode, error) {
	for j := i + 1; j < len(nodes); j++ {
		if nodes[j] != nil {
			node := nodes[j]
			nodes[j] = nil
			return node, nil
		}
	}
	return nil, fmt.Errorf("Unexpected end of input")
}

func ParseNumber(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

type AstNode interface {
	Eval() (float64, error)
}

type numberNode struct {
	value float64
}

func (n *numberNode) Eval() (float64, error) {
	return n.value, nil
}

type plusNode struct {
	left  AstNode
	right AstNode
}

func (n *plusNode) Eval() (float64, error) {
	if n.left == nil {
		return 0, fmt.Errorf("Left side of + is nil")
	}
	a, err := n.left.Eval()
	if err != nil {
		return 0, err
	}
	if n.right == nil {
		return 0, fmt.Errorf("Right side of + is nil")
	}
	b, err := n.right.Eval()
	if err != nil {
		return 0, err
	}
	return a + b, nil
}

type minusNode struct {
	left  AstNode
	right AstNode
}

func (n *minusNode) Eval() (float64, error) {
	if n.left == nil {
		return 0, fmt.Errorf("Left side of - is nil")
	}
	a, err := n.left.Eval()
	if err != nil {
		return 0, err
	}
	if n.right == nil {
		return 0, fmt.Errorf("Right side of - is nil")
	}
	b, err := n.right.Eval()
	if err != nil {
		return 0, err
	}
	return a - b, nil
}

type timesNode struct {
	left  AstNode
	right AstNode
}

func (n *timesNode) Eval() (float64, error) {
	if n.left == nil {
		return 0, fmt.Errorf("Left side of * is nil")
	}
	a, err := n.left.Eval()
	if err != nil {
		return 0, err
	}
	if n.right == nil {
		return 0, fmt.Errorf("Right side of * is nil")
	}
	b, err := n.right.Eval()
	if err != nil {
		return 0, err
	}
	return a * b, nil
}

type divideNode struct {
	left  AstNode
	right AstNode
}

func (n *divideNode) Eval() (float64, error) {
	if n.left == nil {
		return 0, fmt.Errorf("Left side of / is nil")
	}
	a, err := n.left.Eval()
	if err != nil {
		return 0, err
	}
	if n.right == nil {
		return 0, fmt.Errorf("Right side of / is nil")
	}
	b, err := n.right.Eval()
	if err != nil {
		return 0, err
	}
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}

type bracketNode struct {
	subTree AstNode
}

func (n *bracketNode) Eval() (float64, error) {
	return n.subTree.Eval()
}
