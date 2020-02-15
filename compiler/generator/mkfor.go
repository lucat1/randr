package generator

import (
	"errors"
	"go/ast"
	"go/token"
	"log"
)

type forParserState int

const (
	forParserIdle = iota
	forParserFor
	forParserKey
	forParserValue
	forParserRange
	forParserSlice
)

// makeFor generates a for loop in golang ast
// from a string, also doing some typechecking
func makeFor(expr *node) (ast.Node, []ast.Stmt, error) {
	_res := random(10)
	res := makeLit(_res)
	input := expr.value[2 : len(expr.value)-1]

	// Extremely rudimental parsing (unfortunately very strict)
	state := forParserIdle
	var forVal, keyVal, valVal, sliceVal string
	for _, c := range input {
		switch c {
		case ',':
			if state == forParserIdle || state == forParserKey {
				state = forParserValue
				break
			}

		case ':':
			if state == forParserIdle {
				state = forParserRange
				break
			}

		case ' ':
			// from for we go into key
			if state == forParserFor {
				state = forParserKey
				break
			}

			// from key we go into val
			if state == forParserKey {
				state = forParserValue
				break
			}

			// from val we go into range
			if state == forParserValue && valVal != "" {
				state = forParserRange
				break
			}

		default:
			// we handle this separately since its not
			// possible to handle it a particular char 'cause
			// it can be also inserted into vars and we wanna
			// catch that
			if c == 'e' && state == forParserRange {
				state = forParserSlice
				break
			}

			// we handle this here because even if it's
			// the start of our `for` keyword we wanna capture this
			if c == 'f' && state == forParserIdle {
				state = forParserFor
			}

			switch state {
			case forParserIdle:
				log.Fatal("For parsing, unhandled: " + string(c))

			case forParserRange:
				// ignore
				break

			case forParserFor:
				forVal += string(c)

			case forParserKey:
				keyVal += string(c)

			case forParserValue:
				valVal += string(c)

			case forParserSlice:
				sliceVal += string(c)
			}
		}
	}

	// kinda dumb checks since our parsing technique is SO bad
	if forVal == "for" && keyVal != "" && valVal != "" && sliceVal != "" {
		// Build the for loop and its body(expr.children if any)
		body := &ast.BlockStmt{}
		if len(expr.children) > 0 {
			// nodes are the values returned by the expressions
			// extras are the generated expressions themselves
			// we lastly just wanna append the `nodes`(values) to
			// our return value (_res)
			nodes, extras := toNodes(expr.children)
			body.List = append(body.List, extras...)
			left := makeIdent(_res)
			for _, node := range nodes {
				right := makeIdent(node.(*ast.BasicLit).Value)
				body.List = append(body.List, makeAssign(left, right))
			}
		}

		return res, []ast.Stmt{
			&ast.DeclStmt{
				Decl: &ast.GenDecl{
					Tok: token.VAR,
					Specs: []ast.Spec{
						&ast.ValueSpec{
							Type: makeIdent("string"),
							Names: []*ast.Ident{
								makeIdent(_res),
							},
						},
					},
				},
			},
			&ast.RangeStmt{
				Key:   makeLit(keyVal).(ast.Expr),
				Value: makeLit(valVal).(ast.Expr),
				X:     makeLit(sliceVal).(ast.Expr),
				Tok:   token.DEFINE,
				Body:  body,
			},
		}, nil
	}

	return nil, nil, errors.New("Invalid for loop: " + input)
}
