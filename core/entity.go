package core

type IEntity interface {
  String() string
  GetName() string
  IsDefined() bool
  IsPrivate() bool
  IsConstant() bool
  Refered()
  IsRefered() bool
  GetNumRefered() int
  GetTypeNode() ITypeNode
  GetTypeRef() ITypeRef
}

type IFunction interface {
  IEntity
}

type IVariable interface {
  IEntity
}
