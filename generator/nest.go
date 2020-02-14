package generator

type nodeType int

const (
	rawType = iota
	exprType
)

type node struct {
	kind     nodeType
	value    string
	children []*node
}

func nest(parts []string) []*node {
	depth := 0
	res := []*node{}
	var latest *node
	for _, part := range parts {
		// Ignore empty strings
		if part == "" {
			continue
		}
		curr := gen(part)
		// decrease the depth and ignore the ending tag
		if curr.kind == exprType && part[1] == '/' {
			depth--
			continue
		}

		if depth != 0 {
			// we are inside a nesting, we handle whatever
			// comes as normal but we put it inside the latest
			// expression's children slice
			latest.children = append(latest.children, curr)
		} else {
			res = append(res, curr)
		}

		// increase the depth and handle the expression
		if curr.kind == exprType && part[1] == '#' {
			depth++
		}

		// if what we had was an expression we save it
		// as it will be used later to nest its children
		if curr.kind == exprType {
			latest = curr
		}
	}

	return res
}

func gen(part string) *node {
	isExpr := len(part) > 1 && part[0] == '{' && part[len(part)-1] == '}'
	if !isExpr {
		// Handle a static string easily
		return &node{
			kind:  rawType,
			value: part,
		}
	}

	return &node{
		kind:  exprType,
		value: part,
	}
}
