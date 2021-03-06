package ast

import (
  "testing"
  "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestDoWhile(t *testing.T) {
/*
  do {
    b(a);
  } while (a < 100);
 */
  x := NewDoWhileNode(
    loc(0,0),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "b"), []core.IExprNode { NewVariableNode(loc(0,0), "a") })),
    NewBinaryOpNode(loc(0,0), "<", NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "100")),
  )
  s := `{
  "ClassName": "ast.DoWhileNode",
  "Location": "[:0,0]",
  "Body": {
    "ClassName": "ast.ExprStmtNode",
    "Location": "[:0,0]",
    "Expr": {
      "ClassName": "ast.FuncallNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.VariableNode",
        "Location": "[:0,0]",
        "Name": "b",
        "Entity": null
      },
      "Args": [
        {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "a",
          "Entity": null
        }
      ]
    }
  },
  "Cond": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "\u003c",
    "Left": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "a",
      "Entity": null
    },
    "Right": {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int",
        "Type": null
      },
      "Value": 100
    },
    "Type": null
  }
}`
  xt.AssertStringEqualsDiff(t, "DoWhileNode", xt.JSON(x), s)
}
