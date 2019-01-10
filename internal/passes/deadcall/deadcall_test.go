package deadcall

import (
	"strings"
	"testing"

	"github.com/nickng/migo/v3"
	"github.com/nickng/migo/v3/parser"
)

func TestRemoveUndefined(t *testing.T) {
	s := `
	def main():
		spawn a();
		send v;
	def a():
		call b();
		call c();
	def c():
		recv x;
	def d():
		tau;
	`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
	}
	if want, got := 4, len(prog.Funcs); want != got {
		t.Errorf("expecting %d functions but got %d", want, got)
	}
	if want, got := prog.Funcs[1].Name, "a"; want != got {
		t.Errorf("expecting function 1 to be `def %s` but got `def %s`", want, got)
	}
	if want, got := 2, len(prog.Funcs[1].Stmts); want != got {
		t.Errorf("expecting %d statements (before removal) but got %d:\n%s", want, got, prog.Funcs[1])
	}
	Remove(prog)
	if want, got := 4, len(prog.Funcs); want != got {
		t.Errorf("expecting %d functions but got %d", want, got)
	}
	if want, got := prog.Funcs[1].Name, "a"; want != got {
		t.Errorf("expecting function 1 to be `def %s` but got `def %s`", want, got)
	}
	if want, got := 1, len(prog.Funcs[1].Stmts); want != got {
		t.Errorf("expecting %d statements (after removal) but got %d:\n%s", want, got, prog.Funcs[1])
	}
}

func TestRemoveUndefined2(t *testing.T) {
	s := `
	def main():
		spawn a();
		send v;
	def a():
		if call b(); else call c(); call e(); endif;
	def c():
		recv x;
	def d():
		tau;
	`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
	}
	if want, got := 4, len(prog.Funcs); want != got {
		t.Errorf("expecting %d functions but got %d", want, got)
	}
	if want, got := prog.Funcs[1].Name, "a"; want != got {
		t.Errorf("expecting function 1 to be `def %s` but got `def %s`", want, got)
	}
	if want, got := 1, len(prog.Funcs[1].Stmts); want != got {
		t.Errorf("expecting %d statements (before removal) but got %d:\n%s", want, got, prog.Funcs[1])
	}
	if ite, ok := prog.Funcs[1].Stmts[0].(*migo.IfStatement); !ok {
		t.Errorf("expecting a if-then statement but got %T", prog.Funcs[1].Stmts)
		t.FailNow()
	} else {
		if want, got := 1, len(ite.Then); want != got {
			t.Errorf("expecting %d statement in then-statement but got %d:\n%s", want, got, ite.Then)
		}
		if want, got := 2, len(ite.Else); want != got {
			t.Errorf("expecting %d statement in else-statement but got %d:\n%s", want, got, ite.Else)
		}
	}
	Remove(prog)
	if want, got := 4, len(prog.Funcs); want != got {
		t.Errorf("expecting %d functions but got %d", want, got)
	}
	if want, got := prog.Funcs[1].Name, "a"; want != got {
		t.Errorf("expecting function 1 to be `def %s` but got `def %s`", want, got)
	}
	if want, got := 1, len(prog.Funcs[1].Stmts); want != got {
		t.Errorf("expecting %d statements (after removal) but got %d:\n%s", want, got, prog.Funcs[1])
	}
	if ite, ok := prog.Funcs[1].Stmts[0].(*migo.IfStatement); !ok {
		t.Errorf("expecting a if-then statement but got %T", prog.Funcs[1].Stmts)
		t.FailNow()
	} else {
		if want, got := 1, len(ite.Then); want != got {
			// becomes tau
			if _, ok := ite.Then[0].(*migo.TauStatement); !ok {
				t.Errorf("expecting statement to be reduced to tau but got %s", ite.Then[0])
			}
			t.Errorf("expecting %d statement in then-statement but got %d:\n%s", want, got, ite.Then)
		}
		if want, got := 1, len(ite.Else); want != got {
			t.Errorf("expecting %d statement in else-statement but got %d:\n%s", want, got, ite.Else)
		}
	}
}

func TestRemoveUndefined3(t *testing.T) {
	s := `
	def main():
		spawn a();
		send v;
	def a():
		if call b(); else call c(); call e(); endif;
	def d():
		tau;
	`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
	}
	if want, got := 3, len(prog.Funcs); want != got {
		t.Errorf("expecting %d functions but got %d", want, got)
	}
	if want, got := prog.Funcs[1].Name, "a"; want != got {
		t.Errorf("expecting function 1 to be `def %s` but got `def %s`", want, got)
	}
	if want, got := 1, len(prog.Funcs[1].Stmts); want != got {
		t.Errorf("expecting %d statements (before removal) but got %d:\n%s", want, got, prog.Funcs[1])
	}
	if ite, ok := prog.Funcs[1].Stmts[0].(*migo.IfStatement); !ok {
		t.Errorf("expecting a if-then statement but got %T", prog.Funcs[1].Stmts)
		t.FailNow()
	} else {
		if want, got := 1, len(ite.Then); want != got {
			t.Errorf("expecting %d statement in then-statement but got %d:\n%s", want, got, ite.Then)
		}
		if want, got := 2, len(ite.Else); want != got {
			t.Errorf("expecting %d statement in else-statement but got %d:\n%s", want, got, ite.Else)
		}
	}
	Remove(prog)
	if want, got := 3, len(prog.Funcs); want != got {
		t.Errorf("expecting %d functions but got %d", want, got)
	}
	if want, got := prog.Funcs[1].Name, "a"; want != got {
		t.Errorf("expecting function 1 to be `def %s` but got `def %s`", want, got)
	}
	if want, got := 1, len(prog.Funcs[1].Stmts); want != got {
		t.Errorf("expecting %d statements (after removal) but got %d:\n%s", want, got, prog.Funcs[1])
	}
	if _, ok := prog.Funcs[1].Stmts[0].(*migo.TauStatement); !ok {
		t.Errorf("expecting a tau statement but got %T", prog.Funcs[1].Stmts[0])
		t.FailNow()
	}
}
