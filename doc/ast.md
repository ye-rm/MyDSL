<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# ast

```go
import "awesomeDSL/ast"
```

Package ast contains types that represent the nodes of the abstract syntax tree.

## Index

- [type ArrayLiteral](<#ArrayLiteral>)
  - [func \(al \*ArrayLiteral\) String\(\) string](<#ArrayLiteral.String>)
  - [func \(al \*ArrayLiteral\) TokenLiteral\(\) string](<#ArrayLiteral.TokenLiteral>)
- [type BlockStatement](<#BlockStatement>)
  - [func \(bs \*BlockStatement\) String\(\) string](<#BlockStatement.String>)
  - [func \(bs \*BlockStatement\) TokenLiteral\(\) string](<#BlockStatement.TokenLiteral>)
- [type Boolean](<#Boolean>)
  - [func \(b \*Boolean\) String\(\) string](<#Boolean.String>)
  - [func \(b \*Boolean\) TokenLiteral\(\) string](<#Boolean.TokenLiteral>)
- [type CallExpression](<#CallExpression>)
  - [func \(ce \*CallExpression\) String\(\) string](<#CallExpression.String>)
  - [func \(ce \*CallExpression\) TokenLiteral\(\) string](<#CallExpression.TokenLiteral>)
- [type Expression](<#Expression>)
- [type ExpressionStatement](<#ExpressionStatement>)
  - [func \(es \*ExpressionStatement\) String\(\) string](<#ExpressionStatement.String>)
  - [func \(es \*ExpressionStatement\) TokenLiteral\(\) string](<#ExpressionStatement.TokenLiteral>)
- [type FunctionLiteral](<#FunctionLiteral>)
  - [func \(fl \*FunctionLiteral\) String\(\) string](<#FunctionLiteral.String>)
  - [func \(fl \*FunctionLiteral\) TokenLiteral\(\) string](<#FunctionLiteral.TokenLiteral>)
- [type Identifier](<#Identifier>)
  - [func \(i \*Identifier\) String\(\) string](<#Identifier.String>)
  - [func \(i \*Identifier\) TokenLiteral\(\) string](<#Identifier.TokenLiteral>)
- [type IfExpression](<#IfExpression>)
  - [func \(ie \*IfExpression\) String\(\) string](<#IfExpression.String>)
  - [func \(ie \*IfExpression\) TokenLiteral\(\) string](<#IfExpression.TokenLiteral>)
- [type IndexExpression](<#IndexExpression>)
  - [func \(ie \*IndexExpression\) String\(\) string](<#IndexExpression.String>)
  - [func \(ie \*IndexExpression\) TokenLiteral\(\) string](<#IndexExpression.TokenLiteral>)
- [type InfixExpression](<#InfixExpression>)
  - [func \(ie \*InfixExpression\) String\(\) string](<#InfixExpression.String>)
  - [func \(ie \*InfixExpression\) TokenLiteral\(\) string](<#InfixExpression.TokenLiteral>)
- [type IntegerLiteral](<#IntegerLiteral>)
  - [func \(il \*IntegerLiteral\) String\(\) string](<#IntegerLiteral.String>)
  - [func \(il \*IntegerLiteral\) TokenLiteral\(\) string](<#IntegerLiteral.TokenLiteral>)
- [type LetStatement](<#LetStatement>)
  - [func \(ls \*LetStatement\) String\(\) string](<#LetStatement.String>)
  - [func \(ls \*LetStatement\) TokenLiteral\(\) string](<#LetStatement.TokenLiteral>)
- [type Node](<#Node>)
- [type PrefixExpression](<#PrefixExpression>)
  - [func \(pe \*PrefixExpression\) String\(\) string](<#PrefixExpression.String>)
  - [func \(pe \*PrefixExpression\) TokenLiteral\(\) string](<#PrefixExpression.TokenLiteral>)
- [type Program](<#Program>)
  - [func \(p \*Program\) String\(\) string](<#Program.String>)
  - [func \(p \*Program\) TokenLiteral\(\) string](<#Program.TokenLiteral>)
- [type ReturnStatement](<#ReturnStatement>)
  - [func \(rs \*ReturnStatement\) String\(\) string](<#ReturnStatement.String>)
  - [func \(rs \*ReturnStatement\) TokenLiteral\(\) string](<#ReturnStatement.TokenLiteral>)
- [type Statement](<#Statement>)
- [type StringLiteral](<#StringLiteral>)
  - [func \(sl \*StringLiteral\) String\(\) string](<#StringLiteral.String>)
  - [func \(sl \*StringLiteral\) TokenLiteral\(\) string](<#StringLiteral.TokenLiteral>)


<a name="ArrayLiteral"></a>
## type [ArrayLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L317-L320>)

ArrayLiteral represents form "\[\<expression\>\]"

```go
type ArrayLiteral struct {
    Token    token.Token
    Elements []Expression
}
```

<a name="ArrayLiteral.String"></a>
### func \(\*ArrayLiteral\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L324>)

```go
func (al *ArrayLiteral) String() string
```



<a name="ArrayLiteral.TokenLiteral"></a>
### func \(\*ArrayLiteral\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L323>)

```go
func (al *ArrayLiteral) TokenLiteral() string
```



<a name="BlockStatement"></a>
## type [BlockStatement](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L238-L241>)

BlockStatement represents the statement of the form "\{ \<statement\>... \}"

```go
type BlockStatement struct {
    Token      token.Token
    Statements []Statement
}
```

<a name="BlockStatement.String"></a>
### func \(\*BlockStatement\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L245>)

```go
func (bs *BlockStatement) String() string
```



<a name="BlockStatement.TokenLiteral"></a>
### func \(\*BlockStatement\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L244>)

```go
func (bs *BlockStatement) TokenLiteral() string
```



<a name="Boolean"></a>
## type [Boolean](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L202-L205>)

Boolean represents the expression of the form "\<boolean\>"

```go
type Boolean struct {
    Token token.Token
    Value bool
}
```

<a name="Boolean.String"></a>
### func \(\*Boolean\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L209>)

```go
func (b *Boolean) String() string
```



<a name="Boolean.TokenLiteral"></a>
### func \(\*Boolean\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L208>)

```go
func (b *Boolean) TokenLiteral() string
```



<a name="CallExpression"></a>
## type [CallExpression](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L282-L286>)

CallExpression represents the expression of the form "\<expression\>\(\<arguments\>\)"

```go
type CallExpression struct {
    Token     token.Token
    Function  Expression
    Arguments []Expression
}
```

<a name="CallExpression.String"></a>
### func \(\*CallExpression\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L290>)

```go
func (ce *CallExpression) String() string
```



<a name="CallExpression.TokenLiteral"></a>
### func \(\*CallExpression\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L289>)

```go
func (ce *CallExpression) TokenLiteral() string
```



<a name="Expression"></a>
## type [Expression](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L22-L25>)



```go
type Expression interface {
    Node
    // contains filtered or unexported methods
}
```

<a name="ExpressionStatement"></a>
## type [ExpressionStatement](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L27-L30>)



```go
type ExpressionStatement struct {
    Token      token.Token
    Expression Expression
}
```

<a name="ExpressionStatement.String"></a>
### func \(\*ExpressionStatement\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L42>)

```go
func (es *ExpressionStatement) String() string
```

return the phrase item

<a name="ExpressionStatement.TokenLiteral"></a>
### func \(\*ExpressionStatement\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L37>)

```go
func (es *ExpressionStatement) TokenLiteral() string
```

TokenLiteral return the phrase description

<a name="FunctionLiteral"></a>
## type [FunctionLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L256-L260>)

FunctionLiteral represents the expression of the form "fun\(\<parameters\>\) \<block statement\>"

```go
type FunctionLiteral struct {
    Token      token.Token
    Parameters []*Identifier
    Body       *BlockStatement
}
```

<a name="FunctionLiteral.String"></a>
### func \(\*FunctionLiteral\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L264>)

```go
func (fl *FunctionLiteral) String() string
```



<a name="FunctionLiteral.TokenLiteral"></a>
### func \(\*FunctionLiteral\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L263>)

```go
func (fl *FunctionLiteral) TokenLiteral() string
```



<a name="Identifier"></a>
## type [Identifier](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L90-L93>)

Identifier represents the expression of the form "\<identifier\>"

```go
type Identifier struct {
    Token token.Token
    Value string
}
```

<a name="Identifier.String"></a>
### func \(\*Identifier\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L115>)

```go
func (i *Identifier) String() string
```



<a name="Identifier.TokenLiteral"></a>
### func \(\*Identifier\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L112>)

```go
func (i *Identifier) TokenLiteral() string
```



<a name="IfExpression"></a>
## type [IfExpression](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L212-L217>)

IfExpression represents the expression of the form "if \<condition\> \<consequence\> else \<alternative\>"

```go
type IfExpression struct {
    Token       token.Token
    Condition   Expression
    Consequence *BlockStatement
    Alternative *BlockStatement
}
```

<a name="IfExpression.String"></a>
### func \(\*IfExpression\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L221>)

```go
func (ie *IfExpression) String() string
```



<a name="IfExpression.TokenLiteral"></a>
### func \(\*IfExpression\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L220>)

```go
func (ie *IfExpression) TokenLiteral() string
```



<a name="IndexExpression"></a>
## type [IndexExpression](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L340-L344>)

IndexExpression represents form "\<expression\>\[\<expression\>\]"

```go
type IndexExpression struct {
    Token token.Token
    Left  Expression
    Index Expression
}
```

<a name="IndexExpression.String"></a>
### func \(\*IndexExpression\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L348>)

```go
func (ie *IndexExpression) String() string
```



<a name="IndexExpression.TokenLiteral"></a>
### func \(\*IndexExpression\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L347>)

```go
func (ie *IndexExpression) TokenLiteral() string
```



<a name="InfixExpression"></a>
## type [InfixExpression](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L180-L185>)

InfixExpression represents the expression of the form "\<expression\>\<infix operator\>\<expression\>"

```go
type InfixExpression struct {
    Token    token.Token
    Left     Expression
    Operator string
    Right    Expression
}
```

<a name="InfixExpression.String"></a>
### func \(\*InfixExpression\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L189>)

```go
func (ie *InfixExpression) String() string
```



<a name="InfixExpression.TokenLiteral"></a>
### func \(\*InfixExpression\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L188>)

```go
func (ie *InfixExpression) TokenLiteral() string
```



<a name="IntegerLiteral"></a>
## type [IntegerLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L144-L147>)

IntegerLiteral represents the expression of the form "\<integer\>"

```go
type IntegerLiteral struct {
    Token token.Token
    Value int64
}
```

<a name="IntegerLiteral.String"></a>
### func \(\*IntegerLiteral\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L153>)

```go
func (il *IntegerLiteral) String() string
```



<a name="IntegerLiteral.TokenLiteral"></a>
### func \(\*IntegerLiteral\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L150>)

```go
func (il *IntegerLiteral) TokenLiteral() string
```



<a name="LetStatement"></a>
## type [LetStatement](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L77-L81>)

LetStatement represents the statement of the form "let \<identifier\> = \<expression\>;"

```go
type LetStatement struct {
    Token token.Token
    Name  *Identifier
    Value Expression
}
```

<a name="LetStatement.String"></a>
### func \(\*LetStatement\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L95>)

```go
func (ls *LetStatement) String() string
```



<a name="LetStatement.TokenLiteral"></a>
### func \(\*LetStatement\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L84>)

```go
func (ls *LetStatement) TokenLiteral() string
```



<a name="Node"></a>
## type [Node](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L12-L15>)

Node is the interface that all nodes in the AST implement. TokenLiteral is used only for debugging and testing.

```go
type Node interface {
    TokenLiteral() string
    String() string
}
```

<a name="PrefixExpression"></a>
## type [PrefixExpression](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L158-L162>)

PrefixExpression represents the expression of the form "\<prefix operator\>\<expression\>"

```go
type PrefixExpression struct {
    Token    token.Token
    Operator string
    Right    Expression
}
```

<a name="PrefixExpression.String"></a>
### func \(\*PrefixExpression\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L168>)

```go
func (pe *PrefixExpression) String() string
```



<a name="PrefixExpression.TokenLiteral"></a>
### func \(\*PrefixExpression\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L165>)

```go
func (pe *PrefixExpression) TokenLiteral() string
```



<a name="Program"></a>
## type [Program](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L52-L54>)

Program is the root node of every AST our parser produces. Every valid program is a series of statements.

```go
type Program struct {
    Statements []Statement
}
```

<a name="Program.String"></a>
### func \(\*Program\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L66>)

```go
func (p *Program) String() string
```

get the phrase description

<a name="Program.TokenLiteral"></a>
### func \(\*Program\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L57>)

```go
func (p *Program) TokenLiteral() string
```

TokenLiteral get the phrase item

<a name="ReturnStatement"></a>
## type [ReturnStatement](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L120-L123>)

ReturnStatement represents the statement of the form "return \<expression\>;"

```go
type ReturnStatement struct {
    Token       token.Token
    ReturnValue Expression
}
```

<a name="ReturnStatement.String"></a>
### func \(\*ReturnStatement\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L129>)

```go
func (rs *ReturnStatement) String() string
```



<a name="ReturnStatement.TokenLiteral"></a>
### func \(\*ReturnStatement\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L126>)

```go
func (rs *ReturnStatement) TokenLiteral() string
```



<a name="Statement"></a>
## type [Statement](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L17-L20>)



```go
type Statement interface {
    Node
    // contains filtered or unexported methods
}
```

<a name="StringLiteral"></a>
## type [StringLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L307-L310>)

StringLiteral represents form "\<string literal\>"

```go
type StringLiteral struct {
    Token token.Token
    Value string
}
```

<a name="StringLiteral.String"></a>
### func \(\*StringLiteral\) [String](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L314>)

```go
func (sl *StringLiteral) String() string
```



<a name="StringLiteral.TokenLiteral"></a>
### func \(\*StringLiteral\) [TokenLiteral](<https://github.com/ye-rm/awesomeDSL/blob/master/ast/ast.go#L313>)

```go
func (sl *StringLiteral) TokenLiteral() string
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
