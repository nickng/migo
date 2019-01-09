package parser

// Helper functions for yacc parser
// These functions wrap MiGo AST

import "github.com/nickng/migo"

func sendStmt(ch string) *migo.SendStatement {
	return &migo.SendStatement{Chan: ch}
}

func recvStmt(ch string) *migo.RecvStatement {
	return &migo.RecvStatement{Chan: ch}
}

func tauStmt() *migo.TauStatement {
	return &migo.TauStatement{}
}

func newchanStmt(name, ch string, size int) migo.Statement {
	return &migo.NewChanStatement{
		Name: &plainNamedVar{s: name},
		Chan: ch,
		Size: int64(size),
	}
}

func readStmt(name string) migo.Statement {
	return &migo.MemRead{Name: name}
}

func writeStmt(name string) migo.Statement {
	return &migo.MemWrite{Name: name}
}

func newmemStmt(name string) migo.Statement {
	return &migo.NewMem{Name: name}
}

func closeStmt(ch string) migo.Statement {
	return &migo.CloseStatement{Chan: ch}
}

func callStmt(fn string, params []*migo.Parameter) migo.Statement {
	return &migo.CallStatement{Name: fn, Params: params}
}

func spawnStmt(fn string, params []*migo.Parameter) migo.Statement {
	return &migo.SpawnStatement{Name: fn, Params: params}
}

func params(p ...*migo.Parameter) []*migo.Parameter {
	return p
}

func plainParam(name string) *migo.Parameter {
	return &migo.Parameter{Caller: &plainNamedVar{s: name}, Callee: &plainNamedVar{s: name}}
}

func ifStmt(iftrue, iffalse []migo.Statement) migo.Statement {
	return &migo.IfStatement{Then: iftrue, Else: iffalse}
}

func selectStmt(cases [][]migo.Statement) migo.Statement {
	return &migo.SelectStatement{Cases: cases}
}

func cases(c ...[]migo.Statement) [][]migo.Statement {
	return c
}

func stmts(s ...migo.Statement) []migo.Statement {
	return s
}
