package gomorphy

import (
	"strings"
	"testing"
)

func TestParsesReturnsAllParsesWithTaggedLexeme(t *testing.T) {
	a, err := Default()
	if err != nil {
		t.Fatal(err)
	}
	parses := a.Parses("стали")
	lemmas := map[string]bool{}
	for _, p := range parses {
		if p.Tag == "" || len(p.Lexeme) == 0 {
			t.Fatalf("empty parse: %+v", p)
		}
		for _, f := range p.Lexeme {
			if f.Tag == "" {
				t.Fatalf("form without tag in lexeme of %q: %+v", p.Lexeme[0].Word, f)
			}
		}
		lemmas[p.Lexeme[0].Word] = true
	}
	if !lemmas["сталь"] || !lemmas["стать"] {
		t.Fatalf("expected lemmas сталь and стать, got %v", lemmas)
	}
}

func TestParsesUnknownWord(t *testing.T) {
	a, err := Default()
	if err != nil {
		t.Fatal(err)
	}
	if got := a.Parses("  "); got != nil {
		t.Fatalf("expected nil for blank input, got %v", got)
	}
}

func TestParsesFormTagsMatchCase(t *testing.T) {
	a, err := Default()
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range a.Parses("кошку") {
		if !strings.HasPrefix(p.Tag, "NOUN") {
			continue
		}
		for _, f := range p.Lexeme {
			if f.Word == "кошкой" && !strings.Contains(f.Tag, "ablt") {
				t.Fatalf("кошкой should be ablt, got %q", f.Tag)
			}
		}
		return
	}
	t.Fatal("no NOUN parse for кошку")
}
