// goexpr attempts to parse each command line argument as a go expression. This
// is a basic wrapper around `go/parser.ParseExpr`. Each argument has it's AST
// written to stdout, even if `ParseExpr` returns an error. The status of the
// parse result is written to stderr for each expression as well.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage: goexpr EXPRESSION ...")
		os.Exit(1)
	}
	for _, arg := range args {
		expr, err := parser.ParseExpr(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "goexpr:fail:", arg)
			fmt.Fprintln(os.Stderr, "goexpr:error:", err)
		} else {
			fmt.Fprintln(os.Stderr, "goexpr:pass:", arg)
		}
		ast.Print(nil, expr)
	}
}
