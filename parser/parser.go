// Code generated by goyacc -o parser.go grammar.y. DO NOT EDIT.

//line grammar.y:1

package parser

import __yyfmt__ "fmt"

//line grammar.y:3

// each object has a value and a type.
type value struct {
	v string // a string in go that produce the value of the object
	t string // a string representing the gotype of the object
	c int    // the code returned by lexer is stored here. Always set by the lexer, even for variables (set as IDENTIFIER). A valid go type, without spaces.
}

//line grammar.y:35
type yySymType struct {
	yys    int
	value  value
	list   []string
	values []value
}

const MULTI = 57346
const DIV = 57347
const MOD = 57348
const PLUS = 57349
const MINUS = 57350
const PLUSPLUS = 57351
const MINUSMINUS = 57352
const LTE = 57353
const GTE = 57354
const LT = 57355
const GT = 57356
const EQ = 57357
const NEQ = 57358
const COLON = 57359
const SEMICOLON = 57360
const DOT = 57361
const COMMA = 57362
const LBRACKET = 57363
const RBRACKET = 57364
const LPAREN = 57365
const RPAREN = 57366
const LBRACE = 57367
const RBRACE = 57368
const AND = 57369
const OR = 57370
const NOT = 57371
const DOTDOT = 57372
const ASSIGN = 57373
const QUESTION = 57374
const REGEXMATCH = 57375
const REGEXNOTMATCH = 57376
const LOWER = 57377
const UPPER = 57378
const FORMAT = 57379
const INTTYPE = 57380
const BOOLTYPE = 57381
const STRINGTYPE = 57382
const BINTYPE = 57383
const FOR = 57384
const RETURN = 57385
const WAITFOR = 57386
const OPTIONS = 57387
const IGNORE = 57388
const HEADLESS = 57389
const TIMEOUT = 57390
const DISTINCT = 57391
const FILTER = 57392
const CURRENT = 57393
const SORT = 57394
const LIMIT = 57395
const LET = 57396
const COLLECT = 57397
const ASC = 57398
const DESC = 57399
const NIL = 57400
const TRUE = 57401
const FALSE = 57402
const USE = 57403
const INTO = 57404
const KEEP = 57405
const WITH = 57406
const COUNT = 57407
const ALL = 57408
const ANY = 57409
const AGGREGATE = 57410
const EVENT = 57411
const LIKE = 57412
const IN = 57413
const WHILE = 57414
const BOOL = 57415
const AT = 57416
const IDENTIFIER = 57417
const IGNOREID = 57418
const STRING = 57419
const NUMBER = 57420
const NAMESPACESEPARATOR = 57421
const SELECT = 57422
const CLICK = 57423
const DOCUMENT = 57424
const PAGE = 57425
const CONTAINS = 57426

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
	"MINUSMINUS",
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
	"LOWER",
	"UPPER",
	"FORMAT",
	"INTTYPE",
	"BOOLTYPE",
	"STRINGTYPE",
	"BINTYPE",
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
	"NIL",
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

//line grammar.y:192

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 174

var yyAct = [...]int8{
	58, 59, 60, 55, 57, 56, 74, 64, 62, 63,
	61, 65, 66, 79, 32, 44, 31, 24, 89, 19,
	5, 77, 46, 67, 68, 69, 58, 59, 60, 55,
	57, 56, 26, 64, 62, 63, 61, 65, 66, 28,
	84, 29, 76, 84, 91, 83, 82, 88, 71, 67,
	68, 69, 58, 59, 60, 55, 57, 56, 52, 64,
	62, 63, 61, 65, 66, 75, 36, 13, 33, 11,
	34, 35, 72, 9, 78, 67, 68, 69, 2, 22,
	70, 27, 43, 85, 45, 1, 37, 38, 39, 23,
	10, 51, 53, 30, 80, 17, 12, 73, 54, 32,
	21, 31, 86, 37, 38, 39, 70, 42, 47, 49,
	48, 50, 8, 40, 41, 0, 32, 4, 31, 20,
	7, 0, 3, 0, 42, 6, 25, 0, 14, 0,
	40, 41, 70, 18, 16, 0, 15, 81, 0, 0,
	0, 0, 0, 87, 0, 0, 0, 0, 0, 0,
	90, 36, 0, 33, 0, 34, 35, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 36, 0,
	33, 0, 34, 35,
}

var yyPact = [...]int16{
	-54, -1000, -54, 53, -1000, -56, 53, -1000, -1000, 53,
	-1000, -1000, -58, 53, 1, 95, 95, -60, 95, 70,
	-1000, -1000, -1000, 38, -1000, -1000, 95, 48, -1000, 27,
	-7, 95, 95, 19, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 48, -50, 48, -1000, -1000, -1000, -1000,
	-1000, 70, -62, 48, 95, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 95, 27, 22, 23, 48, 78, 95, 25, -1000,
	-1000, -4, -1000, -1000, 95, 20, -1000, 48, -1000, -1000,
	48, -1000,
}

var yyPgo = [...]int8{
	0, 65, 39, 41, 98, 93, 22, 89, 6, 85,
	78, 122, 112, 117, 73, 90, 69, 67,
}

var yyR1 = [...]int8{
	0, 9, 9, 11, 10, 10, 13, 6, 6, 6,
	6, 6, 12, 12, 14, 14, 16, 16, 16, 15,
	15, 7, 7, 17, 17, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 5, 5, 5, 5, 5, 5, 1, 1, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	8, 8,
}

var yyR2 = [...]int8{
	0, 3, 2, 0, 2, 1, 3, 1, 1, 1,
	1, 3, 2, 1, 1, 2, 3, 2, 2, 2,
	2, 1, 3, 4, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	2, 3, 4, 3, 4, 3, 1, 1, 1, 1,
	1, 3,
}

var yyChk = [...]int16{
	-1000, -9, -10, -11, -13, 74, -11, -13, -12, -14,
	-15, -16, 43, -17, 75, 83, 81, 42, 80, 75,
	-12, -15, -16, -7, 75, -12, 31, -1, -2, -3,
	-5, 23, 21, 75, 77, 78, 73, 8, 9, 10,
	35, 36, 29, -1, 75, -1, -6, 38, 40, 39,
	41, 21, 20, -1, -4, 7, 9, 8, 4, 5,
	6, 14, 12, 13, 11, 15, 16, 27, 28, 29,
	84, 21, -3, -1, -8, -1, 23, 71, -6, 75,
	-2, -1, 24, 22, 20, -8, 24, -1, 22, 22,
	-1, 24,
}

var yyDef = [...]int8{
	3, -2, 3, 0, 5, 0, 0, 4, 2, 0,
	13, 14, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 12, 15, 19, 21, 20, 0, 17, 47, 49,
	0, 0, 0, 56, 57, 58, 59, 41, 42, 43,
	44, 45, 46, 18, 0, 24, 6, 7, 8, 9,
	10, 0, 0, 16, 0, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 0, 50, 0, 0, 60, 0, 0, 0, 22,
	48, 0, 51, 53, 0, 0, 55, 23, 11, 52,
	61, 54,
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
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84,
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
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:73
		{
			yylex.(*myLexer).finalize()
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:74
		{
			yylex.(*myLexer).finalize()
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:78
		{
			yylex.(*myLexer).incOut()
			yylex.(*myLexer).addLines("{")
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:88
		{
			yylex.(*myLexer).declInputParam(yyDollar[2].value.v, yyDollar[3].value.v)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:95
		{
			yyVAL.value.v = "[]byte"
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:96
		{
			yyVAL.value.v = "[]" + yyDollar[2].value.v
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:103
		{
			yylex.(*myLexer).addLines("}")
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:104
		{
			yylex.(*myLexer).addLines("}")
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:108
		{ /* todo */
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:109
		{ /* todo */
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:113
		{
			yylex.(*myLexer).vSetVar(yyDollar[1].value.v, yyDollar[3].value)
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:114
		{ /* todo */
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:115
		{ /* todo */
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:119
		{
			yylex.(*myLexer).declOutputParams(yyDollar[2].list)
			yylex.(*myLexer).saveOut()
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:120
		{ /* */
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:124
		{
			yyVAL.list = append(yyVAL.list, yyDollar[1].value.v)
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:125
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].value.v)
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:129
		{
			yylex.(*myLexer).forNameInExpression(yyDollar[2].value.v, yyDollar[4].value)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:130
		{
			yylex.(*myLexer).selectExpression(yyDollar[2].value)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:137
		{
			yyVAL.value = yyDollar[1].value
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:138
		{
			yyVAL.value = yyDollar[1].value
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:139
		{
			yyVAL.value = yyDollar[1].value
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:140
		{
			yyVAL.value = yyDollar[1].value
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:141
		{
			yyVAL.value = yyDollar[1].value
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:142
		{
			yyVAL.value = yyDollar[1].value
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:143
		{
			yyVAL.value = yyDollar[1].value
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:144
		{
			yyVAL.value = yyDollar[1].value
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:145
		{
			yyVAL.value = yyDollar[1].value
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:146
		{
			yyVAL.value = yyDollar[1].value
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:147
		{
			yyVAL.value = yyDollar[1].value
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:148
		{
			yyVAL.value = yyDollar[1].value
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:149
		{
			yyVAL.value = yyDollar[1].value
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:150
		{
			yyVAL.value = yyDollar[1].value
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:151
		{
			yyVAL.value = yyDollar[1].value
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:152
		{
			yyVAL.value = yyDollar[1].value
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:156
		{
			yyVAL.value = yyDollar[1].value
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:157
		{
			yyVAL.value = yyDollar[1].value
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:158
		{
			yyVAL.value = yyDollar[1].value
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:159
		{
			yyVAL.value = yyDollar[1].value
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:160
		{
			yyVAL.value = yyDollar[1].value
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:161
		{
			yyVAL.value = yyDollar[1].value
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:165
		{
			yyVAL.value = yyDollar[1].value
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:166
		{
			yyVAL.value = yylex.(*myLexer).vOpe2(yyDollar[2].value.c, yyDollar[1].value, yyDollar[3].value)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:170
		{
			yyVAL.value = yyDollar[1].value
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:171
		{
			yyVAL.value = yylex.(*myLexer).vOpe1(yyDollar[1].value.c, yyDollar[2].value)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:175
		{
			yyVAL.value = yylex.(*myLexer).vParen(yyDollar[2].value)
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:176
		{
			yyVAL.value = yylex.(*myLexer).vGetElementOf(yyDollar[1].value, yyDollar[3].value)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:177
		{
			yyVAL.value = yylex.(*myLexer).vMakeArray(yyDollar[2].values)
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:178
		{ /* TODO - function call computing and returning a value */
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:179
		{ /* TODO - function call computing and returning a value - empty input params */
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:180
		{
			yyVAL.value = yylex.(*myLexer).vGetVar(yyDollar[1].value.v)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:181
		{
			yyVAL.value = yyDollar[1].value
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:182
		{
			yyVAL.value = yyDollar[1].value
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:183
		{
			yyVAL.value = yyDollar[1].value
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:187
		{
			yyVAL.values = []value{yyDollar[1].value}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:188
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].value)
		}
	}
	goto yystack /* stack new state and value */
}
