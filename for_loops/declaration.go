package main

func varDeclerationStatement() *AstNode {
	matchToken(T_VAR, "v")
	matchIdent()
	id := addGlobalSymbol(LastScannedIdentifier)
	genGlobalSymbol()

	if T.token == T_ASSIGN {
		matchToken(T_ASSIGN, "=")
		var left, right, tree *AstNode
		left = makeLeaf(A_ASSIGNVAL, -1, id)
		right = binExpr(0)
		tree = makeAstNode(A_ASSIGN, left, nil, right, 0, -1)
		return tree
	}
	return nil
}
