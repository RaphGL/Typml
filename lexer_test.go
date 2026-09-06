package typml

import (
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	for _, s := range []string{
		"\"some random string\"",
		"   \t\n  \"2304324\" \t\r\n",
	} {
		l := NewLexer(s)
		s = strings.Trim(strings.TrimSpace(s), "\"")

		tok, ok := l.Eat()
		if !ok || tok.typ != TokenTypeString || tok.value != s {
			t.Errorf("Failed to lex string: `%v`", s)
		}
	}
}

func TestIdentifiers(t *testing.T) {
	for _, i := range []string{
		"\t\t\twhatever",
		"\n\n\nyes",
		"m",
	} {
		l := NewLexer(i)
		i := strings.TrimSpace(i)

		tok, ok := l.Eat()
		if !ok || (tok.typ != TokenTypeIdent && tok.value != i) {
			t.Errorf("Failed to lex identifier: `%v`", i)
		}
	}

	for _, i := range []string{
		"whatever5",
		"m1",
		"_whatever",
		"some_thing",
		"0whatever",
	} {
		l := NewLexer(i)

		tok, ok := l.Eat()
		if ok && tok.typ == TokenTypeIdent && tok.value == i {
			t.Errorf("Lexed erroneous identifier: `%v`", i)
		}
	}
}

func TestNumbers(t *testing.T) {
	for _, n := range []string{
		"50",
		"\t\r\t\t\n78",
		"1500",
	} {
		l := NewLexer(n)
		n = strings.TrimSpace(n)

		tok, ok := l.Eat()
		if !ok || tok.typ != TokenTypeNumber || tok.value != n {
			t.Errorf("Failed to lex number: `%v`", n)
		}
	}
}
