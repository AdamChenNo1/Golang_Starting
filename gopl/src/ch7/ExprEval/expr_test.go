/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-20 22:55:04
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-21 01:02:16
 * @FilePath: \GoPath\src\ExprEval\expr_test.go
 */
package expr

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x,3) + pow(y,3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x,3) + pow(y,3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}

	var prevExpr string
	for _, test := range tests {
		//仅在表达式变更的时候才输出
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // 解析出错
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q,want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
