package generator

import (
	"go/ast"
	"go/token"
)

// add creates a golang ast made up of binary addition exprs
func add(nodes []ast.Node) ast.Node {
	var (
		root   *ast.BinaryExpr = nil
		latest *ast.BinaryExpr = nil
	)

	for i, node := range nodes {
		// Ignore empty strings
		if n, ok := node.(*ast.BasicLit); ok && n.Value == "``" {
			continue
		}

		if i == len(nodes)-1 {
			latest.X = node.(ast.Expr)
			break
		}

		this := &ast.BinaryExpr {
			Op: token.ADD,
			Y: node.(ast.Expr),
		}
		if latest != nil {
			latest.X = this
		}

		latest = this
		if root == nil {
			root = this
		}
	}

	return root
}