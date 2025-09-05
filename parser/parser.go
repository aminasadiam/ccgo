package parser

import (
	"fmt"

	"github.com/aminasadiam/ccgo/lexer"
)

type NodeType int

type Node struct {
	Type     NodeType
	Value    string
	Left     *Node
	Right    *Node
	Children []*Node
}

const (
	ProgramNode NodeType = iota
	FunctionNode
	ReturnNode
	BinaryExprNode
	NumberNode
)

func Parse(tokens []lexer.Token) (*Node, error) {
	var pos int

	// Expecting: int main() { return <expr>; }
	if pos >= len(tokens) || tokens[pos].TokenType != lexer.IntKeyword {
		return nil, fmt.Errorf("expected 'int' at line %d", tokens[pos].Line)
	}
	pos++

	if pos >= len(tokens) || tokens[pos].TokenType != lexer.Ident || tokens[pos].Literal != "main" {
		return nil, fmt.Errorf("expected 'main' at line %d", tokens[pos].Line)
	}
	pos++

	if pos >= len(tokens) || tokens[pos].TokenType != lexer.LParen {
		return nil, fmt.Errorf("expected '(' at line %d", tokens[pos].Line)
	}
	pos++

	if pos >= len(tokens) || tokens[pos].TokenType != lexer.RParen {
		return nil, fmt.Errorf("expected ')' at line %d", tokens[pos].Line)
	}
	pos++

	if pos >= len(tokens) || tokens[pos].TokenType != lexer.LBrace {
		return nil, fmt.Errorf("expected '{' at line %d", tokens[pos].Line)
	}
	pos++

	// Parse return statement
	if pos >= len(tokens) || tokens[pos].TokenType != lexer.ReturnKeyword {
		return nil, fmt.Errorf("expected 'return' at line %d", tokens[pos].Line)
	}
	pos++

	// Parse simple expression (e.g., number or binary operation)
	expr, newPos, err := parseExpr(tokens, pos)
	if err != nil {
		return nil, err
	}
	pos = newPos

	if pos >= len(tokens) || tokens[pos].TokenType != lexer.Semicolon {
		return nil, fmt.Errorf("expected ';' at line %d", tokens[pos].Line)
	}
	pos++

	if pos >= len(tokens) || tokens[pos].TokenType != lexer.RBrace {
		return nil, fmt.Errorf("expected '}' at line %d", tokens[pos].Line)
	}
	pos++

	if pos >= len(tokens) || tokens[pos].TokenType != lexer.EOF {
		return nil, fmt.Errorf("unexpected tokens after function end at line %d", tokens[pos].Line)
	}

	// Build AST
	program := &Node{Type: ProgramNode}
	function := &Node{Type: FunctionNode, Value: "main"}
	returnStmt := &Node{Type: ReturnNode, Children: []*Node{expr}}
	function.Children = []*Node{returnStmt}
	program.Children = []*Node{function}
	return program, nil
}

func parseExpr(tokens []lexer.Token, pos int) (*Node, int, error) {
	if pos >= len(tokens) {
		return nil, pos, fmt.Errorf("unexpected end of input")
	}

	if tokens[pos].TokenType == lexer.Number {
		return &Node{Type: NumberNode, Value: tokens[pos].Literal}, pos + 1, nil
	}

	if tokens[pos].TokenType == lexer.LParen {
		pos++
		left, newPos, err := parseExpr(tokens, pos)
		if err != nil {
			return nil, pos, err
		}
		pos = newPos

		if pos >= len(tokens) {
			return nil, pos, fmt.Errorf("expected operator")
		}

		op := tokens[pos]
		if op.TokenType != lexer.Plus && op.TokenType != lexer.Minus {
			return nil, pos, fmt.Errorf("expected '+' or '-' at line %d", op.Line)
		}
		pos++

		right, newPos, err := parseExpr(tokens, pos)
		if err != nil {
			return nil, pos, err
		}
		pos = newPos

		if pos >= len(tokens) || tokens[pos].TokenType != lexer.RParen {
			return nil, pos, fmt.Errorf("expected ')' at line %d", tokens[pos].Line)
		}
		pos++

		return &Node{
			Type:  BinaryExprNode,
			Value: op.Literal,
			Left:  left,
			Right: right,
		}, pos, nil
	}

	return nil, pos, fmt.Errorf("invalid expression at line %d", tokens[pos].Line)
}
