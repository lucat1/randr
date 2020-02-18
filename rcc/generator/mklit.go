package generator

import (
	"go/ast"
	"go/token"
)

// makeLit creates a variable in the golang ast
func makeLit(v string) ast.Node {
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: v,
	}
}

func makeIdent(v string) *ast.Ident {
	return &ast.Ident{
		Name: v,
	}
}
