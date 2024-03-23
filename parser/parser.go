// Code generated by goyacc -o parser.go grammar.y. DO NOT EDIT.

//line grammar.y:1

package parser

import __yyfmt__ "fmt"

//line grammar.y:3

//line grammar.y:28
type yySymType struct {
	yys   int
	value string // a string in go that produce the value of the object
	gtype string // a string representing the gotype of the object
}

const MULTI = 57346
const DIV = 57347
const MOD = 57348
const PLUS = 57349
const MINUS = 57350
const PLUSPLUS = 57351
const LTE = 57352
const GTE = 57353
const LT = 57354
const GT = 57355
const EQ = 57356
const NEQ = 57357
const COLON = 57358
const SEMICOLON = 57359
const DOT = 57360
const COMMA = 57361
const LBRACKET = 57362
const RBRACKET = 57363
const LPAREN = 57364
const RPAREN = 57365
const LBRACE = 57366
const RBRACE = 57367
const AND = 57368
const OR = 57369
const NOT = 57370
const DOTDOT = 57371
const ASSIGN = 57372
const QUESTION = 57373
const REGEXMATCH = 57374
const REGEXNOTMATCH = 57375
const FOR = 57376
const RETURN = 57377
const WAITFOR = 57378
const OPTIONS = 57379
const IGNORE = 57380
const HEADLESS = 57381
const TIMEOUT = 57382
const DISTINCT = 57383
const FILTER = 57384
const CURRENT = 57385
const SORT = 57386
const LIMIT = 57387
const LET = 57388
const COLLECT = 57389
const ASC = 57390
const DESC = 57391
const NONE = 57392
const NULL = 57393
const TRUE = 57394
const FALSE = 57395
const USE = 57396
const INTO = 57397
const KEEP = 57398
const WITH = 57399
const COUNT = 57400
const ALL = 57401
const ANY = 57402
const AGGREGATE = 57403
const EVENT = 57404
const LIKE = 57405
const IN = 57406
const WHILE = 57407
const BOOL = 57408
const AT = 57409
const IDENTIFIER = 57410
const IGNOREID = 57411
const STRING = 57412
const NUMBER = 57413
const NAMESPACESEPARATOR = 57414
const SELECT = 57415
const CLICK = 57416
const DOCUMENT = 57417
const PAGE = 57418
const CONTAINS = 57419

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"MULTI",
	"DIV",
	"MOD",
	"PLUS",
	"MINUS",
	"PLUSPLUS",
	"LTE",
	"GTE",
	"LT",
	"GT",
	"EQ",
	"NEQ",
	"COLON",
	"SEMICOLON",
	"DOT",
	"COMMA",
	"LBRACKET",
	"RBRACKET",
	"LPAREN",
	"RPAREN",
	"LBRACE",
	"RBRACE",
	"AND",
	"OR",
	"NOT",
	"DOTDOT",
	"ASSIGN",
	"QUESTION",
	"REGEXMATCH",
	"REGEXNOTMATCH",
	"FOR",
	"RETURN",
	"WAITFOR",
	"OPTIONS",
	"IGNORE",
	"HEADLESS",
	"TIMEOUT",
	"DISTINCT",
	"FILTER",
	"CURRENT",
	"SORT",
	"LIMIT",
	"LET",
	"COLLECT",
	"ASC",
	"DESC",
	"NONE",
	"NULL",
	"TRUE",
	"FALSE",
	"USE",
	"INTO",
	"KEEP",
	"WITH",
	"COUNT",
	"ALL",
	"ANY",
	"AGGREGATE",
	"EVENT",
	"LIKE",
	"IN",
	"WHILE",
	"BOOL",
	"AT",
	"IDENTIFIER",
	"IGNOREID",
	"STRING",
	"NUMBER",
	"NAMESPACESEPARATOR",
	"SELECT",
	"CLICK",
	"DOCUMENT",
	"PAGE",
	"CONTAINS",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:173

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 34,
	4, 36,
	5, 36,
	6, 36,
	8, 36,
	26, 57,
	27, 57,
	-2, 26,
}

const yyPrivate = 57344

const yyLast = 182

var yyAct = [...]int8{
	3, 24, 25, 16, 11, 10, 80, 41, 7, 33,
	21, 30, 36, 45, 100, 39, 42, 43, 37, 22,
	46, 77, 38, 46, 44, 72, 31, 46, 23, 27,
	68, 69, 40, 71, 75, 76, 55, 8, 12, 104,
	5, 2, 79, 14, 15, 1, 13, 11, 10, 82,
	83, 84, 85, 86, 87, 29, 88, 89, 90, 91,
	92, 93, 94, 95, 96, 97, 98, 78, 11, 10,
	30, 73, 26, 20, 35, 107, 81, 9, 41, 108,
	33, 12, 6, 19, 28, 0, 14, 15, 18, 13,
	32, 99, 101, 102, 103, 53, 54, 55, 56, 57,
	4, 70, 12, 17, 0, 74, 0, 14, 15, 109,
	13, 0, 53, 54, 55, 56, 57, 0, 61, 63,
	60, 62, 58, 59, 0, 64, 65, 0, 36, 0,
	34, 105, 33, 35, 106, 0, 46, 66, 67, 50,
	52, 49, 51, 47, 48, 64, 65, 53, 54, 55,
	56, 57, 104, 53, 54, 55, 0, 66, 67, 0,
	0, 53, 54, 55, 56, 57, 105, 61, 63, 60,
	62, 58, 59, 46, 0, 0, 50, 52, 49, 51,
	47, 48,
}

var yyPact = [...]int16{
	-30, -1000, -30, -1000, -1000, 34, -1000, -60, -49, -1000,
	62, -50, -8, 10, 10, 10, -1000, -1000, -1000, -1000,
	5, -1000, -55, -1000, 166, 157, 131, -1000, 62, -1000,
	3, -1000, 62, -1000, -1000, -1000, -1000, -43, 62, 20,
	10, -1000, 20, 20, -64, -1000, -61, 10, 10, 10,
	10, 10, 10, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, -54, -54, -54, -54, 129, 108,
	111, 149, 3, -1000, -1000, 166, 157, 10, -1000, 16,
	-1000, -1000, 20, 20, 20, 20, 20, 20, 30, 30,
	-1000, 149, 149, 91, 91, 91, 91, 91, 91, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 143, 13, -1000,
}

var yyPgo = [...]int8{
	0, 28, 1, 73, 29, 55, 2, 26, 72, 45,
	41, 0, 100, 40, 82, 77,
}

var yyR1 = [...]int8{
	0, 9, 9, 10, 10, 12, 12, 3, 3, 11,
	11, 13, 13, 15, 15, 15, 15, 14, 14, 1,
	1, 1, 2, 2, 2, 4, 4, 6, 6, 6,
	6, 6, 6, 6, 6, 5, 5, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 7, 7,
}

var yyR2 = [...]int8{
	0, 2, 1, 2, 1, 2, 3, 1, 3, 2,
	1, 1, 2, 3, 2, 2, 2, 2, 5, 1,
	1, 1, 1, 3, 3, 1, 1, 1, 3, 3,
	3, 3, 3, 2, 3, 1, 1, 1, 3, 3,
	3, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 1,
}

var yyChk = [...]int16{
	-1000, -9, -10, -11, -12, -13, -14, 38, 67, -15,
	35, 34, 68, 76, 73, 74, -11, -12, -14, -15,
	-3, 70, 68, -1, -2, -6, -8, -4, 22, -5,
	8, -7, 28, 70, 68, 71, 66, 68, 30, -2,
	22, 68, -2, -2, 19, 68, 7, 14, 15, 12,
	10, 13, 11, 4, 5, 6, 7, 8, 14, 15,
	12, 10, 13, 11, 14, 15, 26, 27, -2, -6,
	-8, -6, 22, 68, -8, -2, -6, 64, -1, -2,
	70, -4, -2, -2, -2, -2, -2, -2, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -7,
	68, -7, -7, -7, 23, 23, 23, -6, -2, -11,
}

var yyDef = [...]int8{
	0, -2, 0, 2, 4, 0, 10, 0, 0, 11,
	0, 0, 0, 0, 0, 0, 1, 3, 9, 12,
	5, 7, 0, 17, 19, 20, 21, 22, 0, 27,
	0, 37, 0, 25, -2, 35, 56, 0, 0, 14,
	0, 26, 15, 16, 0, 6, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 33, 0, 36, 42, 0, 0, 0, 13, 0,
	8, 23, 43, 44, 45, 46, 47, 48, 28, 29,
	30, 31, 32, 49, 50, 51, 52, 53, 54, 38,
	57, 39, 40, 41, 24, 34, 55, 0, 0, 18,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77,
}

var yyTok3 = [...]int8{
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
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
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
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
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
	yyn = int(yyPact[yystate])
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
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
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
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
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
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
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

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:60
		{
			yylex.(*myLexer).finalize()
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:61
		{
			yylex.(*myLexer).finalize()
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:72
		{
			yylex.(*myLexer).addLines("rt.Ignore(" + yyDollar[2].value + ")")
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:75
		{ // @ paramName paramType
			yylex.(*myLexer).setParam(yyDollar[2].value, yyDollar[3].value)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:82
		{
			yyVAL.value = yyDollar[1].value + "," + yyDollar[3].value
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:88
		{ /* todo */
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:89
		{ /* todo */
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:93
		{ /* todo */
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:94
		{ /* todo */
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:98
		{
			yylex.(*myLexer).setVar(yyDollar[1].value, yyDollar[3].value)
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:99
		{ /* todo */
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:100
		{ /* todo */
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:101
		{ /* todo */
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:105
		{ /* todo */
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:106
		{ /* todo */
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:120
		{
			yyVAL.value = "(" + yyDollar[1].value + ")+(" + yyDollar[3].value + ")"
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:121
		{
			yyVAL.value = "(" + yyDollar[2].value + ")"
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:130
		{
			yyVAL.value = yyDollar[1].value
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:132
		{
			yyVAL.value = "(" + yyDollar[1].value + ")*(" + yyDollar[3].value + ")"
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:133
		{
			yyVAL.value = "(" + yyDollar[1].value + ")/(" + yyDollar[3].value + ")"
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:134
		{
			yyVAL.value = "(" + yyDollar[1].value + ")%(" + yyDollar[3].value + ")"
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:135
		{
			yyVAL.value = "(" + yyDollar[1].value + ")+(" + yyDollar[3].value + ")"
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:136
		{
			yyVAL.value = "(" + yyDollar[1].value + ")-(" + yyDollar[3].value + ")"
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:137
		{
			yyVAL.value = "-(" + yyDollar[2].value + ")"
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:138
		{
			yyVAL.value = "(" + yyDollar[2].value + ")"
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:148
		{
			yyVAL.value = "(" + yyDollar[1].value + "==" + yyDollar[3].value + ")"
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:149
		{
			yyVAL.value = "(" + yyDollar[1].value + "!=" + yyDollar[3].value + ")"
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:150
		{
			yyVAL.value = "(" + yyDollar[1].value + "&&" + yyDollar[3].value + ")"
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:151
		{
			yyVAL.value = "(" + yyDollar[1].value + "||" + yyDollar[3].value + ")"
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:152
		{
			yyVAL.value = "!(" + yyDollar[2].value + ")"
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:153
		{
			yyVAL.value = "(" + yyDollar[1].value + "==" + yyDollar[3].value + ")"
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:154
		{
			yyVAL.value = "(" + yyDollar[1].value + "!=" + yyDollar[3].value + ")"
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:155
		{
			yyVAL.value = "(" + yyDollar[1].value + "<" + yyDollar[3].value + ")"
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:156
		{
			yyVAL.value = "(" + yyDollar[1].value + "<=" + yyDollar[3].value + ")"
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:157
		{
			yyVAL.value = "(" + yyDollar[1].value + ">" + yyDollar[3].value + ")"
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:158
		{
			yyVAL.value = "(" + yyDollar[1].value + ">=" + yyDollar[3].value + ")"
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:159
		{
			yyVAL.value = "(" + yyDollar[1].value + "==" + yyDollar[3].value + ")"
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:160
		{
			yyVAL.value = "(" + yyDollar[1].value + "!=" + yyDollar[3].value + ")"
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:161
		{
			yyVAL.value = "(" + yyDollar[1].value + "<" + yyDollar[3].value + ")"
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:162
		{
			yyVAL.value = "(" + yyDollar[1].value + "<=" + yyDollar[3].value + ")"
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:163
		{
			yyVAL.value = "(" + yyDollar[1].value + ">" + yyDollar[3].value + ")"
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:164
		{
			yyVAL.value = "(" + yyDollar[1].value + ">=" + yyDollar[3].value + ")"
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:165
		{
			yyVAL.value = "(" + yyDollar[2].value + ")"
		}
	}
	goto yystack /* stack new state and value */
}
