package lexer

import "github.com/nao-18/interpreter-with-golang/src/monkey/token"

type Lexer struct {
	input        string
	posting      int  // 入力における現在の位置(現在の文字を指し示す)
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
	l.posting = l.readPosition
	// 次の文字インデックスを更新
	l.readPosition += 1
}

// トークンを取得
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

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
	case '}':
		tok = newToken(token.LBRACE, l.ch)
	case '{':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Litertal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

// トークン作成
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Litertal: string(ch)}
}