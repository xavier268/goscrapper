// Code generated by goyacc -o parser.go grammar.y. DO NOT EDIT.

//line grammar.y:1

package parser

import __yyfmt__ "fmt"

//line grammar.y:3

import (
	"fmt"
)

// each object has a value and a type.
type value struct {
	v string // a string in go that produce the value of the object
	t string // a string representing the gotype of the object
	c int    // the code returned by lexer is stored here. Always set by the lexer, even for variables (set as IDENTIFIER). A valid go type, without spaces.
}

var _zv = value{} // zero value

var lx *myLexer // shorthand for lx

// options for SELECT ONE, ANY, ALL
type selopt struct {
	from  value   // can be *rod.Page or *rod.Element
	css   value   // css selector
	loopv string  // loop variable identifier
	where []value // list of where conditions, applied on loopv
	limit value
	cases []casopt
}

// cases for select ANY
type casopt struct {
	def bool
	e1  value
	e2  value
}

//line grammar.y:64
type yySymType struct {
	yys     int
	value   value
	list    []string
	values  []value
	mvalue  map[string]value
	selopt  selopt
	casopt  casopt
	casopts []casopt
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
const FORMAT = 57378
const NOW = 57379
const TEXT = 57380
const HREF = 57381
const ATTRIBUTE = 57382
const LEFT = 57383
const RIGHT = 57384
const MIDDLE = 57385
const INTTYPE = 57386
const BOOLTYPE = 57387
const STRINGTYPE = 57388
const BINTYPE = 57389
const FOR = 57390
const RETURN = 57391
const WAITFOR = 57392
const OPTIONS = 57393
const IGNORE = 57394
const HEADLESS = 57395
const TIMEOUT = 57396
const TRUE = 57397
const FALSE = 57398
const EVENT = 57399
const LIKE = 57400
const IN = 57401
const WHILE = 57402
const BOOL = 57403
const AT = 57404
const IDENTIFIER = 57405
const IGNOREID = 57406
const STRING = 57407
const NUMBER = 57408
const NAMESPACESEPARATOR = 57409
const SELECT = 57410
const ALL = 57411
const ANY = 57412
const ONE = 57413
const AS = 57414
const FROM = 57415
const WHERE = 57416
const LIMIT = 57417
const DISTINCT = 57418
const SORT = 57419
const ASC = 57420
const DESC = 57421
const DEFAULT = 57422
const CASE = 57423
const CLICK = 57424
const INPUT = 57425
const DOCUMENT = 57426
const PAGE = 57427
const CONTAINS = 57428
const PRINT = 57429
const SLOW = 57430

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
	"FORMAT",
	"NOW",
	"TEXT",
	"HREF",
	"ATTRIBUTE",
	"LEFT",
	"RIGHT",
	"MIDDLE",
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
	"TRUE",
	"FALSE",
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
	"ALL",
	"ANY",
	"ONE",
	"AS",
	"FROM",
	"WHERE",
	"LIMIT",
	"DISTINCT",
	"SORT",
	"ASC",
	"DESC",
	"DEFAULT",
	"CASE",
	"CLICK",
	"INPUT",
	"DOCUMENT",
	"PAGE",
	"CONTAINS",
	"PRINT",
	"SLOW",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:309

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 756

var yyAct = [...]uint8{
	30, 145, 133, 53, 52, 98, 147, 146, 154, 155,
	134, 35, 148, 34, 31, 36, 99, 125, 103, 49,
	50, 100, 88, 51, 27, 22, 6, 94, 41, 32,
	62, 106, 29, 105, 119, 84, 86, 67, 68, 69,
	64, 66, 65, 73, 71, 72, 70, 74, 75, 20,
	13, 60, 40, 95, 37, 153, 38, 39, 137, 76,
	77, 78, 96, 83, 15, 121, 108, 122, 82, 21,
	81, 120, 107, 80, 89, 90, 91, 58, 101, 109,
	61, 59, 102, 18, 19, 12, 14, 11, 16, 17,
	110, 111, 112, 113, 114, 115, 25, 9, 24, 10,
	3, 54, 56, 55, 57, 23, 92, 124, 2, 5,
	126, 4, 28, 8, 1, 7, 150, 131, 132, 79,
	144, 87, 97, 85, 26, 33, 136, 135, 138, 139,
	140, 141, 63, 0, 142, 143, 0, 0, 149, 0,
	0, 0, 0, 0, 0, 0, 151, 152, 0, 0,
	0, 0, 0, 0, 157, 158, 159, 160, 67, 68,
	69, 64, 66, 65, 73, 71, 72, 70, 74, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	76, 77, 78, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 80, 67, 68, 69, 64, 66,
	65, 73, 71, 72, 70, 74, 75, 0, 162, 0,
	0, 0, 0, 0, 0, 0, 0, 76, 77, 78,
	0, 0, 0, 116, 118, 117, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	79, 0, 67, 68, 69, 64, 66, 65, 73, 71,
	72, 70, 74, 75, 0, 161, 0, 0, 0, 0,
	0, 0, 0, 0, 76, 77, 78, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 79, 80, 67,
	68, 69, 64, 66, 65, 73, 71, 72, 70, 74,
	75, 156, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 76, 77, 78, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 0, 79, 0, 67, 68, 69, 64,
	66, 65, 73, 71, 72, 70, 74, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 76, 77,
	78, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 79, 80, 67, 68, 69, 64, 66, 65, 73,
	71, 72, 70, 74, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 76, 77, 78, 0, 0,
	0, 0, 0, 0, 134, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 0, 79, 0,
	0, 67, 68, 69, 64, 66, 65, 73, 71, 72,
	70, 74, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 130, 76, 77, 78, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 79, 0, 80, 67, 68,
	69, 64, 66, 65, 73, 71, 72, 70, 74, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	76, 77, 78, 0, 0, 0, 0, 0, 0, 0,
	129, 0, 0, 0, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 79, 0, 0, 67, 68, 69, 64,
	66, 65, 73, 71, 72, 70, 74, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 128, 76, 77,
	78, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	79, 0, 80, 67, 68, 69, 64, 66, 65, 73,
	71, 72, 70, 74, 75, 0, 0, 0, 0, 0,
	123, 0, 0, 0, 0, 76, 77, 78, 0, 0,
	0, 0, 0, 0, 0, 127, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 0, 79, 0,
	67, 68, 69, 64, 66, 65, 73, 71, 72, 70,
	74, 75, 0, 0, 0, 0, 0, 0, 0, 104,
	0, 0, 76, 77, 78, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 79, 80, 67, 68, 69,
	64, 66, 65, 73, 71, 72, 70, 74, 75, 0,
	0, 0, 0, 0, 42, 43, 0, 0, 0, 76,
	77, 78, 0, 0, 0, 0, 35, 0, 34, 0,
	36, 0, 0, 80, 46, 0, 0, 0, 0, 0,
	44, 45, 79, 41, 48, 0, 0, 0, 0, 0,
	0, 0, 93, 67, 68, 69, 64, 66, 65, 73,
	71, 72, 70, 74, 75, 0, 0, 40, 0, 37,
	0, 38, 39, 0, 0, 76, 77, 78, 0, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 80,
	0, 47, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 79,
}

var yyPact = [...]int16{
	-1000, -1000, -36, -36, 1, -1000, -38, 1, -1000, -1000,
	1, -1000, -1000, -39, 1, 2, 626, -1000, 626, 626,
	-40, -69, 57, -1000, -1000, -1000, 32, -1000, -1000, 626,
	669, -1000, 50, -9, 626, 626, -41, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 33,
	613, -32, 626, -1000, -1000, -1000, -1000, -1000, 57, -47,
	-42, -1000, 669, 626, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 626, -45, 50, 576, 12, 669, 47, 63, 626,
	626, 626, 626, 626, 626, 154, 13, 46, -1000, 51,
	-1000, -1000, 529, -1000, -1000, -1000, 626, -1000, -46, 626,
	492, 444, 407, 669, 359, 669, 626, 626, -62, -1000,
	-1000, -47, 57, -1000, 669, 42, 669, 626, 626, 626,
	626, 322, 322, -74, -51, -1000, -1000, 626, 669, 669,
	669, 669, -1000, -1000, -74, -1000, 626, 39, -1000, 669,
	-66, -1000, 275, 626, 626, 626, 626, 238, 669, 669,
	191, -1000, -1000,
}

var yyPgo = [...]uint8{
	0, 0, 14, 29, 132, 125, 3, 124, 123, 122,
	5, 2, 121, 1, 120, 116, 114, 108, 100, 111,
	97, 109, 99, 87, 85, 86, 80,
}

var yyR1 = [...]int8{
	0, 16, 16, 17, 19, 18, 18, 21, 6, 6,
	6, 6, 6, 6, 10, 9, 9, 20, 20, 22,
	22, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 23, 23, 26, 7, 7, 25,
	25, 25, 25, 14, 14, 13, 13, 11, 15, 15,
	15, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 5, 5,
	5, 5, 5, 5, 5, 1, 1, 2, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 8,
	8, 12, 12, 12, 12,
}

var yyR2 = [...]int8{
	0, 4, 3, 0, 0, 2, 1, 3, 1, 1,
	1, 1, 3, 3, 3, 1, 3, 2, 1, 1,
	2, 3, 2, 1, 2, 4, 4, 4, 4, 6,
	6, 6, 4, 6, 2, 3, 0, 1, 3, 4,
	7, 6, 6, 1, 2, 5, 4, 2, 0, 3,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 2, 3,
	4, 3, 3, 3, 1, 1, 1, 1, 1, 1,
	3, 3, 5, 1, 3,
}

var yyChk = [...]int16{
	-1000, -16, -17, -18, -19, -21, 62, -19, -21, -20,
	-22, -23, -24, 49, -25, 63, 87, 88, 82, 83,
	48, 68, 63, -20, -23, -24, -7, 63, -20, 30,
	-1, -2, -3, -5, 22, 20, 24, 63, 65, 66,
	61, 37, 8, 9, 34, 35, 28, 85, 38, -1,
	-1, 63, 73, -6, 44, 46, 45, 47, 20, 24,
	19, -26, -1, -4, 7, 9, 8, 4, 5, 6,
	13, 11, 12, 10, 14, 15, 26, 27, 28, 86,
	40, 20, 18, -3, -1, -8, -1, -12, 63, 41,
	42, 43, 73, 59, 59, -1, -6, -9, -10, 63,
	63, -2, -1, 63, 23, 21, 19, 25, 19, 16,
	-1, -1, -1, -1, -1, -1, 69, 71, 70, 21,
	25, 19, 16, 21, -1, 63, -1, 73, 73, 73,
	73, -1, -1, -11, 72, -10, -6, 16, -1, -1,
	-1, -1, -11, -11, -14, -13, 81, 80, 63, -1,
	-15, -13, -1, 16, 74, 75, 16, -1, -1, -1,
	-1, 17, 17,
}

var yyDef = [...]int8{
	3, -2, 4, 4, 0, 6, 0, 0, 5, 2,
	0, 18, 19, 0, 0, 0, 0, 23, 0, 0,
	0, 0, 0, 1, 17, 20, 34, 37, 36, 0,
	22, 75, 77, 0, 0, 0, 0, 84, 85, 86,
	87, 88, 68, 69, 70, 71, 72, 73, 74, 24,
	0, 0, 0, 7, 8, 9, 10, 11, 0, 0,
	0, 35, 21, 0, 51, 52, 53, 54, 55, 56,
	57, 58, 59, 60, 61, 62, 63, 64, 65, 66,
	67, 0, 0, 78, 0, 0, 89, 0, 93, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 15, 0,
	38, 76, 0, 81, 79, 82, 0, 83, 0, 0,
	25, 26, 27, 28, 32, 39, 0, 0, 0, 12,
	13, 0, 0, 80, 90, 94, 91, 0, 0, 0,
	0, 0, 0, 0, 0, 16, 14, 0, 29, 30,
	31, 33, 48, 41, 42, 43, 0, 0, 47, 92,
	40, 44, 0, 0, 0, 0, 0, 0, 49, 50,
	0, 46, 45,
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
	82, 83, 84, 85, 86, 87, 88,
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
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:113
		{
			lx.finalize()
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:114
		{
			lx.finalize()
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:118
		{
			lx = yylex.(*myLexer)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:122
		{
			lx.incOut()
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:132
		{
			lx.declInputParam(yyDollar[2].value.v, yyDollar[3].value.v)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:136
		{
			yyVAL.value.v = "int"
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:137
		{
			yyVAL.value.v = "string"
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:138
		{
			yyVAL.value.v = "bool"
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:139
		{
			yyVAL.value.v = "[]byte"
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:140
		{
			yyVAL.value.v = "[]" + yyDollar[2].value.v
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:141
		{
			yyVAL.value.v = lx.objectType(yyDollar[2].values)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:145
		{
			yyVAL.value = value{v: yyDollar[1].value.v, t: yyDollar[3].value.v}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:148
		{
			yyVAL.values = []value{yyDollar[1].value}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:149
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].value)
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:155
		{
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:156
		{
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:160
		{ /* todo */
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:161
		{ /* todo */
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:165
		{
			lx.vSetVar(yyDollar[1].value.v, yyDollar[3].value)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:166
		{
			lx.addImport("fmt")
			lx.addLines(fmt.Sprintf("fmt.Println(%s)", yyDollar[2].value.v))
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:167
		{
			lx.addImport("rt")
			lx.addLines("rt.Slow(_ctx)")
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:169
		{
			lx.Click(yyDollar[2].value, _zv, _zv)
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:170
		{
			lx.Click(yyDollar[2].value, yyDollar[3].value, yyDollar[4].value)
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:171
		{
			lx.Click(yyDollar[2].value, yyDollar[3].value, yyDollar[4].value)
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:172
		{
			lx.Click(yyDollar[2].value, yyDollar[3].value, yyDollar[4].value)
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:174
		{
			lx.ClickFrom(yyDollar[2].value, _zv, _zv, yyDollar[4].value)
		}
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:175
		{
			lx.ClickFrom(yyDollar[2].value, yyDollar[3].value, yyDollar[4].value, yyDollar[6].value)
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:176
		{
			lx.ClickFrom(yyDollar[2].value, yyDollar[3].value, yyDollar[4].value, yyDollar[6].value)
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:177
		{
			lx.ClickFrom(yyDollar[2].value, yyDollar[3].value, yyDollar[4].value, yyDollar[6].value)
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:179
		{
			lx.input(yyDollar[2].value, yyDollar[4].value)
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:180
		{
			lx.inputFrom(yyDollar[2].value, yyDollar[4].value, yyDollar[6].value)
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:184
		{
			lx.declOutputParams(yyDollar[2].list)
			lx.saveOut()
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:185
		{ /* */
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:189
		{
			lx.addLines("}")
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:192
		{
			yyVAL.list = append(yyVAL.list, yyDollar[1].value.v)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:193
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].value.v)
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:197
		{
			lx.forNameInExpression(yyDollar[2].value.v, yyDollar[4].value)
		}
	case 40:
		yyDollar = yyS[yypt-7 : yypt+1]
//line grammar.y:200
		{
			opt := yyDollar[7].selopt
			opt.from = yyDollar[3].value
			opt.css = yyDollar[5].value
			opt.loopv = yyDollar[6].value.v
			lx.selectAll(opt)
		}
	case 41:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:201
		{
			lx.selectOne(yyDollar[3].value, yyDollar[5].value, yyDollar[6].value)
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
//line grammar.y:202
		{
			lx.selectAny(yyDollar[3].value, yyDollar[5].value, yyDollar[6].casopts)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:206
		{
			yyVAL.casopts = []casopt{yyDollar[1].casopt}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:207
		{
			yyVAL.casopts = append(yyDollar[1].casopts, yyDollar[2].casopt)
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:215
		{
			yyVAL.casopt = casopt{e1: yyDollar[2].value, e2: yyDollar[4].value}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:216
		{
			yyVAL.casopt = casopt{def: true, e2: yyDollar[3].value}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:220
		{
			yyVAL.value = yyDollar[2].value
			if typ, ok := lx.vars[yyDollar[2].value.v]; ok {
				lx.errorf("variable %s was already declared (type : %s), cannot be redeclared as SELECT loop variable", yyDollar[2].value.v, typ)
			}
			lx.vars[yyDollar[2].value.v] = "*rod.Element"
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:230
		{
			yyVAL.selopt = selopt{}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:231
		{
			yyVAL.selopt = yyDollar[1].selopt
			yyVAL.selopt.where = append(yyVAL.selopt.where, yyDollar[3].value)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:232
		{
			yyVAL.selopt = yyDollar[1].selopt
			yyVAL.selopt.limit = yyDollar[3].value
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:241
		{
			yyVAL.value = yyDollar[1].value
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:242
		{
			yyVAL.value = yyDollar[1].value
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:243
		{
			yyVAL.value = yyDollar[1].value
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:244
		{
			yyVAL.value = yyDollar[1].value
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:245
		{
			yyVAL.value = yyDollar[1].value
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:246
		{
			yyVAL.value = yyDollar[1].value
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:247
		{
			yyVAL.value = yyDollar[1].value
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:248
		{
			yyVAL.value = yyDollar[1].value
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:249
		{
			yyVAL.value = yyDollar[1].value
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:250
		{
			yyVAL.value = yyDollar[1].value
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:251
		{
			yyVAL.value = yyDollar[1].value
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:252
		{
			yyVAL.value = yyDollar[1].value
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:253
		{
			yyVAL.value = yyDollar[1].value
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:254
		{
			yyVAL.value = yyDollar[1].value
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:255
		{
			yyVAL.value = yyDollar[1].value
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:256
		{
			yyVAL.value = yyDollar[1].value
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:257
		{
			yyVAL.value = yyDollar[1].value
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:261
		{
			yyVAL.value = yyDollar[1].value
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:262
		{
			yyVAL.value = yyDollar[1].value
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:263
		{
			yyVAL.value = yyDollar[1].value
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:264
		{
			yyVAL.value = yyDollar[1].value
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:265
		{
			yyVAL.value = yyDollar[1].value
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:266
		{
			yyVAL.value = yyDollar[1].value
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:267
		{
			yyVAL.value = yyDollar[1].value
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:272
		{
			yyVAL.value = yyDollar[1].value
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:273
		{
			yyVAL.value = lx.vOpe2(yyDollar[2].value.c, yyDollar[1].value, yyDollar[3].value)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:277
		{
			yyVAL.value = yyDollar[1].value
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:278
		{
			yyVAL.value = lx.vOpe1(yyDollar[1].value.c, yyDollar[2].value)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:282
		{
			yyVAL.value = lx.vParen(yyDollar[2].value)
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:283
		{
			yyVAL.value = lx.vGetElementOf(yyDollar[1].value, yyDollar[3].value)
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:284
		{
			yyVAL.value = lx.vAccessObject(yyDollar[1].value, yyDollar[3].value.v)
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:285
		{
			yyVAL.value = lx.vMakeArray(yyDollar[2].values)
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:286
		{
			yyVAL.value = lx.vMakeObject(yyDollar[2].mvalue)
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:289
		{
			yyVAL.value = lx.vGetVar(yyDollar[1].value.v)
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:290
		{
			yyVAL.value = value{v: yyDollar[1].value.v, t: "string"}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:291
		{
			yyVAL.value = value{v: yyDollar[1].value.v, t: "int"}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:292
		{
			yyVAL.value = value{v: yyDollar[1].value.v, t: "bool"}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:293
		{
			yyVAL.value = value{v: "time.Now()", t: "time.Time"}
			lx.addImport("time")
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:297
		{
			yyVAL.values = []value{yyDollar[1].value}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:298
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].value)
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:302
		{
			yyVAL.mvalue = map[string]value{yyDollar[1].value.v: yyDollar[3].value}
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:303
		{
			yyVAL.mvalue = yyDollar[1].mvalue
			yyVAL.mvalue[yyDollar[3].value.v] = yyDollar[5].value
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:305
		{
			yyVAL.mvalue = map[string]value{yyDollar[1].value.v: lx.vGetVar(yyDollar[1].value.v)}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:306
		{
			yyVAL.mvalue = yyDollar[1].mvalue
			yyVAL.mvalue[yyDollar[3].value.v] = lx.vGetVar(yyDollar[3].value.v)
		}
	}
	goto yystack /* stack new state and value */
}
