/*
 * File: \src\ch12\format\format_test.go                                       *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 17:13:28                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 18:46:05                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package display

import (
	expr "go_start/src/ch7/ExprEval"
)

func ExampleDisplay() {
	e, _ := expr.Parse("sqrt(A / pi)")
	Display("e", e)

	// Output:
	// Display e (expr.call):
	// e.fn = "sqrt"
	// e.args[0].type = expr.binary
	// e.args[0].value.op = 47
	// e.args[0].value.x.type = expr.Var
	// e.args[0].value.x.value = "A"
	// e.args[0].value.y.type = expr.Var
	// e.args[0].value.y.value = "pi"
}
