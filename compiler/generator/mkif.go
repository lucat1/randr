package generator

import (
	"errors"
	"go/ast"
	"go/token"
	"reflect"
)

// makeFor generates a for loop in golang ast
// from a string, also doing some typechecking
// via the integrated golang parser (parsing a fake expr/stmt)
func makeIf(expr *node) (ast.Node, []ast.Stmt, error) {
	_res := random(10)
	res := makeLit(_res)
	input := expr.value[2 : len(expr.value)-1]
	stmt, err := parseStmt(input + " {}")
	if err != nil {
		return nil, nil, errors.New("Could not parse if statement: " + err.Error())
	}
	forLoop := reflect.Indirect(reflect.ValueOf(stmt))

	// Build the for loop body(expr.children if any)
	body := makeBlock(_res, expr.children)
	
	// Set the body value
	if f := forLoop.FieldByName("Body"); f.IsValid() && f.CanSet() {
		f.Set(reflect.ValueOf(body))

		return res, []ast.Stmt{
			&ast.DeclStmt{
				Decl: &ast.GenDecl{
					Tok: token.VAR,
					Specs: []ast.Spec{
						&ast.ValueSpec{
							Type: makeIdent("string"),
							Names: []*ast.Ident{
								makeIdent(_res),
							},
						},
					},
				},
			},
			stmt,
		}, nil
	}

	return nil, nil, errors.New("Cannot set if body")
}