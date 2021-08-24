package main

import (
	"bufio"
	"os"
)

var Line int
var Column int
var Putback rune
var InputFilePtr *os.File
var InputFile *bufio.Reader
var OutputFilePtr *os.File
var OutputFile *bufio.Writer
var TokenStr = []string{"+", "-", "*", "/", "intlit"}
var T Token
var OpPrecedence = []int{
	0, 10, 10, // T_EOF, T_PLUS, T_MINUS
	20, 20, // T_STAR, T_SLASH
	30, 30, // T_EQ, T_NEQ
	40, 40, 40, 40, // T_LT, T_GT, T_LE, T_GE
}
var LastScannedIdentifier string

const (
	T_EOF = iota
	T_PLUS
	T_MINUS
	T_STAR
	T_SLASH
	T_EQ
	T_NEQ
	T_LT
	T_GT
	T_LE
	T_GE
	//no precedence
	T_NEWLINE
	T_INTLIT
	T_PRINT
	T_ASSIGN
	T_LBRACE
	T_RBRACE
	T_LPAREN
	T_RPAREN
	//keywords
	T_VAR
	T_IDENT
	T_IF
	T_ELSE
)

//Ast ops
const (
	A_EOF_PLACEHOLDER = iota //used to T_TOKEN == A_OP in getAstType
	A_ADD
	A_SUBTRACT
	A_MULTIPLY
	A_DIVIDE
	A_EQ
	A_NEQ
	A_LT
	A_GT
	A_LE
	A_GE
	A_INTLIT
	A_ASSIGNVAL
	A_ASSIGN
	A_IDENT
	A_PRINT
	A_GLUETO
	A_IF
)
