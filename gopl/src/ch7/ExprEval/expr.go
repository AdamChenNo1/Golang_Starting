/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-20 22:29:18
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-20 22:54:16
 * @FilePath: \GoPath\src\ExprEval\expr.go
 */
//算术表达式求值器
package expr

import (
	"fmt"
	"math"
)

//Expr:算数运算表达式
type Expr interface {
	//Eval 返回表达式在env上下文下的值
	Eval(env Env) float64
}

//Var 表示一个变量，比如x
type Var string

//literal 是一个数字常量，比如3.141
type literal float64

//unary 表示一元操作符表达式，如-x
type unary struct {
	op rune //'+','-'中的一个
	x  Expr
}

//binary 表示二元操作符表达式，如x+y
type binary struct {
	op   rune //'+','-','*','/'中的一个
	x, y Expr
}

//call 表示函数调用表达式，如sin(x)
type call struct {
	fn   string //"pow","sin","sqrt"中的一个
	args []Expr
}

//Env 上下文，用来将变量映射到数值
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {

	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {

	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "cos":
		return math.Cos(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}
