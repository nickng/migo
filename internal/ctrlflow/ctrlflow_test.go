package ctrlflow_test

import (
	"strings"
	"testing"

	"github.com/nickng/migo/internal/ctrlflow"
	"github.com/nickng/migo/parser"
)

func TestEmpty(t *testing.T) {
	s := `def main(): tau;`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	g := ctrlflow.NewGraph(prog)
	if want, got := 1, len(g.Nodes); want != got {
		t.Errorf("expected %d node but got %d", want, got)
	}
	if want, got := 0, len(g.Nodes[0].Succs); want != got {
		t.Errorf("expected %d successors but got %d", want, got)
	}
	if want, got := 0, len(g.Nodes[0].Preds); want != got {
		t.Errorf("expected %d predecessors but got %d", want, got)
	}
}

func TestChain(t *testing.T) {
	s := `
def main.main(): call a(); call b();
def a(): call c();
def b(): call c();
def c(): spawn d();
def d(): spawn e();
def e(): tau;
`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	g := ctrlflow.NewGraph(prog)
	if want, got := 6, len(g.Nodes); want != got {
		t.Errorf("expected %d nodes but got %d", want, got)
	}
	// Order:    main, a, c, d, e, b
	preds := [6]int{0, 1, 2, 1, 1, 1}
	succs := [6]int{2, 1, 1, 1, 0, 1}
	for i := 0; i < 6; i++ {
		if want, got := succs[i], len(g.Nodes[i].Succs); want != got {
			t.Errorf("node[%d]: expected %d successors but got %d: %v", i, want, got, g.Nodes[i])
		}
		if want, got := preds[i], len(g.Nodes[i].Preds); want != got {
			t.Errorf("node[%d]: expected %d predecessors but got %d: %v", i, want, got, g.Nodes[i])
		}
	}
}

func TestComplex(t *testing.T) {
	s := `
def main.main():
	let ch1 = newchan T, 0;
	call sel(ch1);
	let ch2 = newchan T, 0;
	call a(ch1, ch2);
	spawn f(ch1, ch2);
def a(a, b):
	tau;
def f(a, b):
	call mkchan();
	call g(b, a);
def mkchan():
	let ch = newchan T, 0;
def g(b,c):
	send b;
	recv c;
def sel(a):
	select
	case recv a; send a;
	endselect;
def h():
	tau;
	tau;
`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	g := ctrlflow.NewGraph(prog)
	// Order:
	//   main       .
	//   sel           .
	//   a                .
	//   f                   .
	//   mkchan                 .
	//   g                         .
	//   h                            .
	preds := [7]int{0, 1, 1, 1, 1, 1, 0}
	succs := [7]int{3, 0, 0, 2, 0, 0, 0}
	for i := 0; i < 7; i++ {
		if want, got := succs[i], len(g.Nodes[i].Succs); want != got {
			t.Errorf("node[%d]: expected %d successors but got %d: %v", i, want, got, g.Nodes[i])
		}
		if want, got := preds[i], len(g.Nodes[i].Preds); want != got {
			t.Errorf("node[%d]: expected %d predecessors but got %d: %v", i, want, got, g.Nodes[i])
		}
	}
}

func TestRecursive(t *testing.T) {
	s := `
def multicall():
	call multicall();
	spawn multicall();
def calling_inf_loop():
	call infloop(a);
def calling_inf_loop2():
	call infloop2(a);
def infloop2(a):
	call infloop2(a);
def infloop(a):
	call infloop(a);
	send a;
def mutual_recursive_a():
	call mutual_recursive_b();
def mutual_recursive_b():
	call mutual_recursive_a();
`
	prog, err := parser.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	g := ctrlflow.NewGraph(prog)
	if want, got := 7, len(g.Nodes); want != got {
		t.Errorf("expected %d node but got %d", want, got)
	}
	// Order:
	//  multicall  .
	//  calling_inf_loop
	//  infloop           .
	//  calling_inf_loop2    .
	//  infloop2                .
	//  mutual_recursive_a         .
	//  mutual_recursive_b            .
	preds := [7]int{1, 0, 2, 0, 2, 1, 1}
	succs := [7]int{1, 1, 1, 1, 1, 1, 1}
	for i := 0; i < 7; i++ {
		if want, got := succs[i], len(g.Nodes[i].Succs); want != got {
			t.Errorf("node[%d]: expected %d successors but got %d: %v", i, want, got, g.Nodes[i])
		}
		if want, got := preds[i], len(g.Nodes[i].Preds); want != got {
			t.Errorf("node[%d]: expected %d predecessors but got %d: %v", i, want, got, g.Nodes[i])
		}
	}
}
