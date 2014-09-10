package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/typesys"
)

// StructNode
type StructNode struct {
  ClassName string
  Location duck.ILocation
  TypeNode duck.ITypeNode
  Name string
  Members []Slot
}

func NewStructNode(loc duck.ILocation, ref duck.ITypeRef, name string, membs []Slot) StructNode {
  if loc == nil { panic("location is nil") }
  if ref == nil { panic("ref is nil") }
  return StructNode { "ast.StructNode", loc, NewTypeNode(loc, ref), name, membs }
}

func (self StructNode) String() string {
  return fmt.Sprintf("<ast.StructNode Name=%s location=%s typeNode=%s members=%s>", self.Name, self.Location, self.TypeNode, self.Members)
}

func (self StructNode) IsTypeDefinition() bool {
  return true
}

func (self StructNode) GetLocation() duck.ILocation {
  return self.Location
}

func (self StructNode) GetTypeRef() duck.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self StructNode) DefiningType() duck.IType {
  var membs []duck.ISlot
  for i := range self.Members {
    membs[i] = self.Members[i]
  }
  return typesys.NewStructType(self.Name, membs, self.Location)
}

// UnionNode
type UnionNode struct {
  ClassName string
  Location duck.ILocation
  TypeNode duck.ITypeNode
  Name string
  Members []Slot
}

func NewUnionNode(loc duck.ILocation, ref duck.ITypeRef, name string, membs []Slot) UnionNode {
  if loc == nil { panic("location is nil") }
  if ref == nil { panic("ref is nil") }
  return UnionNode { "ast.UnionNode", loc, NewTypeNode(loc, ref), name, membs }
}

func (self UnionNode) String() string {
  return fmt.Sprintf("<ast.UnionNode name=%s location=%s typeNode=%s members=%s>", self.Name, self.Location, self.TypeNode, self.Members)
}

func (self UnionNode) IsTypeDefinition() bool {
  return true
}

func (self UnionNode) GetLocation() duck.ILocation {
  return self.Location
}

func (self UnionNode) GetTypeRef() duck.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self UnionNode) DefiningType() duck.IType {
  var membs []duck.ISlot
  for i := range self.Members {
    membs[i] = self.Members[i]
  }
  return typesys.NewUnionType(self.Name, membs, self.Location)
}
