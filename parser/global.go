package main

import (
	"bufio"
	"os"
)

var Line int
var Column int
var Putback rune
var FilePtr *os.File
var File *bufio.Reader
var TokenStr = []string{"+", "-", "*", "/", "intlit"}
var T Token
