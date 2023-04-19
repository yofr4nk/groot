package token

import "groot/pkg/domain"

var keywords = map[string]domain.TokenType{
	"fn":  domain.FUNCTION,
	"let": domain.LET,
}

func LookupIdent(ident string) domain.TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return domain.IDENT
}
