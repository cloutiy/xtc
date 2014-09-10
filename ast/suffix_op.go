package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// SuffixOpNode
type SuffixOpNode struct {
  ClassName string
  Location core.Location
  Operator string
  Expr core.IExprNode
}

func NewSuffixOpNode(loc core.Location, operator string, expr core.IExprNode) *SuffixOpNode {
  if expr == nil { panic("expr is nil") }
  return &SuffixOpNode { "ast.SuffixOpNode", loc, operator, expr }
}

func (self SuffixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ %s 1)", self.Expr)
    case "--": return fmt.Sprintf("(- %s 1)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

func (self SuffixOpNode) IsExprNode() bool {
  return true
}

func (self SuffixOpNode) GetLocation() core.Location {
  return self.Location
}