package generator

import (
	"go/ast"
	"go/token"
)

func makeAssign(left ast.Expr, right ast.Expr) ast.Stmt {
	return &ast.AssignStmt{
		Tok: token.ADD_ASSIGN,
		Lhs: []ast.Expr{left},
		Rhs: []ast.Expr{right},
	}
}