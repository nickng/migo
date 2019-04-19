package migo_test

import (
	"strings"
	"testing"

	"github.com/nickng/migo/v3"
	"github.com/nickng/migo/v3/parser"
)

// Test simplifying name.
func TestSimpleName(t *testing.T) {
	fullGoName := "(github.com/nickng/migo).String#1"
	simpleName := "github.com_nickng_migo.String#1"
	s := &migo.CallStatement{Name: fullGoName, Params: []*migo.Parameter{}}
	if s.SimpleName() != simpleName {
		t.Errorf("SimplifyName of %s should be %s but got %s",
			fullGoName, simpleName, s.SimpleName())
	}
}

// High level syntax consistency tests
// These tests make sure Statement.String() are consistent with input language,
// more detailed parsing checks can be found in parser/parser_test.go

func TestBasicSyntax(t *testing.T) {
	s := `def main():
    let ch = newchan T, 0;
    send ch;
    recv ch;
    select
      case send ch;
      case recv ch;
      case tau;
    endselect;
    if close ch; else spawn f(ch); endif;
def f(ch):
    recv ch;
`
	r := strings.NewReader(s)
	parsed, err := parser.Parse(r)
	if err != nil {
		t.Error(err)
	}
	if want, got := s, parsed.String(); want != got {
		t.Errorf("syntax mismatch, want:\n%s\ngot:\n%s", want, got)
	}
}

func TestMemSyntax(t *testing.T) {
	s := `def main():
    letmem x;
    spawn fn();
    write x;
`
	r := strings.NewReader(s)
	parsed, err := parser.Parse(r)
	if err != nil {
		t.Error(err)
	}
	if want, got := s, parsed.String(); want != got {
		t.Errorf("syntax mismatch, want:\n%s\ngot:\n%s", want, got)
	}
}

func TestLockSyntax(t *testing.T) {
	s := `def main():
    letsync mu mutex;
    letsync rmu rwmutex;
    lock mu;
    rlock rmu;
    runlock mu;
    unlock mu;
`
	r := strings.NewReader(s)
	parsed, err := parser.Parse(r)
	if err != nil {
		t.Error(err)
	}
	if want, got := s, parsed.String(); want != got {
		t.Errorf("syntax mismatch, want:\n%s\ngot:\n%s", want, got)
	}
}
