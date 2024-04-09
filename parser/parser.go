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
	yys     int
	tok     tok     // token read from lexer
	node    Node    // default for statements and expression
	nodes   Nodes   // default for lists of expressions or statements
	nodemap NodeMap // default set of Node, with string keys, using valid id syntax.
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
const RAW = 57356
const SLOW = 57357
const LEFT = 57358
const RIGHT = 57359
const MIDDLE = 57360
const RETURN = 57361
const COMMA = 57362
const FOR = 57363
const SELECT = 57364
const AS = 57365
const FROM = 57366
const WHERE = 57367
const LIMIT = 57368
const LPAREN = 57369
const RPAREN = 57370
const LBRACKET = 57371
const RBRACKET = 57372
const LBRACE = 57373
const RBRACE = 57374
const DOT = 57375
const LEN = 57376
const PLUS = 57377
const MINUS = 57378
const PLUSPLUS = 57379
const MINUSMINUS = 57380
const MULTI = 57381
const DIV = 57382
const MOD = 57383
const ABS = 57384
const NOT = 57385
const AND = 57386
const OR = 57387
const XOR = 57388
const EQ = 57389
const NEQ = 57390
const LT = 57391
const LTE = 57392
const GT = 57393
const GTE = 57394
const CONTAINS = 57395
const FIND = 57396
const PATH = 57397
const WITH = 57398
const JOIN = 57399
const PAGE = 57400
const COLON = 57401
const TEXT = 57402
const ATTR = 57403
const OF = 57404
const DISTINCT = 57405
const AT = 57406
const DOTDOT = 57407
const QUESTION = 57408
const BANG = 57409
const RANGE = 57410

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
	"RAW",
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

//line grammar.y:279

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 262

var yyAct = [...]uint8{
	40, 131, 67, 55, 68, 42, 52, 41, 27, 26,
	25, 30, 109, 73, 21, 126, 57, 36, 37, 127,
	74, 75, 76, 125, 110, 61, 21, 108, 132, 133,
	99, 24, 38, 32, 100, 33, 70, 83, 84, 107,
	86, 87, 102, 77, 88, 89, 90, 83, 84, 98,
	83, 84, 92, 93, 96, 97, 94, 95, 105, 83,
	84, 83, 84, 83, 84, 63, 80, 104, 31, 132,
	133, 69, 81, 91, 83, 84, 106, 115, 103, 77,
	60, 114, 101, 20, 78, 113, 112, 79, 80, 60,
	116, 117, 27, 26, 25, 30, 65, 35, 80, 134,
	118, 124, 120, 59, 111, 119, 121, 58, 54, 69,
	123, 122, 3, 56, 62, 24, 53, 32, 64, 33,
	44, 34, 50, 45, 46, 47, 48, 43, 128, 85,
	49, 51, 135, 136, 137, 27, 26, 25, 30, 5,
	82, 6, 130, 129, 16, 39, 17, 8, 18, 72,
	71, 2, 31, 27, 26, 25, 30, 29, 24, 66,
	32, 4, 33, 1, 23, 50, 45, 46, 47, 48,
	28, 22, 0, 49, 51, 0, 24, 0, 32, 0,
	33, 0, 0, 50, 45, 46, 47, 48, 0, 0,
	0, 49, 51, 0, 0, 31, 27, 26, 25, 30,
	0, 27, 26, 25, 30, 0, 0, 0, 0, 0,
	0, 0, 0, 31, 0, 0, 0, 0, 0, 24,
	0, 32, 0, 33, 24, 0, 32, 9, 33, 0,
	10, 11, 0, 12, 0, 13, 0, 0, 0, 7,
	0, 14, 15, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 19, 31, 0, 0, 0,
	0, 31,
}

var yyPact = [...]int16{
	-1000, -1000, 220, -1000, 220, -1000, -1000, 192, 220, 89,
	197, 197, 131, 99, 106, 149, -1000, -1000, 98, 197,
	60, -1000, -1000, -1000, 149, -1000, -1000, -1000, -1000, -1000,
	-1000, 107, 88, 64, -1000, 149, 4, 72, 78, 149,
	17, 5, -1000, 149, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 1, -1000, -1000, 70, -1000, 19, -1000, 69,
	197, 30, -1000, 46, -1000, -1000, 7, -1000, -47, -1000,
	15, 95, 4, -1000, -1000, -1000, -1000, -1000, 197, -1000,
	149, 68, 149, -1000, -1000, 149, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 149,
	102, 149, 106, -1000, -1000, -1000, -1000, -1000, 102, 149,
	-1000, -1000, -1000, 92, 17, -1000, 5, -1000, -7, -1000,
	6, -5, -1000, 17, -1000, -1000, -1000, 149, 3, 90,
	44, -1000, 149, 149, -1000, -1000, 17, 17,
}

var yyPgo = [...]uint8{
	0, 171, 170, 0, 141, 164, 2, 4, 32, 163,
	112, 161, 139, 159, 157, 151, 6, 150, 149, 13,
	148, 83, 147, 3, 143, 142, 1, 140, 7, 129,
	5, 127, 120, 116,
}

var yyR1 = [...]int8{
	0, 9, 15, 10, 10, 11, 11, 4, 4, 4,
	4, 4, 4, 17, 17, 18, 18, 19, 19, 19,
	19, 12, 12, 12, 20, 20, 21, 21, 22, 22,
	23, 24, 24, 25, 25, 26, 26, 5, 5, 3,
	3, 28, 28, 30, 30, 32, 32, 16, 16, 16,
	1, 1, 1, 1, 1, 33, 33, 2, 2, 8,
	8, 14, 14, 13, 13, 6, 31, 31, 31, 31,
	31, 31, 31, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 27, 27, 7,
}

var yyR2 = [...]int8{
	0, 2, 0, 2, 1, 1, 2, 4, 4, 5,
	3, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	1, 3, 4, 2, 0, 1, 1, 3, 5, 8,
	1, 0, 1, 1, 2, 2, 2, 1, 2, 3,
	1, 3, 1, 2, 1, 1, 1, 1, 1, 3,
	1, 1, 1, 1, 1, 4, 3, 3, 2, 1,
	3, 2, 3, 1, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -9, -15, -10, -11, -12, -4, 19, -22, 7,
	10, 11, 13, 15, 21, 22, -12, -4, -20, 63,
	-21, -16, -1, -5, 27, 6, 5, 4, -2, -14,
	7, 64, 29, 31, -10, 8, -16, -16, -8, 14,
	-3, -28, -30, -31, -32, 35, 36, 37, 38, 42,
	34, 43, -16, -33, 9, -23, 7, -3, 9, -21,
	20, -3, 7, -8, 30, 32, -13, -6, -7, 7,
	-3, -17, -18, -19, 16, 17, 18, -16, 12, 9,
	20, -8, -27, 44, 45, -29, 35, 36, 39, 40,
	41, 68, 47, 48, 51, 52, 49, 50, -30, 29,
	33, 12, 23, 9, -16, 28, 30, 32, 20, 59,
	9, 9, -19, -16, -3, 9, -28, -30, -3, -7,
	-3, -23, -6, -3, 9, 30, 9, 24, -3, -24,
	-25, -26, 25, 26, 9, -26, -3, -3,
}

var yyDef = [...]int8{
	2, -2, 0, 1, 0, 4, 5, 24, 0, 0,
	0, 0, 0, 0, 0, 0, 3, 6, 0, 0,
	25, 26, 47, 48, 0, 50, 51, 52, 53, 54,
	37, 0, 0, 0, 23, 0, 13, 0, 0, 0,
	59, 40, 42, 0, 44, 66, 67, 68, 69, 70,
	71, 72, 45, 46, 12, 0, 30, 0, 21, 0,
	0, 0, 38, 0, 58, 61, 0, 63, 0, 87,
	0, 0, 14, 15, 17, 18, 19, 20, 0, 10,
	0, 0, 0, 85, 86, 0, 73, 74, 75, 76,
	77, 78, 79, 80, 81, 82, 83, 84, 43, 0,
	0, 0, 0, 22, 27, 49, 57, 62, 0, 0,
	7, 8, 16, 0, 60, 11, 39, 41, 0, 56,
	0, 0, 64, 65, 9, 55, 28, 0, 31, 0,
	32, 33, 0, 0, 29, 34, 35, 36,
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
	62, 63, 64, 65, 66, 67, 68,
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
//line grammar.y:86
		{
			yyVAL.nodes = yyDollar[2].nodes
			lx.root = yyVAL.nodes
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:90
		{
			lx = yylex.(*myLexer)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:94
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[2].nodes...)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:95
		{
			yyVAL.nodes = yyDollar[1].nodes
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:99
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:100
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[2].node)
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:104
		{
			yyVAL.node = lx.newNodeAssign(yyDollar[1].tok, yyDollar[3].node)
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:105
		{ /*todo*/
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:106
		{ /*todo*/
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:109
		{
			yyVAL.node = nodePrint{nodes: yyDollar[2].nodes, raw: false}
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:110
		{
			yyVAL.node = nodePrint{nodes: yyDollar[3].nodes, raw: true}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:111
		{ /*todo*/
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:115
		{ /*todo*/
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:116
		{ /*todo*/
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:120
		{ /*todo*/
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:121
		{ /*todo*/
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:125
		{ /*todo*/
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:126
		{ /*todo*/
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:127
		{ /*todo*/
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:128
		{ /*todo*/
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:132
		{ /*todo*/
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:133
		{ /*todo*/
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:134
		{ /*todo*/
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:138
		{ /*todo*/
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:139
		{ /*todo*/
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:143
		{ /*todo*/
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:144
		{ /*todo*/
		}
	case 28:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:147
		{ /*todo*/
		}
	case 29:
		yyDollar = yyS[yypt-8 : yypt+1]
//line grammar.y:148
		{ /*todo*/
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:152
		{ /*todo*/
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:156
		{ /*todo*/
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:157
		{ /*todo*/
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:161
		{ /*todo*/
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:162
		{ /*todo*/
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:166
		{ /*todo*/
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:167
		{ /*todo*/
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:171
		{
			yyVAL.node = lx.newNodeVariable(yyDollar[1].tok, false)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:172
		{
			yyVAL.node = lx.newNodeVariable(yyDollar[2].tok, true)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:183
		{ /*todo*/
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:184
		{ /*todo*/
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:193
		{ /*todo*/
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:194
		{ /*todo*/
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:198
		{ /*todo*/
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:199
		{ /*todo*/
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:203
		{ /*todo*/
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:204
		{ /*todo*/
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:205
		{ /*todo*/
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:209
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:210
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:211
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:212
		{
			yyVAL.node = yyDollar[1].node
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:213
		{
			yyVAL.node = yyDollar[1].nodemap
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:217
		{ /*todo*/
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:218
		{ /*todo*/
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:222
		{
			yyVAL.node = yyDollar[2].nodes
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:223
		{
			yyVAL.node = Nodes(nil)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:227
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:228
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[3].node)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:232
		{
			yyVAL.nodemap = lx.newNodeMap(nil, nil)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:233
		{
			yyVAL.nodemap = yyDollar[2].nodemap
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:237
		{
			yyVAL.nodemap = lx.newNodeMap(nil, yyDollar[1].node)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:238
		{
			yyVAL.nodemap = lx.newNodeMap(yyDollar[1].nodemap, yyDollar[3].node)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:242
		{
			yyVAL.node = lx.newNodeKeyValue(yyDollar[1].node, yyDollar[3].node)
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:246
		{ /*todo*/
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:247
		{ /*todo*/
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:248
		{ /*todo*/
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:249
		{ /*todo*/
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:250
		{ /*todo*/
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:251
		{ /*todo*/
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:252
		{ /*todo*/
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:257
		{ /*todo*/
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:258
		{ /*todo*/
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:259
		{ /*todo*/
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:260
		{ /*todo*/
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:261
		{ /*todo*/
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:262
		{ /*todo*/
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:263
		{ /*todo*/
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:264
		{ /*todo*/
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:265
		{ /*todo*/
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:266
		{ /*todo*/
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:267
		{ /*todo*/
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:268
		{ /*todo*/
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:272
		{ /*todo*/
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:273
		{ /*todo*/
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:277
		{
			yyVAL.node = lx.newNodeKey(yyDollar[1].tok)
		}
	}
	goto yystack /* stack new state and value */
}
