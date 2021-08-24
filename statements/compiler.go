package main

import (
	"fmt"
	"log"
)

var AstOps = []string{"+", "-", "*", "/"}

// Given an AST, interpret the
// operators in it and return
// a final value.
func interpretAST(n *AstNode) string {
	var leftval, rightval string

	// Get the left and right sub-tree values
	if n.left != nil {
		leftval = interpretAST(n.left)
	}
	if n.right != nil {
		rightval = interpretAST(n.right)
	}

	// Debug: Print what we are about to do
	if n.op == A_INTLIT {
		fmt.Printf("int %d\n", n.intval)
	} else {
		fmt.Printf("%s %s %s\n", leftval, AstOps[n.op], rightval)
	}
	switch n.op {
	case A_ADD:
		return genAdd(leftval, rightval)
	case A_SUBTRACT:
		return genSub(leftval, rightval)
	case A_MULTIPLY:
		return genMul(leftval, rightval)
	case A_DIVIDE:
		return genDiv(leftval, rightval)
	case A_INTLIT:
		return genNumber(n)
	default:
		log.Fatalf("Unknown AST operator %d\n", n.op)
		panic("Unknown AST operator")
	}
}
