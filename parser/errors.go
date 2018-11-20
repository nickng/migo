package parser

import "fmt"

// ErrParse is a parse error.
type ErrParse struct {
	Pos TokenPos
	Err string // Error string returned from parser.
}

func (e *ErrParse) Error() string {
	return fmt.Sprintf("Parse failed at %s: %s", e.Pos, e.Err)
}
