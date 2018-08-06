package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token // これから読見込むtokenリテラル
	peekToken token.Token // 次に読見込むtokenリテラル
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つのtokenを読み込む。curTokenとpeekTokenの両方をsetされる
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken      // 読み込むtokenリテラルに移動する
	p.peekToken = p.l.NextToken() // 先読みtokenリテラルを一つ進める
}

func (p *Parser) ParseProgram() *ast.Program {
	pg := &ast.Program{}
	pg.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			pg.Statements = append(pg.Statements, stmt)
		}
		p.nextToken()
	}

	return pg
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) { // let = 5 を想定しているので次に読み込むtokenは x, y のようなIdentifierであることを期待する。
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}

}
