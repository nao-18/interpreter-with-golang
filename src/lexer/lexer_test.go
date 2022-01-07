package lexer

import (
	"testing"

	"github.com/nao-18/interpreter-with-golang/src/monkey/token"
)

func TestNextToken(t *testing.T) {
	// 入力されるソースコードを定義
	input := `=+(),;`

	// テストで期待する型とリテラル値を定義
	tests := []struct {
		expectedType    token.TokenType
		expextedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	// レキサー作成
	l := New(input)

	for i, tt := range tests {
		// トークン列に変換
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expextedLiteral, tok.Litertal)
		}

		if tok.Literal != tt.expextedLiteral {
			t.Fatalf("tests[%d] - litetral wrong, expected=%q, got=%q", i, tt.expextedLiteral, tok.Literal)
		}
	}
}
