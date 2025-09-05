package lexer

import (
	"fmt"
	"unicode"
)

type TokenType int

type Token struct {
	TokenType TokenType
	Literal   string
	Line      int
}

const (
	Ident TokenType = iota
	Number
	Plus
	Minus
	Star
	Slash
	Semicolon
	LBrace
	RBrace
	LParen
	RParen
	IntKeyword
	ReturnKeyword
	EOF
)

func Tokenize(source string) ([]Token, error) {
	var tokens []Token
	line := 1
	i := 0

	for i < len(source) {
		char := rune(source[i])

		if unicode.IsSpace(char) {
			if char == '\n' {
				line++
			}
			i++
			continue
		}

		if unicode.IsLetter(char) {
			start := i
			for i < len(source) && (unicode.IsLetter(rune(source[i]))) || unicode.IsDigit((rune(source[i]))) {
				i++
			}
			literal := source[start:i]
			tokenType := Ident

			switch literal {
			case "int":
				tokenType = IntKeyword
			case "return":
				tokenType = ReturnKeyword
			}

			tokens = append(tokens, Token{TokenType: tokenType, Literal: literal, Line: line})
			continue
		}

		if unicode.IsDigit(char) {
			start := i
			for i < len(source) && unicode.IsDigit(rune(source[i])) {
				i++
			}
			tokens = append(tokens, Token{TokenType: Number, Literal: source[start:i], Line: line})
			continue
		}

		switch char {
		case '+':
			tokens = append(tokens, Token{TokenType: Plus, Literal: "+", Line: line})
		case '-':
			tokens = append(tokens, Token{TokenType: Minus, Literal: "-", Line: line})
		case '*':
			tokens = append(tokens, Token{TokenType: Star, Literal: "*", Line: line})
		case '/':
			tokens = append(tokens, Token{TokenType: Slash, Literal: "/", Line: line})
		case ';':
			tokens = append(tokens, Token{TokenType: Semicolon, Literal: ";", Line: line})
		case '{':
			tokens = append(tokens, Token{TokenType: LBrace, Literal: "{", Line: line})
		case '}':
			tokens = append(tokens, Token{TokenType: RBrace, Literal: "}", Line: line})
		case '(':
			tokens = append(tokens, Token{TokenType: LParen, Literal: "(", Line: line})
		case ')':
			tokens = append(tokens, Token{TokenType: RParen, Literal: ")", Line: line})
		default:
			return nil, fmt.Errorf("unexpected character '%c' at line %d", char, line)
		}
		i++
	}

	tokens = append(tokens, Token{TokenType: EOF, Literal: "", Line: line})
	return tokens, nil
}
