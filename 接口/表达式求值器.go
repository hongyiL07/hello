package main

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

func main() {
	fmt.Println("error 接口")
}

// Expr：算术表达式
type Expr interface {
	Eval(env Env) float64 //返回表达式在 env 上上下文的值

	//给 Expr 方法加上另外一个方法，Check 方法用于在表达式语法树上检查静态错误
	Check(vars map[Var]bool) error
}

type Var string      // Var表示一个变量
type literal float64 //一元操作符表达式
type unary struct {
	op   rune // +   - 中的一个
	x, y Expr
}
type binary struct { //二元操作符表达式
	op   rune // +-*/ 中的一个
	x, y Expr
}

type call struct { //函数调用表达式
	fn   string // pow sin sqrt 中的一个
	args []Expr
}

//要对包含变量的表达式求值，需要一个上下文来把变量映射到数值
type Env map[Var]float64

// Var 的 Eval 方法从上下文中查询结果，如果变量不存在返回 0 ，literal 的 Eval 方法则直接返回本身的值
func (v Var) Eval(env Env) float64 {
	return env[v]
}
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// unary 和 binary 的 Eval 方法首先对它们的操作数递归求值，然后应用 op 操作
//最后 call 方法先对 pow sin 或者 sqrt 函数的参数求值，再调用 math 包中对应的函数

//给 Expr 方法加上另外一个方法，Check 方法用于在表达式语法树上检查静态错误

// parseAndCheck 函数组合了解析和检查步骤
func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if V != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable : %s", v)
		}
	}
	return expl, nil
}

//要完成这个 Web 应用，仅需要加下面的 plot 函数，其函数签名与 htpp.HandlerFunc 类似
func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

//一个接口类型，包含一个返回错误消息的方法
type error interface {
	Error() string
}

//构造 error 最简单的方法是调用 errors.New，会返回一个包含指定错误消息的新 error 实例
//直接调用 errors.New 比较罕见，有一个更易用的封装函数 fmt.Errorf,它还额外提供了字符串格式化功能
