package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置(現在の文字)
	readPosition int  // これから読み込む位置（これから読み込む文字の位置）
	ch           byte // 現在検査中の文字
}

// New Initializer
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

// readChar ソースコードの文字列のなかで読み込む位置を進める
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// readIdentifier 識別子を読み込んで非英字に到達するまで字句解析を続ける
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
	/*
		一つ読み進める前後。以下のような挙動になる
		str := "abc"
		i := str[0:1] => "a"が得られる。
	*/
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// skipWhitespaces 空白、タブ、改行を読み込んだら無視する -> 無視してひとつ先を読み進める
func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// NextToken ソースコードからTokenを読み込んで生成する。
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// 空白と改行を無視する。
	l.skipWhitespaces()

	switch l.ch {
	case '=':
		// ひとつ先のリテラルが "=" だったとき
		if l.peekChar() == '=' {
			ch := l.ch

			// 1つ進める
			l.readChar()

			// = + = => "==" を生成する。 = １つ文字を進めたときのリテラルを追加する
			lit := string(ch) + string(l.ch)

			// string -> byteにすることもできるが、メモリコピーが走るので好ましくないので使わない。
			tok = token.Token{Type: token.EQ, Literal: lit}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: lit}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // 読み込む位置を一つ進める
	return tok
}

// newToken tokenを生成する
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
