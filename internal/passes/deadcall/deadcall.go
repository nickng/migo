// Package deadcall defines a transformation pass to remove dead function calls.
//
// Dead functions calls are calls (or spawns) to functions that are not defined.
package deadcall

import "github.com/nickng/migo/v3"

// Remove removes undefined function calls and spawns.
func Remove(prog *migo.Program) {
	rmvr := undefRemover{prog: prog}
	for i := range prog.Funcs {
		rmvr.traverse(&prog.Funcs[i].Stmts)
	}
}

type undefRemover struct {
	prog *migo.Program
}

func (r undefRemover) traverse(stmts *[]migo.Statement) {
	ss := *stmts
	for i := 0; i < len(ss); i++ {
		switch stmt := (ss)[i].(type) {
		case *migo.IfStatement:
			r.traverse(&stmt.Then)
			r.traverse(&stmt.Else)
			if len(stmt.Then) == 1 {
				_, isThenTau := stmt.Then[0].(*migo.TauStatement)
				if len(stmt.Else) == 1 {
					_, isElseTau := stmt.Else[0].(*migo.TauStatement)
					if isThenTau && isElseTau { // if tau; else tau; endif;
						ss[i] = nil
						ss = append(ss[:i], ss[i+1:]...)
						i--
					}
				}
			}
		case *migo.SpawnStatement:
			if _, found := r.prog.Function(stmt.Name); !found {
				ss[i] = nil
				ss = append(ss[:i], ss[i+1:]...)
				i--
			}
		case *migo.CallStatement:
			if _, found := r.prog.Function(stmt.Name); !found {
				ss[i] = nil
				ss = append(ss[:i], ss[i+1:]...)
				i--
			}
		}
	}
	if len(ss) == 0 {
		ss = []migo.Statement{&migo.TauStatement{}}
	}
	*stmts = ss
}
