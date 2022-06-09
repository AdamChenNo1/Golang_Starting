/*
 * File: \src\ch7\ExprEval\parse.go                                            *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 17:54:13                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 18:20:49                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package expr

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// ---- parser ----

// Parse parses the input string as an arithmetic expression.
//
//   expr = num                         a literal number, e.g., 3.14159
//        | id                          a variable name, e.g., x
//        | id '(' expr ',' ... ')'     a function call
//        | '-' expr                    a unary operator (+-)
//        | expr '+' expr               a binary operator (+-*/)
//
func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			panic(x)
		}
	}()

	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next() //initial lookahead
	e := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

type lexer struct {
	scan scanner.Scanner

	token rune //current lookahead token
}

func (l *lexer) next() {
	l.token = l.scan.Scan()
}

func (l *lexer) text() string {
	return l.scan.TokenText()
}

func (l *lexer) describe() string {
	switch l.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifer %s", l.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", l.text())
	}

	return fmt.Sprintf("%q", rune(l.token))
}

func parseExpr(l *lexer) Expr {
	return parseBinary(l, 1)
}

// binary = unary ('+' binary)*
// parseBinary stops when it encounters an
// operator of lower precedence than prec1.
func parseBinary(l *lexer, prec1 int) Expr {
	lhs := parseUnary(l)
	for prec := precedence(l.token); prec >= prec1; prec-- {
		for precedence(l.token) == prec {
			op := l.token
			l.next()
			rhs := parseBinary(l, prec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

func parseUnary(l *lexer) Expr {
	if l.token == '+' || l.token == '-' {
		op := l.token
		l.next()
		return unary{op, parseUnary(l)}
	}

	return parsePrimary(l)
}

// primary = id
//         | id '(' expr ',' ... ',' expr ')'
//         | num
//         | '(' expr ')'
func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case scanner.Ident:
		id := lex.text()
		lex.next() // consume Ident
		if lex.token != '(' {
			return Var(id)
		}
		lex.next() // consume '('
		var args []Expr
		if lex.token != ')' {
			for {
				args = append(args, parseExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next() // consume ','
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("got %s, want ')'", lex.describe())
				panic(lexPanic(msg))
			}
		}
		lex.next() // consume ')'
		return call{id, args}

	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next() // consume number
		return literal(f)

	case '(':
		lex.next() // consume '('
		e := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next() // consume ')'
		return e
	}
	msg := fmt.Sprintf("unexpected %s", lex.describe())
	panic(lexPanic(msg))
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}

	return 0
}

type lexPanic string
