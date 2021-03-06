package ast

import (
  "testing"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestLogicalAndNode(t *testing.T) {
  x := NewLogicalAndNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewLogicalAndNode(loc(0,0), NewVariableNode(loc(0,0), "b"), NewVariableNode(loc(0,0), "c")))
  s := `{
  "ClassName": "ast.LogicalAndNode",
  "Location": "[:0,0]",
  "Left": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a",
    "Entity": null
  },
  "Right": {
    "ClassName": "ast.LogicalAndNode",
    "Location": "[:0,0]",
    "Left": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "b",
      "Entity": null
    },
    "Right": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "c",
      "Entity": null
    },
    "Type": null
  },
  "Type": null
}`
  xt.AssertStringEqualsDiff(t, "LogicalAndNode", xt.JSON(x), s)
}
