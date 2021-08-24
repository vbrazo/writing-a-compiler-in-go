package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var VariableCounter = 0

func openOutputFile(path string) {
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	file, err := os.Create(path) // For read access.
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	OutputFilePtr = file
	OutputFile = bufio.NewWriter(file)
}

func genMainFuncStart() {
	fmt.Fprintf(OutputFile, "package main \n\n import \"fmt\" \n\nfunc main(){\n")
}

func genPrint() {
	fmt.Fprintf(OutputFile, "fmt.Printf(\"%%d\\n\",%s)\n", getLastGenVariable())
}

func genMainFuncEnd() {
	fmt.Fprintf(OutputFile, "}")
}

func getLastGenVariable() string {
	return "v" + strconv.Itoa(VariableCounter)
}

func generateVariable() string {
	VariableCounter++
	return getLastGenVariable()
}

func genMathExpression(operator string, left, right string) string {
	variableName := generateVariable()
	fmt.Fprintf(OutputFile, "%s := (%v %s %v)\n", variableName, left, operator, right)
	return variableName
}

func genAdd(left, right string) string {
	return genMathExpression("+", left, right)
}
func genSub(left, right string) string {
	return genMathExpression("-", left, right)
}
func genMul(left, right string) string {
	return genMathExpression("*", left, right)
}
func genDiv(left, right string) string {
	return genMathExpression("/", left, right)
}

func genNumber(node *AstNode) string {
	variableName := generateVariable()
	fmt.Fprintf(OutputFile, "%s := %s\n", variableName, strconv.Itoa(node.intval))
	return variableName
}
