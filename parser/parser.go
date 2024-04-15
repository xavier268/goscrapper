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
const SLOW = 57355
const LEFT = 57356
const RIGHT = 57357
const MIDDLE = 57358
const RETURN = 57359
const COMMA = 57360
const FOR = 57361
const SELECT = 57362
const AS = 57363
const FROM = 57364
const TO = 57365
const STEP = 57366
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
const NAND = 57389
const EQ = 57390
const NEQ = 57391
const LT = 57392
const LTE = 57393
const GT = 57394
const GTE = 57395
const CONTAINS = 57396
const FIND = 57397
const PATH = 57398
const WITH = 57399
const JOIN = 57400
const PAGE = 57401
const COLON = 57402
const TEXT = 57403
const ATTR = 57404
const OF = 57405
const DISTINCT = 57406
const AT = 57407
const DOTDOT = 57408
const QUESTION = 57409
const NOW = 57410
const VERSION = 57411
const FILE_SEPARATOR = 57412
const IF = 57413
const THEN = 57414
const ELSE = 57415
const ASSERT = 57416
const FAIL = 57417
const PRINT = 57418
const FORMAT = 57419
const RAW = 57420
const GO = 57421
const JSON = 57422
const GSC = 57423
const NL = 57424
const DOLLAR = 57425
const NIL = 57426
const LAST = 57427
const RANGE = 57428

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
	"IF",
	"THEN",
	"ELSE",
	"ASSERT",
	"FAIL",
	"PRINT",
	"FORMAT",
	"RAW",
	"GO",
	"JSON",
	"GSC",
	"NL",
	"DOLLAR",
	"NIL",
	"LAST",
	"RANGE",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:365

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 477

var yyAct = [...]uint8{
	31, 175, 9, 120, 10, 177, 71, 121, 49, 137,
	6, 180, 3, 33, 75, 25, 156, 83, 79, 80,
	81, 82, 84, 85, 86, 32, 155, 174, 122, 61,
	60, 59, 62, 143, 134, 143, 93, 94, 96, 95,
	154, 93, 94, 96, 95, 113, 91, 153, 112, 114,
	146, 145, 56, 118, 57, 4, 58, 115, 84, 41,
	36, 37, 38, 39, 190, 178, 179, 40, 42, 144,
	76, 127, 88, 128, 129, 116, 87, 70, 133, 178,
	179, 126, 26, 125, 43, 77, 44, 25, 141, 29,
	63, 27, 147, 65, 66, 67, 122, 74, 93, 94,
	96, 95, 172, 45, 46, 48, 47, 68, 64, 69,
	28, 149, 124, 123, 150, 130, 131, 78, 148, 2,
	89, 90, 151, 93, 94, 96, 95, 157, 158, 5,
	92, 97, 161, 34, 24, 162, 176, 160, 93, 94,
	96, 95, 8, 164, 165, 141, 163, 166, 15, 167,
	52, 17, 18, 119, 22, 30, 74, 171, 135, 170,
	173, 73, 136, 11, 61, 60, 59, 62, 14, 181,
	55, 72, 50, 182, 1, 184, 54, 35, 51, 186,
	187, 53, 185, 189, 192, 188, 193, 56, 0, 57,
	117, 58, 195, 194, 41, 36, 37, 38, 39, 0,
	0, 0, 40, 42, 0, 0, 0, 15, 0, 0,
	17, 18, 168, 22, 0, 20, 19, 21, 0, 43,
	0, 44, 0, 0, 16, 63, 191, 14, 65, 66,
	67, 61, 60, 59, 62, 0, 169, 0, 45, 46,
	48, 47, 68, 64, 69, 0, 93, 94, 96, 95,
	93, 94, 96, 95, 56, 0, 57, 0, 58, 0,
	0, 41, 36, 37, 38, 39, 0, 0, 0, 40,
	42, 23, 0, 0, 20, 19, 21, 0, 183, 61,
	60, 59, 62, 16, 0, 0, 43, 0, 44, 138,
	139, 140, 63, 0, 142, 65, 66, 67, 93, 94,
	96, 95, 56, 0, 0, 45, 46, 48, 47, 68,
	64, 69, 0, 61, 60, 59, 62, 0, 0, 93,
	94, 96, 95, 138, 139, 140, 93, 94, 96, 95,
	152, 15, 0, 0, 17, 18, 56, 22, 0, 0,
	63, 0, 0, 65, 66, 67, 93, 94, 96, 95,
	15, 14, 132, 17, 18, 0, 22, 68, 64, 69,
	7, 15, 12, 13, 17, 18, 0, 22, 0, 0,
	14, 0, 0, 0, 63, 0, 0, 65, 66, 67,
	159, 14, 93, 94, 96, 95, 0, 0, 0, 0,
	0, 68, 64, 69, 0, 23, 0, 0, 20, 19,
	21, 93, 94, 96, 95, 0, 0, 16, 0, 0,
	0, 0, 0, 0, 23, 0, 0, 20, 19, 21,
	0, 0, 0, 0, 0, 0, 16, 0, 20, 19,
	21, 0, 0, 0, 98, 100, 99, 16, 101, 102,
	103, 0, 0, 0, 0, 0, 0, 104, 105, 108,
	109, 106, 107, 0, 0, 0, 0, 0, 0, 0,
	0, 110, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 111,
}

var yyPact = [...]int16{
	-1000, -1000, 343, -1000, 343, -1000, 73, 25, 68, -1000,
	-1000, -1000, 149, 227, 200, 77, 110, 227, 227, 227,
	227, 227, 227, 227, -1000, 67, -1000, 63, 227, 227,
	28, 338, 399, -1000, 227, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 16,
	-1000, -1000, -1000, -1000, -1000, -1000, 227, 160, 21, -1000,
	-1000, -1000, -1000, 106, 105, -1000, -1000, -1000, -1000, -1000,
	343, 59, 227, 227, -1000, 94, 324, 227, 26, 275,
	282, 338, 338, 15, 338, 338, -3, -1000, -1000, 42,
	41, 227, 227, -1000, -1000, -1000, -1000, 227, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 227, 89, 302, 17, -1000, -1000, 8,
	-1000, -44, -1000, -1000, -1000, -1000, 227, 227, 357, 338,
	90, 227, -1000, 338, 227, -1000, 309, -1000, -1000, -1000,
	-1000, -1000, 227, 227, 141, -1000, -1000, 338, 399, -1000,
	206, -1000, -1000, -1000, -1000, 89, 227, 79, 338, 227,
	5, 54, 338, -1000, 338, 338, -62, -1000, 227, -1000,
	-1000, 338, 227, 254, 227, -1000, 40, -1000, 227, 227,
	141, -8, 202, 227, 54, -1000, 338, 338, -1000, -1000,
	354, 227, 338, -1000, -62, 338,
}

var yyPgo = [...]uint8{
	0, 181, 178, 8, 0, 25, 13, 177, 176, 3,
	7, 174, 129, 172, 170, 10, 2, 4, 163, 9,
	162, 158, 17, 12, 55, 155, 91, 153, 150, 142,
	136, 1, 5, 133, 131, 130, 6, 119,
}

var yyR1 = [...]int8{
	0, 11, 37, 23, 23, 24, 24, 15, 15, 15,
	18, 18, 17, 17, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 21, 21, 20, 20,
	19, 19, 19, 19, 12, 12, 12, 12, 26, 26,
	25, 25, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 36, 31, 31, 30, 30, 32, 32, 8, 8,
	8, 4, 4, 5, 5, 6, 6, 7, 7, 7,
	7, 3, 3, 3, 3, 1, 1, 1, 13, 13,
	2, 2, 22, 22, 28, 28, 27, 27, 9, 14,
	14, 14, 14, 14, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33, 33, 33, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 35, 35, 35, 35, 10,
}

var yyR2 = [...]int8{
	0, 2, 0, 2, 1, 2, 3, 1, 1, 1,
	4, 4, 6, 6, 3, 3, 4, 3, 4, 1,
	2, 2, 1, 2, 1, 2, 0, 1, 1, 2,
	1, 1, 1, 1, 3, 4, 4, 3, 0, 1,
	1, 3, 1, 8, 7, 6, 5, 4, 3, 7,
	5, 1, 0, 1, 1, 2, 2, 2, 1, 2,
	2, 3, 1, 3, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 1, 1, 1, 4, 3,
	3, 2, 1, 3, 2, 3, 1, 3, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -11, -37, -23, -24, -12, -15, 17, -29, -16,
	-17, -18, 19, 20, 27, 7, 83, 10, 11, 75,
	74, 76, 13, 71, -12, -15, 9, -26, 85, 64,
	-25, -4, -5, -6, -33, -7, 35, 36, 37, 38,
	42, 34, 43, 59, 61, 78, 79, 81, 80, -3,
	-13, -2, -28, -1, -8, -14, 27, 29, 31, 6,
	5, 4, 7, 65, 83, 68, 69, 70, 82, 84,
	9, -36, 22, 12, 7, -4, -24, 8, 7, -4,
	-4, -4, -4, -22, -4, -4, -4, 9, 9, -26,
	-26, 18, -35, 44, 45, 47, 46, -34, 35, 37,
	36, 39, 40, 41, 48, 49, 52, 53, 50, 51,
	62, 77, -6, 29, 33, -4, -22, 30, 32, -27,
	-9, -10, 7, 7, 7, -23, 22, 12, -4, -4,
	21, 22, 28, -4, 8, -21, -20, -19, 14, 15,
	16, -3, 12, 18, 72, 9, 9, -4, -5, -6,
	-4, -10, 28, 30, 32, 18, 60, -4, -4, 23,
	-36, -4, -4, -19, -4, -4, -16, -17, 71, 30,
	-9, -4, 23, -4, 22, -31, -30, -32, 25, 26,
	73, -4, -4, 24, -4, -32, -4, -4, -17, -16,
	72, 24, -4, -31, -16, -4,
}

var yyDef = [...]int8{
	2, -2, 0, 1, 0, 4, 0, 38, 0, 7,
	8, 9, 42, 0, 0, 0, 0, 0, 0, 19,
	0, 22, 24, 0, 3, 0, 5, 0, 38, 38,
	39, 40, 62, 64, 0, 66, 94, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 67,
	68, 69, 70, 71, 72, 73, 0, 0, 0, 75,
	76, 77, 58, 0, 0, 89, 90, 91, 92, 93,
	0, 0, 0, 0, 51, 0, 0, 0, 0, 26,
	0, 20, 21, 23, 82, 25, 0, 6, 34, 0,
	0, 0, 0, 121, 122, 123, 124, 0, 107, 108,
	109, 110, 111, 112, 113, 114, 115, 116, 117, 118,
	119, 120, 65, 0, 0, 0, 0, 81, 84, 0,
	86, 0, 125, 59, 60, 37, 0, 0, 0, 48,
	0, 0, 14, 15, 0, 17, 27, 28, 30, 31,
	32, 33, 0, 0, 0, 35, 36, 41, 61, 63,
	0, 79, 74, 80, 85, 0, 0, 0, 47, 0,
	0, 52, 16, 29, 18, 83, 10, 11, 0, 78,
	87, 88, 0, 46, 0, 50, 53, 54, 0, 0,
	0, 0, 45, 0, 52, 55, 56, 57, 12, 13,
	0, 0, 44, 49, 0, 43,
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
	82, 83, 84, 85, 86,
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
//line grammar.y:105
		{
			yyVAL.node = nodeProgram{req: yyDollar[2].nodes, invars: lx.ParamsList()}
			lx.root = yyVAL.node
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:109
		{
			lx = yylex.(*myLexer)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:113
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[2].node)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:114
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:118
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:119
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[2].node)
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:129
		{
			yyVAL.node = nodeIf{cond: yyDollar[2].node, t: yyDollar[4].node}
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:131
		{
			yyVAL.node = nodeIf{cond: yyDollar[2].node, t: yyDollar[4].node}
		}
	case 12:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:135
		{
			yyVAL.node = nodeIf{cond: yyDollar[2].node, t: yyDollar[4].node, e: yyDollar[6].node}
		}
	case 13:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:136
		{
			yyVAL.node = nodeIf{cond: yyDollar[2].node, t: yyDollar[4].node, e: yyDollar[6].node}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:140
		{
			yyVAL.node = yyDollar[2].nodes
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:142
		{
			yyVAL.node = lx.newNodeAssign(yyDollar[1].tok, yyDollar[3].node, false)
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:143
		{
			yyVAL.node = lx.newNodeAssign(yyDollar[2].tok, yyDollar[4].node, true)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:145
		{
			yyVAL.node = lx.mergeNodeClick(nodeClick{element: yyDollar[2].node}, yyDollar[3].node)
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:146
		{
			yyVAL.node = nodeInput{yyDollar[2].node, yyDollar[4].node}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:148
		{
			yyVAL.node = nodeFail{}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:149
		{
			yyVAL.node = nodeFail{yyDollar[2].node}
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:150
		{
			yyVAL.node = nodeAssert{yyDollar[2].node}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:152
		{
			yyVAL.node = nodePrint{nil}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:153
		{
			yyVAL.node = nodePrint{yyDollar[2].nodes}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:156
		{
			yyVAL.node = nodeSlow{m: nil}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:157
		{
			yyVAL.node = nodeSlow{m: yyDollar[1].tok}
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:161
		{
			yyVAL.node = nodeClick{}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:166
		{
			yyVAL.node = yyDollar[1].node
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:167
		{
			yyVAL.node = lx.mergeNodeClick(yyDollar[1].node, yyDollar[2].node)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:171
		{
			yyVAL.node = nodeClick{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:172
		{
			yyVAL.node = nodeClick{right: true}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:173
		{
			yyVAL.node = nodeClick{middle: true}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:174
		{
			yyVAL.node = nodeClick{count: yyDollar[1].node}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:178
		{
			yyVAL.node = nodeReturn{what: yyDollar[2].nodes}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:179
		{
			yyVAL.node = nodeReturn{what: yyDollar[3].nodes, last: true}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:180
		{
			yyVAL.node = nodeReturn{what: yyDollar[3].nodes, distinct: true}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:181
		{
			yyVAL.node = yyDollar[1].nodeWithBody.appendBody(yyDollar[3].nodes)
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:185
		{
			yyVAL.nodes = Nodes{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:190
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:191
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[3].node)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:194
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(nil, nil, nil, nil)
		}
	case 43:
		yyDollar = yyS[yypt-8 : yypt+1]
//line grammar.y:197
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(yyDollar[2].tok, yyDollar[4].node, yyDollar[6].node, yyDollar[8].node)
		}
	case 44:
		yyDollar = yyS[yypt-7 : yypt+1]
//line grammar.y:199
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(nil, yyDollar[3].node, yyDollar[5].node, yyDollar[7].node)
		}
	case 45:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:202
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(yyDollar[2].tok, yyDollar[4].node, yyDollar[6].node, nil)
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:204
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(nil, yyDollar[3].node, yyDollar[5].node, nil)
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:206
		{
			yyVAL.nodeWithBody = lx.newNodeForArray(yyDollar[2].tok, yyDollar[4].node)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:207
		{
			yyVAL.nodeWithBody = lx.newNodeForArray(nil, yyDollar[3].node)
		}
	case 49:
		yyDollar = yyS[yypt-7 : yypt+1]
//line grammar.y:210
		{
			yyVAL.nodeWithBody = lx.newNodeSelect(yyDollar[4].tok, yyDollar[2].node, yyDollar[6].node, yyDollar[7].nodeWithBody)
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:212
		{
			yyVAL.nodeWithBody = lx.newNodeSelect(nil, yyDollar[2].node, yyDollar[4].node, yyDollar[5].nodeWithBody)
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:220
		{
			yyVAL.nodeWithBody = nodeSelect{}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:221
		{
			yyVAL.nodeWithBody = yyDollar[1].nodeWithBody
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:225
		{
			yyVAL.nodeWithBody = yyDollar[1].nodeWithBody
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:226
		{
			yyVAL.nodeWithBody = yyDollar[1].nodeWithBody.(nodeSelect).mergeOptions(yyDollar[2].nodeWithBody)
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:230
		{
			yyVAL.nodeWithBody = nodeSelect{where: []Node{yyDollar[2].node}}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:231
		{
			yyVAL.nodeWithBody = nodeSelect{limit: yyDollar[2].node}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:235
		{
			yyVAL.node = lx.newNodeVariable(yyDollar[1].tok, false, true, false)
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:236
		{
			yyVAL.node = lx.newNodeVariable(yyDollar[2].tok, true, false, false)
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:237
		{
			yyVAL.node = lx.newNodeVariable(yyDollar[2].tok, false, true, true)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:245
		{
			yyVAL.node = lx.newNodeOpe2Bool(yyDollar[1].node, yyDollar[2].tok, yyDollar[3].node)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:250
		{
			yyVAL.node = lx.newNodeOpe2(yyDollar[1].node, yyDollar[2].tok, yyDollar[3].node)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:255
		{
			yyVAL.node = lx.newNodeOpe1(yyDollar[1].tok, yyDollar[2].node)
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:262
		{
			yyVAL.node = yyDollar[1].node
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:263
		{
			yyVAL.node = yyDollar[1].nodemap
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:270
		{
			yyVAL.node = yyDollar[2].node
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:274
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:275
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:276
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:280
		{
			yyVAL.node = nodeArrayAccess{a: yyDollar[1].node, i: yyDollar[3].node}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:281
		{
			yyVAL.node = nodeMapAccess{m: yyDollar[1].node, k: yyDollar[3].node}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:285
		{
			yyVAL.node = yyDollar[2].nodes
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:286
		{
			yyVAL.node = Nodes(nil)
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:290
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:291
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[3].node)
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:295
		{
			yyVAL.nodemap = lx.newNodeMap(nil, nil)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:296
		{
			yyVAL.nodemap = yyDollar[2].nodemap
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:300
		{
			yyVAL.nodemap = lx.newNodeMap(nil, yyDollar[1].node)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:301
		{
			yyVAL.nodemap = lx.newNodeMap(yyDollar[1].nodemap, yyDollar[3].node)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:305
		{
			yyVAL.node = lx.newNodeKeyValue(yyDollar[1].node, yyDollar[3].node)
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:309
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:310
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:311
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:312
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:313
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:363
		{
			yyVAL.node = lx.newNodeKey(yyDollar[1].tok)
		}
	}
	goto yystack /* stack new state and value */
}
