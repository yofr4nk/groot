package lexer_test

import (
	"groot/pkg/domain"
	"groot/pkg/lexer"
	"testing"
)

func TestNextTokenWithKnownTokens(t *testing.T) {
	input := `=+(){},;`
	tokens := []domain.Token{
		{domain.ASSIGN, "="},
		{domain.PLUS, "+"},
		{domain.LPAREN, "("},
		{domain.RPAREN, ")"},
		{domain.LBRACE, "{"},
		{domain.RBRACE, "}"},
		{domain.COMMA, ","},
		{domain.SEMICOLON, ";"},
		{domain.EOF, ""},
	}

	l := lexer.New(input)

	for i, tok := range tokens {
		nt := l.NextToken()
		if nt.Type != tok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tok.Type, nt.Type)
		}
		if nt.Literal != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tok.Literal, nt.Literal)
		}
	}
}

func TestNextTokenDetectingVariableDeclaration(t *testing.T) {
	input := "let var_declared = 5;"

	tokens := []domain.Token{
		{domain.LET, "let"},
		{domain.IDENT, "var_declared"},
		{domain.ASSIGN, "="},
		{domain.INT, "5"},
		{domain.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	for i, tok := range tokens {
		nt := l.NextToken()
		if nt.Type != tok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tok.Type, nt.Type)
		}
		if nt.Literal != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tok.Literal, nt.Literal)
		}
	}
}

func TestNextTokenDetectingFunctionDeclaration(t *testing.T) {
	input := `let func_declared = fn(a, b) {
			a + b;
		};`

	tokens := []domain.Token{
		{domain.LET, "let"},
		{domain.IDENT, "func_declared"},
		{domain.ASSIGN, "="},
		{domain.FUNCTION, "fn"},
		{domain.LPAREN, "("},
		{domain.IDENT, "a"},
		{domain.COMMA, ","},
		{domain.IDENT, "b"},
		{domain.RPAREN, ")"},
		{domain.LBRACE, "{"},
		{domain.IDENT, "a"},
		{domain.PLUS, "+"},
		{domain.IDENT, "b"},
		{domain.SEMICOLON, ";"},
		{domain.RBRACE, "}"},
		{domain.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	for i, tok := range tokens {
		nt := l.NextToken()
		if nt.Type != tok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tok.Type, nt.Type)
		}
		if nt.Literal != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tok.Literal, nt.Literal)
		}
	}
}

func TestNextTokenDetectingFunctionDeclarationAndReferenced(t *testing.T) {
	input := `
		let first_number = 1;
		let second_number = 2;
		let func_declared = fn(a, b) {
			a + b;
		};
		let result = func_declared(first_number, second_number);
	`

	tokens := []domain.Token{
		{domain.LET, "let"},
		{domain.IDENT, "first_number"},
		{domain.ASSIGN, "="},
		{domain.INT, "1"},
		{domain.SEMICOLON, ";"},
		{domain.LET, "let"},
		{domain.IDENT, "second_number"},
		{domain.ASSIGN, "="},
		{domain.INT, "2"},
		{domain.SEMICOLON, ";"},
		{domain.LET, "let"},
		{domain.IDENT, "func_declared"},
		{domain.ASSIGN, "="},
		{domain.FUNCTION, "fn"},
		{domain.LPAREN, "("},
		{domain.IDENT, "a"},
		{domain.COMMA, ","},
		{domain.IDENT, "b"},
		{domain.RPAREN, ")"},
		{domain.LBRACE, "{"},
		{domain.IDENT, "a"},
		{domain.PLUS, "+"},
		{domain.IDENT, "b"},
		{domain.SEMICOLON, ";"},
		{domain.RBRACE, "}"},
		{domain.SEMICOLON, ";"},
		{domain.LET, "let"},
		{domain.IDENT, "result"},
		{domain.ASSIGN, "="},
		{domain.IDENT, "func_declared"},
		{domain.LPAREN, "("},
		{domain.IDENT, "first_number"},
		{domain.COMMA, ","},
		{domain.IDENT, "second_number"},
		{domain.RPAREN, ")"},
		{domain.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	for i, tok := range tokens {
		nt := l.NextToken()
		if nt.Type != tok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tok.Type, nt.Type)
		}
		if nt.Literal != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tok.Literal, nt.Literal)
		}
	}
}

func TestNextTokenInvalidToken(t *testing.T) {
	input := "¡"

	l := lexer.New(input)

	if l.NextToken().Type != domain.ILLEGAL {
		t.Fatalf("tokentype wrong. expected=%q, got=%q",
			domain.ILLEGAL, l.NextToken().Type)
	}
}
