package migo

//go:generate go tool yacc -p migo -o parser.y.go migo.y

import "io"

// Lexer for migo.
type Lexer struct {
	scanner *Scanner
	Errors  chan error
}

// NewLexer returns a new yacc-compatible lexer.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{scanner: NewScanner(r), Errors: make(chan error, 1)}
}

// Lex is provided for yacc-compatible parser.
func (l *Lexer) Lex(yylval *migoSymType) int {
	token := l.scanner.Scan()
	switch token := token.(type) {
	case *DigitsToken:
		yylval.num = token.num
	case *IdentToken:
		yylval.str = token.str
	}
	return int(token.Tok())
}

// Error handles error.
func (l *Lexer) Error(err string) {
	l.Errors <- &ErrParse{Err: err, Pos: l.scanner.pos}
}
