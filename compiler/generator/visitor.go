package generator

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"

	"github.com/lucat1/randr/compiler/parser"
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

	// latestFuncDecl
	var lfd *ast.FuncDecl

	return func(c *astutil.Cursor) bool {
		n := c.Node()

		switch x := n.(type) {
		case *ast.FuncDecl:
			lfd = c.Node().(*ast.FuncDecl)
			break

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
					// Insert any nodes if we got any expressions (extras)
					if _, ok := c.Parent().(*ast.BlockStmt); ok {
						fmt.Println("adding that way")
						for _, extra := range extras {
							c.InsertBefore(extra)
						}
					} else {
						// Find the node inside the parent function
						// we cannot do deep nestings right now, we'll
						// have to refactor this code in order to do so
						index := -1
						for i, child := range lfd.Body.List {
							if child == c.Parent() || child == c.Node() {
								index = i
							}
						}

						if index == -1 {
							log.Fatal(`Deep nesting of randr.HTML is not yet supported,
							please place your call at the root of the function`)
						}

						lfd.Body.List = append(lfd.Body.List[:index], append(extras, lfd.Body.List[index:]...)...)
					}
				}

				c.Replace(final)
			}
		}/*
		switch x := n.(type) {
		case *ast.FuncDecl:
			found, extras, index := false, []ast.Stmt{}, 0
			f := subVisitor(fset, importName, &found, &extras, &index)
			astutil.Apply(x, f, nil)
			if found {
				// we got some extras to append at the
				// beginning of the function's body
				x.Body.List = append(extras, x.Body.List...)
			}
		}*/
		return true
	}
}
/*
func subVisitor(fset *token.FileSet, importName string, found *bool, e *[]ast.Stmt, index *int) astutil.ApplyFunc {
	return func(c *astutil.Cursor) bool {
		n := c.Node()
		c.InsertBefore()
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
*/