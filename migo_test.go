package migo_test

import (
	"strings"
	"testing"

	"github.com/nickng/migo"
	"github.com/nickng/migo/parser"
)

// Test simplifying name.
func TestSimpleName(t *testing.T) {
	fullGoName := "(github.com/nickng/migo).String#1"
	simpleName := "github.com_nickng_migo.String#1"
	s := &migo.CallStatement{Name: fullGoName, Params: []*migo.Parameter{}}
	if s.SimpleName() != simpleName {
		t.Errorf("SimplifyName of %s should be %s but got %s",
			fullGoName, simpleName, s.SimpleName())
	}
}
