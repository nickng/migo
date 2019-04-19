package parser

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

// Scanner is a lexical scanner.
type Scanner struct {
	r   *bufio.Reader
	pos TokenPos
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r), pos: TokenPos{Char: 0, Lines: []int{}}}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if reached the end or error occurs.
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	if ch == '\n' {
		s.pos.Lines = append(s.pos.Lines, s.pos.Char)
		s.pos.Char = 0
	} else {
		s.pos.Char++
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
	if s.pos.Char == 0 {
		s.pos.Char = s.pos.Lines[len(s.pos.Lines)-1]
		s.pos.Lines = s.pos.Lines[:len(s.pos.Lines)-1]
	} else {
		s.pos.Char--
	}
}

// Scan returns the next token and parsed value.
func (s *Scanner) Scan() Token {
	var startPos, endPos TokenPos
	ch := s.read()

	if isWhitespace(ch) {
		s.skipWhitespace()
		ch = s.read()
	}
	if isIdent(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Track token positions.
	startPos = s.pos
	defer func() { endPos = s.pos }()

	switch ch {
	case eof:
		return &ConstToken{t: 0, start: startPos, end: endPos}
	case ':':
		return &ConstToken{t: tCOLON, start: startPos, end: endPos}
	case ';':
		return &ConstToken{t: tSEMICOLON, start: startPos, end: endPos}
	case ',':
		return &ConstToken{t: tCOMMA, start: startPos, end: endPos}
	case '(':
		return &ConstToken{t: tLPAREN, start: startPos, end: endPos}
	case ')':
		return &ConstToken{t: tRPAREN, start: startPos, end: endPos}
	case '=':
		return &ConstToken{t: tEQ, start: startPos, end: endPos}
	case '-':
		if ch2 := s.read(); ch2 == '-' {
			s.unread()
			s.unread()
			s.skipComment()
			return s.Scan()
		}
	}
	return &ConstToken{t: tILLEGAL, start: startPos, end: endPos}
}

func (s *Scanner) scanIdent() Token {
	var startPos, endPos TokenPos
	var buf bytes.Buffer

	startPos = s.pos
	defer func() { endPos = s.pos }()

	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isIdent(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	switch buf.String() {
	case "def":
		return &ConstToken{t: tDEF, start: startPos, end: endPos}
	case "call":
		return &ConstToken{t: tCALL, start: startPos, end: endPos}
	case "spawn":
		return &ConstToken{t: tSPAWN, start: startPos, end: endPos}
	case "case":
		return &ConstToken{t: tCASE, start: startPos, end: endPos}
	case "close":
		return &ConstToken{t: tCLOSE, start: startPos, end: endPos}
	case "else":
		return &ConstToken{t: tELSE, start: startPos, end: endPos}
	case "endif":
		return &ConstToken{t: tENDIF, start: startPos, end: endPos}
	case "endselect":
		return &ConstToken{t: tENDSELECT, start: startPos, end: endPos}
	case "if":
		return &ConstToken{t: tIF, start: startPos, end: endPos}
	case "let":
		return &ConstToken{t: tLET, start: startPos, end: endPos}
	case "newchan":
		return &ConstToken{t: tNEWCHAN, start: startPos, end: endPos}
	case "select":
		return &ConstToken{t: tSELECT, start: startPos, end: endPos}
	case "send":
		return &ConstToken{t: tSEND, start: startPos, end: endPos}
	case "recv":
		return &ConstToken{t: tRECV, start: startPos, end: endPos}
	case "tau":
		return &ConstToken{t: tTAU, start: startPos, end: endPos}
	case "letmem":
		return &ConstToken{t: tLETMEM, start: startPos, end: endPos}
	case "read":
		return &ConstToken{t: tREAD, start: startPos, end: endPos}
	case "write":
		return &ConstToken{t: tWRITE, start: startPos, end: endPos}
	case "letsync":
		return &ConstToken{t: tLETSYNC, start: startPos, end: endPos}
	case "mutex":
		return &ConstToken{t: tMUTEX, start: startPos, end: endPos}
	case "rwmutex":
		return &ConstToken{t: tRWMUTEX, start: startPos, end: endPos}
	case "lock":
		return &ConstToken{t: tLOCK, start: startPos, end: endPos}
	case "unlock":
		return &ConstToken{t: tUNLOCK, start: startPos, end: endPos}
	case "rlock":
		return &ConstToken{t: tRLOCK, start: startPos, end: endPos}
	case "runlock":
		return &ConstToken{t: tRUNLOCK, start: startPos, end: endPos}
	}

	if i, err := strconv.Atoi(buf.String()); err == nil {
		return &DigitsToken{num: i, start: startPos, end: endPos}
	}
	return &IdentToken{str: buf.String(), start: startPos, end: endPos}
}

func (s *Scanner) skipComment() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == '\n' {
			break
		}
	}
}

func (s *Scanner) skipWhitespace() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		}
	}
}
