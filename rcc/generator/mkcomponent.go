package generator

import (
	"encoding/json"
	"errors"
	"go/ast"
	"go/token"
	"log"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/lucat1/randr/rcc/parser"
	"golang.org/x/net/html"
)

func makeComponent(expr *node) (ast.Node, []ast.Stmt, error) {
	input := expr.value[8:len(expr.value)-1]
	parts := strings.Split(input, " ")

	// Once again a rudimental parsing, but to be fair
	// custom component should be handled much more carefully.
	// it's just the current design the problem here
	// But anyway this expression should only be used internall
	// so we can assume it's okay to have a bad parser in this case
	name, _attr := parts[0], strings.Join(parts[1:], " ")
	var attr []html.Attribute
	if err := json.Unmarshal([]byte(_attr), &attr); err != nil {
		return nil, nil, errors.New(`Cannot parse custom component arguments,
		illegal usage of the internal #randr expression: ` + err.Error())
	}
	name = strcase.ToCamel(name)

	var (
		extras []ast.Stmt
		children []ast.Node
	)

	if len(expr.children) > 0 {
		// If we have children inside the expression
		// it means that they are indeed the children
		// of the custom element, so we just compute
		// them before the actual render call
		children, extras = toNodes(expr.children)
	}

	// Build a struct for the given props
	props := &ast.CompositeLit{
		Type: makeIdent(name + "Props"),
		Elts: []ast.Expr{},
	}
	for _, prop := range attr {
		key := strcase.ToCamel(prop.Key)
		// We gotta parse the value as a whole new template
		// because theoretically it could contain anything
		raws, exprs := parser.Text(prop.Val, false, []string{}, []string{})
		value, e, err := Generate(raws, exprs)
		if err != nil {
			log.Fatal("Generation error inside custom component props: " + err.Error())
		}
		extras = append(extras, e...)
		props.Elts = append(props.Elts, &ast.KeyValueExpr{
			Key: makeIdent(key),
			Value: value.(ast.Expr),
		})
	}

	if len(children) > 0 {
		// Apply any children as the item Children
		// inside the props struct
		props.Elts = append(props.Elts, &ast.KeyValueExpr{
			Key: makeIdent("Children"),
			Value: add(children).(ast.Expr),
		})
	}

	return &ast.CallExpr{
		Args: []ast.Expr{
			makeIdent(name),
			makeIdent("ctx"), // SHOULDN'T BE HARDCODED :/
			&ast.UnaryExpr{
				Op: token.AND,
				X: props,
			},
		},
		Fun: &ast.SelectorExpr{
			X: makeIdent("randr"), // TODO: Custom import name :/
			Sel: makeIdent("MustRender"),
		},
	}, extras, nil
}