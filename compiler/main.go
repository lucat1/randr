package main

import (
	"flag"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"

	"github.com/lucat1/randr/compiler/generator"
	cparser "github.com/lucat1/randr/compiler/parser"
	"golang.org/x/tools/go/ast/astutil"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("Wrong usage, must provide at least one file source: randr <file.go>")
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, flag.Arg(0), nil, parser.ParseComments)
	if err != nil {
		log.Fatal("Could not parse input file:" + err.Error())
	}
	cparser.CheckComment(node, fset)

	//ast.Print(fset, node)
	astutil.Apply(node, generator.Visit(fset, node), nil)

	printer.Fprint(os.Stdout, fset, node)
}