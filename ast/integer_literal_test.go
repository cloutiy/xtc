package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestDecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "int",
    "Type": null
  },
  "Value": 12345
}`
  xt.AssertStringEqualsDiff(t, "DecimalIntegerLiteralNode", xt.JSON(x), s)
}

func TestOctalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0755")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "int",
    "Type": null
  },
  "Value": 493
}`
  xt.AssertStringEqualsDiff(t, "OctalIntegerLiteralNode", xt.JSON(x), s)
}

func TestHexadecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0xFFFF")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "int",
    "Type": null
  },
  "Value": 65535
}`
  xt.AssertStringEqualsDiff(t, "HexadecimalIntegerLiteralNode", xt.JSON(x), s)
}

func TestCharacterIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "'a'")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "char",
    "Type": null
  },
  "Value": 97
}`
  xt.AssertStringEqualsDiff(t, "CharacterIntegerLiteralNode", xt.JSON(x), s)
}

func TestUnsignedIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345U")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "unsigned int",
    "Type": null
  },
  "Value": 12345
}`
  xt.AssertStringEqualsDiff(t, "unsigned literal", xt.JSON(x), s)
}

func TestLongIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345L")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "long",
    "Type": null
  },
  "Value": 12345
}`
  xt.AssertStringEqualsDiff(t, "long literal", xt.JSON(x), s)
}

func TestUnsignedLongIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345UL")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "unsigned long",
    "Type": null
  },
  "Value": 12345
}`
  xt.AssertStringEqualsDiff(t, "unsigned long literal", xt.JSON(x), s)
}
