//line parse.y:2
package main

import __yyfmt__ "fmt"

//line parse.y:2
//line parse.y:6
type yySymType struct {
	yys       int
	numval    float64
	ident     string
	lval      Lvalue
	expr      Expression
	stmt      Statement
	argList   []Expression
	stmtBlock []Statement
}

const IF = 57346
const THEN = 57347
const ELSE = 57348
const END = 57349
const WHILE = 57350
const DO = 57351
const PRINT = 57352
const AND = 57353
const OR = 57354
const NOT = 57355
const NUMBER = 57356
const IDENTIFIER = 57357

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IF",
	"THEN",
	"ELSE",
	"END",
	"WHILE",
	"DO",
	"PRINT",
	"AND",
	"OR",
	"NOT",
	"NUMBER",
	"IDENTIFIER",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'='",
	"','",
	"'<'",
	"'>'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parse.y:100

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 0,
	-1, 15,
	11, 23,
	12, 23,
	-2, 16,
}

const yyNprod = 33
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 112

var yyAct = [...]int{

	20, 2, 16, 11, 35, 36, 37, 38, 32, 55,
	33, 34, 43, 56, 23, 18, 35, 36, 37, 38,
	30, 41, 17, 21, 13, 56, 15, 64, 25, 45,
	27, 46, 47, 19, 66, 40, 63, 21, 13, 42,
	37, 38, 44, 57, 31, 24, 11, 26, 48, 49,
	50, 51, 52, 53, 54, 31, 31, 7, 11, 58,
	6, 62, 5, 11, 35, 36, 37, 38, 32, 14,
	33, 34, 21, 13, 4, 35, 36, 37, 38, 8,
	22, 9, 19, 59, 60, 10, 3, 12, 8, 39,
	9, 1, 13, 65, 10, 0, 12, 8, 8, 9,
	9, 13, 61, 10, 10, 12, 12, 0, 28, 29,
	13, 13,
}
var yyPact = [...]int{

	-1000, -1000, 96, -1000, -1000, -1000, -1000, -1000, -1000, 9,
	9, -6, 23, -1000, 25, -1000, 97, 58, 48, 9,
	-1000, -1000, 12, 23, -9, 59, 23, -1000, 58, 58,
	-1000, -1000, 23, 23, 23, 23, 23, 23, 23, -16,
	-12, -1000, 59, 23, 0, 77, -1000, -1000, 59, 59,
	59, 22, 22, -1000, -1000, -1000, -1000, 95, 59, -1000,
	32, 19, 86, -1000, -1000, 30, -1000,
}
var yyPgo = [...]int{

	0, 91, 0, 15, 69, 26, 2, 86, 74, 62,
	60, 57, 45, 1,
}
var yyR1 = [...]int{

	0, 1, 13, 13, 7, 7, 7, 7, 7, 8,
	8, 9, 10, 11, 12, 12, 4, 4, 4, 4,
	5, 5, 5, 6, 6, 2, 3, 3, 3, 3,
	3, 3, 3,
}
var yyR2 = [...]int{

	0, 1, 0, 2, 1, 1, 1, 1, 1, 8,
	6, 6, 3, 2, 1, 3, 1, 3, 3, 2,
	3, 3, 3, 1, 3, 1, 1, 1, 3, 3,
	3, 3, 3,
}
var yyChk = [...]int{

	-1000, -1, -13, -7, -8, -9, -10, -11, 2, 4,
	8, -2, 10, 15, -4, -5, -6, 13, -3, 24,
	-2, 14, -4, 20, -12, -3, 24, 5, 11, 12,
	-6, -5, 20, 22, 23, 16, 17, 18, 19, -4,
	-3, 9, -3, 21, -3, -13, -6, -6, -3, -3,
	-3, -3, -3, -3, -3, 25, 25, -13, -3, 6,
	7, 7, -13, 4, 8, 7, 4,
}
var yyDef = [...]int{

	2, -2, -2, 3, 4, 5, 6, 7, 8, 0,
	0, 0, 0, 25, 0, -2, 0, 0, 0, 0,
	26, 27, 0, 0, 13, 14, 0, 2, 0, 0,
	19, 23, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 2, 12, 0, 0, 0, 17, 18, 20, 21,
	22, 29, 30, 31, 32, 24, 28, 0, 15, 2,
	0, 0, 0, 10, 11, 0, 9,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	24, 25, 18, 16, 21, 17, 3, 19, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	22, 20, 23,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:34
		{
			yylex.(*Lexer).program = yyDollar[1].stmtBlock
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:38
		{
			yyVAL.stmtBlock = []Statement{}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:39
		{
			yyVAL.stmtBlock = append(yyDollar[1].stmtBlock, yyDollar[2].stmt)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:43
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:44
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:45
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:46
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:47
		{
			yyVAL.stmt = nil
		}
	case 9:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parse.y:50
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[4].stmtBlock, yyDollar[6].stmtBlock}
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y:51
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[4].stmtBlock, nil}
		}
	case 11:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y:55
		{
			yyVAL.stmt = &WhileStmt{yyDollar[2].expr, yyDollar[4].stmtBlock}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:59
		{
			yyVAL.stmt = &AssignStmt{yyDollar[1].lval, yyDollar[3].expr}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:63
		{
			yyVAL.stmt = &PrintStmt{yyDollar[2].argList}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:67
		{
			yyVAL.argList = []Expression{yyDollar[1].expr}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:68
		{
			yyVAL.argList = append(yyDollar[1].argList, yyDollar[3].expr)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:72
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:73
		{
			yyVAL.expr = &LogicExpr{AND, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:74
		{
			yyVAL.expr = &LogicExpr{OR, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:75
		{
			yyVAL.expr = &LogicExpr{NOT, yyDollar[2].expr, nil}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:78
		{
			yyVAL.expr = &RelExpr{'=', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:79
		{
			yyVAL.expr = &RelExpr{'<', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:80
		{
			yyVAL.expr = &RelExpr{'>', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:84
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:85
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:88
		{
			yyVAL.lval = Identifier(yyDollar[1].ident)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:91
		{
			yyVAL.expr = yyDollar[1].lval
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:92
		{
			yyVAL.expr = Number(yyDollar[1].numval)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:93
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:94
		{
			yyVAL.expr = &ArithExpr{'+', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:95
		{
			yyVAL.expr = &ArithExpr{'-', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:96
		{
			yyVAL.expr = &ArithExpr{'*', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:97
		{
			yyVAL.expr = &ArithExpr{'/', yyDollar[1].expr, yyDollar[3].expr}
		}
	}
	goto yystack /* stack new state and value */
}
