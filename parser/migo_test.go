package parser

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	want := `def main():
    let ch = newchan T, 0;
    send ch;
`
	p, err := Parse(strings.NewReader(`   def main(): let ch = newchan T, 0;
	send ch;   `))
	if err != nil {
		t.Errorf("cannot parse: %v", err)
	}
	if got := p.String(); want != got {
		t.Errorf("unexpected parsed migo, want:\n%sgot:\n%s", want, got)
	}
}
