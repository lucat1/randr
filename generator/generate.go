package generator

import (
	"go/ast"
)

// Generate generates a golang ast from two arrays
// of raw strings and template experssions
func Generate(raws []string, exprs []string) (ast.Node, []ast.Stmt, error) {
	if len(exprs) == 0 {
		// Entirely static template, just handle it simply
		return makeStrs(raws), []ast.Stmt{}, nil
	}

	rawNodes := group(raws, exprs)
	nested := nest(rawNodes)
	nodes, extras := toNodes(nested)

	return add(nodes), extras, nil
}
