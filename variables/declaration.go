package main

func varDeclerationStatement() {
	matchToken(T_VAR, "v")
	matchIdent()
	id := addGlobalSymbol(LastScannedIdentifier)
	genGlobalSymbol()

	if T.token == T_EQ {
		matchToken(T_EQ, "=")
		var left, right, tree *AstNode
		left = makeLeaf(A_ASSIGNVAL, -1, id)
		right = binExpr(0)
		tree = makeAstNode(A_ASSIGN, left, right, 0, -1)
		interpretAST(tree)
	}
}
