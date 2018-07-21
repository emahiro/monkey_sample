package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string // StatementがもつLiteralの実際の値（string）
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		// statement(節) が無いときは字句解析で何も返さなかった =　コードが書かれていないということ
		return ""
	}
}

// LetStatement
//  ex) let five = 5
//  let => token.Let
//  five => token.IDENT
//  5 => あらゆる式の表現が入る
type LetStatement struct {
	Token token.Token // token.LET のトークン
	Name  *Identifier // 変数の識別子
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // token.IDENT のトークン
	Value string      // 変数の実際の値
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
