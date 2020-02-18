package generator

// Taken from:
// https://golang.hotexamples.com/site/file?hash=0xeae472f8a2c86a3371230cd9c3c44240f6f7a57ecbbfd523f8bfa2777706e5d1&fullName=Godeps/_workspace/src/github.com/0xfaded/eval.go&project=philipmulcahy/godebug

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
)

// A convenience function for parsing a Stmt by wrapping it in a func literal
func parseStmt(input string) (ast.Stmt, error) {
	// parser.ParseExpr has some quirks, for example it won't parse unused map literals as
	// ExprStmts. Unused map literals are technically illegal, but it would be
	// nice to check them at a later stage. Therefore, we want to parse expressions
	// as expressions, and stmts and stmts. We try both.
	// However, there is a bug in parser.ParseExpr that it does not detect excess input.
	// Therefore, the _ of _ = 1 will be parsed as an expression. To avoid this, attempt
	// to parse the input as a statement first, and fall back to an expression
	expr := "func(){" + input + ";}"
	e, err := parser.ParseExpr(expr)
	if err != nil {
		if e, err := parser.ParseExpr(input); err == nil {
			return &ast.ExprStmt{X: e}, nil
		}

		errs := err.(scanner.ErrorList)
		for i := range errs {
			errs[i].Pos.Offset -= 7
			errs[i].Pos.Column -= 7
		}

		return nil, errs
	}

	node := e.(*ast.FuncLit).Body.List[0]
	stmt, ok := node.(ast.Stmt)
	if !ok {
		return nil, fmt.Errorf("%T not supported", node)
	}
	
	return stmt, nil
}