package migo

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nickng/migo/mock_migo"
)

// ErrFuncNotExists is Error if function does not exist in program which is
// should.
type ErrFuncNotExist struct {
	f string
}

func (e *ErrFuncNotExist) Error() string {
	return fmt.Sprintf("Expects %s() in program, but it does not exist", e.f)
}

// ErrFuncExist is Error if function exist in a program which is should not.
type ErrFuncExist struct {
	f string
}

func (e *ErrFuncExist) Error() string {
	return fmt.Sprintf("Expects %s() removed from program, but it exists", e.f)
}

// Tests CleanUp with simple send/recv/work functions.
//
// def main():
//   let ch = newchan ch_instance, 0
//   spawn send(ch)
//   spawn recv(ch)
//   spawn work()
//   recv ch
//   recv ch
// def send(sch):
//   send sch
// def recv(rch):
//   recv rch
//   send rch
// def work:
//
// main, send, recv should remain after CleanUp
//
func TestCleanUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mainCh := mock_migo.NewMockNamedVar(ctrl)
	mainCh.EXPECT().Name().AnyTimes().Return("ch")
	mainCh.EXPECT().String().AnyTimes().Return("ch_instance")
	sendCh := mock_migo.NewMockNamedVar(ctrl)
	sendCh.EXPECT().Name().AnyTimes().Return("sch")
	recvCh := mock_migo.NewMockNamedVar(ctrl)
	recvCh.EXPECT().Name().AnyTimes().Return("rch")

	p := NewProgram()
	mainFunc := NewFunction("main.main")
	mainFunc.AddStmts(
		&NewChanStatement{Name: mainCh, Chan: "ch", Size: 0},
		&SpawnStatement{Name: "send", Params: []*Parameter{
			&Parameter{Caller: mainCh}},
		},
		&SpawnStatement{Name: "recv", Params: []*Parameter{
			&Parameter{Caller: mainCh}},
		},
		&RecvStatement{Chan: "ch"},
		&RecvStatement{Chan: "ch"},
	)
	sendFunc := NewFunction("send")
	sendFunc.AddParams(&Parameter{Caller: mainCh, Callee: sendCh})
	sendFunc.AddStmts(
		&SendStatement{Chan: "sch"},
	)
	recvFunc := NewFunction("recv")
	sendFunc.AddParams(&Parameter{Caller: mainCh, Callee: recvCh})
	recvFunc.AddStmts(
		&RecvStatement{Chan: "rch"},
		&SendStatement{Chan: "rch"},
	)
	workFunc := NewFunction("work")
	p.AddFunction(mainFunc)
	p.AddFunction(sendFunc)
	p.AddFunction(recvFunc)
	p.AddFunction(workFunc)

	if len(p.Funcs) != 4 {
		t.Errorf("Expects 4 functions in program, but got %d", len(p.Funcs))
	}
	for _, nonempty := range []*Function{mainFunc, sendFunc, recvFunc} {
		if nonempty.IsEmpty() {
			t.Errorf("Expects %s() to be non-empty, but got %t",
				nonempty.Name, nonempty.IsEmpty())
		}
	}
	for _, empty := range []*Function{workFunc} {
		if !empty.IsEmpty() {
			t.Errorf("Expects %s() to be empty, but got %t",
				empty.Name, empty.IsEmpty())
		}
	}
	// These should exist
	for _, exist := range []string{"main.main", "send", "recv", "work"} {
		if _, ok := p.Function(exist); !ok {
			t.Error(&ErrFuncNotExist{f: exist})
		}
	}
	p.CleanUp()
	if len(p.Funcs) != 3 {
		t.Errorf("Expects 3 functions in program, but got %d", len(p.Funcs))
	}
	// These should remain
	for _, remain := range []string{"main.main", "send", "recv"} {
		if _, ok := p.Function(remain); !ok {
			t.Error(&ErrFuncNotExist{f: remain})
		}
	}
	// These should be removed
	for _, removed := range []string{"work"} {
		if _, ok := p.Function(removed); ok {
			t.Error(&ErrFuncExist{f: removed})
		}
	}
}

// Tests CleanUp with calls to empty functions.
//
// def main():
//   let ch = newchan ch_instance, 1
//   call work(ch)
// def work(ch):
//   call workwork()
//   spawn work$1(ch)
//   call work$2(ch)
// def workwork():
//   call workworkwork()
// def workworkwork():
// def work$1(ch):
// def work$2(ch):
//   call work$3(ch)
// def work$3(ch):
//   recv ch
//   send ch
//
// main, work, work$2, work$3 should remain after CleanUp
//
func TestCleanUp2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mainCh := mock_migo.NewMockNamedVar(ctrl)
	mainCh.EXPECT().Name().AnyTimes().Return("ch")
	mainCh.EXPECT().String().AnyTimes().Return("ch_instance")
	xCh := mock_migo.NewMockNamedVar(ctrl)
	xCh.EXPECT().Name().AnyTimes().Return("ch")
	xCh.EXPECT().String().AnyTimes().Return("ch_instance")

	p := NewProgram()
	mainFunc := NewFunction("main.main")
	mainFunc.AddStmts(
		&NewChanStatement{Name: mainCh, Chan: "ch", Size: 1},
		&CallStatement{Name: "work", Params: []*Parameter{
			&Parameter{Caller: mainCh, Callee: xCh}},
		},
	)
	workFunc := NewFunction("work")
	workFunc.AddStmts(
		&CallStatement{Name: "workwork"},
		&SpawnStatement{Name: "work$1", Params: []*Parameter{
			&Parameter{Caller: mainCh, Callee: xCh}},
		},
		&SpawnStatement{Name: "work$2", Params: []*Parameter{
			&Parameter{Caller: mainCh, Callee: xCh}},
		},
	)
	workworkFunc := NewFunction("workwork")
	workworkFunc.AddStmts(
		&CallStatement{Name: "workworkwork"},
	)
	workworkworkFunc := NewFunction("workworkwork")
	workClosure := NewFunction("work$1")
	workClosure.AddParams(&Parameter{Caller: mainCh, Callee: xCh})
	workClosure2 := NewFunction("work$2")
	workClosure2.AddParams(&Parameter{Caller: mainCh, Callee: xCh})
	workClosure2.AddStmts(
		&CallStatement{Name: "work$3", Params: []*Parameter{
			&Parameter{Caller: mainCh, Callee: xCh}},
		},
	)
	workClosure3 := NewFunction("work$3")
	workClosure3.AddParams(&Parameter{Caller: mainCh, Callee: xCh})
	workClosure3.AddStmts(
		&SendStatement{Chan: xCh.Name()},
		&RecvStatement{Chan: xCh.Name()},
	)
	p.AddFunction(mainFunc)
	p.AddFunction(workFunc)
	p.AddFunction(workworkFunc)
	p.AddFunction(workworkworkFunc)
	p.AddFunction(workClosure)
	p.AddFunction(workClosure2)
	p.AddFunction(workClosure3)

	if len(p.Funcs) != 7 {
		t.Errorf("Expects 7 functions in program, but got %d", len(p.Funcs))
	}
	for _, nonEmpty := range []*Function{mainFunc, workFunc, workworkFunc, workClosure2, workClosure3} {
		if nonEmpty.IsEmpty() {
			t.Errorf("Expects %s() to be non-empty, but got %t",
				nonEmpty.Name, nonEmpty.IsEmpty())
		}
	}
	for _, empty := range []*Function{workworkworkFunc, workClosure} {
		if !empty.IsEmpty() {
			t.Errorf("Expects %s() to empty, but got %t",
				empty.Name, empty.IsEmpty())

		}
	}
	for _, exist := range []string{"main.main", "work", "workwork", "workworkwork", "work$1", "work$2", "work$3"} {
		if _, ok := p.Function(exist); !ok {
			t.Error(&ErrFuncNotExist{f: exist})
		}
	}
	p.CleanUp()
	if len(p.Funcs) != 4 {
		t.Errorf("Expects 4 functions in program, but got %d", len(p.Funcs))
	}
	// These should remain
	for _, remain := range []string{"main.main", "work", "work$2", "work$3"} {
		if _, ok := p.Function(remain); !ok {
			t.Error(&ErrFuncNotExist{f: remain})
		}
	}
	// These should be removed
	for _, removed := range []string{"workwork", "workworkwork", "work$1"} {
		if _, ok := p.Function(removed); ok {
			t.Error(&ErrFuncExist{f: removed})
		}
	}
}

// Test simplifying name.
func TestSimpleName(t *testing.T) {
	fullGoName := "(github.com/nickng/migo).String#1"
	simpleName := "github.com_nickng_migo.String#1"
	s := &CallStatement{Name: fullGoName, Params: []*Parameter{}}
	if s.SimpleName() != simpleName {
		t.Errorf("SimplifyName of %s should be %s but got %s",
			fullGoName, simpleName, s.SimpleName())
	}
}
