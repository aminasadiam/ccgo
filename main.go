package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aminasadiam/ccgo/codegen"
	"github.com/aminasadiam/ccgo/lexer"
	"github.com/aminasadiam/ccgo/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ccgo <source_file.c>")
		os.Exit(1)
	}

	startTime := time.Now()

	sourceFile := os.Args[1]
	source, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	tokens, err := lexer.Tokenize(string(source))
	if err != nil {
		fmt.Printf("Lexing error: %v\n", err)
		os.Exit(1)
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		fmt.Printf("Parsing error: %v\n", err)
		os.Exit(1)
	}

	err = codegen.Generate(ast, "out.asm")
	if err != nil {
		fmt.Printf("Codegen error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Compilation successful! Output: out.asm")
	fmt.Printf("Compile Time: %s", time.Since(startTime))
}
