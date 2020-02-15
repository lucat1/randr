package parser

import (
	"go/ast"
	"go/token"
	"log"
	"strings"
)

// CheckComment checks the golang program if it
// contains a `// +build ignore` comment, required
// to make the file ignored, so that we can output the
// compiled file and be sure it worn't collide with the
// source file
func CheckComment(node *ast.File, fset *token.FileSet) {
	found := false
	for i, comment := range node.Comments {
		if strings.TrimSpace(comment.Text()) == "+build ignore" {
			found = true

			// Remove the comment from the final file
			node.Comments = append(node.Comments[:i], node.Comments[i+1:]...)
		}
	}

	if !found {
		log.Fatal(`You must append a comment stating "// +build ignore" at the top
		of your file in order to prevent it from being compiled, so that
		the compiler can output the processed component result`)
	}
}