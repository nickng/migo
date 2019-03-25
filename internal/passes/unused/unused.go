// Package unused defines a transformation pass to remove unused functions.
//
// The transformation works by recursively removing control flow graph
// nodes that have no predecessor (caller), until all nodes in the graph
// are called. A caveat is cycles in the graph are not removed.
package unused

import (
	"github.com/nickng/migo/internal/ctrlflow"
	"github.com/nickng/migo/v3"
)

// Remove removes all unused functions from Program prog except entry.
func Remove(prog *migo.Program, entry *migo.Function) {
	removeQ := findUnusedToplevel(prog)
	var n *ctrlflow.Node
	for len(removeQ) > 0 {
		n, removeQ = removeQ[0], removeQ[1:]
		if n.Func() == entry { // skip
			continue
		}
		// remove this node from the successors
		for _, s := range n.Succs {
			for i, p := range s.Preds {
				if p.Func() == n.Func() {
					// remove n from s.Preds
					s.Preds = append(s.Preds[:i], s.Preds[i+1:]...)
				}
			}
			// If n is the only predecessor of successor s, remove s too
			if len(s.Preds) == 0 {
				removeQ = append(removeQ, s)
			}
		}
		// remove function
		for i := range prog.Funcs {
			if prog.Funcs[i] == n.Func() {
				prog.Funcs = append(prog.Funcs[:i], prog.Funcs[i+1:]...)
				break
			}
		}
	}
}

// findUnusedToplevel finds all unused toplevel functions.
func findUnusedToplevel(prog *migo.Program) []*ctrlflow.Node {
	var emptyNodes []*ctrlflow.Node
	graph := ctrlflow.NewGraph(prog)
	for _, node := range graph.Nodes {
		if len(node.Preds) == 0 {
			emptyNodes = append(emptyNodes, node)
		}
	}
	return emptyNodes
}
