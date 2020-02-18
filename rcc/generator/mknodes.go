package generator

import (
	"go/ast"
	"log"
	"strings"
)

func toNodes(nodes []*node) ([]ast.Node, []ast.Stmt) {
	res := []ast.Node{}
	extras := []ast.Stmt{}
	for _, node := range nodes {
		if node.kind == rawType {
			res = append(res, makeStr(node.value))
		} else {
			if node.value[1] != '#' && node.value[1] != '/' && len(node.children) == 0 {
				// Handle simple inline values (es. {var})
				// loops and blocks-based expressions are handled later
				res = append(res, makeLit(node.value[1:len(node.value)-1]))
				continue
			}

			input := node.value[2 : len(node.value)-1]
			rule := strings.Split(input, " ")[0]
			var (
				tok ast.Node
				extra []ast.Stmt
				err error
			)

			switch rule {
			case "for":
				tok, extra, err = makeFor(node)
				break
			
			case "if":
				tok, extra, err = makeIf(node)
				break

			case "randr":
				tok, extra, err = makeComponent(node)
				break

			default:
				log.Fatal("Unhandled expression: " + node.value)
			}

			if err != nil {
				log.Fatal("Error while parsing complex expr node: " + err.Error())
			}
			res = append(res, tok)
			extras = append(extras, extra...)
		}
	}
	return res, extras
}
