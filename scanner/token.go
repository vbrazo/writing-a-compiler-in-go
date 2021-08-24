package main

const (
	T_PLUS = iota
	T_MINUS
	T_STAR
	T_SLASH
	T_INTLIT
)

type Token struct {
	token    int
	intvalue int
}
