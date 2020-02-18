package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/lucat1/randr/rcc/generator"
	"golang.org/x/tools/go/ast/astutil"
)

var (
	recursive     *bool
	watch         *bool
	originalInput string
)

func init() {
	recursive = flag.Bool("recursive", false, "Compile files in subfolders of the input directory")
	watch = flag.Bool("watch", false, "Watch for file changes on a file or a directory")
	flag.Parse()
}

func main() {
	if flag.NArg() < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [opts] <input> <output>\n\nOptions:\n", os.Args[0])

		flag.PrintDefaults()
		os.Exit(1)
	}

	input, output := flag.Arg(0), flag.Arg(1)
	originalInput = appendSlash(input)
	fi, err := os.Stat(input)
	if err != nil {
		log.Fatal("Could not stat input source: " + err.Error())
	}

	if fi.Mode().IsDir() {
		handleDirectory(input, output)
	} else {
		checkCompile(compileFile(input, output))
	}
}

func handleDirectory(input string, output string) {
	files, err := ioutil.ReadDir(input)
	if err != nil {
		log.Fatal("Could not read input directory: " + err.Error())
	}

	for _, file := range files {
		in := path.Join(input, file.Name())
		if file.Mode().IsDir() {
			if *recursive {
				handleDirectory(in, output)
			}
			continue
		}

		checkCompile(compileFile(in, output))
	}
}

func checkCompile(err error) {
	if err != nil {
		log.Fatal("Error during compilation: " + err.Error())
	}
}

func appendSlash(path string) string {
	if path[len(path)-1] != '/' {
		path += "/"
	}

	return path
}

func compileFile(input string, output string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, input, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	astutil.Apply(node, generator.Visit(fset, node), nil)
	var compiled bytes.Buffer
	if err = printer.Fprint(&compiled, fset, node); err != nil {
		return err
	}
	// generate the output path
	var filename string
	if input == originalInput {
		// Single file compilation
		filename = strings.Replace(path.Base(input), ".go", "", 1)
	} else {
		filename = strings.Replace(strings.Replace(input, originalInput, "", 1), ".go", "", 1)
	}

	out := strings.Replace(output, "[name]", filename, -1)
	dir := filepath.Dir(out)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println("Creating output directory: " + dir)
		if err = os.Mkdir(dir, 0755); err != nil {
			return err
		}
	}

	log.Printf("Writing file to %s", out)
	return ioutil.WriteFile(out, compiled.Bytes(), 0644)
}