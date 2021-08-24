package main

import (
	"log"
	"os"
)

func singleStatement() *AstNode {
	var tree *AstNode = nil
	switch T.token {
	case T_PRINT:
		tree = printStatement()
		break
	case T_VAR:
		tree = varDeclerationStatement()
		break
	case T_FUNC:
		tree = functionDeclaration()
		break
	case T_IDENT:
		tree = assignStatement()
		break
	case T_IF:
		tree = ifStatement()
		break
	case T_WHILE:
		tree = whileStatement()
		break
	case T_FOR:
		tree = forStatement()
		break
	default:
		log.Fatalf("Unknown token Line %d Column %d", Line, Column)
	}
	return tree
}

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
		tree = singleStatement()

		if tree != nil {
			if left == nil {
				left = tree
			} else {
				left = makeAstNode(A_GLUETO, left, nil, tree, -1, -1)
			}
		}

		for isCurrentTokenNewLine() {
			matchNewLine()
		}

		if T.token == T_RBRACE {
			matchRBrace()
			return left
		}

	}
	return nil
}

func forStatement() *AstNode {
	var conditionAst, bodyAst, preOpAst, postOpAst, tree *AstNode
	matchToken(T_FOR, "for")
	matchLParen()
	preOpAst = singleStatement()
	matchSemi()

	conditionAst = binExpr(0)
	if conditionAst.op < A_EQ || conditionAst.op > A_GE {
		log.Fatalf("Condition is not returning a boolean on Line %d", Line)
	}

	matchSemi()

	postOpAst = singleStatement()

	matchRParen()
	bodyAst = compundStatement()

	// Glue the compound statement and the postop tree
	tree = makeAstNode(A_GLUETO, bodyAst, nil, postOpAst, -1, -1)

	// Make a WHILE loop with the condition and this new body
	tree = makeAstNode(A_WHILE, conditionAst, tree, nil, -1, -1)

	// And glue the preop tree to the A_WHILE tree
	return (makeAstNode(A_GLUETO, preOpAst, nil, tree, -1, -1))

}

func whileStatement() *AstNode {
	var conditionAst, bodyAst *AstNode
	matchToken(T_WHILE, "while")
	matchLParen()
	conditionAst = binExpr(0)

	if conditionAst.op < A_EQ || conditionAst.op > A_GE {
		log.Fatalf("Condition is not returning a boolean on Line %d", Line)
	}
	matchRParen()
	bodyAst = compundStatement()
	return makeAstNode(A_WHILE, conditionAst, bodyAst, nil, -1, -1)
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

	return tree

}
