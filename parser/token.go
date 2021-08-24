package main

const (
	T_EOF = iota
	T_PLUS
	T_MINUS
	T_STAR
	T_SLASH
	T_INTLIT
)

type Token struct {
	token    int
	intvalue int
}
