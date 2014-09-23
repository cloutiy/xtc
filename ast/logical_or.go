package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// LogicalOrNode
type LogicalOrNode struct {
  ClassName string
  Location core.Location
  Left core.IExprNode
  Right core.IExprNode
  t core.IType
}

func NewLogicalOrNode(loc core.Location, left core.IExprNode, right core.IExprNode) *LogicalOrNode {
  if left == nil { panic("left is nil") }
  if right == nil { panic("right is nil") }
  return &LogicalOrNode { "ast.LogicalOrNode", loc, left, right, nil }
}

func (self LogicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
}

func (self LogicalOrNode) IsExprNode() bool {
  return true
}

func (self LogicalOrNode) GetLocation() core.Location {
  return self.Location
}

func (self LogicalOrNode) GetLeft() core.IExprNode {
  return self.Left
}

func (self LogicalOrNode) GetRight() core.IExprNode {
  return self.Right
}

func (self LogicalOrNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *LogicalOrNode) SetType(t core.IType) {
  self.t = t
}

func (self LogicalOrNode) IsConstant() bool {
  return false
}

func (self LogicalOrNode) IsParameter() bool {
  return false
}

func (self LogicalOrNode) IsLvalue() bool {
  return false
}

func (self LogicalOrNode) IsAssignable() bool {
  return false
}

func (self LogicalOrNode) IsLoadable() bool {
  return false
}

func (self LogicalOrNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self LogicalOrNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
