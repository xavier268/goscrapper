// Code generated by goyacc -o parser.go grammar.y. DO NOT EDIT.

//line grammar.y:1

package parser

import __yyfmt__ "fmt"

//line grammar.y:3

// each object has a value and a type.
type value struct {
	v string // a string in go that produce the value of the object
	t string // a string representing the gotype of the object
	c int    // the ope code is stored here
}

//line grammar.y:35
type yySymType struct {
	yys   int
	value value
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
const LOWER = 57376
const UPPER = 57377
const INTTYPE = 57378
const BOOLTYPE = 57379
const STRINGTYPE = 57380
const FOR = 57381
const RETURN = 57382
const WAITFOR = 57383
const OPTIONS = 57384
const IGNORE = 57385
const HEADLESS = 57386
const TIMEOUT = 57387
const DISTINCT = 57388
const FILTER = 57389
const CURRENT = 57390
const SORT = 57391
const LIMIT = 57392
const LET = 57393
const COLLECT = 57394
const ASC = 57395
const DESC = 57396
const NONE = 57397
const NULL = 57398
const TRUE = 57399
const FALSE = 57400
const USE = 57401
const INTO = 57402
const KEEP = 57403
const WITH = 57404
const COUNT = 57405
const ALL = 57406
const ANY = 57407
const AGGREGATE = 57408
const EVENT = 57409
const LIKE = 57410
const IN = 57411
const WHILE = 57412
const BOOL = 57413
const AT = 57414
const IDENTIFIER = 57415
const IGNOREID = 57416
const STRING = 57417
const NUMBER = 57418
const NAMESPACESEPARATOR = 57419
const SELECT = 57420
const CLICK = 57421
const DOCUMENT = 57422
const PAGE = 57423
const CONTAINS = 57424

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
	"LOWER",
	"UPPER",
	"INTTYPE",
	"BOOLTYPE",
	"STRINGTYPE",
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

//line grammar.y:181

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 182

var yyAct = [...]int8{
	47, 48, 49, 45, 46, 38, 53, 51, 52, 50,
	54, 55, 36, 66, 39, 11, 10, 63, 65, 64,
	43, 9, 56, 57, 58, 3, 22, 19, 16, 4,
	25, 5, 17, 62, 2, 11, 10, 47, 48, 49,
	45, 46, 24, 53, 51, 52, 50, 54, 55, 12,
	60, 1, 6, 23, 14, 15, 70, 13, 18, 56,
	57, 58, 44, 28, 26, 27, 11, 10, 20, 12,
	7, 69, 0, 0, 14, 15, 0, 13, 59, 47,
	48, 49, 45, 46, 29, 53, 51, 52, 50, 54,
	55, 35, 0, 32, 0, 33, 34, 72, 24, 8,
	12, 56, 57, 58, 0, 14, 15, 0, 13, 0,
	30, 31, 0, 0, 21, 59, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 37, 0, 0, 40, 41,
	42, 0, 0, 0, 0, 0, 0, 0, 0, 61,
	0, 0, 0, 0, 0, 0, 0, 35, 0, 32,
	0, 33, 34, 0, 67, 0, 0, 59, 68, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 71,
}

var yyPact = [...]int16{
	27, -1000, 27, -1000, -1000, -24, -1000, 76, -61, -1000,
	76, -68, -16, 76, 76, 76, -1000, -1000, -1000, -1000,
	1, 75, -1000, 20, 76, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -19, 75, -56, 76,
	75, 75, 75, 76, 20, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 33, -1000, -1000, -1000, -1000, 76, 75, 75, -1000,
	-1000, -4, -1000,
}

var yyPgo = [...]int8{
	0, 114, 68, 26, 65, 64, 63, 62, 53, 51,
	34, 25, 29, 33, 31, 52, 21, 30,
}

var yyR1 = [...]int8{
	0, 9, 9, 10, 10, 12, 12, 13, 13, 13,
	2, 2, 11, 11, 14, 14, 16, 16, 16, 16,
	15, 15, 15, 5, 4, 6, 17, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 8, 8, 8, 1, 1, 1, 3, 3,
	3, 3, 3,
}

var yyR2 = [...]int8{
	0, 2, 1, 2, 1, 2, 3, 1, 1, 1,
	1, 3, 2, 1, 1, 2, 3, 2, 2, 2,
	2, 1, 5, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 2, 3, 1,
	1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -9, -10, -11, -12, -14, -15, 43, 72, -16,
	40, 39, 73, 81, 78, 79, -11, -12, -15, -16,
	-2, -1, -3, -8, 22, -17, -5, -4, -6, 8,
	34, 35, 73, 75, 76, 71, 73, -1, 73, 30,
	-1, -1, -1, 19, -7, 7, 8, 4, 5, 6,
	13, 11, 12, 10, 14, 15, 26, 27, 28, 82,
	-3, -1, -13, 36, 38, 37, 69, -1, -1, -3,
	23, -1, -11,
}

var yyDef = [...]int8{
	0, -2, 0, 2, 4, 0, 13, 0, 0, 14,
	21, 0, 0, 0, 0, 0, 1, 3, 12, 15,
	5, 10, 45, 0, 0, 49, 50, 51, 52, 42,
	43, 44, 26, 23, 24, 25, 0, 20, 0, 0,
	17, 18, 19, 0, 0, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	47, 0, 6, 7, 8, 9, 0, 16, 11, 46,
	48, 0, 22,
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
	82,
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
//line grammar.y:68
		{
			yylex.(*myLexer).finalize()
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:69
		{
			yylex.(*myLexer).finalize()
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:80
		{ /* todo */
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:81
		{ /* todo */
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:92
		{ /*todo*/
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:98
		{ /* todo */
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:99
		{ /* todo */
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:103
		{ /* todo */
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:104
		{ /* todo */
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:108
		{ /*todo*/
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:109
		{ /* todo */
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:110
		{ /* todo */
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:111
		{ /* todo */
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:115
		{ /* todo */
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:116
		{ /* todo */
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:117
		{ /* todo */
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:166
		{
			yyVAL.value = yyDollar[1].value
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:167
		{
			yyVAL.value = yylex.(*myLexer).Ope2(yyDollar[2].value.c, yyDollar[1].value, yyDollar[3].value)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:168
		{
			yyVAL.value = yylex.(*myLexer).Ope1(yyDollar[1].value.c, yyDollar[2].value)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:173
		{
			yyVAL.value = yylex.(*myLexer).Paren(yyDollar[2].value)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:174
		{ /* todo */
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:175
		{ /* todo */
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:176
		{ /* todo */
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:177
		{ /* todo */
		}
	}
	goto yystack /* stack new state and value */
}
