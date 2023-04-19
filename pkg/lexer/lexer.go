package lexer

import (
	"groot/pkg/domain"
	"groot/pkg/token"
	"regexp"
)

type Lexer struct {
	input         string
	position      int
	readPosition  int
	ch            byte
	letterPattern *regexp.Regexp
}

func New(input string) *Lexer {
	l := &Lexer{
		input:         input,
		letterPattern: regexp.MustCompile(`[a-zA-Z_]+`),
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition

	l.readPosition += 1
}

func (l *Lexer) NextToken() domain.Token {
	var tok domain.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(domain.ASSIGN, l.ch)
	case ';':
		tok = newToken(domain.SEMICOLON, l.ch)
	case '(':
		tok = newToken(domain.LPAREN, l.ch)
	case ')':
		tok = newToken(domain.RPAREN, l.ch)
	case ',':
		tok = newToken(domain.COMMA, l.ch)
	case '+':
		tok = newToken(domain.PLUS, l.ch)
	case '{':
		tok = newToken(domain.LBRACE, l.ch)
	case '}':
		tok = newToken(domain.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = domain.EOF
	default:
		if l.isLetter(l.ch) {
			tok.Literal = l.getIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = domain.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(domain.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType domain.TokenType, ch byte) domain.Token {
	return domain.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) getIdentifier() string {
	position := l.position
	for l.isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) isLetter(ch byte) bool {
	return l.letterPattern.MatchString(string(ch))
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
