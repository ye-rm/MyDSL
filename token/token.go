// content: token type and token struct
package token

type TokenType string

// Token struct
type Token struct {
	Type    TokenType // token type
	Literal string    // token literal
}

// token type
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

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

	STATE  = "STATE"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBARCE = "}"

	RESPOND  = "RESPOND"
	CATCH    = "CATCH"
	GOTO     = "GOTO"
	HAVE     = "HAVE"
	IF       = "IF"
	ELSE     = "ELSE"
	ELSEIF   = "ELSEIF"
	LET      = "LET"
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)
