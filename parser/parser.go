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

//line grammar.y:54
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
const INTTYPE = 57382
const BOOLTYPE = 57383
const STRINGTYPE = 57384
const BINTYPE = 57385
const FOR = 57386
const RETURN = 57387
const WAITFOR = 57388
const OPTIONS = 57389
const IGNORE = 57390
const HEADLESS = 57391
const TIMEOUT = 57392
const TRUE = 57393
const FALSE = 57394
const EVENT = 57395
const LIKE = 57396
const IN = 57397
const WHILE = 57398
const BOOL = 57399
const AT = 57400
const IDENTIFIER = 57401
const IGNOREID = 57402
const STRING = 57403
const NUMBER = 57404
const NAMESPACESEPARATOR = 57405
const SELECT = 57406
const ALL = 57407
const ANY = 57408
const ONE = 57409
const AS = 57410
const FROM = 57411
const WHERE = 57412
const LIMIT = 57413
const DISTINCT = 57414
const SORT = 57415
const ASC = 57416
const DESC = 57417
const DEFAULT = 57418
const CASE = 57419
const CLICK = 57420
const DOCUMENT = 57421
const PAGE = 57422
const CONTAINS = 57423

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
	"DOCUMENT",
	"PAGE",
	"CONTAINS",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:273

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 385

var yyAct = [...]uint8{
	27, 130, 129, 49, 48, 90, 63, 64, 65, 60,
	62, 61, 69, 67, 68, 66, 70, 71, 132, 131,
	138, 139, 118, 17, 13, 41, 42, 58, 72, 73,
	74, 124, 79, 81, 28, 128, 127, 32, 15, 31,
	125, 33, 91, 18, 117, 45, 113, 95, 92, 87,
	83, 43, 44, 47, 40, 34, 35, 16, 88, 24,
	19, 6, 86, 26, 107, 56, 109, 103, 106, 104,
	105, 32, 108, 31, 39, 33, 36, 94, 37, 38,
	100, 121, 98, 75, 97, 110, 99, 102, 40, 34,
	35, 77, 12, 76, 93, 101, 54, 46, 11, 112,
	55, 57, 114, 22, 115, 116, 14, 5, 39, 21,
	36, 8, 37, 38, 120, 119, 50, 52, 51, 53,
	10, 9, 126, 4, 29, 3, 2, 7, 133, 20,
	1, 135, 136, 137, 134, 135, 25, 82, 89, 141,
	142, 143, 63, 64, 65, 60, 62, 61, 69, 67,
	68, 66, 70, 71, 140, 78, 80, 23, 30, 84,
	85, 59, 0, 0, 72, 73, 74, 63, 64, 65,
	60, 62, 61, 69, 67, 68, 66, 70, 71, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	73, 74, 63, 64, 65, 60, 62, 61, 69, 67,
	68, 66, 70, 71, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 73, 74, 0, 0, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 123, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 122, 63, 64, 65,
	60, 62, 61, 69, 67, 68, 66, 70, 71, 75,
	0, 0, 0, 0, 111, 0, 0, 0, 0, 72,
	73, 74, 63, 64, 65, 60, 62, 61, 69, 67,
	68, 66, 70, 71, 0, 0, 0, 0, 0, 0,
	0, 96, 0, 0, 72, 73, 74, 63, 64, 65,
	60, 62, 61, 69, 67, 68, 66, 70, 71, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	73, 74, 0, 0, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 75,
}

var yyPact = [...]int16{
	-1000, -1000, 3, 3, -21, -1000, 1, -21, -1000, -1000,
	-21, -1000, -1000, 0, -21, 33, 17, -6, -65, 76,
	-1000, -1000, -1000, 46, -1000, -1000, 17, 303, -1000, 73,
	51, 17, 17, -9, 51, 51, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 7, 17, -1000,
	-1000, -1000, -1000, -1000, 76, -17, -11, -1000, 303, 17,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 17, -12, 73, 278,
	63, 303, 61, 79, 73, 73, 17, 2, 43, 47,
	-1000, 69, -1000, -1000, 253, -1000, -1000, -1000, 17, -1000,
	-13, 17, 303, 17, 17, -15, -46, -1000, -1000, -17,
	76, -1000, 303, 65, 303, 188, 163, -35, -19, -1000,
	-1000, 17, -23, -24, -58, -58, 303, -1000, -1000, -58,
	-1000, 17, 17, -58, -50, -1000, 138, 303, 17, 17,
	17, 303, 303, 303,
}

var yyPgo = [...]uint8{
	0, 0, 34, 124, 161, 158, 3, 157, 156, 138,
	5, 137, 1, 2, 134, 130, 126, 125, 123, 121,
	107, 120, 98, 92, 106, 101,
}

var yyR1 = [...]int8{
	0, 15, 15, 16, 18, 17, 17, 20, 6, 6,
	6, 6, 6, 6, 10, 9, 9, 19, 19, 21,
	21, 23, 23, 22, 22, 25, 7, 7, 24, 24,
	24, 24, 24, 14, 14, 14, 13, 13, 12, 12,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 5, 5, 5, 5,
	5, 5, 1, 1, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 8, 8,
	11, 11, 11, 11,
}

var yyR2 = [...]int8{
	0, 4, 3, 0, 0, 2, 1, 3, 1, 1,
	1, 1, 3, 3, 3, 1, 3, 2, 1, 1,
	2, 3, 2, 2, 3, 0, 1, 3, 4, 8,
	7, 7, 7, 0, 3, 3, 1, 2, 4, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 2, 3, 4, 3, 3,
	3, 2, 2, 1, 1, 1, 1, 1, 1, 3,
	3, 5, 1, 3,
}

var yyChk = [...]int16{
	-1000, -15, -16, -17, -18, -20, 58, -18, -20, -19,
	-21, -22, -23, 45, -24, 59, 78, 44, 64, 59,
	-19, -22, -23, -7, 59, -19, 30, -1, -2, -3,
	-5, 22, 20, 24, 38, 39, 59, 61, 62, 57,
	37, 8, 9, 34, 35, 28, 80, 59, 69, -6,
	40, 42, 41, 43, 20, 24, 19, -25, -1, -4,
	7, 9, 8, 4, 5, 6, 13, 11, 12, 10,
	14, 15, 26, 27, 28, 81, 20, 18, -3, -1,
	-8, -1, -11, 59, -3, -3, 55, -1, -6, -9,
	-10, 59, 59, -2, -1, 59, 23, 21, 19, 25,
	19, 16, -1, 65, 67, 68, 66, 21, 25, 19,
	16, 21, -1, 59, -1, -1, -1, 59, 68, -10,
	-6, 16, 68, 68, 66, 59, -1, 59, 59, -13,
	-12, 77, 76, -13, -14, -12, -1, -1, 70, 71,
	16, -1, -1, -1,
}

var yyDef = [...]int8{
	3, -2, 4, 4, 0, 6, 0, 0, 5, 2,
	0, 18, 19, 0, 0, 0, 0, 0, 0, 0,
	1, 17, 20, 23, 26, 25, 0, 22, 62, 64,
	0, 0, 0, 0, 0, 0, 73, 74, 75, 76,
	77, 56, 57, 58, 59, 60, 61, 0, 0, 7,
	8, 9, 10, 11, 0, 0, 0, 24, 21, 0,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 0, 0, 65, 0,
	0, 78, 0, 82, 71, 72, 0, 0, 0, 0,
	15, 0, 27, 63, 0, 68, 66, 69, 0, 70,
	0, 0, 28, 0, 0, 0, 0, 12, 13, 0,
	0, 67, 79, 83, 80, 0, 0, 0, 0, 16,
	14, 0, 0, 0, 0, 0, 81, 33, 30, 31,
	36, 0, 0, 32, 29, 37, 0, 39, 0, 0,
	0, 34, 35, 38,
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
//line grammar.y:101
		{
			lx.finalize()
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:102
		{
			lx.finalize()
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:106
		{
			lx = yylex.(*myLexer)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:110
		{
			lx.incOut()
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:120
		{
			lx.declInputParam(yyDollar[2].value.v, yyDollar[3].value.v)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:124
		{
			yyVAL.value.v = "int"
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:125
		{
			yyVAL.value.v = "string"
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:126
		{
			yyVAL.value.v = "bool"
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:127
		{
			yyVAL.value.v = "[]byte"
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:128
		{
			yyVAL.value.v = "[]" + yyDollar[2].value.v
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:129
		{
			yyVAL.value.v = lx.objectType(yyDollar[2].values)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:133
		{
			yyVAL.value = value{v: yyDollar[1].value.v, t: yyDollar[3].value.v}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:136
		{
			yyVAL.values = []value{yyDollar[1].value}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:137
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].value)
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:143
		{
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:144
		{
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:148
		{ /* todo */
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:149
		{ /* todo */
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:153
		{
			lx.vSetVar(yyDollar[1].value.v, yyDollar[3].value)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:154
		{ /* todo */
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:158
		{
			lx.declOutputParams(yyDollar[2].list)
			lx.saveOut()
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:159
		{ /* */
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:163
		{
			lx.addLines("}")
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:166
		{
			yyVAL.list = append(yyVAL.list, yyDollar[1].value.v)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:167
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].value.v)
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:171
		{
			lx.forNameInExpression(yyDollar[2].value.v, yyDollar[4].value)
		}
	case 29:
		yyDollar = yyS[yypt-8 : yypt+1]
//line grammar.y:176
		{
			opt := yyDollar[8].selopt
			opt.from = yyDollar[3].value
			opt.css = yyDollar[5].value
			opt.loopv = yyDollar[7].value.v
			lx.selectAll(opt)
		}
	case 30:
		yyDollar = yyS[yypt-7 : yypt+1]
//line grammar.y:177
		{
			lx.addLines("{// select TODO")
		}
	case 31:
		yyDollar = yyS[yypt-7 : yypt+1]
//line grammar.y:179
		{
			lx.addLines("{// select TODO")
		}
	case 32:
		yyDollar = yyS[yypt-7 : yypt+1]
//line grammar.y:180
		{
			lx.addLines("{// select TODO")
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:184
		{
			yyVAL.selopt = selopt{}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:185
		{
			yyVAL.selopt = yyDollar[1].selopt
			yyVAL.selopt.where = append(yyVAL.selopt.where, yyDollar[3].value)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:186
		{
			yyVAL.selopt = yyDollar[1].selopt
			yyVAL.selopt.limit = yyDollar[3].value
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:190
		{
			yyVAL.casopts = []casopt{}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:191
		{
			yyVAL.casopts = append(yyDollar[1].casopts, yyDollar[2].casopt)
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:196
		{
			yyVAL.casopt = casopt{e1: yyDollar[2].value, e2: yyDollar[4].value}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:197
		{
			yyVAL.casopt = casopt{def: true, e2: yyDollar[2].value}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:205
		{
			yyVAL.value = yyDollar[1].value
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:206
		{
			yyVAL.value = yyDollar[1].value
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:207
		{
			yyVAL.value = yyDollar[1].value
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:208
		{
			yyVAL.value = yyDollar[1].value
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:209
		{
			yyVAL.value = yyDollar[1].value
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:210
		{
			yyVAL.value = yyDollar[1].value
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:211
		{
			yyVAL.value = yyDollar[1].value
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:212
		{
			yyVAL.value = yyDollar[1].value
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:213
		{
			yyVAL.value = yyDollar[1].value
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:214
		{
			yyVAL.value = yyDollar[1].value
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:215
		{
			yyVAL.value = yyDollar[1].value
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:216
		{
			yyVAL.value = yyDollar[1].value
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:217
		{
			yyVAL.value = yyDollar[1].value
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:218
		{
			yyVAL.value = yyDollar[1].value
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:219
		{
			yyVAL.value = yyDollar[1].value
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:220
		{
			yyVAL.value = yyDollar[1].value
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:224
		{
			yyVAL.value = yyDollar[1].value
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:225
		{
			yyVAL.value = yyDollar[1].value
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:226
		{
			yyVAL.value = yyDollar[1].value
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:227
		{
			yyVAL.value = yyDollar[1].value
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:228
		{
			yyVAL.value = yyDollar[1].value
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:229
		{
			yyVAL.value = yyDollar[1].value
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:234
		{
			yyVAL.value = yyDollar[1].value
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:235
		{
			yyVAL.value = lx.vOpe2(yyDollar[2].value.c, yyDollar[1].value, yyDollar[3].value)
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:239
		{
			yyVAL.value = yyDollar[1].value
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:240
		{
			yyVAL.value = lx.vOpe1(yyDollar[1].value.c, yyDollar[2].value)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:244
		{
			yyVAL.value = lx.vParen(yyDollar[2].value)
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:245
		{
			yyVAL.value = lx.vGetElementOf(yyDollar[1].value, yyDollar[3].value)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:246
		{
			yyVAL.value = lx.vAccessObject(yyDollar[1].value, yyDollar[3].value.v)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:247
		{
			yyVAL.value = lx.vMakeArray(yyDollar[2].values)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:248
		{
			yyVAL.value = lx.vMakeObject(yyDollar[2].mvalue)
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:250
		{ /*todo*/
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:251
		{ /* */
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:253
		{
			yyVAL.value = lx.vGetVar(yyDollar[1].value.v)
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:254
		{
			yyVAL.value = value{v: yyDollar[1].value.v, t: "string"}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:255
		{
			yyVAL.value = value{v: yyDollar[1].value.v, t: "int"}
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:256
		{
			yyVAL.value = value{v: yyDollar[1].value.v, t: "bool"}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:257
		{
			yyVAL.value = value{v: "time.Now()", t: "time.Time"}
			lx.imports["time"] = true
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:261
		{
			yyVAL.values = []value{yyDollar[1].value}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:262
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].value)
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:266
		{
			yyVAL.mvalue = map[string]value{yyDollar[1].value.v: yyDollar[3].value}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar.y:267
		{
			yyVAL.mvalue = yyDollar[1].mvalue
			yyVAL.mvalue[yyDollar[3].value.v] = yyDollar[5].value
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:269
		{
			yyVAL.mvalue = map[string]value{yyDollar[1].value.v: lx.vGetVar(yyDollar[1].value.v)}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:270
		{
			yyVAL.mvalue = yyDollar[1].mvalue
			yyVAL.mvalue[yyDollar[3].value.v] = lx.vGetVar(yyDollar[3].value.v)
		}
	}
	goto yystack /* stack new state and value */
}
