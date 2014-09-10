package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestMemberNode(t *testing.T) {
  x := NewMemberNode(loc(0,0), NewVariableNode(loc(0,0), "a"), "b")
  s := `{
  "ClassName": "ast.MemberNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Member": "b"
}`
  xt.AssertStringEqualsDiff(t, "MemberNode", xt.JSON(x), s)
}