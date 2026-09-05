package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenTypeComma TokenType = iota
	TokenTypeColon
	TokenTypeSemicolon
	TokenTypeIdent
	TokenTypeNumber
	TokenTypeLeftParen
	TokenTypeRightParen
	TokenTypeLeftBrace
	TokenTypeRightBrace
	TokenTypeBar
	TokenTypeSingleQuote
	TokenTypeDoubleQuote
	TokenTypeGreaterThan
	TokenTypeLessThan
	TokenTypeDot
	TokenTypeMinus
	TokenTypeHat
	TokenTypePlus
	TokenTypeMult
	TokenTypeDiv
	TokenTypeEqual
	TokenTypeBackslash
	TokenTypeUnderscore
)

type Token struct {
	typ   TokenType
	value string
}

type Lexer struct {
	sr *strings.Reader
}

func NewLexer(s string) Lexer {
	return Lexer{
		sr: strings.NewReader(s),
	}
}

func (l *Lexer) lexNumber() (token Token, ok bool) {
	l.sr.UnreadRune()
	var sb strings.Builder
	for {
		r, _, err := l.sr.ReadRune()
		if err != nil || !unicode.IsDigit(r) {
			l.sr.UnreadRune()
			break
		}
		sb.WriteRune(r)
	}
	token.value = sb.String()
	token.typ = TokenTypeNumber
	ok = true
	return token, sb.Len() != 0
}

func (l *Lexer) lexIdent() (token Token, ok bool) {
	l.sr.UnreadRune()

	var sb strings.Builder
	r, _, err := l.sr.ReadRune()
	if err != nil || !unicode.IsLetter(r) {
		l.sr.UnreadRune()
		ok = false
		return
	}
	sb.WriteRune(r)

	isIdentRune := func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
	}

	for {
		r, _, err := l.sr.ReadRune()
		if err != nil || !isIdentRune(r) {
			l.sr.UnreadRune()
			break
		}
		sb.WriteRune(r)
	}

	token.value = sb.String()
	token.typ = TokenTypeIdent
	ok = true
	return
}

func (l *Lexer) skipWhitespace() {
	for {
		r, _, err := l.sr.ReadRune()
		if err != nil {
			break
		}
		if !unicode.IsSpace(r) {
			l.sr.UnreadRune()
			break
		}
	}
}

func (l *Lexer) Eat() (token Token, ok bool) {
	l.skipWhitespace()

	r, _, err := l.sr.ReadRune()
	if err != nil {
		return
	}

	token.typ = -1

	switch r {
	case ':':
		token.typ = TokenTypeColon
	case ';':
		token.typ = TokenTypeSemicolon
	case '(':
		token.typ = TokenTypeLeftParen
	case ')':
		token.typ = TokenTypeRightParen
	case '{':
		token.typ = TokenTypeLeftBrace
	case '}':
		token.typ = TokenTypeRightBrace
	case '|':
		token.typ = TokenTypeBar
	case '\'':
		token.typ = TokenTypeSingleQuote
	case '"':
		token.typ = TokenTypeDoubleQuote
	case '>':
		token.typ = TokenTypeGreaterThan
	case '<':
		token.typ = TokenTypeLessThan
	case '.':
		token.typ = TokenTypeDot
	case '-':
		token.typ = TokenTypeMinus
	case '^':
		token.typ = TokenTypeHat
	case '+':
		token.typ = TokenTypePlus
	case '*':
		token.typ = TokenTypeMult
	case '/':
		token.typ = TokenTypeDiv
	case '=':
		token.typ = TokenTypeEqual
	case '\\':
		token.typ = TokenTypeBackslash
	case '_':
		token.typ = TokenTypeUnderscore
	}

	if token.typ == -1 {
		switch {
		case unicode.IsDigit(r):
			return l.lexNumber()
		case unicode.IsLetter(r):
			return l.lexIdent()
		}
	}

	token.value = string(r)
	ok = true
	return
}

func main() {
	l := NewLexer("5 + (7 * num) / 2")
	for {
		tok, ok := l.Eat()
		if !ok {
			break
		}
		fmt.Print(tok.value)
	}
}
