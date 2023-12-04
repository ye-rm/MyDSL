package parser

import (
	"awesomeDSL/ast"
	"awesomeDSL/lexer"
	"awesomeDSL/token"
	"fmt"
	"strconv"
)

// provide a precedence for each token
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type PrefixParseFn func() ast.Expression
type InfixParseFn func(ast.Expression) ast.Expression

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string

	PrefixParseFns map[token.TokenType]PrefixParseFn
	InfixParseFns  map[token.TokenType]InfixParseFn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn PrefixParseFn) {
	p.PrefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn InfixParseFn) {
	p.InfixParseFns[tokenType] = fn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	//init cur&peek Token
	p.nextToken()
	p.nextToken()

	//init map
	p.PrefixParseFns = make(map[token.TokenType]PrefixParseFn)
	p.InfixParseFns = make(map[token.TokenType]InfixParseFn)
	//add parse function for each token
	p.registerPrefix(token.IDENT, p.phraseIdentifier)
	p.registerPrefix(token.INT, p.phraseIntegerLiteral)
	p.registerPrefix(token.BANG, p.phrasePrefixExpression)
	p.registerPrefix(token.MINUS, p.phrasePrefixExpression)
	p.registerPrefix(token.TRUE, p.phraseBoolean)
	p.registerPrefix(token.FALSE, p.phraseBoolean)
	p.registerPrefix(token.LPAREN, p.phraseGroupedExpression)
	p.registerPrefix(token.IF, p.phraseIfExpression)
	p.registerPrefix(token.FUNCTION, p.phraseFunctionLiteral)

	p.registerInfix(token.LPAREN, p.phraseCallExpression)
	p.registerInfix(token.MINUS, p.phraseInfixExpression)
	p.registerInfix(token.PLUS, p.phraseInfixExpression)
	p.registerInfix(token.DIVIDE, p.phraseInfixExpression)
	p.registerInfix(token.MULTI, p.phraseInfixExpression)
	p.registerInfix(token.EQ, p.phraseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.phraseInfixExpression)
	p.registerInfix(token.LT, p.phraseInfixExpression)
	p.registerInfix(token.RT, p.phraseInfixExpression)

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) phraseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.phraseLetStatement()
	case token.RETURN:
		return p.phraseReturnStatement()
	default:
		return p.phraseExpressionStatement()
	}
}

func (p *Parser) phraseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.phraseExpression(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) phraseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	stmt.ReturnValue = p.phraseExpression(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) phraseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) phraseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.phraseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) phraseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)

		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value

	return lit
}

func (p *Parser) phrasePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.phraseExpression(PREFIX)

	return expression
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// if the next token is t, then consume it and advance token
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)

	p.errors = append(p.errors, msg)
}

func (p *Parser) phraseExpression(precedence int) ast.Expression {
	//get prefix parse function
	prefix := p.PrefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()
	//if next token is not semicolon and precedence is lower than next token's precedence
	//then get infix parse function
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.InfixParseFns[p.peekToken.Type]

		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.phraseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

var precedences = map[token.TokenType]int{
	token.EQ:     EQUALS,
	token.NOT_EQ: EQUALS,
	token.LT:     LESSGREATER,
	token.RT:     LESSGREATER,
	token.PLUS:   SUM,
	token.MINUS:  SUM,
	token.DIVIDE: PRODUCT,
	token.MULTI:  PRODUCT,
	token.LPAREN: CALL,
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) phraseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	//get operator precedence
	precedence := p.curPrecedence()
	p.nextToken()
	//recursively call phraseExpression to get right expression
	expression.Right = p.phraseExpression(precedence)

	return expression
}

func (p *Parser) phraseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) phraseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.phraseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return exp
}

func (p *Parser) phraseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()
	expression.Condition = p.phraseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expression.Consequence = p.phraseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}
		expression.Alternative = p.phraseBlockStatement()
	}
	return expression
}

func (p *Parser) phraseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()
	for !p.curTokenIs(token.RBARCE) && !p.curTokenIs(token.EOF) {
		stmt := p.phraseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) phraseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	lit.Parameters = p.phraseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	lit.Body = p.phraseBlockStatement()

	return lit
}

func (p *Parser) phraseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}
	p.nextToken()
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.SEPARATE) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return identifiers
}

func (p *Parser) phraseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.phraseCallArguments()
	return exp
}

func (p *Parser) phraseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	//no arguments
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}
	p.nextToken()
	args = append(args, p.phraseExpression(LOWEST))

	for p.peekTokenIs(token.SEPARATE) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.phraseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}
