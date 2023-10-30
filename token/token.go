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
	BANG = "!"

	EQ     = "=="
	NOT_EQ = "!="

	LT = "<"
	RT = ">"

	SEPERATE=","
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
	RETURN ="RETURN"
)
