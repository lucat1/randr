package generator

import "go/ast"

// makeBlock handles a number of nodes and creates an ast
// with their expressions, assigning them all to the given
// `assignee` variable
func makeBlock(assignee string, nodes []*node) *ast.BlockStmt {
	body := &ast.BlockStmt{}
	if len(nodes) > 0 {
		// nodes are the values returned by the expressions
		// extras are the generated expressions themselves
		// we lastly just wanna append the `nodes`(values) to
		// our return value (_res)
		nodes, extras := toNodes(nodes)
		body.List = append(body.List, extras...)
		left := makeIdent(assignee)
		body.List = append(body.List, makeAssign(left, add(nodes).(ast.Expr)))
	}
	return body
}