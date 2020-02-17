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

	var (
		body *ast.BlockStmt = nil
		els  *ast.BlockStmt = nil
		foundElse = false
		elseIndex = -1
	)
	for i, child := range expr.children {
		if child.kind == exprType && child.value == "{else}" {
			foundElse = true
			elseIndex = i
		}
	}

	if foundElse {
		// build both the body and the else block
		bodyChildren := expr.children[:elseIndex]
		elseChildren := expr.children[elseIndex+1:]
		body = makeBlock(_res, bodyChildren)
		els = makeBlock(_res, elseChildren)
	} else {
		// Build the if body(expr.children if any) with all the nodes
		body = makeBlock(_res, expr.children)
	}
	
	// Set the body value
	if b := forLoop.FieldByName("Body"); b.IsValid() && b.CanSet() {
		if e := forLoop.FieldByName("Else"); e.IsValid() && e.CanSet() {
			if body != nil {
				b.Set(reflect.ValueOf(body))
			}

			if els != nil {
				e.Set(reflect.ValueOf(els))
			}

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

		return nil, nil, errors.New("Cannot set else body")
	}

	return nil, nil, errors.New("Cannot set if body")
}