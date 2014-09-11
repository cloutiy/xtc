package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

// IntegerLiteralNode
type IntegerLiteralNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  Value int
}

func NewIntegerLiteralNode(loc core.Location, literal string) *IntegerLiteralNode {
  var ref core.ITypeRef
  var value int
  var err error
  if ( startsWith(literal, "'") && endsWith(literal, "'") ) && 2 < len(literal) {
    _, err = fmt.Sscanf(literal[1:len(literal)-1], "%c", &value)
    ref = typesys.NewIntegerTypeRef(loc, "char")
  } else {
    if ( startsWith(literal, "0X") || startsWith(literal, "0x") ) && 2 < len(literal) {
      // hexadecimal
      _, err = fmt.Sscanf(literal[2:], "%x", &value)
    } else {
      if startsWith(literal, "0") && 1 < len(literal) {
        // octal
        _, err = fmt.Sscanf(literal[1:], "%o", &value)
      } else {
        // decimal
        _, err = fmt.Sscanf(literal, "%d", &value)
      }
    }
    if endsWith(literal, "UL") {
      ref = typesys.NewIntegerTypeRef(loc, "unsigned long")
    } else if endsWith(literal, "L") {
      ref = typesys.NewIntegerTypeRef(loc, "long")
    } else if endsWith(literal, "U") {
      ref = typesys.NewIntegerTypeRef(loc, "unsigned int")
    } else {
      ref = typesys.NewIntegerTypeRef(loc, "int")
    }
  }
  if err != nil {
    panic(err)
  }
  return &IntegerLiteralNode { "ast.IntegerLiteralNode", loc, NewTypeNode(loc, ref), value }
}

func startsWith(s, prefix string) bool {
  return len(prefix) <= len(s) && strings.Index(s, prefix) == 0
}

func endsWith(s, suffix string) bool {
  return len(suffix) <= len(s) && strings.LastIndex(s, suffix) == len(s)-len(suffix)
}

func (self IntegerLiteralNode) String() string {
  return fmt.Sprintf("%d", self.Value)
}

func (self IntegerLiteralNode) IsExprNode() bool {
  return true
}

func (self IntegerLiteralNode) GetLocation() core.Location {
  return self.Location
}

func (self IntegerLiteralNode) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}
