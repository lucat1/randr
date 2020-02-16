package generator

import (
	"errors"
	"go/ast"
	"go/token"
	"reflect"
)

type forParserState int

const (
	forParserIdle = iota
	forParserFor
	forParserKey
	forParserValue
	forParserRange
	forParserSlice
)

// makeFor generates a for loop in golang ast
// from a string, also doing some typechecking
// via the integrated golang parser (parsing a fake expr/stmt)
func makeFor(expr *node) (ast.Node, []ast.Stmt, error) {
	_res := random(10)
	res := makeLit(_res)
	input := expr.value[2 : len(expr.value)-1]
	stmt, err := parseStmt(input + " {}")
	if err != nil {
		return nil, nil, err
	}
	forLoop := reflect.Indirect(reflect.ValueOf(stmt))

	// Build the for loop body(expr.children if any)
	body := &ast.BlockStmt{}
	if len(expr.children) > 0 {
		// nodes are the values returned by the expressions
		// extras are the generated expressions themselves
		// we lastly just wanna append the `nodes`(values) to
		// our return value (_res)
		nodes, extras := toNodes(expr.children)
		body.List = append(body.List, extras...)
		left := makeIdent(_res)
		body.List = append(body.List, makeAssign(left, add(nodes).(ast.Expr)))
	}
	
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

	return nil, nil, errors.New("Cannot set for loop body")
}
