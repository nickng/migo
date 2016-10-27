package migo

// Plain NamedVar for use with parser.

type plainNamedVar struct {
	s string
}

func (v *plainNamedVar) Name() string {
	return v.s
}

func (v *plainNamedVar) String() string {
	return v.s
}
