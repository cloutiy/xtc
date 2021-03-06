package ast

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

// CondExprNode
type CondExprNode struct {
  ClassName string
  Location core.Location
  Cond core.IExprNode
  ThenExpr core.IExprNode
  ElseExpr core.IExprNode
  Type core.IType
}

func NewCondExprNode(loc core.Location, cond core.IExprNode, thenExpr core.IExprNode, elseExpr core.IExprNode) *CondExprNode {
  if cond == nil { panic("cond is nil") }
  if thenExpr == nil { panic("thenExpr is nil") }
  if elseExpr == nil { panic("elseExpr is nil") }
  return &CondExprNode { "ast.CondExprNode", loc, cond, thenExpr, elseExpr, nil }
}

func (self CondExprNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenExpr, self.ElseExpr)
}

func (self *CondExprNode) AsExprNode() core.IExprNode {
  return self
}

func (self CondExprNode) GetLocation() core.Location {
  return self.Location
}

func (self *CondExprNode) GetCond() core.IExprNode {
  return self.Cond
}

func (self *CondExprNode) GetThenExpr() core.IExprNode {
  return self.ThenExpr
}

func (self *CondExprNode) SetThenExpr(expr core.IExprNode) {
  self.ThenExpr = expr
}

func (self *CondExprNode) GetElseExpr() core.IExprNode {
  return self.ElseExpr
}

func (self *CondExprNode) SetElseExpr(expr core.IExprNode) {
  self.ElseExpr = expr
}

func (self *CondExprNode) GetType() core.IType {
  if self.Type == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.Type
}

func (self *CondExprNode) SetType(t core.IType) {
  if self.Type != nil {
    panic("#SetType called twice")
  }
  self.Type = t
}

func (self *CondExprNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self *CondExprNode) IsConstant() bool {
  return false
}

func (self *CondExprNode) IsParameter() bool {
  return false
}

func (self *CondExprNode) IsLvalue() bool {
  return false
}

func (self *CondExprNode) IsAssignable() bool {
  return false
}

func (self *CondExprNode) IsLoadable() bool {
  return false
}

func (self *CondExprNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self *CondExprNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
