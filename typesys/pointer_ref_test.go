package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestPointerTypeRefToString1(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewPointerTypeRef(NewCharTypeRef(loc))
  xt.AssertEquals(t, "char*", x.String(), "char*")
}

func TestPointerTypeRefToString2(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewPointerTypeRef(NewPointerTypeRef(NewCharTypeRef(loc)))
  xt.AssertEquals(t, "char**", x.String(), "char**")
}