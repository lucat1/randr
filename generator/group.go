package generator

import "go/ast"

// group groups bot the raws and the exprs
// arrays into a single array of ast nodes
func group(raws []ast.Node, exprs []ast.Node) []ast.Node {
	res := []ast.Node{}

	for i, raw := range raws {
		if i < len(exprs) {
			res = append(res, raw, exprs[i])
		} else {
			res = append(res, raw)
		}
	}

	return reverse(res)
}

// reverse reverses the given slice
func reverse(s []ast.Node) []ast.Node {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
	}

	return s
}