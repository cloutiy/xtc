package entity

import (
  "bitbucket.org/yyuu/xtc/core"
)

type Parameter struct {
  *DefinedVariable
}

func NewParameter(t core.ITypeNode, name string) *Parameter {
  definedVariable := NewDefinedVariable(true, t, name, nil)
  definedVariable.ClassName = "entity.Parameter"
  return &Parameter { definedVariable }
}

func NewParameters(xs...*Parameter) []*Parameter {
  if 0 < len(xs) {
    return xs
  } else {
    return []*Parameter { }
  }
}

func AsParameter(x core.IEntity) *Parameter {
  return x.(*Parameter)
}

func (self *Parameter) String() string {
  return self.DefinedVariable.String()
}

func (self *Parameter) IsPrivate() bool {
  return false
}

func (self *Parameter) IsConstant() bool {
  return false
}

func (self *Parameter) IsParameter() bool {
  return true
}

func (self *Parameter) IsVariable() bool {
  return true
}

func (self *Parameter) GetInitializer() core.IExprNode {
  return nil
}

func (self *Parameter) SetInitializer(e core.IExprNode) {
  // noop
}

func (self *Parameter) HasInitializer() bool {
  return false
}

func (self *Parameter) GetNumRefered() int {
  return self.DefinedVariable.GetNumRefered()
}

func (self *Parameter) IsRefered() bool {
  return self.DefinedVariable.IsRefered()
}

func (self *Parameter) Refered() {
  self.DefinedVariable.Refered()
}

func (self *Parameter) IsDefined() bool {
  return true
}

func (self *Parameter) GetTypeNode() core.ITypeNode {
  return self.DefinedVariable.TypeNode
}

func (self *Parameter) GetTypeRef() core.ITypeRef {
  return self.DefinedVariable.TypeNode.GetTypeRef()
}

func (self *Parameter) GetType() core.IType {
  return self.DefinedVariable.TypeNode.GetType()
}

func (self *Parameter) GetName() string {
  return self.DefinedVariable.Name
}

func (self *Parameter) GetValue() core.IExprNode {
  panic("Parameter#GetValue called")
}

func (self *Parameter) SymbolString() string {
  return self.DefinedVariable.SymbolString()
}

func (self *Parameter) GetMemref() core.IMemoryReference {
  return self.DefinedVariable.GetMemref()
}

func (self *Parameter) SetMemref(memref core.IMemoryReference) {
  self.DefinedVariable.SetMemref(memref)
}

func (self *Parameter) GetAddress() core.IOperand {
  return self.DefinedVariable.GetAddress()
}

func (self *Parameter) SetAddress(address core.IOperand) {
  self.DefinedVariable.SetAddress(address)
}
