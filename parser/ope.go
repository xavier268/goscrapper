package parser

import "fmt"

// binary operators on expressions
func (m *myLexer) vOpe2(ope int, left value, right value) value {
	switch ope {

	case PLUS:
		if left.t == right.t && (left.t == "int" || left.t == "string") {
			return value{t: left.t, v: fmt.Sprintf("((%s) + (%s))", left.v, right.v)}
		}
	case MINUS:
		if left.t == right.t && (left.t == "int") {
			return value{t: left.t, v: fmt.Sprintf("((%s) - (%s))", left.v, right.v)}
		}

	case MULTI:
		if left.t == right.t && (left.t == "int") {
			return value{t: left.t, v: fmt.Sprintf("((%s) * (%s))", left.v, right.v)}
		}

	case DIV:
		if left.t == right.t && (left.t == "int") {
			return value{t: left.t, v: fmt.Sprintf("((%s) / (%s))", left.v, right.v)}
		}
	case MOD:
		if left.t == right.t && (left.t == "int") {
			return value{t: left.t, v: fmt.Sprintf("((%s) %% (%s))", left.v, right.v)}
		}
	case AND:
		if left.t == right.t && (left.t == "bool") {
			return value{t: left.t, v: fmt.Sprintf("((%s) && (%s))", left.v, right.v)}
		}

	case OR:
		if left.t == right.t && (left.t == "bool") {
			return value{t: left.t, v: fmt.Sprintf("((%s) || (%s))", left.v, right.v)}
		}

	case CONTAINS:
		if left.t == right.t && (left.t == "string") {
			m.imports["strings"] = true
			return value{t: "bool", v: fmt.Sprintf("(strings.Contains(%s,%s))", left.v, right.v)}
		}

	case EQ:
		if left.t == right.t { // only compare same types
			return value{t: "bool", v: fmt.Sprintf("((%s) == (%s))", left.v, right.v)}
		}
	case NEQ:
		if left.t == right.t { // only compare same types
			return value{t: "bool", v: fmt.Sprintf("((%s) != (%s))", left.v, right.v)}
		}

	case GT:
		if left.t == right.t && (left.t == "int") { // only compare ints
			return value{t: "bool", v: fmt.Sprintf("((%s) > (%s))", left.v, right.v)}
		}
	case GTE:
		if left.t == right.t && (left.t == "int") { // only compare ints
			return value{t: "bool", v: fmt.Sprintf("((%s) >= (%s))", left.v, right.v)}
		}
	case LT:
		if left.t == right.t && (left.t == "int") { // only compare ints
			return value{t: "bool", v: fmt.Sprintf("((%s) < (%s))", left.v, right.v)}
		}
	case LTE:
		if left.t == right.t && (left.t == "int") { // only compare ints
			return value{t: "bool", v: fmt.Sprintf("((%s) <= (%s))", left.v, right.v)}
		}

	default:
		m.errorf("unknown operator : %s", TokenAsString(ope))
		return value{}
	}

	// we found an operator, but could not apply it.
	m.errorf("types %s and %s are mismatched for %s", left.t, right.t, TokenAsString(ope))
	return value{}

}

// unary operators on expressions
func (m *myLexer) vOpe1(ope int, v value) value {
	switch ope {
	case NOT:
		if v.t == "bool" {
			m.imports["fmt"] = true
			return value{t: "bool", v: fmt.Sprintf("!(%s)", v.v)}
		}
	case MINUS:
		if v.t == "int" {
			m.imports["fmt"] = true
			return value{t: "int", v: fmt.Sprintf("-(%s)", v.v)}
		}
	case LOWER:
		if v.t == "string" {
			m.imports["strings"] = true
			return value{t: "string", v: fmt.Sprintf("strings.ToLower(%s)", v.v)}
		}
	case UPPER:
		if v.t == "string" {
			m.imports["strings"] = true
			return value{t: "string", v: fmt.Sprintf("strings.ToUpper(%s)", v.v)}
		}
	default:
		m.errorf("unknown operator : %s", TokenAsString(ope))
		return value{}
	}

	// we found an operator, but could not apply it.
	m.errorf("type %s is mismatched for %s unary", v.t, TokenAsString(ope))
	return value{}
}

// encapsulate value in parenthesis, keep same type.
func (m *myLexer) vParen(v value) value {
	return value{t: v.t, v: "(" + v.v + ")"}
}
