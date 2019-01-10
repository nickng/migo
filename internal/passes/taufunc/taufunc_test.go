package taufunc

import (
	"strings"
	"testing"

	"github.com/nickng/migo/v3/parser"
)

// main, send, recv should remain after removing empty.
func TestRemoveTau(t *testing.T) {
	s := `
def main():
	let ch = newchan ch_instance, 0;
	spawn s(ch);
	spawn r(ch);
	spawn work();
	recv ch;
	recv ch;
def s(sch):
	send sch;
def r(rch):
	recv rch;
	send rch;
def work():
	tau;
def unreachable():
	tau;
	`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if mainfn, found := prog.Function("main"); !found {
		t.Errorf("main function not found")
	} else {
		if want, got := 5, len(prog.Funcs); want != got {
			t.Errorf("expects %d functions but got %d", want, got)
		}
		Find(prog, RemoveExcept(mainfn))
		if want, got := 3, len(prog.Funcs); want != got {
			t.Errorf("expects %d functions after removing {work,unreachable} but got %d", want, got)
		}
	}
}

// Tests if non-taus are propagated backwards in the control-flow graph.
func TestRemoveTauPropagate(t *testing.T) {
	s := `
def main():
	send a;
	recv b;
	call c();
def c():
	call d();
def d():
	send x;
def e(): tau;
	`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if mainfn, found := prog.Function("main"); !found {
		t.Errorf("main function not found")
	} else {
		Find(prog, RemoveExcept(mainfn))
	}
}

// Tests if non-taus are propagated backwards in the control-flow graph
// (deeply nested calls).
func TestRemoveTauPropagate2(t *testing.T) {
	s := `
def main():
	send a;
	recv b;
	call c();
def c():
	call d();
def d():
	send x;
	call e();
def e(): tau; spawn f();
def f(): call g(); call f();
def g():
	let ch = newchan T, 0;
	call j();
def h(): call i();
def j(): call i();
def i(): spawn j(); call h();`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if mainfn, found := prog.Function("main"); !found {
		t.Errorf("main function not found")
	} else {
		Find(prog, RemoveExcept(mainfn))
	}
}

func TestReduceMainTau(t *testing.T) {
	s := `
def main():
	call b();
	call c();
def b(): tau;
def c(): tau;
def d(): let ch = newchan T, 0; spawn e(ch);
def e(ch): send ch;`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if mainfn, found := prog.Function("main"); !found {
		t.Errorf("main function not found")
	} else {
		if want, got := 5, len(prog.Funcs); want != got {
			t.Errorf("expects %d functions but got %d", want, got)
		}
		Find(prog, RemoveExcept(mainfn)) // main, d, e
		if want, got := 3, len(prog.Funcs); want != got {
			t.Errorf("expects %d functions after removing {b,c} but got %d", want, got)
		}
		Find(prog, Remove) // d, e
		if want, got := 2, len(prog.Funcs); want != got {
			t.Errorf("expects %d functions after removing {main} but got %d", want, got)
		}
	}
}
