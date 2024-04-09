// Code generated by goyacc -o parser.go grammar.y. DO NOT EDIT.

//line grammar.y:1

package parser

import __yyfmt__ "fmt"

//line grammar.y:3

import (
	"fmt"
)

// keep the compiler happy
var _ = fmt.Println

type tok struct {
	v string // token cvalue
	t string // token type
	c int    // lexer/parser constant code
}

var lx *myLexer // shorthand for lx

//line grammar.y:23
type yySymType struct {
	yys   int
	tok   tok   // token read from lexer
	node  Node  // default for statements and expression
	nodes Nodes // default for lists of expressions or statements
}

const BOOL = 57346
const NUMBER = 57347
const STRING = 57348
const IDENTIFIER = 57349
const ASSIGN = 57350
const SEMICOLON = 57351
const CLICK = 57352
const INPUT = 57353
const IN = 57354
const PRINT = 57355
const SLOW = 57356
const LEFT = 57357
const RIGHT = 57358
const MIDDLE = 57359
const RETURN = 57360
const COMMA = 57361
const FOR = 57362
const SELECT = 57363
const AS = 57364
const FROM = 57365
const WHERE = 57366
const LIMIT = 57367
const LPAREN = 57368
const RPAREN = 57369
const LBRACKET = 57370
const RBRACKET = 57371
const LBRACE = 57372
const RBRACE = 57373
const DOT = 57374
const LEN = 57375
const PLUS = 57376
const MINUS = 57377
const PLUSPLUS = 57378
const MINUSMINUS = 57379
const MULTI = 57380
const DIV = 57381
const MOD = 57382
const ABS = 57383
const NOT = 57384
const AND = 57385
const OR = 57386
const XOR = 57387
const EQ = 57388
const NEQ = 57389
const LT = 57390
const LTE = 57391
const GT = 57392
const GTE = 57393
const CONTAINS = 57394
const FIND = 57395
const PATH = 57396
const WITH = 57397
const JOIN = 57398
const PAGE = 57399
const COLON = 57400
const TEXT = 57401
const ATTR = 57402
const OF = 57403
const DISTINCT = 57404
const AT = 57405
const DOTDOT = 57406
const QUESTION = 57407
const BANG = 57408
const RANGE = 57409

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"BOOL",
	"NUMBER",
	"STRING",
	"IDENTIFIER",
	"ASSIGN",
	"SEMICOLON",
	"CLICK",
	"INPUT",
	"IN",
	"PRINT",
	"SLOW",
	"LEFT",
	"RIGHT",
	"MIDDLE",
	"RETURN",
	"COMMA",
	"FOR",
	"SELECT",
	"AS",
	"FROM",
	"WHERE",
	"LIMIT",
	"LPAREN",
	"RPAREN",
	"LBRACKET",
	"RBRACKET",
	"LBRACE",
	"RBRACE",
	"DOT",
	"LEN",
	"PLUS",
	"MINUS",
	"PLUSPLUS",
	"MINUSMINUS",
	"MULTI",
	"DIV",
	"MOD",
	"ABS",
	"NOT",
	"AND",
	"OR",
	"XOR",
	"EQ",
	"NEQ",
	"LT",
	"LTE",
	"GT",
	"GTE",
	"CONTAINS",
	"FIND",
	"PATH",
	"WITH",
	"JOIN",
	"PAGE",
	"COLON",
	"TEXT",
	"ATTR",
	"OF",
	"DISTINCT",
	"AT",
	"DOTDOT",
	"QUESTION",
	"BANG",
	"RANGE",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:276

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 237

var yyAct = [...]uint8{
	39, 128, 66, 54, 67, 41, 51, 40, 27, 26,
	25, 30, 107, 72, 21, 124, 56, 36, 37, 73,
	74, 75, 20, 122, 123, 60, 21, 108, 81, 82,
	24, 97, 32, 59, 33, 98, 69, 81, 82, 38,
	84, 85, 58, 76, 86, 87, 88, 103, 96, 129,
	130, 68, 90, 91, 94, 95, 92, 93, 81, 82,
	100, 81, 82, 81, 82, 106, 102, 31, 81, 82,
	99, 77, 62, 89, 131, 64, 79, 105, 76, 35,
	112, 81, 82, 121, 111, 110, 104, 101, 113, 114,
	27, 26, 25, 30, 129, 130, 78, 59, 115, 109,
	117, 57, 53, 116, 118, 68, 79, 55, 120, 119,
	61, 65, 24, 5, 32, 63, 33, 3, 16, 49,
	44, 45, 46, 47, 29, 125, 34, 48, 50, 132,
	133, 134, 27, 26, 25, 30, 6, 52, 43, 42,
	83, 17, 80, 23, 127, 126, 8, 18, 71, 31,
	27, 26, 25, 30, 24, 70, 32, 2, 33, 4,
	1, 49, 44, 45, 46, 47, 28, 22, 0, 48,
	50, 0, 24, 0, 32, 0, 33, 27, 26, 25,
	30, 9, 0, 0, 10, 11, 0, 12, 13, 0,
	0, 31, 7, 0, 14, 15, 0, 0, 0, 24,
	0, 32, 0, 33, 0, 0, 0, 0, 19, 31,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 31,
}

var yyPact = [...]int16{
	-1000, -1000, 174, -1000, 174, -1000, -1000, 146, 174, 71,
	173, 173, 128, 93, 100, 128, -1000, -1000, 92, 173,
	14, -1000, -1000, -1000, 128, -1000, -1000, -1000, -1000, -1000,
	-1000, 103, 86, 44, -1000, 128, 4, 59, 87, -15,
	6, -1000, 128, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 3, -1000, -1000, 58, -1000, 38, -1000, 78, 173,
	20, -1000, 57, -1000, -1000, 46, -1000, -46, -1000, 18,
	90, 4, -1000, -1000, -1000, -1000, -1000, 173, -1000, 128,
	128, -1000, -1000, 128, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 128, 98, 128,
	100, -1000, -1000, -1000, -1000, -1000, 98, 128, -1000, -1000,
	-1000, 74, -15, 6, -1000, -6, -1000, 15, -8, -1000,
	-15, -1000, -1000, -1000, 128, 25, 65, 70, -1000, 128,
	128, -1000, -1000, -15, -15,
}

var yyPgo = [...]uint8{
	0, 167, 166, 0, 136, 39, 160, 117, 159, 113,
	157, 6, 155, 148, 13, 147, 22, 146, 3, 145,
	144, 1, 143, 4, 142, 7, 140, 5, 139, 138,
	137, 124, 111, 2,
}

var yyR1 = [...]int8{
	0, 6, 10, 7, 7, 8, 8, 4, 4, 4,
	4, 4, 12, 12, 13, 13, 14, 14, 14, 14,
	9, 9, 9, 15, 15, 16, 16, 17, 17, 18,
	19, 19, 20, 20, 21, 21, 22, 22, 23, 3,
	3, 25, 25, 27, 27, 29, 29, 11, 11, 11,
	1, 1, 1, 1, 1, 30, 30, 2, 2, 5,
	5, 31, 31, 32, 32, 33, 28, 28, 28, 28,
	28, 28, 28, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 24, 24,
}

var yyR2 = [...]int8{
	0, 2, 0, 2, 1, 1, 2, 4, 4, 5,
	3, 2, 0, 1, 1, 2, 1, 1, 1, 1,
	3, 4, 2, 0, 1, 1, 3, 5, 8, 1,
	0, 1, 1, 2, 2, 2, 1, 2, 1, 3,
	1, 3, 1, 2, 1, 1, 1, 1, 1, 3,
	1, 1, 1, 1, 1, 4, 3, 3, 2, 1,
	3, 2, 3, 1, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -6, -10, -7, -8, -9, -4, 18, -17, 7,
	10, 11, 13, 14, 20, 21, -9, -4, -15, 62,
	-16, -11, -1, -22, 26, 6, 5, 4, -2, -31,
	7, 63, 28, 30, -7, 8, -11, -11, -5, -3,
	-25, -27, -28, -29, 34, 35, 36, 37, 41, 33,
	42, -11, -30, 9, -18, 7, -3, 9, -16, 19,
	-3, 7, -5, 29, 31, -32, -33, -23, 7, -3,
	-12, -13, -14, 15, 16, 17, -11, 12, 9, 19,
	-24, 43, 44, -26, 34, 35, 38, 39, 40, 67,
	46, 47, 50, 51, 48, 49, -27, 28, 32, 12,
	22, 9, -11, 27, 29, 31, 19, 58, 9, 9,
	-14, -11, -3, -25, -27, -3, -23, -3, -18, -33,
	-3, 9, 29, 9, 23, -3, -19, -20, -21, 24,
	25, 9, -21, -3, -3,
}

var yyDef = [...]int8{
	2, -2, 0, 1, 0, 4, 5, 23, 0, 0,
	0, 0, 0, 0, 0, 0, 3, 6, 0, 0,
	24, 25, 47, 48, 0, 50, 51, 52, 53, 54,
	36, 0, 0, 0, 22, 0, 12, 0, 0, 59,
	40, 42, 0, 44, 66, 67, 68, 69, 70, 71,
	72, 45, 46, 11, 0, 29, 0, 20, 0, 0,
	0, 37, 0, 58, 61, 0, 63, 0, 38, 0,
	0, 13, 14, 16, 17, 18, 19, 0, 10, 0,
	0, 85, 86, 0, 73, 74, 75, 76, 77, 78,
	79, 80, 81, 82, 83, 84, 43, 0, 0, 0,
	0, 21, 26, 49, 57, 62, 0, 0, 7, 8,
	15, 0, 60, 39, 41, 0, 56, 0, 0, 64,
	65, 9, 55, 27, 0, 30, 0, 31, 32, 0,
	0, 28, 33, 34, 35,
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
	62, 63, 64, 65, 66, 67,
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
//line grammar.y:84
		{
			yyVAL.nodes = yyDollar[2].nodes
			lx.root = yyVAL.nodes
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:88
		{
			lx = yylex.(*myLexer)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:92
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[2].nodes...)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:93
		{
			yyVAL.nodes = yyDollar[1].nodes
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:97
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:98
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[2].node)
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:102
		{ /*todo*/
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:103
		{ /*todo*/
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:104
		{ /*todo*/
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:107
		{
			yyVAL.node = nodePrint{yyDollar[2].nodes}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:108
		{ /*todo*/
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:112
		{ /*todo*/
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:113
		{ /*todo*/
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:117
		{ /*todo*/
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:118
		{ /*todo*/
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:122
		{ /*todo*/
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:123
		{ /*todo*/
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:124
		{ /*todo*/
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:125
		{ /*todo*/
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:129
		{ /*todo*/
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:130
		{ /*todo*/
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:131
		{ /*todo*/
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:135
		{ /*todo*/
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:136
		{ /*todo*/
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:140
		{ /*todo*/
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:141
		{ /*todo*/
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:144
		{ /*todo*/
		}
	case 28:
		yyDollar = yyS[yypt-8 : yypt+1]
//line grammar.y:145
		{ /*todo*/
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:149
		{ /*todo*/
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:153
		{ /*todo*/
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:154
		{ /*todo*/
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:158
		{ /*todo*/
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:159
		{ /*todo*/
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:163
		{ /*todo*/
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:164
		{ /*todo*/
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:168
		{ /*todo*/
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:169
		{ /*todo*/
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:173
		{ /*todo*/
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:182
		{ /*todo*/
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:183
		{ /*todo*/
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:192
		{ /*todo*/
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:193
		{ /*todo*/
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:197
		{ /*todo*/
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:198
		{ /*todo*/
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:202
		{ /*todo*/
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:203
		{ /*todo*/
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:204
		{ /*todo*/
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:208
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:209
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:210
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:211
		{
			yyVAL.node = yyDollar[1].node
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:212
		{ /*todo*/
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:216
		{ /*todo*/
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:217
		{ /*todo*/
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:221
		{
			yyVAL.node = yyDollar[2].nodes
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:222
		{
			yyVAL.node = Nodes(nil)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:226
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:227
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[3].node)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:231
		{ /*todo*/
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:232
		{ /*todo*/
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:236
		{ /*todo*/
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:237
		{ /*todo*/
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:241
		{ /*todo*/
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:245
		{ /*todo*/
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:246
		{ /*todo*/
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:247
		{ /*todo*/
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:248
		{ /*todo*/
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:249
		{ /*todo*/
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:250
		{ /*todo*/
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:251
		{ /*todo*/
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:256
		{ /*todo*/
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:257
		{ /*todo*/
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:258
		{ /*todo*/
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:259
		{ /*todo*/
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:260
		{ /*todo*/
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:261
		{ /*todo*/
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:262
		{ /*todo*/
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:263
		{ /*todo*/
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:264
		{ /*todo*/
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:265
		{ /*todo*/
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:266
		{ /*todo*/
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:267
		{ /*todo*/
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:271
		{ /*todo*/
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:272
		{ /*todo*/
		}
	}
	goto yystack /* stack new state and value */
}
