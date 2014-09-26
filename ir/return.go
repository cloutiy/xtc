package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Return struct {
  ClassName string
  Location core.Location
  Expr core.IExpr
}

func NewReturn(loc core.Location, expr core.IExpr) *Return {
  return &Return { "ir.Return", loc, expr }
}

func (self Return) AsStmt() core.IStmt {
  return self
}

func (self Return) GetLocation() core.Location {
  return self.Location
}