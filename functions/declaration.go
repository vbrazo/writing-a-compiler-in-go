package main

func varDeclerationStatement() *AstNode {
	matchToken(T_VAR, "v")
	matchIdent()
	id := addGlobalSymbol(LastScannedIdentifier)
	genGlobalSymbol("int")

	if T.token == T_ASSIGN {
		matchToken(T_ASSIGN, "=")
		var left, right, tree *AstNode
		left = makeLeaf(A_ASSIGNVAL, -1, id)
		right = binExpr(0)
		tree = makeAstNode(A_ASSIGN, left, nil, right, 0, -1)
		return tree
	} else if T.token == T_FUNC {
		return nil
	}
	return nil
}

func functionDeclaration() *AstNode {
	matchToken(T_FUNC, "f")
	matchIdent()
	id := addGlobalSymbol(LastScannedIdentifier)
	genGlobalSymbol("func()")
	matchLParen()
	matchRParen()
	tree := compundStatement()

	return makeAstUnary(A_FUNC, tree, -1, id)

}
