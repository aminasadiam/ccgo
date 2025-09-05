package lexer

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Token
		wantErr  bool
	}{
		{
			name:  "Basic function",
			input: `int main() { return 42; }`,
			expected: []Token{
				{TokenType: IntKeyword, Literal: "int", Line: 1},
				{TokenType: Ident, Literal: "main", Line: 1},
				{TokenType: LParen, Literal: "(", Line: 1},
				{TokenType: RParen, Literal: ")", Line: 1},
				{TokenType: LBrace, Literal: "{", Line: 1},
				{TokenType: ReturnKeyword, Literal: "return", Line: 1},
				{TokenType: Number, Literal: "42", Line: 1},
				{TokenType: Semicolon, Literal: ";", Line: 1},
				{TokenType: RBrace, Literal: "}", Line: 1},
				{TokenType: EOF, Literal: "", Line: 1},
			},
			wantErr: false,
		},
		{
			name:  "Arithmetic expression",
			input: `int main() { return (1 + 2); }`,
			expected: []Token{
				{TokenType: IntKeyword, Literal: "int", Line: 1},
				{TokenType: Ident, Literal: "main", Line: 1},
				{TokenType: LParen, Literal: "(", Line: 1},
				{TokenType: RParen, Literal: ")", Line: 1},
				{TokenType: LBrace, Literal: "{", Line: 1},
				{TokenType: ReturnKeyword, Literal: "return", Line: 1},
				{TokenType: LParen, Literal: "(", Line: 1},
				{TokenType: Number, Literal: "1", Line: 1},
				{TokenType: Plus, Literal: "+", Line: 1},
				{TokenType: Number, Literal: "2", Line: 1},
				{TokenType: RParen, Literal: ")", Line: 1},
				{TokenType: Semicolon, Literal: ";", Line: 1},
				{TokenType: RBrace, Literal: "}", Line: 1},
				{TokenType: EOF, Literal: "", Line: 1},
			},
			wantErr: false,
		},
		{
			name:    "Invalid character",
			input:   `int main() { return @; }`,
			wantErr: true,
		},
		{
			name: "Multiple lines",
			input: `int main()
{
	return 0;
}`,
			expected: []Token{
				{TokenType: IntKeyword, Literal: "int", Line: 1},
				{TokenType: Ident, Literal: "main", Line: 1},
				{TokenType: LParen, Literal: "(", Line: 1},
				{TokenType: RParen, Literal: ")", Line: 1},
				{TokenType: LBrace, Literal: "{", Line: 2},
				{TokenType: ReturnKeyword, Literal: "return", Line: 3},
				{TokenType: Number, Literal: "0", Line: 3},
				{TokenType: Semicolon, Literal: ";", Line: 3},
				{TokenType: RBrace, Literal: "}", Line: 4},
				{TokenType: EOF, Literal: "", Line: 4},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tokenize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tokenize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Tokenize() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
