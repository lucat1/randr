package main

import (
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"

	"github.com/lucat1/randr/compiler/generator"
	"golang.org/x/tools/go/ast/astutil"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "inputs/hello.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	//ast.Print(fset, node)
	astutil.Apply(node, generator.Visit(fset, node), nil)

	printer.Fprint(os.Stdout, fset, node)
}