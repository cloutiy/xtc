package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/core"
)

// CaseNode
type CaseNode struct {
  ClassName string
  Location core.Location
  Values []core.IExprNode
  Body core.IStmtNode
}

func NewCaseNode(loc core.Location, values []core.IExprNode, body core.IStmtNode) *CaseNode {
  if body == nil { panic("body is nil") }
  return &CaseNode { "ast.CaseNode", loc, values, body }
}

func (self CaseNode) String() string {
  sValues := make([]string, len(self.Values))
  for i := range self.Values {
    sValues[i] = fmt.Sprintf("(= switch-cond %s)", self.Values[i])
  }
  switch len(sValues) {
    case 0:  return fmt.Sprintf("(else %s)", self.Body)
    case 1:  return fmt.Sprintf("(%s %s)", sValues[0], self.Body)
    default: return fmt.Sprintf("((or %s) %s)", strings.Join(sValues, " "), self.Body)
  }
}

func (self CaseNode) IsStmtNode() bool {
  return true
}

func (self CaseNode) GetLocation() core.Location {
  return self.Location
}

func (self CaseNode) GetValues() []core.IExprNode {
  return self.Values
}

func (self CaseNode) GetBody() core.IStmtNode {
  return self.Body
}