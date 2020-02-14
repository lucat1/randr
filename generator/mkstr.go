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

// toStrs transforms an array of strings into an
// array of golang ast strings
func toStrs(strs []string) []ast.Node {
	res := []ast.Node{}
	for _, str := range strs {
		res = append(res, makeStr(str))
	}

	return res
}
