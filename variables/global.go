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
var OpPrecedence = []int{0, 0, 10, 10, 20, 20, 0}
var LastScannedIdentifier string

const (
	T_EOF = iota
	T_NEWLINE
	T_PLUS
	T_MINUS
	T_STAR
	T_SLASH
	T_INTLIT
	T_PRINT
	T_EQ
	T_VAR
	T_IDENT
)

//Ast ops
const (
	A_ADD = iota
	A_SUBTRACT
	A_MULTIPLY
	A_DIVIDE
	A_INTLIT
	A_ASSIGNVAL
	A_ASSIGN
	A_IDENT
)
