package ast

import "awesomeDSL/token"

//let <id> = <experssion>

type Node interface {
	//return the phrase item
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Experssion interface {
	Node
	experssionNode()
}

type Program struct {
	Statements []Statement
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Experssion
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Experssion
}

func (l *LetStatement) statementNode() {}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (l *LetStatement) experssionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (rs *ReturnStatement) statementNode() {

}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
