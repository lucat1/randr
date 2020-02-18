package generator

import (
	"go/ast"
	"go/token"
)

// add creates a golang ast made up of binary addition exprs
func add(_nodes []ast.Node) ast.Node {
	if len(_nodes) == 1 {
		// Prevent from attempting to sum only one node
		return _nodes[0]
	}

	nodes := reverse(_nodes)
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

		this := &ast.BinaryExpr{
			Op: token.ADD,
			Y:  node.(ast.Expr),
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

// reverse reverses the given slice
func reverse(s []ast.Node) []ast.Node {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
