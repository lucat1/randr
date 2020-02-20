package parser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/net/html"
)

var curly = regexp.MustCompile("\\{(.*?)\\}")

// Parse returns a golang ast from a string
// containing an html template and some expressions
func Parse(input string) ([]string, []string, error) {
	input = input[1 : len(input)-1]  // remove trailing `` or ""
	input = strings.TrimSpace(input) // remove trailing spaces

	// TODO: Handle <body> inside templates which is currently ignored
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return nil, nil, err
	}
	body := doc.FirstChild.FirstChild.NextSibling
	raws, exprs := loop(body, false, []string{}, []string{})

	return raws, exprs, Validate(raws, exprs)
}

func loop(parent *html.Node, strict bool, raws []string, exprs []string) ([]string, []string) {
	for e := parent.FirstChild; e != nil; e = e.NextSibling {
		switch e.Type {
		case html.TextNode:
			raws, exprs = Text(e.Data, strict, raws, exprs)
			break

		case html.ElementNode:
			tag := e.Data
			if !IsTagValid(tag) {
				var finalTag string
				if !strings.Contains(tag, ".") {
					// Simple tag from the same package
					finalTag = strcase.ToCamel(tag)
				} else {
					parts := strings.Split(tag, ".")
					scopes, name := parts[:len(parts)-1], parts[len(parts)-1]
					// Capitalize only the name
					name = strcase.ToCamel(name)
					finalTag = strings.Join(append(scopes, name), ".")
				}

				attrs, _ := json.Marshal(e.Attr)
				exprs = append(exprs, "{#randr "+finalTag+" "+string(attrs)+"}")
				raws = append(raws, "")
				raws, exprs = loop(e, strict || tag == "pre", raws, exprs)
				exprs, raws = append(exprs, "{/randr}"), append(raws, "")
				continue
			}

			// Append the static string to the previous
			// or create a new one if something has been
			// inserted in between
			raws, exprs = appendOpeningTag(e, raws, exprs)
			raws, exprs = loop(e, strict || e.Data == "pre", raws, exprs)
			raws, exprs = appendClosingTag(e, raws, exprs)
			break
		}
	}

	return raws, exprs
}

// Text handles the parsing of text values
func Text(input string, strict bool, raws []string, exprs []string) ([]string, []string) {
	// Remove \n and \t if we arent side a pre element
	// to check that out we use the `strict` argument
	if !strict {
		input = strings.TrimSpace(input)
	}

	matches := curly.FindAll([]byte(input), -1) // Find all expressions
	if len(matches) > 0 {
		raws, exprs = literate(input, matches, strict, raws, exprs) // Split static strings and expressions
	} else {
		raws, exprs = addToRaw(input, strict, raws, exprs)
	}
	return raws, exprs
}

// literate transforms the input into a slice
// resenbling javascript's template literals.
func literate(input string, matches [][]byte, strict bool, raws []string, exprs []string) ([]string, []string) {
	for i, match := range matches {
		// Map experssions to the exprs slice
		// and order raw strings with them
		parts := strings.SplitN(input, string(match), 2)
		raws, exprs = addToRaw(parts[0], strict, raws, exprs)

		if len(parts) > 1 {
			exprs = append(exprs, string(match))
			if parts[1] != "" {
				input = parts[1] // Dont reiterate on the already parsed part of the string
				if i == len(matches)-1 {
					// finally append the remaining of the string (if any)
					// to the raws array
					raws, exprs = addToRaw(parts[1], strict, raws, exprs)
					//raws = append(raws, parts[1])
				}
			}
		}
	}

	return raws, exprs
}

// addToRaw append the string to the raws array
// in two ways depending on the length of the exprs array
func addToRaw(tba string, strict bool, raws []string, exprs []string) ([]string, []string) {
	if !strict {
		// Trim useless spaces to be able to do some
		// cleaner code output afterwards
		tba = strings.Replace(strings.Replace(tba, "\t", "", -1), "\n", "", -1)
	}

	if len(exprs) >= len(raws) {
		// Something new has been added in between
		// we append it to the array
		raws = append(raws, tba)
	} else {
		// Nothing has been added, we can just append
		// it to the previous string
		raws[len(raws)-1] += tba
	}

	return raws, exprs
}

// appendOpeningTag appends the opening tag element
// and all its attributes preparing the the children
// to be handled in loop afterwards
func appendOpeningTag(node *html.Node, raws []string, exprs []string) ([]string, []string) {
	raws, exprs = addToRaw("<"+node.Data, true, raws, exprs)
	for _, attr := range node.Attr {
		fmt.Println(attr.Val)
		if attr.Val == "" {
			// boolean value, true by default
			raws, exprs = addToRaw(
				" "+attr.Key,
				true, raws, exprs,
			)
			continue
		}

		raws, exprs = addToRaw(
			" "+attr.Key+"=\"",
			true, raws, exprs,
		)
		raws, exprs = Text(attr.Val, true, raws, exprs)

		raws, exprs = addToRaw("\"", true, raws, exprs)
	}
	raws, exprs = addToRaw(">", true, raws, exprs)

	return raws, exprs
}

// appendClosingTag appends the ending part of
// the tag after all children have been handled
func appendClosingTag(node *html.Node, raws []string, exprs []string) ([]string, []string) {
	return addToRaw("</"+node.Data+">", true, raws, exprs)
}
