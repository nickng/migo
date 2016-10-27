%{
package migo

import (
	"io"
)

var prog *Program
%}

%union {
	str    string
	num    int
	prog   *Program
	fun    *Function
	stmt   Statement
	stmts  []Statement
	params []*Parameter
	cases  [][]Statement
}

%token COMMA DEF EQ LPAREN RPAREN COLON SEMICOLON
%token CALL SPAWN CASE CLOSE ELSE ENDIF ENDSELECT IF LET NEWCHAN SELECT SEND RECV TAU
%token <str> IDENT
%token <num> DIGITS
%type <stmt> prefix stmt
%type <fun> def
%type <params> params
%type <stmts> stmts defbody
%type <cases> cases
%type <prog> prog


%%

prog :      def { prog = NewProgram(); $$ = prog; prog.AddFunction($1) }
     | prog def { $1.AddFunction($2) }
     ;

def : DEF IDENT LPAREN params RPAREN COLON defbody { $$ = NewFunction($2); $$.AddParams($4...); $$.AddStmts($7...) }
    ;

params :                    { $$ = []*Parameter{} }
       |              IDENT { $$ = []*Parameter{&Parameter{Caller: &plainNamedVar{s:$1}, Callee: &plainNamedVar{s:$1}}} }
       | params COMMA IDENT { $$ = append($1, &Parameter{Caller: &plainNamedVar{s:$3}, Callee: &plainNamedVar{s:$3}}) }

/* one or more */
defbody :         stmt { $$ = []Statement{$1} }
        | defbody stmt { $$ = append($1, $2) }

/* zero or more */
stmts :            { $$ = []Statement{} }
      | stmts stmt { $$ = append($1, $2) }
      ;

prefix : SEND IDENT { $$ = &SendStatement{Chan: $2} }
       | RECV IDENT { $$ = &RecvStatement{Chan: $2} }
       | TAU        { $$ = &TauStatement{} }
       ;

stmt : LET IDENT EQ NEWCHAN IDENT COMMA DIGITS SEMICOLON { $$ = &NewChanStatement{Name: &plainNamedVar{s:$2}, Chan: $2, Size: int64($7)} }
     | prefix                                  SEMICOLON { $$ = $1 }
     | CLOSE IDENT                             SEMICOLON { $$ = &CloseStatement{Chan: $2} }
     | CALL  IDENT LPAREN params RPAREN        SEMICOLON { $$ = &CallStatement{Name: $2, Params: $4} }
     | SPAWN IDENT LPAREN params RPAREN        SEMICOLON { $$ = &SpawnStatement{Name: $2, Params: $4} }
     | IF stmts ELSE stmts ENDIF               SEMICOLON { $$ = &IfStatement{Then:$2, Else:$4} }
     | SELECT cases ENDSELECT                  SEMICOLON { $$ = &SelectStatement{Cases: $2} }
     ;

cases :                                   { $$ = [][]Statement{} }
      | cases CASE prefix SEMICOLON stmts { $$ = append($1, append([]Statement{$3}, $5...)) }
      ;

%%

// Parse is the entry point to the migo type parser
func Parse(r io.Reader) (*Program, error) {
	l := NewLexer(r)
	migoParse(l)
	select {
	case err := <-l.Errors:
		return nil, err
	default:
		return prog, nil
	}
}
