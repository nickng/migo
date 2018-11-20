// Package parser is a parser for MiGo (mini-go calculus) types.
//
// A MiGo type can be obtained from an io.Reader by calling the Parse function.
//
//    p := parser.Parse(strings.NewReader("   def main(): send ch;   "))
package parser
