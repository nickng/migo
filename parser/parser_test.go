package parser

import (
	"strings"
	"testing"

	"github.com/nickng/migo/v3"
)

func TestParser(t *testing.T) {
	want := `def main():
    let ch = newchan T, 0;
    send ch;
`
	p, err := Parse(strings.NewReader(`   def main(): let ch = newchan T, 0;
	send ch;   `))
	if err != nil {
		t.Errorf("cannot parse: %v", err)
	}
	if got := p.String(); want != got {
		t.Errorf("unexpected parsed migo, want:\n%sgot:\n%s", want, got)
	}
}

func TestParseMem(t *testing.T) {
	s := `def main(): letmem x; read x; spawn f(); write x;
    def f(): read x; write x;`
	parsed, err := Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if want, got := 2, len(parsed.Funcs); want != got {
		t.Errorf("expected %d functions but got %d", want, got)
	}
	fn, found := parsed.Function("main")
	if !found {
		t.Error("cannot find main function")
	}
	if want, got := 4, len(fn.Stmts); want != got {
		t.Errorf("expected %d statements but got %d", want, got)
	}
	stmt0, ok := fn.Stmts[0].(*migo.NewMem)
	if !ok {
		t.Errorf("expecting letmem statement but got %v", fn.Stmts[0])
		t.FailNow()
	}
	if stmt0.Name != "x" {
		t.Errorf("expected letmem x but got letmem %s", stmt0.Name)
	}
	stmt1, ok := fn.Stmts[1].(*migo.MemRead)
	if !ok {
		t.Errorf("expecting read statement but got %v", fn.Stmts[1])
		t.FailNow()
	}
	if stmt1.Name != "x" {
		t.Errorf("expected read x but got letmem %s", stmt1.Name)
	}
	stmt3, ok := fn.Stmts[3].(*migo.MemWrite)
	if !ok {
		t.Errorf("expecting write statement but got %v", fn.Stmts[3])
		t.FailNow()
	}
	if stmt3.Name != "x" {
		t.Errorf("expected write x but got letmem %s", stmt3.Name)
	}
}
