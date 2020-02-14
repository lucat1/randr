package generator

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"

	"github.com/lucat1/randr/parser"
	"golang.org/x/tools/go/ast/astutil"
)

// Visit iterates over all the ast nodes
// using astutil.Apply and modifies any
// randr.HTML calls to static string + expressions
func Visit(fset *token.FileSet, node *ast.File) astutil.ApplyFunc {
	importName := "randr"
	for _, imp := range node.Imports {
		found := false

		if imp.Path.Kind == token.STRING && imp.Path.Value == "\"github.com/lucat1/randr\"" {
			found = true
			if imp.Name != nil {
				importName = imp.Name.Name
			}
		}

		if !found {
			log.Fatal("Cannot compile because it doesn't import the library")
		}
	}

	return func(c *astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.FuncDecl:
			found, extras := false, []ast.Stmt{}
			f := subVisitor(fset, importName, &found, &extras)
			astutil.Apply(x, f, nil)
			if found {
				// we got some extras to append at the
				// beginning of the function's body
				for _, e := range extras {
					if k, ok := e.(*ast.RangeStmt); ok {
						fmt.Println(k.Body)
					}
				}
				x.Body.List = append(extras, x.Body.List...)
			}
		}
		return true
	}
}

func subVisitor(fset *token.FileSet, importName string, found *bool, e *[]ast.Stmt) astutil.ApplyFunc {
	return func(c *astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.CallExpr:
			fun, ok := x.Fun.(*ast.SelectorExpr)
			if !ok {
				break
			}

			scope := fun.X.(*ast.Ident).Name
			name := fun.Sel.Name
			if scope == importName && name == "HTML" {
				arg := x.Args[0].(*ast.BasicLit)
				if len(x.Args) != 1 || arg.Kind != token.STRING {
					log.Fatalf("Can only call randr.HTML with a single string argument. At %v\n", fset.Position(n.Pos()))
				}

				// Everything is correnct, we are sure that this
				// CallExpr is a call to randr.HTML so we replace it
				// with a properly generated ast
				raws, exprs, err := parser.Parse(arg.Value)
				if err != nil {
					log.Fatal("Parse error: " + err.Error())
				}

				final, extras, err := Generate(raws, exprs)
				if err != nil {
					log.Fatal("Generation error: " + err.Error())
				}
				if len(extras) > 0 {
					*found = true
					*e = extras
				}
				c.Replace(final)
			}
		}

		return true
	}
}
