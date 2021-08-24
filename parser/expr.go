package main

import (
	"log"
	"os"
)

func getAstType(tok int) int {
	switch tok {
	case T_PLUS:
		return (A_ADD)
	case T_MINUS:
		return (A_SUBTRACT)
	case T_STAR:
		return (A_MULTIPLY)
	case T_SLASH:
		return (A_DIVIDE)
	default:
		log.Fatalf("unknown token in getAstType() on line %d column %d \n", Line, Column)
		os.Exit(3)
	}
	return -1
}

func primary() AstNode {
	var n AstNode

	switch T.token {
	case T_INTLIT:
		n = makeLeaf(A_INTLIT, T.intvalue)
		scan(&T)
		return n
	default:
		log.Fatalf("Syntax error on line %d column %d ", Line, Column)
		os.Exit(4)
	}
	return n
}

func binExpr() AstNode {
	var n, left, right AstNode
	var nodeType int

	//Get interger literal left
	// + get next token
	left = primary()

	if T.token == T_EOF {
		return left
	}
	//token to node
	nodeType = getAstType(T.token)
	//next token
	scan(&T)

	//get the right node - recursively
	right = binExpr()

	//build the ast tree
	n = makeAstNode(nodeType, &left, &right, 0)
	return n
}
