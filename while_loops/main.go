package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func usage(prog string) {
	log.Printf("Usage: %s <infile> <outfile>\n", prog)
	os.Exit(1)
}

func initVariables() {
	Line = 1
	Putback = '\n'
}
func cleanup() {
	if InputFilePtr != nil {
		fmt.Println("Closing the inputfile")
		InputFilePtr.Close()
	}
	if OutputFilePtr != nil {
		OutputFile.Flush()
		fmt.Println("Closing the outputfile")
		OutputFilePtr.Close()
	}
}

func compileFile(path, tmpPath string) {
	OutputFile.Flush()
	cmd := exec.Command("go", "build", "-o", path, tmpPath)
	log.Printf("Compiling...")
	output, err := cmd.Output()
	log.Printf(string(output))
	if err != nil {
		log.Printf("Compiling finished with error: %v", err)
	}
	log.Printf("Successfully compiled")
	cmd = exec.Command(filepath.Join(filepath.Base(path), "build.exe"))
	cmd.Run()
}

func main() {

	/*x := map[string]interface{}{"x": "5"}
	y := map[string]string{"x": "5"}

	meinTest(x, y)*/

	generateVariable()
	defer cleanup()
	if len(os.Args) != 3 {
		usage(os.Args[0])
	}
	initVariables()
	openInputFile(os.Args[1])
	openOutputFile("./bin/tmp.go")
	scan(&T)
	genMainFuncStart()
	ast := compundStatement()
	writeOutput(interpretAST(ast))
	genMainFuncEnd()
	compileFile(os.Args[2], "./bin/tmp.go")
}

func meinTest(x, y interface{}) {
	print(x.(map[string]interface{})["x"].(string))
	print(x.(map[string]interface{})["x"])
}
