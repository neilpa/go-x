// goexpr attempts to parse each command line argument as a go expresions. This
// is a basic wrapper around `go/parser.ParseExpr`. Each successfully parsed
// argument has it's AST written to stdout. Errors for expressions that fail to
// parse are written to stderr.
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
		if err == nil {
			fmt.Printf("goexpr: %s:\n", arg)
			ast.Print(nil, expr)
		} else {
			fmt.Fprintf(os.Stderr, "goexpr: %s: %v\n", arg, err)
		}
	}
}
