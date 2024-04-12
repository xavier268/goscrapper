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
	yys          int
	tok          tok          // token read from lexer, implements Node.
	node         Node         // default for statements and expression
	nodes        Nodes        // default for lists of expressions or statements, implements Node.
	nodemap      NodeMap      // default set of Node, with string keys, using valid id syntax, implements Node.
	nodeWithBody NodeWithBody // a node that incorporates a set of nodes
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
const TO = 57367
const STEP = 57368
const WHERE = 57369
const LIMIT = 57370
const LPAREN = 57371
const RPAREN = 57372
const LBRACKET = 57373
const RBRACKET = 57374
const LBRACE = 57375
const RBRACE = 57376
const DOT = 57377
const LEN = 57378
const PLUS = 57379
const MINUS = 57380
const PLUSPLUS = 57381
const MINUSMINUS = 57382
const MULTI = 57383
const DIV = 57384
const MOD = 57385
const ABS = 57386
const NOT = 57387
const AND = 57388
const OR = 57389
const XOR = 57390
const NAND = 57391
const EQ = 57392
const NEQ = 57393
const LT = 57394
const LTE = 57395
const GT = 57396
const GTE = 57397
const CONTAINS = 57398
const FIND = 57399
const PATH = 57400
const WITH = 57401
const JOIN = 57402
const PAGE = 57403
const COLON = 57404
const TEXT = 57405
const ATTR = 57406
const OF = 57407
const DISTINCT = 57408
const AT = 57409
const DOTDOT = 57410
const QUESTION = 57411
const NOW = 57412
const VERSION = 57413
const FILE_SEPARATOR = 57414
const RANGE = 57415

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
	"TO",
	"STEP",
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
	"NAND",
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
	"NOW",
	"VERSION",
	"FILE_SEPARATOR",
	"RANGE",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:316

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 271

var yyAct = [...]uint8{
	92, 141, 95, 60, 143, 35, 96, 103, 21, 22,
	126, 55, 56, 23, 59, 90, 64, 47, 46, 45,
	48, 3, 140, 47, 46, 45, 48, 125, 123, 123,
	108, 69, 70, 72, 71, 104, 105, 106, 86, 66,
	122, 124, 42, 89, 43, 91, 44, 111, 42, 31,
	26, 27, 28, 29, 67, 100, 87, 30, 32, 110,
	88, 107, 112, 113, 69, 70, 72, 71, 116, 144,
	145, 97, 58, 109, 33, 99, 34, 65, 117, 53,
	49, 144, 145, 50, 51, 52, 49, 118, 119, 50,
	51, 52, 18, 54, 97, 120, 114, 115, 93, 63,
	69, 70, 72, 71, 98, 142, 102, 101, 107, 128,
	127, 129, 130, 2, 57, 68, 133, 73, 132, 69,
	70, 72, 71, 24, 135, 63, 5, 137, 136, 8,
	62, 16, 139, 47, 46, 45, 48, 38, 6, 146,
	94, 148, 61, 17, 19, 150, 151, 149, 153, 20,
	154, 4, 36, 155, 1, 40, 152, 41, 42, 25,
	43, 37, 44, 39, 138, 31, 26, 27, 28, 29,
	0, 0, 0, 30, 32, 0, 69, 70, 72, 71,
	0, 47, 46, 45, 48, 69, 70, 72, 71, 0,
	33, 147, 34, 104, 105, 106, 49, 0, 0, 50,
	51, 52, 0, 74, 75, 0, 42, 76, 77, 78,
	131, 69, 70, 72, 71, 0, 79, 80, 83, 84,
	81, 82, 134, 69, 70, 72, 71, 0, 0, 0,
	85, 69, 70, 72, 71, 121, 69, 70, 72, 71,
	0, 0, 0, 0, 49, 0, 0, 50, 51, 52,
	0, 69, 70, 72, 71, 9, 0, 0, 10, 11,
	0, 12, 0, 13, 0, 0, 0, 7, 0, 14,
	15,
}

var yyPact = [...]int16{
	-1000, -1000, 248, -1000, 248, -1000, 83, 129, 70, 85,
	129, 129, 58, 129, 118, 129, -1000, 68, -1000, 30,
	34, -15, 166, -1000, 129, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 25, -1000, -1000, -1000, -1000,
	-1000, -1000, 129, 13, 64, -1000, -1000, -1000, -1000, 97,
	-1000, -1000, -1000, 248, 129, 177, 18, 129, -1000, -15,
	35, 129, 129, -1000, 73, -1000, -1000, 129, 129, -1000,
	-1000, -1000, -1000, 129, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 129, 87, 205,
	8, -1000, -15, -1000, 7, -1000, -52, -1000, -1000, -1000,
	-15, -1000, 19, -1000, -1000, -1000, -1000, -1000, 129, 9,
	129, 129, 185, -15, 92, 129, -15, 166, -1000, 190,
	-1000, -1000, -1000, 129, -1000, 87, 129, -1000, -15, 139,
	-15, 129, -2, 54, -1000, -15, -1000, -15, 129, 165,
	129, -1000, 42, -1000, 129, 129, 130, 129, 54, -1000,
	-15, -15, 129, -15, -1000, -15,
}

var yyPgo = [...]uint8{
	0, 163, 161, 5, 0, 9, 13, 159, 157, 138,
	155, 2, 6, 154, 126, 152, 15, 21, 151, 149,
	144, 140, 137, 129, 123, 117, 115, 3, 114, 113,
	107, 106, 7, 1, 105, 4,
}

var yyR1 = [...]int8{
	0, 13, 29, 17, 17, 18, 18, 9, 9, 9,
	9, 9, 9, 28, 28, 30, 30, 31, 31, 32,
	32, 32, 32, 14, 14, 20, 20, 19, 19, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 27, 33,
	33, 34, 34, 35, 35, 10, 10, 4, 4, 5,
	5, 6, 6, 7, 7, 7, 7, 3, 3, 3,
	3, 1, 1, 1, 15, 15, 2, 2, 16, 16,
	22, 22, 21, 21, 11, 8, 8, 8, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 26,
	26, 26, 26, 12,
}

var yyR2 = [...]int8{
	0, 2, 0, 2, 1, 2, 3, 3, 3, 4,
	3, 1, 2, 0, 1, 0, 1, 1, 2, 1,
	1, 1, 1, 3, 3, 0, 1, 1, 3, 1,
	8, 7, 6, 5, 4, 3, 7, 5, 1, 0,
	1, 1, 2, 2, 2, 1, 2, 3, 1, 3,
	1, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 1, 1, 4, 3, 3, 2, 1, 3,
	2, 3, 1, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -13, -29, -17, -18, -14, -9, 19, -23, 7,
	10, 11, 13, 15, 21, 22, -14, -9, 9, -20,
	-19, -4, -5, -6, -24, -7, 37, 38, 39, 40,
	44, 36, 45, 61, 63, -3, -15, -2, -22, -1,
	-10, -8, 29, 31, 33, 6, 5, 4, 7, 67,
	70, 71, 72, 9, 8, -4, -4, -28, 14, -4,
	-27, 24, 12, 7, -4, 9, 9, 20, -26, 46,
	47, 49, 48, -25, 37, 38, 41, 42, 43, 50,
	51, 54, 55, 52, 53, 64, -6, 31, 35, -4,
	-16, 32, -4, 34, -21, -11, -12, 7, 7, -17,
	-4, -30, -31, -32, 16, 17, 18, -3, 12, -16,
	24, 12, -4, -4, 23, 24, -4, -5, -6, -4,
	-12, 30, 32, 20, 34, 20, 62, -32, -4, -4,
	-4, 25, -27, -4, 32, -4, -11, -4, 25, -4,
	24, -33, -34, -35, 27, 28, -4, 26, -4, -35,
	-4, -4, 26, -4, -33, -4,
}

var yyDef = [...]int8{
	2, -2, 0, 1, 0, 4, 0, 25, 0, 0,
	0, 0, 13, 11, 29, 0, 3, 0, 5, 0,
	26, 27, 48, 50, 0, 52, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 53, 54, 55, 56, 57,
	58, 59, 0, 0, 0, 61, 62, 63, 45, 0,
	75, 76, 77, 0, 0, 15, 0, 0, 14, 12,
	0, 0, 0, 38, 0, 6, 23, 0, 0, 99,
	100, 101, 102, 0, 87, 88, 89, 90, 91, 92,
	93, 94, 95, 96, 97, 98, 51, 0, 0, 0,
	0, 67, 68, 70, 0, 72, 0, 103, 46, 24,
	7, 8, 16, 17, 19, 20, 21, 22, 0, 10,
	0, 0, 0, 35, 0, 0, 28, 47, 49, 0,
	65, 60, 66, 0, 71, 0, 0, 18, 9, 0,
	34, 0, 0, 39, 64, 69, 73, 74, 0, 33,
	0, 37, 40, 41, 0, 0, 32, 0, 39, 42,
	43, 44, 0, 31, 36, 30,
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
	72, 73,
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
//line grammar.y:92
		{
			yyVAL.node = nodeProgram{req: yyDollar[2].nodes, invars: lx.ParamsList()}
			lx.root = yyVAL.node
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:96
		{
			lx = yylex.(*myLexer)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:100
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[2].node)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:101
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:105
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:106
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[2].node)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:110
		{
			yyVAL.node = lx.newNodeAssign(yyDollar[1].tok, yyDollar[3].node)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:111
		{ /*todo*/
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:112
		{ /*todo*/
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:115
		{
			yyVAL.node = nodePrint{nodes: yyDollar[3].nodes, raw: (yyDollar[2].tok.c == RAW)}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:116
		{
			yyVAL.node = nodeSlow{m: nil}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:117
		{
			yyVAL.node = nodeSlow{m: yyDollar[1].tok}
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:121
		{
			yyVAL.tok = tok{}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:122
		{
			yyVAL.tok = yyDollar[1].tok
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:126
		{ /*todo*/
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:127
		{ /*todo*/
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:131
		{ /*todo*/
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:132
		{ /*todo*/
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:136
		{ /*todo*/
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:137
		{ /*todo*/
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:138
		{ /*todo*/
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:139
		{ /*todo*/
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:143
		{
			yyVAL.node = nodeReturn{yyDollar[2].nodes}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:144
		{
			yyVAL.node = yyDollar[1].nodeWithBody.appendBody(yyDollar[3].nodes)
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:148
		{
			yyVAL.nodes = Nodes{}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:153
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:154
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[3].node)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:157
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(nil, nil, nil, nil)
		}
	case 30:
		yyDollar = yyS[yypt-8 : yypt+1]
//line grammar.y:160
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(yyDollar[2].tok, yyDollar[4].node, yyDollar[6].node, yyDollar[8].node)
		}
	case 31:
		yyDollar = yyS[yypt-7 : yypt+1]
//line grammar.y:162
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(nil, yyDollar[3].node, yyDollar[5].node, yyDollar[7].node)
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:165
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(yyDollar[2].tok, yyDollar[4].node, yyDollar[6].node, nil)
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:167
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(nil, yyDollar[3].node, yyDollar[5].node, nil)
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:169
		{ /*todo*/
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:170
		{ /*todo*/
		}
	case 36:
		yyDollar = yyS[yypt-7 : yypt+1]
//line grammar.y:173
		{ /*todo*/
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:175
		{ /*todo*/
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:183
		{ /*todo*/
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:184
		{ /*todo*/
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:188
		{ /*todo*/
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:189
		{ /*todo*/
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:193
		{ /*todo*/
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:194
		{ /*todo*/
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:198
		{
			yyVAL.node = lx.newNodeVariable(yyDollar[1].tok, false, true)
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:199
		{
			yyVAL.node = lx.newNodeVariable(yyDollar[2].tok, true, false)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:207
		{
			yyVAL.node = lx.newNodeOpe2Bool(yyDollar[1].node, yyDollar[2].tok, yyDollar[3].node)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:212
		{
			yyVAL.node = lx.newNodeOpe2(yyDollar[1].node, yyDollar[2].tok, yyDollar[3].node)
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:217
		{
			yyVAL.node = lx.newNodeOpe1(yyDollar[1].tok, yyDollar[2].node)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:224
		{
			yyVAL.node = yyDollar[1].node
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:225
		{
			yyVAL.node = yyDollar[1].nodemap
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:232
		{
			yyVAL.node = yyDollar[2].node
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:236
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:237
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:238
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:242
		{
			yyVAL.node = nodeArrayAccess{a: yyDollar[1].node, i: yyDollar[3].node}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:243
		{
			yyVAL.node = nodeMapAccess{m: yyDollar[1].node, k: yyDollar[3].node}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:247
		{
			yyVAL.node = yyDollar[2].nodes
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:248
		{
			yyVAL.node = Nodes(nil)
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:252
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:253
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[3].node)
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:257
		{
			yyVAL.nodemap = lx.newNodeMap(nil, nil)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:258
		{
			yyVAL.nodemap = yyDollar[2].nodemap
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:262
		{
			yyVAL.nodemap = lx.newNodeMap(nil, yyDollar[1].node)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:263
		{
			yyVAL.nodemap = lx.newNodeMap(yyDollar[1].nodemap, yyDollar[3].node)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:267
		{
			yyVAL.node = lx.newNodeKeyValue(yyDollar[1].node, yyDollar[3].node)
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:271
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:272
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:273
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:314
		{
			yyVAL.node = lx.newNodeKey(yyDollar[1].tok)
		}
	}
	goto yystack /* stack new state and value */
}
