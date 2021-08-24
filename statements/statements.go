package main

func parseStatements() {
	var n *AstNode
	for 1 == 1 {
		// match as many newlines as you want before new statement
		for isCurrentTokenNewLine() {
			matchNewLine()
		}
		matchToken(T_PRINT, "print")
		n = binExpr(0)
		interpretAST(n)
		genPrint()
		if T.token == T_EOF {
			return
		}
		matchNewLine()
	}
}
