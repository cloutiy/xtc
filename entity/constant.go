package entity

import (
  "encoding/json"
  "bitbucket.org/yyuu/bs/duck"
)

type Constant struct {
  Name string
  TypeNode duck.ITypeNode
  Value duck.IExprNode
}

func NewConstant(t duck.ITypeNode, name string, value duck.IExprNode) Constant {
  return Constant {
    TypeNode: t,
    Name: name,
    Value: value,
  }
}

func (self Constant) String() string {
  panic("Constant#String called")
}

func (self Constant) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Name string
    TypeNode duck.ITypeNode
    Value duck.IExprNode
  }
  x.ClassName = "entity.Constant"
  x.Name = self.Name
  x.TypeNode = self.TypeNode
  x.Value = self.Value
  return json.Marshal(x)
}

func (self Constant) IsEntity() bool {
  return true
}
