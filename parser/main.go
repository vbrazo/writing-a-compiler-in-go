package main

import (
	"fmt"
	"log"
	"os"
)

func usage(prog string) {
	log.Printf("Usage: %s <infile>\n", prog)
	os.Exit(1)
}

func initVariables() {
	Line = 1
	Putback = '\n'
}
func cleanup() {
	if FilePtr != nil {
		fmt.Println("Closing the file")
		FilePtr.Close()
	}
}

func main() {
	defer cleanup()
	if len(os.Args) != 2 {
		usage(os.Args[0])
	}
	initVariables()
	openFile(os.Args[1])
	scan(&T)
	n := binExpr()
	printAstDepth(n)
	interpretAST(&n)
}
