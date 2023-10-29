package lexer

import (
	"testing"

	"awesomeProject/token"
)

func TestNextToken(t *testing.T) {
	input := `let cash = 0;

state hello{
	respond hello;
	goto tasks;
}
state tasks{
}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "cash"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.STATE, "state"},
		{token.IDENT, "hello"},
		{token.LBRACE, "{"},
		{token.RESPOND, "respond"},
		{token.IDENT, "hello"},
		{token.SEMICOLON, ";"},
		{token.GOTO, "goto"},
		{token.IDENT, "tasks"},
		{token.SEMICOLON, ";"},
		{token.RBARCE, "}"},
		{token.STATE, "state"},
		{token.IDENT, "tasks"},
		{token.LBRACE, "{"},
		{token.RBARCE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
