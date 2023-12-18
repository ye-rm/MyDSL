// Package token content: token type and token struct
package token

type TokenType string

// Token struct
type Token struct {
	Type    TokenType // token type
	Literal string    // token literal
}

// Token type const value
const (
	ILLEGAL = "ILLEGAL" // illegal token
	EOF     = "EOF"     // end of file

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MULTI  = "*"
	DIVIDE = "/"
	MOD    = "%"
	BANG   = "!"

	EQ     = "=="
	NOT_EQ = "!="

	LT = "<"
	RT = ">"

	SEPARATE  = ","
	SEMICOLON = ";"
	IDENT     = "IDENT"

	INT = "INT"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBARCE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	IF       = "IF"
	ELSE     = "ELSE"
	ELSEIF   = "ELSEIF"
	LET      = "LET"
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
	STRING   = "STRING"
)
