//line migo.y:2
package parser

import __yyfmt__ "fmt"

//line migo.y:2
import (
	"github.com/nickng/migo"
	"io"
)

var prog *migo.Program

//line migo.y:12
type migoSymType struct {
	yys    int
	str    string
	num    int
	prog   *migo.Program
	fun    *migo.Function
	stmt   migo.Statement
	stmts  []migo.Statement
	params []*migo.Parameter
	cases  [][]migo.Statement
}

const tCOMMA = 57346
const tDEF = 57347
const tEQ = 57348
const tLPAREN = 57349
const tRPAREN = 57350
const tCOLON = 57351
const tSEMICOLON = 57352
const tCALL = 57353
const tSPAWN = 57354
const tCASE = 57355
const tCLOSE = 57356
const tELSE = 57357
const tENDIF = 57358
const tENDSELECT = 57359
const tIF = 57360
const tLET = 57361
const tNEWCHAN = 57362
const tSELECT = 57363
const tSEND = 57364
const tRECV = 57365
const tTAU = 57366
const tIDENT = 57367
const tDIGITS = 57368

var migoToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tCOMMA",
	"tDEF",
	"tEQ",
	"tLPAREN",
	"tRPAREN",
	"tCOLON",
	"tSEMICOLON",
	"tCALL",
	"tSPAWN",
	"tCASE",
	"tCLOSE",
	"tELSE",
	"tENDIF",
	"tENDSELECT",
	"tIF",
	"tLET",
	"tNEWCHAN",
	"tSELECT",
	"tSEND",
	"tRECV",
	"tTAU",
	"tIDENT",
	"tDIGITS",
}
var migoStatenames = [...]string{}

const migoEofCode = 1
const migoErrCode = 2
const migoInitialStackSize = 16

//line migo.y:75
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

//line yacctab:1
var migoExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const migoPrivate = 57344

const migoLast = 83

var migoAct = [...]int{

	31, 16, 18, 19, 7, 17, 59, 52, 49, 20,
	15, 8, 21, 22, 23, 24, 18, 19, 34, 17,
	40, 33, 30, 20, 15, 29, 21, 22, 23, 24,
	22, 23, 24, 43, 28, 26, 12, 5, 60, 57,
	56, 46, 44, 45, 48, 18, 19, 55, 17, 53,
	39, 47, 20, 15, 58, 21, 22, 23, 24, 42,
	36, 27, 14, 41, 25, 10, 10, 11, 10, 51,
	50, 35, 9, 38, 37, 6, 3, 54, 2, 1,
	4, 32, 13,
}
var migoPact = [...]int{

	71, 71, -1000, 12, -1000, 68, -14, 64, -1000, 58,
	11, 34, -1000, 34, -1000, 10, 51, 9, 0, -3,
	-1000, -1000, -4, -7, -1000, -1000, 65, -1000, 50, 67,
	66, 5, 46, -1000, -1000, 13, -1000, -14, -14, -1000,
	-1000, 41, 8, -17, 62, 61, -9, -1000, 39, 73,
	37, 30, 29, -1000, -20, -1000, -1000, -1000, 34, 28,
	-1000,
}
var migoPgo = [...]int{

	0, 1, 50, 78, 4, 0, 82, 81, 79,
}
var migoR1 = [...]int{

	0, 8, 8, 3, 4, 4, 4, 6, 6, 5,
	5, 1, 1, 1, 2, 2, 2, 2, 2, 2,
	2, 7, 7,
}
var migoR2 = [...]int{

	0, 1, 2, 7, 0, 1, 3, 1, 2, 0,
	2, 2, 2, 1, 8, 2, 3, 6, 6, 6,
	4, 0, 5,
}
var migoChk = [...]int{

	-1000, -8, -3, 5, -3, 25, 7, -4, 25, 8,
	4, 9, 25, -6, -2, 19, -1, 14, 11, 12,
	18, 21, 22, 23, 24, -2, 25, 10, 25, 25,
	25, -5, -7, 25, 25, 6, 10, 7, 7, -2,
	15, 17, 13, 20, -4, -4, -5, 10, -1, 25,
	8, 8, 16, 10, 4, 10, 10, 10, -5, 26,
	10,
}
var migoDef = [...]int{

	0, -2, 1, 0, 2, 0, 4, 0, 5, 0,
	0, 0, 6, 3, 7, 0, 0, 0, 0, 0,
	9, 21, 0, 0, 13, 8, 0, 15, 0, 0,
	0, 0, 0, 11, 12, 0, 16, 4, 4, 10,
	9, 0, 0, 0, 0, 0, 0, 20, 0, 0,
	0, 0, 0, 9, 0, 17, 18, 19, 22, 0,
	14,
}
var migoTok1 = [...]int{

	1,
}
var migoTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26,
}
var migoTok3 = [...]int{
	0,
}

var migoErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	migoDebug        = 0
	migoErrorVerbose = false
)

type migoLexer interface {
	Lex(lval *migoSymType) int
	Error(s string)
}

type migoParser interface {
	Parse(migoLexer) int
	Lookahead() int
}

type migoParserImpl struct {
	lval  migoSymType
	stack [migoInitialStackSize]migoSymType
	char  int
}

func (p *migoParserImpl) Lookahead() int {
	return p.char
}

func migoNewParser() migoParser {
	return &migoParserImpl{}
}

const migoFlag = -1000

func migoTokname(c int) string {
	if c >= 1 && c-1 < len(migoToknames) {
		if migoToknames[c-1] != "" {
			return migoToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func migoStatname(s int) string {
	if s >= 0 && s < len(migoStatenames) {
		if migoStatenames[s] != "" {
			return migoStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func migoErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !migoErrorVerbose {
		return "syntax error"
	}

	for _, e := range migoErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + migoTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := migoPact[state]
	for tok := TOKSTART; tok-1 < len(migoToknames); tok++ {
		if n := base + tok; n >= 0 && n < migoLast && migoChk[migoAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if migoDef[state] == -2 {
		i := 0
		for migoExca[i] != -1 || migoExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; migoExca[i] >= 0; i += 2 {
			tok := migoExca[i]
			if tok < TOKSTART || migoExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if migoExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += migoTokname(tok)
	}
	return res
}

func migolex1(lex migoLexer, lval *migoSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = migoTok1[0]
		goto out
	}
	if char < len(migoTok1) {
		token = migoTok1[char]
		goto out
	}
	if char >= migoPrivate {
		if char < migoPrivate+len(migoTok2) {
			token = migoTok2[char-migoPrivate]
			goto out
		}
	}
	for i := 0; i < len(migoTok3); i += 2 {
		token = migoTok3[i+0]
		if token == char {
			token = migoTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = migoTok2[1] /* unknown char */
	}
	if migoDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", migoTokname(token), uint(char))
	}
	return char, token
}

func migoParse(migolex migoLexer) int {
	return migoNewParser().Parse(migolex)
}

func (migorcvr *migoParserImpl) Parse(migolex migoLexer) int {
	var migon int
	var migoVAL migoSymType
	var migoDollar []migoSymType
	_ = migoDollar // silence set and not used
	migoS := migorcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	migostate := 0
	migorcvr.char = -1
	migotoken := -1 // migorcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		migostate = -1
		migorcvr.char = -1
		migotoken = -1
	}()
	migop := -1
	goto migostack

ret0:
	return 0

ret1:
	return 1

migostack:
	/* put a state and value onto the stack */
	if migoDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", migoTokname(migotoken), migoStatname(migostate))
	}

	migop++
	if migop >= len(migoS) {
		nyys := make([]migoSymType, len(migoS)*2)
		copy(nyys, migoS)
		migoS = nyys
	}
	migoS[migop] = migoVAL
	migoS[migop].yys = migostate

migonewstate:
	migon = migoPact[migostate]
	if migon <= migoFlag {
		goto migodefault /* simple state */
	}
	if migorcvr.char < 0 {
		migorcvr.char, migotoken = migolex1(migolex, &migorcvr.lval)
	}
	migon += migotoken
	if migon < 0 || migon >= migoLast {
		goto migodefault
	}
	migon = migoAct[migon]
	if migoChk[migon] == migotoken { /* valid shift */
		migorcvr.char = -1
		migotoken = -1
		migoVAL = migorcvr.lval
		migostate = migon
		if Errflag > 0 {
			Errflag--
		}
		goto migostack
	}

migodefault:
	/* default state action */
	migon = migoDef[migostate]
	if migon == -2 {
		if migorcvr.char < 0 {
			migorcvr.char, migotoken = migolex1(migolex, &migorcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if migoExca[xi+0] == -1 && migoExca[xi+1] == migostate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			migon = migoExca[xi+0]
			if migon < 0 || migon == migotoken {
				break
			}
		}
		migon = migoExca[xi+1]
		if migon < 0 {
			goto ret0
		}
	}
	if migon == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			migolex.Error(migoErrorMessage(migostate, migotoken))
			Nerrs++
			if migoDebug >= 1 {
				__yyfmt__.Printf("%s", migoStatname(migostate))
				__yyfmt__.Printf(" saw %s\n", migoTokname(migotoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for migop >= 0 {
				migon = migoPact[migoS[migop].yys] + migoErrCode
				if migon >= 0 && migon < migoLast {
					migostate = migoAct[migon] /* simulate a shift of "error" */
					if migoChk[migostate] == migoErrCode {
						goto migostack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if migoDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", migoS[migop].yys)
				}
				migop--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if migoDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", migoTokname(migotoken))
			}
			if migotoken == migoEofCode {
				goto ret1
			}
			migorcvr.char = -1
			migotoken = -1
			goto migonewstate /* try again in the same state */
		}
	}

	/* reduction by production migon */
	if migoDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", migon, migoStatname(migostate))
	}

	migont := migon
	migopt := migop
	_ = migopt // guard against "declared and not used"

	migop -= migoR2[migon]
	// migop is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if migop+1 >= len(migoS) {
		nyys := make([]migoSymType, len(migoS)*2)
		copy(nyys, migoS)
		migoS = nyys
	}
	migoVAL = migoS[migop+1]

	/* consult goto table to find next state */
	migon = migoR1[migon]
	migog := migoPgo[migon]
	migoj := migog + migoS[migop].yys + 1

	if migoj >= migoLast {
		migostate = migoAct[migog]
	} else {
		migostate = migoAct[migoj]
		if migoChk[migostate] != -migon {
			migostate = migoAct[migog]
		}
	}
	// dummy call; replaced with literal code
	switch migont {

	case 1:
		migoDollar = migoS[migopt-1 : migopt+1]
		//line migo.y:37
		{
			prog = migo.NewProgram()
			migoVAL.prog = prog
			prog.AddFunction(migoDollar[1].fun)
		}
	case 2:
		migoDollar = migoS[migopt-2 : migopt+1]
		//line migo.y:38
		{
			migoDollar[1].prog.AddFunction(migoDollar[2].fun)
		}
	case 3:
		migoDollar = migoS[migopt-7 : migopt+1]
		//line migo.y:41
		{
			migoVAL.fun = migo.NewFunction(migoDollar[2].str)
			migoVAL.fun.AddParams(migoDollar[4].params...)
			migoVAL.fun.AddStmts(migoDollar[7].stmts...)
		}
	case 4:
		migoDollar = migoS[migopt-0 : migopt+1]
		//line migo.y:44
		{
			migoVAL.params = params()
		}
	case 5:
		migoDollar = migoS[migopt-1 : migopt+1]
		//line migo.y:45
		{
			migoVAL.params = params(plainParam(migoDollar[1].str))
		}
	case 6:
		migoDollar = migoS[migopt-3 : migopt+1]
		//line migo.y:46
		{
			migoVAL.params = append(migoDollar[1].params, plainParam(migoDollar[3].str))
		}
	case 7:
		migoDollar = migoS[migopt-1 : migopt+1]
		//line migo.y:49
		{
			migoVAL.stmts = stmts(migoDollar[1].stmt)
		}
	case 8:
		migoDollar = migoS[migopt-2 : migopt+1]
		//line migo.y:50
		{
			migoVAL.stmts = append(migoDollar[1].stmts, migoDollar[2].stmt)
		}
	case 9:
		migoDollar = migoS[migopt-0 : migopt+1]
		//line migo.y:53
		{
			migoVAL.stmts = stmts()
		}
	case 10:
		migoDollar = migoS[migopt-2 : migopt+1]
		//line migo.y:54
		{
			migoVAL.stmts = append(migoDollar[1].stmts, migoDollar[2].stmt)
		}
	case 11:
		migoDollar = migoS[migopt-2 : migopt+1]
		//line migo.y:57
		{
			migoVAL.stmt = sendStmt(migoDollar[2].str)
		}
	case 12:
		migoDollar = migoS[migopt-2 : migopt+1]
		//line migo.y:58
		{
			migoVAL.stmt = recvStmt(migoDollar[2].str)
		}
	case 13:
		migoDollar = migoS[migopt-1 : migopt+1]
		//line migo.y:59
		{
			migoVAL.stmt = tauStmt()
		}
	case 14:
		migoDollar = migoS[migopt-8 : migopt+1]
		//line migo.y:62
		{
			migoVAL.stmt = newchanStmt(migoDollar[2].str, migoDollar[5].str, migoDollar[7].num)
		}
	case 15:
		migoDollar = migoS[migopt-2 : migopt+1]
		//line migo.y:63
		{
			migoVAL.stmt = migoDollar[1].stmt
		}
	case 16:
		migoDollar = migoS[migopt-3 : migopt+1]
		//line migo.y:64
		{
			migoVAL.stmt = closeStmt(migoDollar[2].str)
		}
	case 17:
		migoDollar = migoS[migopt-6 : migopt+1]
		//line migo.y:65
		{
			migoVAL.stmt = callStmt(migoDollar[2].str, migoDollar[4].params)
		}
	case 18:
		migoDollar = migoS[migopt-6 : migopt+1]
		//line migo.y:66
		{
			migoVAL.stmt = spawnStmt(migoDollar[2].str, migoDollar[4].params)
		}
	case 19:
		migoDollar = migoS[migopt-6 : migopt+1]
		//line migo.y:67
		{
			migoVAL.stmt = ifStmt(migoDollar[2].stmts, migoDollar[4].stmts)
		}
	case 20:
		migoDollar = migoS[migopt-4 : migopt+1]
		//line migo.y:68
		{
			migoVAL.stmt = selectStmt(migoDollar[2].cases)
		}
	case 21:
		migoDollar = migoS[migopt-0 : migopt+1]
		//line migo.y:71
		{
			migoVAL.cases = cases()
		}
	case 22:
		migoDollar = migoS[migopt-5 : migopt+1]
		//line migo.y:72
		{
			migoVAL.cases = append(migoDollar[1].cases, append(stmts(migoDollar[3].stmt), migoDollar[5].stmts...))
		}
	}
	goto migostack /* stack new state and value */
}
