// Code generated by goyacc -o parser.go grammar.y. DO NOT EDIT.

//line grammar.y:1

package parser

import __yyfmt__ "fmt"

//line grammar.y:3

//line grammar.y:28
type yySymType struct {
	yys   int // ceci redéclare yySymType !
	value string
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
const XOR = 57371
const DOTDOT = 57372
const ASSIGN = 57373
const QUESTION = 57374
const REGEXMATCH = 57375
const REGEXNOTMATCH = 57376
const FOR = 57377
const RETURN = 57378
const WAITFOR = 57379
const OPTIONS = 57380
const IGNORE = 57381
const HEADLESS = 57382
const TIMEOUT = 57383
const DISTINCT = 57384
const FILTER = 57385
const CURRENT = 57386
const SORT = 57387
const LIMIT = 57388
const LET = 57389
const COLLECT = 57390
const ASC = 57391
const DESC = 57392
const NONE = 57393
const NULL = 57394
const TRUE = 57395
const FALSE = 57396
const USE = 57397
const INTO = 57398
const KEEP = 57399
const WITH = 57400
const COUNT = 57401
const ALL = 57402
const ANY = 57403
const AGGREGATE = 57404
const EVENT = 57405
const LIKE = 57406
const NOT = 57407
const IN = 57408
const WHILE = 57409
const BOOL = 57410
const AT = 57411
const IDENTIFIER = 57412
const IGNOREID = 57413
const STRING = 57414
const NUMBER = 57415
const NAMESPACESEPARATOR = 57416
const SELECT = 57417
const CLICK = 57418
const DOCUMENT = 57419
const PAGE = 57420
const CONTAINS = 57421

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
	"XOR",
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
	"NOT",
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

//line grammar.y:185

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 150

var yyAct = [...]int8{
	3, 47, 40, 16, 34, 51, 49, 50, 48, 52,
	53, 81, 11, 10, 21, 47, 7, 31, 88, 51,
	49, 50, 48, 52, 53, 44, 37, 24, 22, 36,
	76, 47, 72, 25, 28, 11, 10, 38, 32, 55,
	56, 39, 41, 42, 47, 90, 8, 12, 65, 66,
	67, 34, 14, 15, 31, 13, 89, 68, 26, 11,
	10, 74, 43, 69, 31, 71, 77, 75, 80, 29,
	12, 9, 78, 54, 83, 14, 15, 19, 13, 29,
	82, 47, 35, 57, 65, 66, 67, 54, 70, 84,
	85, 86, 73, 93, 12, 55, 56, 79, 88, 14,
	15, 6, 13, 87, 92, 4, 91, 18, 17, 46,
	64, 33, 45, 5, 36, 2, 27, 1, 34, 35,
	30, 33, 20, 23, 36, 0, 55, 56, 34, 35,
	61, 59, 60, 58, 62, 63, 0, 0, 0, 0,
	55, 56, 0, 89, 61, 59, 60, 58, 62, 63,
}

var yyPact = [...]int16{
	-23, -1000, -23, -1000, -1000, 0, -1000, -58, -42, -1000,
	46, -44, 6, -21, -21, -21, -1000, -1000, -1000, -1000,
	42, -1000, -45, -1000, 8, 133, 57, -1000, -1000, 56,
	-1000, 9, -1000, 56, -1000, -1000, -1000, -36, 56, 37,
	-21, 37, 37, -61, -1000, -68, -21, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 9, 9, 9, -1000, -1000,
	-1000, -1000, -1000, -1000, -39, -1000, -1000, -1000, -6, 119,
	21, -1000, 9, -1000, 8, 133, -21, 8, 133, 57,
	74, -1000, -1000, 37, -1000, -1000, 88, -1000, -1000, -1000,
	-1000, 32, 24, -1000,
}

var yyPgo = [...]int8{
	0, 123, 27, 122, 34, 120, 33, 38, 58, 117,
	115, 0, 105, 113, 101, 71, 112, 110, 109, 83,
}

var yyR1 = [...]int8{
	0, 9, 9, 10, 10, 12, 12, 3, 3, 11,
	11, 13, 13, 15, 15, 15, 15, 15, 15, 14,
	14, 1, 1, 1, 1, 2, 2, 2, 4, 6,
	6, 6, 6, 6, 5, 8, 8, 8, 8, 8,
	8, 7, 17, 17, 17, 16, 19, 19, 19, 19,
	19, 19, 18, 18, 18, 18, 18, 18, 18,
}

var yyR2 = [...]int8{
	0, 2, 1, 2, 1, 2, 3, 1, 3, 2,
	1, 1, 2, 3, 3, 3, 2, 2, 2, 2,
	5, 1, 1, 1, 1, 1, 3, 3, 1, 1,
	3, 3, 2, 3, 1, 1, 3, 2, 3, 3,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -9, -10, -11, -12, -13, -14, 39, 69, -15,
	36, 35, 70, 78, 75, 76, -11, -12, -14, -15,
	-3, 72, 70, -1, -2, -6, -8, 70, -4, 23,
	-5, 8, -7, 65, 72, 73, 68, 70, 31, -2,
	23, -2, -2, 20, 70, -16, -18, 7, 14, 12,
	13, 11, 15, 16, 79, 7, 8, -19, 14, 12,
	13, 11, 15, 16, -17, 27, 28, 29, -2, -6,
	-8, -6, 23, -8, -2, -6, 66, -2, -6, -8,
	-2, 72, -4, -2, -6, -6, -6, -7, 24, 24,
	24, -6, -2, -11,
}

var yyDef = [...]int8{
	0, -2, 0, 2, 4, 0, 10, 0, 0, 11,
	0, 0, 0, 0, 0, 0, 1, 3, 9, 12,
	5, 7, 0, 19, 21, 22, 23, 24, 25, 0,
	29, 0, 35, 0, 28, 34, 41, 0, 0, 16,
	0, 17, 18, 0, 6, 0, 0, 45, 52, 53,
	54, 55, 56, 57, 58, 0, 0, 0, 46, 47,
	48, 49, 50, 51, 0, 42, 43, 44, 0, 0,
	0, 32, 0, 37, 0, 0, 0, 13, 14, 15,
	0, 8, 26, 38, 30, 31, 39, 36, 27, 33,
	40, 0, 0, 20,
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
	72, 73, 74, 75, 76, 77, 78, 79,
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
//line grammar.y:59
		{
			yylex.(*myLexer).finalize()
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:60
		{
			yylex.(*myLexer).finalize()
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:71
		{
			yylex.(*myLexer).addLines("rt.Ignore(" + yyDollar[2].value + ")")
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:74
		{ // @ paramName paramType
			yylex.(*myLexer).setParam(yyDollar[2].value, yyDollar[3].value)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:81
		{
			yyVAL.value = yyDollar[1].value + "," + yyDollar[3].value
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:87
		{ /* todo */
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:88
		{ /* todo */
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:92
		{ /* todo */
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:93
		{ /* todo */
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:97
		{
			yylex.(*myLexer).setVar(yyDollar[1].value, yyDollar[3].value, "string")
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:98
		{
			yylex.(*myLexer).setVar(yyDollar[1].value, yyDollar[3].value, "int")
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:99
		{
			yylex.(*myLexer).setVar(yyDollar[1].value, yyDollar[3].value, "bool")
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:100
		{ /* todo */
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:101
		{ /* todo */
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:102
		{ /* todo */
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:106
		{ /* todo */
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:107
		{ /* todo */
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:117
		{ /* todo */
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:122
		{ /* todo */
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:123
		{
			yyVAL.value = yyDollar[2].value
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:131
		{
			yyVAL.value = yyDollar[1].value
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:133
		{
			yyVAL.value = "(" + yyDollar[1].value + ")+(" + yyDollar[3].value + ")"
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:134
		{
			yyVAL.value = "(" + yyDollar[1].value + ")-(" + yyDollar[3].value + ")"
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:135
		{
			yyVAL.value = "-(" + yyDollar[2].value + ")"
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:136
		{
			yyVAL.value = "(" + yyDollar[2].value + ")"
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:145
		{ /* todo */
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:146
		{
			yyVAL.value = "!(" + yyDollar[2].value + ")"
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:147
		{ /* todo */
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:148
		{ /* todo */
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:149
		{ /* todo */
		}
	}
	goto yystack /* stack new state and value */
}
