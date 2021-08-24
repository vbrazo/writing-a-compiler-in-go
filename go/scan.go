package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func openInputFile(fileName string) {
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	InputFilePtr = file
	InputFile = bufio.NewReader(file)
}

func next() (rune, error) {
	var c rune
	if Putback > 0 {
		c = Putback
		Putback = 0
		return c, nil
	}

	c, _, err := InputFile.ReadRune()
	Column++
	if err != nil {
		if err == io.EOF {
			return 0, err
		} else {
			log.Fatal(err)
			panic(err)
		}
	}
	if c == '\n' {
		Line++
	}
	return c, nil
}

func putback(c rune) {
	Putback = c
}

func skip() (rune, error) {
	c, err := next()
	if err != nil {
		return 0, err
	}
	for c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f' {
		c, err = next()
		if err != nil {
			return 0, err
		}
	}
	return c, nil
}

func scanint(c rune) int {
	var err error
	var val = 0
	for unicode.IsDigit(c) {
		val = val*10 + int(c-'0')
		c, err = next()
		if err != nil {
			if err != io.EOF {
				log.Fatalf("Could not read integer. Line %d Col $d Char: %q", Line, Column, c)

			} else {
				break
			}
		}
	}
	putback(c)
	return val
}

func scan(t *Token) bool {
	c, err := skip()
	if err != nil {
		t.token = T_EOF
		return false
	}
	switch c {
	case '+':
		t.token = T_PLUS
		break
	case '-':
		t.token = T_MINUS
		break
	case '*':
		t.token = T_STAR
		break
	case '/':
		t.token = T_SLASH
		break
	default:
		if unicode.IsDigit(c) {
			t.intvalue = scanint(c)
			t.token = T_INTLIT
		} else {
			fmt.Printf("Unrecognised character %q on Line %d Col %d \n", c, Line, Column)
			os.Exit(2)
		}
	}
	return true
}
