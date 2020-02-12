package generator

import (
	"go/ast"
)

// Generate generates a golang ast from two arrays
// of raw strings and template experssions
func Generate(raws []string, exprs []string) (ast.Node, error) {
	if len(exprs) == 0 {
		// Entirely static template, just handle it simply
		return makeStrs(raws), nil
	}

	nodes := group(toStrs(raws), toExprs(exprs))

	return add(nodes), nil
}