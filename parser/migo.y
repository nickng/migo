%{
package parser

import (
	"io"

	"github.com/nickng/migo/v3"
)

var prog *migo.Program
%}

%union {
	str    string
	num    int
	prog   *migo.Program
	fun    *migo.Function
	stmt   migo.Statement
	stmts  []migo.Statement
	params []*migo.Parameter
	cases  [][]migo.Statement
}

%token tCOMMA tDEF tEQ tLPAREN tRPAREN tCOLON tSEMICOLON
%token tCALL tSPAWN tCASE tCLOSE tELSE tENDIF tENDSELECT tIF tLET tNEWCHAN tSELECT tSEND tRECV tTAU tLETMEM tREAD tWRITE tLETSYNC tMUTEX tLOCK tUNLOCK
%token <str> tIDENT
%token <num> tDIGITS
%type <stmt> prefix memprefix mutexprefix stmt
%type <fun> def
%type <params> params
%type <stmts> stmts defbody
%type <cases> cases
%type <prog> prog


%%

prog :      def { prog = migo.NewProgram(); $$ = prog; prog.AddFunction($1) }
     | prog def { $1.AddFunction($2) }
     ;

def : tDEF tIDENT tLPAREN params tRPAREN tCOLON defbody { $$ = migo.NewFunction($2); $$.AddParams($4...); $$.AddStmts($7...) }
    ;

params :                      { $$ = params() }
       |               tIDENT { $$ = params(plainParam($1)) }
       | params tCOMMA tIDENT { $$ = append($1, plainParam($3)) }

/* one or more */
defbody :         stmt { $$ = stmts($1) }
        | defbody stmt { $$ = append($1, $2) }

/* zero or more */
stmts :            { $$ = stmts() }
      | stmts stmt { $$ = append($1, $2) }
      ;

prefix : tSEND tIDENT { $$ = sendStmt($2) }
       | tRECV tIDENT { $$ = recvStmt($2) }
       | tTAU         { $$ = tauStmt() }
       ;

memprefix : tREAD  tIDENT { $$ = readStmt($2) }
          | tWRITE tIDENT { $$ = writeStmt($2) }
          ;

mutexprefix : tLOCK   tIDENT { $$ = lockStmt($2) }
            | tUNLOCK tIDENT { $$ = unlockStmt($2) }
            ;

stmt : tLET tIDENT tEQ tNEWCHAN tIDENT tCOMMA tDIGITS tSEMICOLON { $$ = newchanStmt($2, $5, $7) }
     | prefix                                         tSEMICOLON { $$ = $1 }
     | tLETMEM tIDENT                                 tSEMICOLON { $$ = newmemStmt($2) }
     | memprefix                                      tSEMICOLON { $$ = $1 }
     | tLETSYNC tIDENT tMUTEX                         tSEMICOLON { $$ = newMutex($2) }
     | mutexprefix                                    tSEMICOLON { $$ = $1 }
     | tCLOSE tIDENT                                  tSEMICOLON { $$ = closeStmt($2) }
     | tCALL  tIDENT tLPAREN params tRPAREN           tSEMICOLON { $$ = callStmt($2, $4) }
     | tSPAWN tIDENT tLPAREN params tRPAREN           tSEMICOLON { $$ = spawnStmt($2, $4) }
     | tIF stmts tELSE stmts tENDIF                   tSEMICOLON { $$ = ifStmt($2, $4) }
     | tSELECT cases tENDSELECT                       tSEMICOLON { $$ = selectStmt($2) }
     ;

cases :                                     { $$ = cases() }
      | cases tCASE prefix tSEMICOLON stmts { $$ = append($1, append(stmts($3), $5...)) }
      ;

%%

// Parse is the entry point to the migo type parser
func Parse(r io.Reader) (*migo.Program, error) {
	l := NewLexer(r)
	migoParse(l)
	select {
	case err := <-l.Errors:
		return nil, err
	default:
		return prog, nil
	}
}
