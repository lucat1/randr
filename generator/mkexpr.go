package generator

import (
	"errors"
	"go/ast"
	"go/token"
	"log"
	"strings"
)

// astify transforms a template string
// into a golang ast, and returns eventual
// extra ast code to be appended at the 
// top of the (function) scope
func makeExpr(expr string) (ast.Node, error) {
	// remove trailing and edngin { }
	input := expr[1:len(expr)-1]
	if input[0] == '/' {
		// we can safely ignore any closing tag
		// as the depth has already been checked
		return nil, nil
	}

	if input[0] != '#' {
		// this means we're just appending a value
		// to the string
		return &ast.BasicLit{
			Kind: token.STRING,
			Value: input,
		}, nil
	}

	// at this point we can anlyze the
	// expression, removing the trailing #
	input = input[1:]
	parts := strings.Split(input, " ")
	switch parts[0] {
	case "for":
		return makeFor(input)
		break
	}

	return nil, errors.New("Unkown expression: " + input)
}

func toExprs(exprs []string) []ast.Node {
	res := []ast.Node{}
	for _, expr := range exprs {
		node, err := makeExpr(expr)
		if err != nil {
			log.Fatal(err)
		}

		if node != nil {
			res = append(res, node)
		}
	}

	return res
}