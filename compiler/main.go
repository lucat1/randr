package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"

	"github.com/lucat1/randr/compiler/generator"
	cparser "github.com/lucat1/randr/compiler/parser"
	"golang.org/x/tools/go/ast/astutil"
)

var (
	write *bool
	watch *bool
)

func init() {
	write = flag.Bool("write", true, "Write files automatically after compilation")
	watch = flag.Bool("watch", false, "Watch for file changes on a file or a directory")
	flag.Parse()
}

func main() {
	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [opts] <input>\nOpts:\n", os.Args[0])

		flag.PrintDefaults()
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, flag.Arg(0), nil, parser.ParseComments)
	if err != nil {
		log.Fatal("Could not parse input file:" + err.Error())
	}

	cparser.CheckComment(node, fset)
	astutil.Apply(node, generator.Visit(fset, node), nil)
	printer.Fprint(os.Stdout, fset, node)
	
}