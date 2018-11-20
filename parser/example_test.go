package parser

import (
	"fmt"
	"log"
	"strings"
)

// This example demonstrates the use of the Parser.
// The output should be exactly the same as input (but pretty printed).
func ExampleParse() {
	s := `def main.main(): let ch = newchan ch, 0; spawn main.sndr(ch); recv ch;
	def main.sndr(ch): send ch;`
	r := strings.NewReader(s)
	parsed, err := Parse(r)
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
