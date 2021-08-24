package main

const (
	T_EOF = iota
	T_NEWLINE
	T_PLUS
	T_MINUS
	T_STAR
	T_SLASH
	T_INTLIT
	T_PRINT
)

type Token struct {
	token    int
	intvalue int
}
