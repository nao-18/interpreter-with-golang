package lexer

import "github.com/nao-18/interpreter-with-golang/src/monkey/token"

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置(現在の文字を指し示す)
	readPosition int  // これから読み込む位置(現在の文字の次)
	ch           byte // 現在検査中の文字
}

// レキサー作成メソッド
func New(input string) *Lexer {
	// レキサー作成
	l := &Lexer{input: input}
	// 読み込み(構造体LexerのchとreadPositionを更新する)
	l.readChar()
	return l
}

// 文字を読み込む
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 読み込む文字がない場合
		l.ch = 0
	} else {
		// 読み込む文字がある場合
		l.ch = l.input[l.readPosition]
	}
	// 現在検査中の文字インデックス更新
	l.position = l.readPosition
	// 次の文字インデックスを更新
	l.readPosition += 1
}

// トークンを取得
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
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
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// 英字の場合
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// 数字の場合
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// 英字以外の場合
			// 文字がわからないためILLEGALへ保存
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 0~9の数字か判定
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// ホワイトスペース等の不要なコードを読み飛ばす
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// トークン作成
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
