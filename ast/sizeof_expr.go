package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// SizeofExprNode
type SizeofExprNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  TypeNode core.ITypeNode
}

func NewSizeofExprNode(loc core.Location, expr core.IExprNode, t core.ITypeRef) *SizeofExprNode {
  if expr == nil { panic("expr is nil") }
  if t == nil { panic("t is nil") }
  return &SizeofExprNode { "ast.SizeofExprNode", loc, expr, NewTypeNode(loc, t) }
}

func (self SizeofExprNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.Expr)
}

func (self SizeofExprNode) IsExprNode() bool {
  return true
}

func (self SizeofExprNode) GetLocation() core.Location {
  return self.Location
}