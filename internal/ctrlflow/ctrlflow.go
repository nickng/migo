// Package ctrlflow represents and constructs control-flow graph (CFG)
// of MiGo functions in a MiGo program.
package ctrlflow

import (
	"fmt"
	"log"
	"strings"

	"github.com/nickng/migo"
)

// Graph is a control flow graph.
type Graph struct {
	Nodes []*Node

	prog *migo.Program
}

func (g *Graph) addNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

// addEdge creates an edge from Node n1 to Node n2.
func (g *Graph) addEdge(n1, n2 *Node) {
	childExists := false
	for _, c := range n1.Succs {
		if c == n2 {
			childExists = true
			break
		}
	}
	if !childExists {
		n1.Succs = append(n1.Succs, n2)
	}
	parentExists := false
	for _, p := range n2.Preds {
		if p == n1 {
			parentExists = true
			break
		}
	}
	if !parentExists {
		n2.Preds = append(n2.Preds, n1)
	}
}

func (g *Graph) String() string {
	var sb strings.Builder
	for _, n := range g.Nodes {
		sb.WriteString(n.String())
	}
	return sb.String()
}

// DotString returns a string representation the graph in dot format.
func (g *Graph) DotString() string {
	name := make(map[*Node]string)
	printed := make(map[*Node]map[*Node]bool)
	for _, n := range g.Nodes {
		name[n] = strings.Replace(n.fn.SimpleName(), ".", "_", -1)
		printed[n] = make(map[*Node]bool)
	}

	var sb strings.Builder
	sb.WriteString("digraph G {\n")
	for _, n := range g.Nodes {
		sb.WriteString(fmt.Sprintf("%s [label=\"%s\"];\n", name[n], n.fn.Name))
		for _, p := range n.Preds {
			if _, ok := printed[p][n]; !ok {
				sb.WriteString(fmt.Sprintf("%s -> %s;\n", name[p], name[n]))
				printed[p][n] = true
			}
		}
		for _, s := range n.Succs {
			if _, ok := printed[n][s]; !ok {
				sb.WriteString(fmt.Sprintf("%s -> %s;\n", name[n], name[s]))
				printed[n][s] = true
			}
		}
	}
	sb.WriteString("}\n")
	return sb.String()
}

// Node is a CFG node (function) for a migo program.
type Node struct {
	Preds []*Node
	Succs []*Node

	fn *migo.Function
}

// Func returns the function the Node n
// represents in the CFG.
func (n Node) Func() *migo.Function {
	return n.fn
}

func (n *Node) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s:\n", n.fn.SimpleName()))
	preds := make([]string, len(n.Preds))
	for i := range n.Preds {
		preds[i] = n.Preds[i].fn.SimpleName()
	}
	if n.Preds != nil {
		sb.WriteString(fmt.Sprintf("\tFROM → %s\n", strings.Join(preds, ", ")))
	}
	succs := make([]string, len(n.Succs))
	for i := range n.Succs {
		succs[i] = n.Succs[i].fn.SimpleName()
	}
	if n.Succs != nil {
		sb.WriteString(fmt.Sprintf("\tTO   ⇒ %s\n", strings.Join(succs, ", ")))
	}
	return sb.String()
}

// NewGraph returns a new CFG given MiGo program prog.
func NewGraph(prog *migo.Program) *Graph {
	b := builder{
		graph:   &Graph{prog: prog},
		visited: make(map[*migo.Function]bool),
	}
	for _, f := range b.graph.prog.Funcs {
		b.visit(f)
	}
	return b.graph
}

// builder is a data structure to build a CFG.
type builder struct {
	graph *Graph

	// temporary data for building graph.
	nodes   map[*migo.Function]*Node
	visited map[*migo.Function]bool
}

func (b *builder) visit(fn *migo.Function) {
	if completed, started := b.visited[fn]; started { // key exists, visit started
		_ = completed
		return
	}
	b.visited[fn] = false // visit started
	if b.nodes == nil {
		b.nodes = make(map[*migo.Function]*Node)
	}
	n := &Node{fn: fn}
	b.nodes[fn] = n
	b.graph.addNode(n)
	b.visitStmts(fn, fn.Stmts) // visit body
	b.visited[fn] = true       // visit complete
}

func (b *builder) visitStmts(parent *migo.Function, stmts []migo.Statement) {
	for _, stmt := range stmts {
		switch stmt := stmt.(type) {
		case *migo.NewChanStatement, *migo.CloseStatement:
			// channel operations

		case *migo.SendStatement, *migo.RecvStatement:
			// message passing operations

		case *migo.TauStatement:
			// no-op

		case *migo.SelectStatement:
			for _, stmts := range stmt.Cases {
				b.visitStmts(parent, stmts)
			}

		case *migo.IfStatement:
			b.visitStmts(parent, stmt.Then)
			b.visitStmts(parent, stmt.Else)

		case *migo.IfForStatement:
			b.visitStmts(parent, stmt.Then)
			b.visitStmts(parent, stmt.Else)

		case *migo.CallStatement:
			if fn, found := b.graph.prog.Function(stmt.Name); found {
				b.visit(fn)
				b.graph.addEdge(b.nodes[parent], b.nodes[fn])
			}

		case *migo.SpawnStatement:
			if fn, found := b.graph.prog.Function(stmt.Name); found {
				b.visit(fn)
				b.graph.addEdge(b.nodes[parent], b.nodes[fn])
			}

		case *migo.NewMem, *migo.MemRead, *migo.MemWrite:
			// no-op

		default:
			log.Fatal(fmt.Errorf("ctrlflow: statement kind not found: %T", stmt))
		}
	}
}
