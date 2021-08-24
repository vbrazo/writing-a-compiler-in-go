package main

import (
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

	switch n.op {
	case A_ADD:
		return genAdd(leftval, rightval)
	case A_SUBTRACT:
		return genSub(leftval, rightval)
	case A_MULTIPLY:
		return genMul(leftval, rightval)
	case A_DIVIDE:
		return genDiv(leftval, rightval)
	case A_EQ:
		return genEq(leftval, rightval)
	case A_NEQ:
		return genNeq(leftval, rightval)
	case A_GT:
		return genGt(leftval, rightval)
	case A_GE:
		return genGe(leftval, rightval)
	case A_LT:
		return genLt(leftval, rightval)
	case A_LE:
		return genLe(leftval, rightval)
	case A_INTLIT:
		return genNumber(n)
	case A_IDENT:
		return genIdent(n)
	case A_ASSIGNVAL:
		return genAssignVal(n)
	case A_ASSIGN:
		return genAssign(leftval, rightval)
	default:
		log.Fatalf("Unknown AST operator %d\n", n.op)
		panic("Unknown AST operator")
	}
}
