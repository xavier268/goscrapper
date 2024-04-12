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
const IF = 57415
const THEN = 57416
const ELSE = 57417
const ASSERT = 57418
const FAIL = 57419
const RANGE = 57420

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
	"IF",
	"THEN",
	"ELSE",
	"ASSERT",
	"FAIL",
	"RANGE",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:346

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 343

var yyAct = [...]uint8{
	103, 157, 9, 106, 10, 159, 61, 42, 28, 121,
	107, 30, 162, 101, 65, 29, 139, 68, 69, 70,
	71, 3, 74, 75, 156, 138, 54, 53, 52, 55,
	6, 136, 98, 136, 15, 24, 99, 16, 17, 137,
	20, 78, 21, 97, 108, 135, 7, 73, 12, 13,
	100, 49, 77, 50, 102, 51, 14, 76, 38, 33,
	34, 35, 36, 113, 114, 60, 37, 39, 118, 15,
	25, 104, 16, 17, 108, 20, 125, 21, 112, 129,
	160, 161, 110, 40, 67, 41, 127, 4, 64, 56,
	111, 14, 57, 58, 59, 130, 131, 24, 173, 132,
	22, 109, 66, 19, 18, 158, 80, 81, 83, 82,
	133, 5, 140, 141, 120, 119, 23, 144, 80, 81,
	83, 82, 143, 80, 81, 83, 82, 146, 125, 2,
	145, 147, 72, 148, 172, 149, 64, 151, 19, 18,
	153, 63, 152, 155, 54, 53, 52, 55, 79, 84,
	163, 128, 31, 62, 8, 164, 45, 166, 105, 26,
	27, 168, 169, 11, 167, 171, 174, 170, 175, 49,
	48, 50, 134, 51, 177, 176, 38, 33, 34, 35,
	36, 43, 150, 1, 37, 39, 47, 32, 80, 81,
	83, 82, 54, 53, 52, 55, 80, 81, 83, 82,
	44, 40, 46, 41, 122, 123, 124, 56, 0, 0,
	57, 58, 59, 0, 0, 0, 15, 49, 0, 16,
	17, 0, 20, 15, 21, 0, 16, 17, 0, 20,
	0, 21, 0, 0, 80, 81, 83, 82, 14, 117,
	15, 0, 0, 16, 17, 14, 20, 0, 21, 160,
	161, 54, 53, 52, 55, 56, 0, 165, 57, 58,
	59, 0, 14, 122, 123, 124, 0, 0, 80, 81,
	83, 82, 115, 116, 0, 0, 49, 80, 81, 83,
	82, 154, 22, 0, 0, 19, 18, 0, 0, 22,
	0, 0, 19, 18, 0, 80, 81, 83, 82, 0,
	0, 126, 80, 81, 83, 82, 142, 85, 86, 19,
	18, 87, 88, 89, 56, 0, 0, 57, 58, 59,
	90, 91, 94, 95, 92, 93, 0, 80, 81, 83,
	82, 0, 0, 0, 96, 80, 81, 83, 82, 80,
	81, 83, 82,
}

var yyPact = [...]int16{
	-1000, -1000, 27, -1000, 27, -1000, 61, 140, 56, -1000,
	-1000, -1000, 129, 140, 216, 76, 140, 140, 140, 140,
	33, 140, 140, -1000, 48, -1000, 43, 21, 293, 270,
	-1000, 140, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 1, -1000, -1000, -1000, -1000, -1000, -1000, 140,
	22, 37, -1000, -1000, -1000, -1000, 94, -1000, -1000, -1000,
	27, 66, 140, 140, -1000, 249, 209, 140, 188, 289,
	293, 293, 140, -1000, 293, 77, -1000, -1000, 140, 140,
	-1000, -1000, -1000, -1000, 140, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 140, 67,
	142, 13, -1000, 293, -1000, 5, -1000, -46, -1000, -1000,
	-1000, 140, 140, 281, 293, 81, 140, -1000, 293, -1000,
	247, -1000, -1000, -1000, -1000, -1000, 140, 11, 62, 293,
	270, -1000, 150, -1000, -1000, -1000, 140, -1000, 67, 140,
	256, 293, 140, 0, 222, -1000, 293, -63, -1000, 140,
	-1000, 293, -1000, 293, 140, 231, 140, -1000, 53, -1000,
	140, 140, 62, 60, 72, 140, 222, -1000, 293, 293,
	-1000, -1000, 233, 140, 293, -1000, -63, 293,
}

var yyPgo = [...]uint8{
	0, 202, 200, 7, 0, 15, 11, 187, 186, 3,
	10, 183, 111, 181, 170, 30, 2, 4, 163, 13,
	21, 87, 160, 159, 158, 156, 154, 152, 149, 148,
	6, 132, 129, 115, 114, 9, 1, 105, 5,
}

var yyR1 = [...]int8{
	0, 11, 32, 20, 20, 21, 21, 15, 15, 15,
	18, 18, 17, 17, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 31, 31, 33, 33, 34, 34,
	35, 35, 35, 35, 12, 12, 23, 23, 22, 22,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 30,
	36, 36, 37, 37, 38, 38, 8, 8, 4, 4,
	5, 5, 6, 6, 7, 7, 7, 7, 3, 3,
	3, 3, 1, 1, 1, 13, 13, 2, 2, 19,
	19, 25, 25, 24, 24, 9, 14, 14, 14, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	29, 29, 29, 29, 10,
}

var yyR2 = [...]int8{
	0, 2, 0, 2, 1, 2, 3, 1, 1, 1,
	4, 4, 6, 6, 3, 3, 3, 4, 1, 2,
	2, 3, 1, 2, 0, 1, 0, 1, 1, 2,
	1, 1, 1, 1, 3, 3, 0, 1, 1, 3,
	1, 8, 7, 6, 5, 4, 3, 7, 5, 1,
	0, 1, 1, 2, 2, 2, 1, 2, 3, 1,
	3, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 1, 1, 4, 3, 3, 2, 1,
	3, 2, 3, 1, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -11, -32, -20, -21, -12, -15, 19, -26, -16,
	-17, -18, 21, 22, 29, 7, 10, 11, 77, 76,
	13, 15, 73, -12, -15, 9, -23, -22, -4, -5,
	-6, -27, -7, 37, 38, 39, 40, 44, 36, 45,
	61, 63, -3, -13, -2, -25, -1, -8, -14, 29,
	31, 33, 6, 5, 4, 7, 67, 70, 71, 72,
	9, -30, 24, 12, 7, -4, -21, 8, -4, -4,
	-4, -4, -31, 14, -4, -4, 9, 9, 20, -29,
	46, 47, 49, 48, -28, 37, 38, 41, 42, 43,
	50, 51, 54, 55, 52, 53, 64, -6, 31, 35,
	-4, -19, 32, -4, 34, -24, -9, -10, 7, 7,
	-20, 24, 12, -4, -4, 23, 24, 30, -4, -33,
	-34, -35, 16, 17, 18, -3, 12, -19, 74, -4,
	-5, -6, -4, -10, 30, 32, 20, 34, 20, 62,
	-4, -4, 25, -30, -4, -35, -4, -16, -17, 73,
	32, -4, -9, -4, 25, -4, 24, -36, -37, -38,
	27, 28, 75, -4, -4, 26, -4, -38, -4, -4,
	-17, -16, 74, 26, -4, -36, -16, -4,
}

var yyDef = [...]int8{
	2, -2, 0, 1, 0, 4, 0, 36, 0, 7,
	8, 9, 40, 0, 0, 0, 0, 0, 18, 0,
	24, 22, 0, 3, 0, 5, 0, 37, 38, 59,
	61, 0, 63, 89, 90, 91, 92, 93, 94, 95,
	96, 97, 64, 65, 66, 67, 68, 69, 70, 0,
	0, 0, 72, 73, 74, 56, 0, 86, 87, 88,
	0, 0, 0, 0, 49, 0, 0, 0, 26, 0,
	19, 20, 0, 25, 23, 0, 6, 34, 0, 0,
	110, 111, 112, 113, 0, 98, 99, 100, 101, 102,
	103, 104, 105, 106, 107, 108, 109, 62, 0, 0,
	0, 0, 78, 79, 81, 0, 83, 0, 114, 57,
	35, 0, 0, 0, 46, 0, 0, 14, 15, 16,
	27, 28, 30, 31, 32, 33, 0, 21, 0, 39,
	58, 60, 0, 76, 71, 77, 0, 82, 0, 0,
	0, 45, 0, 0, 50, 29, 17, 10, 11, 0,
	75, 80, 84, 85, 0, 44, 0, 48, 51, 52,
	0, 0, 0, 0, 43, 0, 50, 53, 54, 55,
	12, 13, 0, 0, 42, 47, 0, 41,
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
	72, 73, 74, 75, 76, 77, 78,
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
//line grammar.y:98
		{
			yyVAL.node = nodeProgram{req: yyDollar[2].nodes, invars: lx.ParamsList()}
			lx.root = yyVAL.node
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:102
		{
			lx = yylex.(*myLexer)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:106
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[2].node)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:107
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:111
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:112
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[2].node)
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:122
		{
			yyVAL.node = nodeIf{cond: yyDollar[2].node, t: yyDollar[4].node}
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:124
		{
			yyVAL.node = nodeIf{cond: yyDollar[2].node, t: yyDollar[4].node}
		}
	case 12:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:128
		{
			yyVAL.node = nodeIf{cond: yyDollar[2].node, t: yyDollar[4].node, e: yyDollar[6].node}
		}
	case 13:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:129
		{
			yyVAL.node = nodeIf{cond: yyDollar[2].node, t: yyDollar[4].node, e: yyDollar[6].node}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:133
		{
			yyVAL.node = yyDollar[2].nodes
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:134
		{
			yyVAL.node = lx.newNodeAssign(yyDollar[1].tok, yyDollar[3].node)
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:135
		{ /*todo*/
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:136
		{ /*todo*/
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:138
		{
			yyVAL.node = nodeFail{}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:139
		{
			yyVAL.node = nodeFail{yyDollar[2].node}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:140
		{
			yyVAL.node = nodeAssert{yyDollar[2].node}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:143
		{
			yyVAL.node = nodePrint{nodes: yyDollar[3].nodes, raw: (yyDollar[2].tok.c == RAW)}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:146
		{
			yyVAL.node = nodeSlow{m: nil}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:147
		{
			yyVAL.node = nodeSlow{m: yyDollar[1].tok}
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:151
		{
			yyVAL.tok = tok{}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:152
		{
			yyVAL.tok = yyDollar[1].tok
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:156
		{ /*todo*/
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:157
		{ /*todo*/
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:161
		{ /*todo*/
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:162
		{ /*todo*/
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:166
		{ /*todo*/
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:167
		{ /*todo*/
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:168
		{ /*todo*/
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:169
		{ /*todo*/
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:173
		{
			yyVAL.node = nodeReturn{yyDollar[2].nodes}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:174
		{
			yyVAL.node = yyDollar[1].nodeWithBody.appendBody(yyDollar[3].nodes)
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:178
		{
			yyVAL.nodes = Nodes{}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:183
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:184
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[3].node)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:187
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(nil, nil, nil, nil)
		}
	case 41:
		yyDollar = yyS[yypt-8 : yypt+1]
//line grammar.y:190
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(yyDollar[2].tok, yyDollar[4].node, yyDollar[6].node, yyDollar[8].node)
		}
	case 42:
		yyDollar = yyS[yypt-7 : yypt+1]
//line grammar.y:192
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(nil, yyDollar[3].node, yyDollar[5].node, yyDollar[7].node)
		}
	case 43:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:195
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(yyDollar[2].tok, yyDollar[4].node, yyDollar[6].node, nil)
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:197
		{
			yyVAL.nodeWithBody = lx.newNodeForLoop(nil, yyDollar[3].node, yyDollar[5].node, nil)
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:199
		{
			yyVAL.nodeWithBody = lx.newNodeForArray(yyDollar[2].tok, yyDollar[4].node)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:200
		{
			yyVAL.nodeWithBody = lx.newNodeForArray(nil, yyDollar[3].node)
		}
	case 47:
		yyDollar = yyS[yypt-7 : yypt+1]
//line grammar.y:203
		{ /*todo*/
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:205
		{ /*todo*/
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:213
		{ /*todo*/
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:214
		{ /*todo*/
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:218
		{ /*todo*/
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:219
		{ /*todo*/
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:223
		{ /*todo*/
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:224
		{ /*todo*/
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:228
		{
			yyVAL.node = lx.newNodeVariable(yyDollar[1].tok, false, true)
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:229
		{
			yyVAL.node = lx.newNodeVariable(yyDollar[2].tok, true, false)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:237
		{
			yyVAL.node = lx.newNodeOpe2Bool(yyDollar[1].node, yyDollar[2].tok, yyDollar[3].node)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:242
		{
			yyVAL.node = lx.newNodeOpe2(yyDollar[1].node, yyDollar[2].tok, yyDollar[3].node)
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:247
		{
			yyVAL.node = lx.newNodeOpe1(yyDollar[1].tok, yyDollar[2].node)
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:254
		{
			yyVAL.node = yyDollar[1].node
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:255
		{
			yyVAL.node = yyDollar[1].nodemap
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:262
		{
			yyVAL.node = yyDollar[2].node
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:266
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:267
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:268
		{
			yyVAL.node = lx.newNodeLitteral(yyDollar[1].tok)
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:272
		{
			yyVAL.node = nodeArrayAccess{a: yyDollar[1].node, i: yyDollar[3].node}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:273
		{
			yyVAL.node = nodeMapAccess{m: yyDollar[1].node, k: yyDollar[3].node}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:277
		{
			yyVAL.node = yyDollar[2].nodes
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:278
		{
			yyVAL.node = Nodes(nil)
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:282
		{
			yyVAL.nodes = Nodes{yyDollar[1].node}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:283
		{
			yyVAL.nodes = append(yyDollar[1].nodes, yyDollar[3].node)
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:287
		{
			yyVAL.nodemap = lx.newNodeMap(nil, nil)
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:288
		{
			yyVAL.nodemap = yyDollar[2].nodemap
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:292
		{
			yyVAL.nodemap = lx.newNodeMap(nil, yyDollar[1].node)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:293
		{
			yyVAL.nodemap = lx.newNodeMap(yyDollar[1].nodemap, yyDollar[3].node)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:297
		{
			yyVAL.node = lx.newNodeKeyValue(yyDollar[1].node, yyDollar[3].node)
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:301
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:302
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:303
		{
			yyVAL.node = nodeOpe0(yyDollar[1].tok)
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:344
		{
			yyVAL.node = lx.newNodeKey(yyDollar[1].tok)
		}
	}
	goto yystack /* stack new state and value */
}
