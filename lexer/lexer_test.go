package lexer

import (
	"testing"

	"awesomeDSL/token"
)

func TestNextToken(t *testing.T) {
	input := `let cash = 0;

state hello{
	respond hello;
	goto tasks;
}

state tasks{
}
if cash != 0 {


}
"foobar"
"foo bar"
[1,2];
`

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
		{token.IF, "if"},
		{token.IDENT, "cash"},
		{token.NOT_EQ, "!="},
		{token.INT, "0"},
		{token.LBRACE, "{"},
		{token.RBARCE, "}"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.SEPARATE, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
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
