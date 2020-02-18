package generator

import (
	"go/ast"
	"strings"
)

func makeStrs(raws []string) ast.Node {
	return makeStr(strings.Join(raws, ""))
}

func makeStr(str string) ast.Node {
	return makeLit("`" + str + "`")
}