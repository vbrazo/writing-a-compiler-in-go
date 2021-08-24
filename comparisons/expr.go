package main

import (
	"log"
	"os"
)

func opPrecedence(tokenType int) int {
	prec := OpPrecedence[tokenType]
	if prec == 0 {
		log.Fatalf("Syntax / Precedence error on line %d column %d, token %d\n", Line, Column, tokenType)
		os.Exit(6)
	}
	return prec
}

// AST_OP and Token have to be at the smae plae
func getAstType(tok int) int {
	if tok > T_EOF && tok < T_INTLIT {
		return tok
	} else {
		log.Fatalf("unknown token in getAstType() on line %d column %d \n", Line, Column)
		os.Exit(3)
	}
	return -1
}

func primary() *AstNode {
	var n *AstNode

	switch T.token {
	case T_INTLIT:
		n = makeLeaf(A_INTLIT, T.intvalue, -1)
		break
	case T_IDENT:
		id, err := findGlobalSymbol(LastScannedIdentifier)
		if err != nil {
			log.Fatalf("Uknown variable %s. %v", LastScannedIdentifier, err)
		}
		n = makeLeaf(A_IDENT, -1, id)
		break
	default:
		log.Fatalf("Syntax error on line %d column %d ", Line, Column)
		os.Exit(4)
	}
	scan(&T)
	return n
}

func binExpr(tokenPrecedence int) *AstNode {
	var left, right *AstNode
	var tokenType int

	//Get interger literal left
	// + get next token
	left = primary()
	tokenType = T.token

	if tokenType == T_NEWLINE || tokenType == T_EOF {
		return left
	}
	// While the precedence of this token is
	// more than that of the previous token precedence
	for opPrecedence(tokenType) > tokenPrecedence {
		//next integer token -> current one cant be integer
		scan(&T)

		//get the right node - recursively
		right = binExpr(OpPrecedence[tokenType])

		// Join that sub-tree with ours. Convert the token
		// into an AST operation at the same time.
		left = makeAstNode(getAstType(tokenType), left, right, 0, -1)

		// Update the details of the current token.
		// If no tokens left, return just the left node
		tokenType = T.token
		if tokenType == T_NEWLINE || tokenType == T_EOF {
			return left
		}
	}
	return left
}
