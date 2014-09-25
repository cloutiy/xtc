package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// AddressNode
type AddressNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  t core.IType
}

func NewAddressNode(loc core.Location, expr core.IExprNode) *AddressNode {
  if expr == nil { panic("expr is nil") }
  return &AddressNode { "ast.AddressNode", loc, expr, nil }
}

func (self AddressNode) String() string {
  return fmt.Sprintf("<ast.AddressNode location=%s expr=%s>", self.Location, self.Expr)
}

func (self *AddressNode) AsExprNode() core.IExprNode {
  return self
}

func (self AddressNode) GetLocation() core.Location {
  return self.Location
}

func (self AddressNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self AddressNode) GetType() core.IType {
  if self.t == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.t
}

func (self *AddressNode) SetType(t core.IType) {
  self.t = t
}

func (self AddressNode) IsConstant() bool {
  return false
}

func (self AddressNode) IsParameter() bool {
  return false
}

func (self AddressNode) IsLvalue() bool {
  return false
}

func (self AddressNode) IsAssignable() bool {
  return false
}

func (self AddressNode) IsLoadable() bool {
  return false
}

func (self AddressNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self AddressNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
