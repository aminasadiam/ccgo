package parser

import (
	"reflect"
	"testing"

	"github.com/aminasadiam/ccgo/lexer"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    []lexer.Token
		expected *Node
		wantErr  bool
	}{
		{
			name: "Basic function with number",
			input: []lexer.Token{
				{TokenType: lexer.IntKeyword, Literal: "int", Line: 1},
				{TokenType: lexer.Ident, Literal: "main", Line: 1},
				{TokenType: lexer.LParen, Literal: "(", Line: 1},
				{TokenType: lexer.RParen, Literal: ")", Line: 1},
				{TokenType: lexer.LBrace, Literal: "{", Line: 1},
				{TokenType: lexer.ReturnKeyword, Literal: "return", Line: 1},
				{TokenType: lexer.Number, Literal: "42", Line: 1},
				{TokenType: lexer.Semicolon, Literal: ";", Line: 1},
				{TokenType: lexer.RBrace, Literal: "}", Line: 1},
				{TokenType: lexer.EOF, Literal: "", Line: 1},
			},
			expected: &Node{
				Type: ProgramNode,
				Children: []*Node{
					{
						Type:  FunctionNode,
						Value: "main",
						Children: []*Node{
							{
								Type: ReturnNode,
								Children: []*Node{
									{Type: NumberNode, Value: "42"},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Arithmetic expression",
			input: []lexer.Token{
				{TokenType: lexer.IntKeyword, Literal: "int", Line: 1},
				{TokenType: lexer.Ident, Literal: "main", Line: 1},
				{TokenType: lexer.LParen, Literal: "(", Line: 1},
				{TokenType: lexer.RParen, Literal: ")", Line: 1},
				{TokenType: lexer.LBrace, Literal: "{", Line: 1},
				{TokenType: lexer.ReturnKeyword, Literal: "return", Line: 1},
				{TokenType: lexer.LParen, Literal: "(", Line: 1},
				{TokenType: lexer.Number, Literal: "2", Line: 1},
				{TokenType: lexer.Plus, Literal: "+", Line: 1},
				{TokenType: lexer.Number, Literal: "3", Line: 1},
				{TokenType: lexer.RParen, Literal: ")", Line: 1},
				{TokenType: lexer.Semicolon, Literal: ";", Line: 1},
				{TokenType: lexer.RBrace, Literal: "}", Line: 1},
				{TokenType: lexer.EOF, Literal: "", Line: 1},
			},
			expected: &Node{
				Type: ProgramNode,
				Children: []*Node{
					{
						Type:  FunctionNode,
						Value: "main",
						Children: []*Node{
							{
								Type: ReturnNode,
								Children: []*Node{
									{
										Type:  BinaryExprNode,
										Value: "+",
										Left:  &Node{Type: NumberNode, Value: "2"},
										Right: &Node{Type: NumberNode, Value: "3"},
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Missing int keyword",
			input: []lexer.Token{
				{TokenType: lexer.Ident, Literal: "main", Line: 1},
				{TokenType: lexer.LParen, Literal: "(", Line: 1},
				{TokenType: lexer.RParen, Literal: ")", Line: 1},
				{TokenType: lexer.LBrace, Literal: "{", Line: 1},
				{TokenType: lexer.ReturnKeyword, Literal: "return", Line: 1},
				{TokenType: lexer.Number, Literal: "42", Line: 1},
				{TokenType: lexer.Semicolon, Literal: ";", Line: 1},
				{TokenType: lexer.RBrace, Literal: "}", Line: 1},
				{TokenType: lexer.EOF, Literal: "", Line: 1},
			},
			wantErr: true,
		},
		{
			name: "Invalid expression",
			input: []lexer.Token{
				{TokenType: lexer.IntKeyword, Literal: "int", Line: 1},
				{TokenType: lexer.Ident, Literal: "main", Line: 1},
				{TokenType: lexer.LParen, Literal: "(", Line: 1},
				{TokenType: lexer.RParen, Literal: ")", Line: 1},
				{TokenType: lexer.LBrace, Literal: "{", Line: 1},
				{TokenType: lexer.ReturnKeyword, Literal: "return", Line: 1},
				{TokenType: lexer.Plus, Literal: "+", Line: 1},
				{TokenType: lexer.Semicolon, Literal: ";", Line: 1},
				{TokenType: lexer.RBrace, Literal: "}", Line: 1},
				{TokenType: lexer.EOF, Literal: "", Line: 1},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Parse() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
