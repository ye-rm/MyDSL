package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MULTI  = "*"
	DIVIDE = "/"
	MOD    = "%"

	SEMICOLON = ";"
	IDENT     = "IDENT"

	INT = "INT"

	STATE  = "STATE"
	LBRACE = "{"
	RBARCE = "}"

	RESPOND = "RESPOND"
	CATCH   = "CATCH"
	GOTO    = "GOTO"
	HAVE    = "HAVE"
	IF      = "IF"
	ELSE    = "ELSE"
	ELSEIF  = "ELSEIF"
	LET     = "LET"
)
