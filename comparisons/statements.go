package main

import (
	"fmt"
	"log"
	"os"
)

func parseStatements() {
	for 1 == 1 {
		// match as many newlines as you want before new statement
		for isCurrentTokenNewLine() {
			matchNewLine()
		}
		switch T.token {
		case T_PRINT:
			printStatement()
			break
		case T_VAR:
			varDeclerationStatement()
			break
		case T_IDENT:
			assignStatement()
			break
		case T_EOF:
			return
		default:
			log.Fatalf("Unown token Line %d Column %d", Line, Column)
			os.Exit(6)
		}
		if T.token == T_EOF {
			return
		}
		OutputFile.Flush()
		matchNewLine()
	}
}

func printStatement() {
	var n *AstNode
	matchToken(T_PRINT, "print")
	n = binExpr(0)
	genPrint(interpretAST(n))
}

func assignStatement() {
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
	tree = makeAstNode(A_ASSIGN, left, right, 0, -1)
	fmt.Fprint(OutputFile, interpretAST(tree))
	OutputFile.Flush()

}
