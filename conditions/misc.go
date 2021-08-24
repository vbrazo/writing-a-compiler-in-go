package main

// Match a left brace and fetch the next token
func matchLBrace() {
	matchToken(T_LBRACE, "{")
}

// Match a right brace and fetch the next token
func matchRBrace() {
	matchToken(T_RBRACE, "}")
}

// Match a left parenthesis and fetch the next token
func matchLParen() {
	matchToken(T_LPAREN, "(")
}

// Match a right parenthesis and fetch the next token
func matchRParen() {
	matchToken(T_RPAREN, ")")
}

func matchIdent() {
	matchToken(T_IDENT, "identifier")
}

func matchNewLine() {
	matchToken(T_NEWLINE, "\\n")
}
