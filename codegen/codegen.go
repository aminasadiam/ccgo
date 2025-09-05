package codegen

import (
	"fmt"
	"os"

	"github.com/aminasadiam/ccgo/parser"
)

// Generate produces x86-64 assembly from AST
func Generate(ast *parser.Node, outputFile string) error {
	var asmCode string
	asmCode += "global _start\n"
	asmCode += "_start:\n"
	asmCode += "    call main\n"
	asmCode += "    mov rdi, rax\n"
	asmCode += "    mov rax, 60\n"
	asmCode += "    syscall\n\n"

	for _, child := range ast.Children {
		if child.Type == parser.FunctionNode {
			asmCode += "main:\n"
			asmCode += "    push rbp\n"
			asmCode += "    mov rbp, rsp\n"
			for _, stmt := range child.Children {
				if stmt.Type == parser.ReturnNode {
					exprCode := generateExpr(stmt.Children[0])
					asmCode += exprCode
					asmCode += "    mov rsp, rbp\n"
					asmCode += "    pop rbp\n"
					asmCode += "    ret\n"
				}
			}
		}
	}

	return os.WriteFile(outputFile, []byte(asmCode), 0644)
}

func generateExpr(node *parser.Node) string {
	if node.Type == parser.NumberNode {
		return fmt.Sprintf("    mov rax, %s\n", node.Value)
	}

	if node.Type == parser.BinaryExprNode {
		leftCode := generateExpr(node.Left)
		rightCode := generateExpr(node.Right)
		opCode := ""
		if node.Value == "+" {
			opCode = "    add rax, rbx\n"
		} else if node.Value == "-" {
			opCode = "    sub rax, rbx\n"
		}
		return fmt.Sprintf("%s    mov rbx, rax\n%s%s", leftCode, rightCode, opCode)
	}

	return ""
}
