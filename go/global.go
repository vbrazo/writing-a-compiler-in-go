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
var OpPrecedence = []int{0, 10, 10, 20, 20, 0}
