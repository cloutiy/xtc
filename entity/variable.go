package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type DefinedVariable struct {
  ClassName string
  Private bool
  Name string
  TypeNode duck.ITypeNode
  NumRefered int
  Initializer duck.IExprNode
}

func NewDefinedVariable(isPrivate bool, t duck.ITypeNode, name string, init duck.IExprNode) DefinedVariable {
  return DefinedVariable { "entity.DefinedVariable", isPrivate, name, t, 0, init }
}

func (self DefinedVariable) String() string {
  return fmt.Sprintf("<entity.DefinedVariable Name=%s Private=%v TypeNode=%s NumRefered=%d Initializer=%s>", self.Name, self.Private, self.TypeNode, self.NumRefered, self.Initializer)
}

func (self DefinedVariable) IsEntity() bool {
  return true
}

func (self DefinedVariable) IsVariable() bool {
  return true
}

func (self DefinedVariable) IsDefinedVariable() bool {
  return true
}

func (self DefinedVariable) HasInitializer() bool {
  return self.Initializer != nil
}

func (self DefinedVariable) IsPrivate() bool {
  return self.Private
}

func (self DefinedVariable) GetName() string {
  return self.Name
}

func (self DefinedVariable) GetTypeNode() duck.ITypeNode {
  return self.TypeNode
}

func (self DefinedVariable) GetNumRefered() int {
  return self.NumRefered
}

func (self DefinedVariable) GetInitializer() duck.IExprNode {
  return self.Initializer
}

type UndefinedVariable struct {
  ClassName string
  Private bool
  Name string
  TypeNode duck.ITypeNode
}

func NewUndefinedVariable(t duck.ITypeNode, name string) UndefinedVariable {
  return UndefinedVariable { "entity.UndefinedVariable", false, name, t }
}

func (self UndefinedVariable) String() string {
  return fmt.Sprintf("<entity.UndefinedVariable Name=%s Private=%v TypeNode=%s>", self.Name, self.Private, self.TypeNode)
}

func (self UndefinedVariable) IsEntity() bool {
  return true
}

func (self UndefinedVariable) IsVariable() bool {
  return true
}

func (self UndefinedVariable) IsUndefinedVariable() bool {
  return true
}
