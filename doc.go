// Copyright © 2016 Nicholas Ng <nickng@projectfate.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package migo is a library for working with MiGo (mini-go calculus) types.
// MiGo is a process calculi/type that captures the core concurrency features of
// Go.
//
// MiGo types syntax
//
// This is the output format of MiGo Types in EBNF.
//
//    identifier = [a-zA-Z0-9_.,#/]
//    digit      = [0-9]
//    program    = definition* ;
//    definition = "def " identifier "(" param ")" ":" def-body ;
//    param      =
//               | params
//               ;
//    params     = identifier
//               | params "," identifier
//               ;
//    def-body   = def-stmt+
//               ;
//    prefix     = "send" identifier
//               | "recv" identifier
//               | "tau"
//               ;
//    memprefix  = "read"  identifier
//               | "write" identifier
//               ;
//    def-stmt   = "let" identifier = "newchan" identifier, digit+ ";"
//               | prefix ";"
//               | "letmem" identifier ";"
//               | memprefix ";"
//               | "close" identifier ";"
//               | "call"  identifier "(" params ")" ";"
//               | "spawn" identifier "(" params ")" ";"
//               | "if" def-stmt+ "else" def-stmt+ "endif" ";"
//               | "select" ( "case" prefix ";" def-stmt* )* "endselect" ";"
//               ;
//
// A MiGo type can be obtained by calling String() function of the Program,
// see examples below.
//
//    p := NewProgram()
//    // ... add functions
//    migoType := p.String()
//
package migo // import "github.com/nickng/migo"
