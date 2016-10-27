package migo_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/nickng/migo"
)

func ExampleStmtsStack() {
	b := []migo.Statement{}
	s := migo.NewStmtsStack() // Create a new stack
	s.Push(b)                 // Push to empty stack
	b, err := s.Pop()         // Pop from stack (stack is empty again)
	if err != nil {
		fmt.Println("error:", err)
	}
	// Output:
}

// The example demonstrates the usage of the migo API for building MiGo programs.
func ExampleProgram() {
	p := migo.NewProgram()
	f := migo.NewFunction("F")
	SendXStmt := &migo.SendStatement{Chan: "x"}                              // send x
	callGStmt := &migo.CallStatement{Name: "G", Params: []*migo.Parameter{}} // call G()
	f.AddStmts(SendXStmt, callGStmt)                                         // F()
	g := migo.NewFunction("G")
	g.AddParams()                    // G()
	g.AddStmts(&migo.TauStatement{}) // tau
	p.AddFunction(f)
	p.AddFunction(g)
	fmt.Print(p.String())
	// Output:
	// def F():
	//     send x;
	//     call G();
	// def G():
	//     tau;
}

// This example demonstrates the usage of the CleanUp function to remove
// unwanted empty functions.
func ExampleProgram_CleanUp() {
	p := migo.NewProgram()
	f := migo.NewFunction("F")
	SendXStmt := &migo.SendStatement{Chan: "x"}                              // send x
	callGStmt := &migo.CallStatement{Name: "G", Params: []*migo.Parameter{}} // call G()
	f.AddStmts(SendXStmt, callGStmt)                                         // F()
	g := migo.NewFunction("G")
	g.AddParams()                    // G()
	g.AddStmts(&migo.TauStatement{}) // tau
	p.AddFunction(f)                 // Note that calling G() will be removed.
	p.AddFunction(g)                 // Note that G() is an empty function.
	p.CleanUp()
	fmt.Print(p.String())
	// Output:
	// def F():
	//     send x;
}

// This example demonstrates the use of the Parser.
// The output should be exactly the same as input (but pretty printed).
func ExampleParse() {
	s := `def main.main(): let ch = newchan ch, 0; spawn main.sndr(ch); recv ch;
	def main.sndr(ch): send ch;`
	r := strings.NewReader(s)
	parsed, err := migo.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parsed.String())
	// Output:
	// def main.main():
	//     let ch = newchan ch, 0;
	//     spawn main.sndr(ch);
	//     recv ch;
	// def main.sndr(ch):
	//     send ch;
}
