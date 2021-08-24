package main

import (
	"fmt"
	"log"
	"os"
)

var AstOps = []string{"+", "-", "*", "/"}

// Given an AST, interpret the
// operators in it and return
// a final value.
func interpretAST(n *AstNode) int {
	var leftval, rightval int

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
		fmt.Printf("%d %s %d\n", leftval, AstOps[n.op], rightval)
	}

	switch n.op {
	case A_ADD:
		return (leftval + rightval)
	case A_SUBTRACT:
		return (leftval - rightval)
	case A_MULTIPLY:
		return (leftval * rightval)
	case A_DIVIDE:
		return (leftval / rightval)
	case A_INTLIT:
		return n.intval
	default:
		log.Fatalf("Unknown AST operator %d\n", n.op)
		os.Exit(5)
	}
	return -1
}
