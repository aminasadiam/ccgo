package codegen

import (
	"os"
	"path"
	"testing"

	"github.com/aminasadiam/ccgo/parser"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		input    *parser.Node
		expected string
		wantErr  bool
	}{
		{
			name: "Return number",
			input: &parser.Node{
				Type: parser.ProgramNode,
				Children: []*parser.Node{
					{
						Type:  parser.FunctionNode,
						Value: "main",
						Children: []*parser.Node{
							{
								Type: parser.ReturnNode,
								Children: []*parser.Node{
									{Type: parser.NumberNode, Value: "42"},
								},
							},
						},
					},
				},
			},
			expected: `global _start
_start:
    call main
    mov rdi, rax
    mov rax, 60
    syscall

main:
    push rbp
    mov rbp, rsp
    mov rax, 42
    mov rsp, rbp
    pop rbp
    ret
`,
			wantErr: false,
		},
		{
			name: "Arithmetic expression",
			input: &parser.Node{
				Type: parser.ProgramNode,
				Children: []*parser.Node{
					{
						Type:  parser.FunctionNode,
						Value: "main",
						Children: []*parser.Node{
							{
								Type: parser.ReturnNode,
								Children: []*parser.Node{
									{
										Type:  parser.BinaryExprNode,
										Value: "+",
										Left:  &parser.Node{Type: parser.NumberNode, Value: "2"},
										Right: &parser.Node{Type: parser.NumberNode, Value: "3"},
									},
								},
							},
						},
					},
				},
			},
			expected: `global _start
_start:
    call main
    mov rdi, rax
    mov rax, 60
    syscall

main:
    push rbp
    mov rbp, rsp
    mov rax, 2
    mov rbx, rax
    mov rax, 3
    add rax, rbx
    mov rsp, rbp
    pop rbp
    ret
`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := path.Join(os.TempDir(), "test_output.asm")
			err := Generate(tt.input, tmpFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got, err := os.ReadFile(tmpFile)
				if err != nil {
					t.Errorf("Failed to read output file: %v", err)
					return
				}
				if string(got) != tt.expected {
					t.Errorf("Generate() got = %v, want %v", string(got), tt.expected)
				}
				os.Remove(tmpFile)
			}
		})
	}
}
