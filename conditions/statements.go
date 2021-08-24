package main

import (
	"fmt"
	"log"
	"os"
)

func compundStatement() *AstNode {
	var left *AstNode = nil
	var tree *AstNode

	// match as many newlines as you want before new statement
	for isCurrentTokenNewLine() {
		matchNewLine()
	}
	matchLBrace()

	for 1 == 1 {
		// match as many newlines as you want before new statement
		for isCurrentTokenNewLine() {
			matchNewLine()
		}
		switch T.token {
		case T_PRINT:
			tree = printStatement()
			break
		case T_VAR:
			varDeclerationStatement()
			tree = nil
			break
		case T_IDENT:
			tree = assignStatement()
			break
		case T_IF:
			tree = ifStatement()
			break
		case T_RBRACE:
			// When we hit a right curly bracket,
			// skip past it and return the AST
			matchRBrace()
			return (left)
		default:
			log.Fatalf("Unown token Line %d Column %d", Line, Column)
			os.Exit(6)
		}
		OutputFile.Flush()

		if tree != nil {
			if left == nil {
				left = tree
			} else {
				left = makeAstNode(A_GLUETO, left, nil, tree, -1, -1)
			}
		}
	}
	return nil
}

func ifStatement() *AstNode {
	var conditionAst, trueAst, falseAst *AstNode
	matchToken(T_IF, "if")
	matchLParen()
	conditionAst = binExpr(0)

	if conditionAst.op < A_EQ || conditionAst.op > A_GE {
		log.Fatalf("Condition is not returning a boolean on Line %d", Line)
	}
	matchRParen()
	trueAst = compundStatement()
	if T.token == T_ELSE {
		scan(&T)
		falseAst = compundStatement()
	}
	return makeAstNode(A_IF, conditionAst, trueAst, falseAst, -1, -1)
}

func printStatement() *AstNode {
	var tree *AstNode
	matchToken(T_PRINT, "print")
	tree = binExpr(0)
	return makeAstUnary(A_PRINT, tree, -1, -1)
}

func assignStatement() *AstNode {
	var left, right, tree *AstNode
	matchIdent()
	id, err := findGlobalSymbol(LastScannedIdentifier)
	if err != nil {
		log.Fatalf("Couldn't find global Symbol %s", LastScannedIdentifier)
		os.Exit(7)
	}
	left = makeLeaf(A_ASSIGNVAL, -1, id)
	matchToken(T_ASSIGN, "=")
	right = binExpr(0)
	tree = makeAstNode(A_ASSIGN, left, nil, right, 0, -1)
	fmt.Fprint(OutputFile, interpretAST(tree))
	OutputFile.Flush()

	return tree

}
