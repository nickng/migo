// Package taufunc defines a transformation pass to remove τ functions.
//
// τ functions are functions where the bodies are empty or just τ-actions.
// The transformation algorithm is as follows:
//
//     Build control flow graph of given program
//     Foreach function:
//     	Mark τ if function body does not contain non control flow primitives
//     Repeat until no changes:
//     	Foreach non-τ function:
//     		If CFG parent is τ: Mark parent non-τ
//     Remove all function definitions marked as τ
//
// Whether a primitive is considered a τ or not is defined by the isTau method.
//
// Usage
//
// To remove all tau functions:
//
//     taufunc.Find(prog, taufunc.Remove)
//
package taufunc

// This file contains the implementation of transformation which
// removes empty (i.e. no communication) migo Functions from migo Programs.

import (
	"fmt"
	"log"

	"github.com/nickng/migo"
	"github.com/nickng/migo/internal/ctrlflow"
)

// Find finds function definitions from Program prog
// if its body is reducible to tau (no communication).
//
// visitTauFn is then applied to each functions found,
// the return value indicates if fn should be removed.
//
// If fn is marked remove, prog will be updated accordingly.
func Find(prog *migo.Program, visitTauFn func(fn *migo.Function) bool) {
	tff := &tauFuncFinder{
		graph: ctrlflow.NewGraph(prog),
		istau: make(map[*ctrlflow.Node]bool),
	}
	func2node := make(map[*migo.Function]*ctrlflow.Node)
	for _, node := range tff.graph.Nodes {
		func2node[node.Func()] = node
		tff.taintTau(node)
	}
	tff.propagate()
	for i := 0; i < len(prog.Funcs); i++ {
		if tff.istau[func2node[prog.Funcs[i]]] {
			if visitTauFn != nil {
				if remove := visitTauFn(prog.Funcs[i]); remove {
					prog.Funcs[i] = nil
					prog.Funcs = append(prog.Funcs[:i], prog.Funcs[i+1:]...)
					i--
				}
			}
		}
	}
}

// Remove marks taufn to be removed from its parent Program.
func Remove(taufn *migo.Function) (remove bool) {
	return true
}

// RemoveExcept returns a custom RemoveAll function that
// marks taufn to be removed unless taufn is exceptfn (an exception).
func RemoveExcept(exceptfn *migo.Function) func(taufn *migo.Function) (remove bool) {
	return func(taufn *migo.Function) (removed bool) {
		return exceptfn != taufn // only remove if is not excluded function
	}
}

type tauFuncFinder struct {
	graph *ctrlflow.Graph
	istau map[*ctrlflow.Node]bool
}

func (t *tauFuncFinder) taintTau(n *ctrlflow.Node) {
	t.istau[n] = t.isTau(n, n.Func().Stmts)
}

// isTau inspects Statements stmts and
// returns true if all statements can be reduced to tau.
func (t *tauFuncFinder) isTau(n *ctrlflow.Node, stmts []migo.Statement) bool {
	var istainted bool
	for _, stmt := range stmts {
		switch stmt := stmt.(type) {
		case *migo.NewChanStatement, *migo.CloseStatement:
			istainted = true

		case *migo.SendStatement, *migo.RecvStatement:
			istainted = true

		case *migo.SelectStatement:
			// no need to traverse into cases
			istainted = true

		case *migo.TauStatement:

		case *migo.IfStatement:
			istainted = istainted || !t.isTau(n, stmt.Then)
			istainted = istainted || !t.isTau(n, stmt.Else)

		case *migo.IfForStatement:
			istainted = istainted || !t.isTau(n, stmt.Then)
			istainted = istainted || !t.isTau(n, stmt.Else)

		case *migo.CallStatement, *migo.SpawnStatement:
			// skip for now

		case *migo.NewMem, *migo.MemRead, *migo.MemWrite:
			istainted = true

		default:
			log.Fatal(fmt.Errorf("passes/taufunc: statement kind not found: %T", stmt))
		}
	}
	return !istainted
}

// progpagate taints caller (CFG parent) of non-tau functions to be
// non-tau function. Repeats until all non-tau functions are found.
func (t *tauFuncFinder) propagate() {
	var changed bool
	for {
		for _, node := range t.graph.Nodes {
			if !t.istau[node] {
				for _, pred := range node.Preds {
					if t.istau[pred] {
						t.istau[pred] = false
						changed = true
					}
				}
			}
		}
		if !changed {
			break // done
		}
		changed = false
	}
}
