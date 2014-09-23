package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// ArefNode
type ArefNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Index core.IExprNode
  t core.IType
}

func NewArefNode(loc core.Location, expr core.IExprNode, index core.IExprNode) *ArefNode {
  if expr == nil { panic("expr is nil") }
  if index == nil { panic("index is nil") }
  return &ArefNode { "ast.ArefNode", loc, expr, index, nil }
}

func (self ArefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

func (self ArefNode) IsExprNode() bool {
  return true
}

func (self ArefNode) GetLocation() core.Location {
  return self.Location
}

func (self ArefNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self ArefNode) GetIndex() core.IExprNode {
  return self.Index
}

func (self ArefNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *ArefNode) SetType(t core.IType) {
  self.t = t
}

func (self ArefNode) IsConstant() bool {
  return false
}

func (self ArefNode) IsParameter() bool {
  return false
}

func (self ArefNode) IsLvalue() bool {
  return true
}

func (self ArefNode) IsAssignable() bool {
  return true
}

func (self ArefNode) IsLoadable() bool {
  return false
}

func (self ArefNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self ArefNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
