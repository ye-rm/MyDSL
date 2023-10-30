package parser

import(
	"awesomeDSL/ast"
	"awesomeDSL/lexer"
	"awesomeDSL/token"
)

type Parser struct{
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser{
	p := &Parser{l: l}

	//init cur&peek Token
	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser)NextToken(){
	p.curToken=p.peekToken
	p.peekToken=p.l.NextToken()
}

