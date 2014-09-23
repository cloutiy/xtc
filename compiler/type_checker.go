package compiler

import (
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

type TypeChecker struct {
  errorHandler *core.ErrorHandler
  typeTable *typesys.TypeTable
  currentFunction *entity.DefinedFunction
}

func NewTypeChecker(errorHandler *core.ErrorHandler, table *typesys.TypeTable) *TypeChecker {
  return &TypeChecker { errorHandler, table, nil }
}

func (self *TypeChecker) Check(a *ast.AST) {
  vs := a.GetDefinedVariables()
  for i := range vs {
    self.checkVariable(vs[i])
  }
  fs := a.GetDefinedFunctions()
  for i := range fs {
    self.currentFunction = fs[i]
    self.checkReturnType(fs[i])
    self.checkParamTypes(fs[i])
    ast.VisitStmt(self, fs[i].GetBody())
  }
}

func (self *TypeChecker) checkVariable(v *entity.DefinedVariable) {
  if self.isInvalidVariableType(v.GetType()) {
    self.errorHandler.Errorf("invalid variable type\n")
  }
  if v.HasInitializer() {
    if self.isInvalidLHSType(v.GetType()) {
      self.errorHandler.Errorf("invalid LHS type: %s\n", v.GetType())
    }
    ast.VisitExpr(self, v.GetInitializer())
    v.SetInitializer(self.implicitCast(v.GetType(), v.GetInitializer()))
  }
}

func (self *TypeChecker) isInvalidVariableType(t core.IType) bool {
  return t.IsVoid() || (t.IsArray() && ! t.IsAllocatedArray())
}

func (self *TypeChecker) isInvalidLHSType(t core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsVoid() || t.IsArray()
}

func (self *TypeChecker) isInvalidRHSType(t core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsVoid()
}

func (self *TypeChecker) implicitCast(t core.IType, expr core.IExprNode) core.IExprNode {
  if expr.GetType().IsSameType(t) {
    return expr
  } else {
    if expr.GetType().IsCastableTo(t) {
      if ! expr.GetType().IsCompatible(t) && ! self.isSafeIntegerCast(expr, t) {
        self.errorHandler.Warnf("%s incompatible inplicit cast from %s to %s\n", expr.GetLocation(), expr.GetType(), t)
      }
      typeNode := ast.NewTypeNode(expr.GetLocation(), typesys.NewVoidTypeRef(expr.GetLocation()))
      typeNode.SetType(t)
      return ast.NewCastNode(expr.GetLocation(), typeNode, expr)
    } else {
      self.errorHandler.Errorf("invalid cast error: %s to %s", expr.GetType(), t)
      return expr
    }
  }
}

func (self *TypeChecker) isSafeIntegerCast(node core.INode, t core.IType) bool {
  if ! t.IsInteger() {
    return false
  } else {
    i, ok := t.(typesys.IntegerType)
    if ! ok {
      return false
    }
    n, ok := node.(ast.IntegerLiteralNode)
    if ! ok {
      return false
    }
    return i.IsInDomain(n.GetValue())
  }
}

func (self *TypeChecker) checkReturnType(f *entity.DefinedFunction) {
  if self.isInvalidReturnType(f.GetReturnType()) {
    self.errorHandler.Errorf("returns invalid type: %s", f.GetReturnType())
  }
}

func (self *TypeChecker) isInvalidReturnType(t core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsArray()
}

func (self *TypeChecker) checkParamTypes(f *entity.DefinedFunction) {
  params := f.GetParameters()
  for i := range params {
    param := params[i]
    if self.isInvalidParameterType(param.GetType()) {
      self.errorHandler.Errorf("invalid parameter type: %s", param.GetType())
    }
  }
}

func (self *TypeChecker) isInvalidParameterType(t core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsVoid() || t.IsIncompleteArray()
}

func (self *TypeChecker) isInvalidStatementType(t core.IType) bool {
  return t.IsStruct() || t.IsUnion()
}

func (self *TypeChecker) mustBeInteger(expr core.IExprNode, op string) bool {
  if ! expr.GetType().IsInteger() {
    self.errorHandler.Errorf("%s wrong operand type for %s: %s\n", expr.GetLocation(), op, expr.GetType())
    return false
  } else {
    return true
  }
}

func (self *TypeChecker) mustBeScalar(expr core.IExprNode, op string) bool {
  if ! expr.GetType().IsScalar() {
    self.errorHandler.Errorf("%s wrong operand type for %s: %s\n", expr.GetLocation(), op, expr.GetType())
    return false
  } else {
    return true
  }
}

func (self *TypeChecker) checkCond(cond core.IExprNode) {
  self.mustBeScalar(cond, "condition expression")
}

func (self *TypeChecker) expectsComparableScalars(node core.IBinaryOpNode) {
  if ! self.mustBeScalar(node.GetLeft(), node.GetOperator()) {
    return
  }
  if ! self.mustBeScalar(node.GetRight(), node.GetOperator()) {
    return
  }
  if node.GetLeft().GetType().IsPointer() {
    right := self.forcePointerType(node.GetLeft(), node.GetRight())
    node.SetRight(right)
    node.SetType(node.GetLeft().GetType())
    return
  }
  if node.GetRight().GetType().IsPointer() {
    left := self.forcePointerType(node.GetRight(), node.GetLeft())
    node.SetLeft(left)
    node.SetType(node.GetRight().GetType())
    return
  }
  self.arithmeticImplicitCast(node)
}

func (self *TypeChecker) forcePointerType(master core.IExprNode, slave core.IExprNode) core.IExprNode {
  if master.GetType().IsCompatible(slave.GetType()) {
    return slave
  } else {
    self.errorHandler.Warnf("incompatible implicit cast from %s to %s\n", slave.GetType(), master.GetType())
    typeNode := ast.NewTypeNode(master.GetLocation(), typesys.NewVoidTypeRef(master.GetLocation()))
    typeNode.SetType(master.GetType())
    return ast.NewCastNode(master.GetLocation(), typeNode, slave)
  }
}

func (self *TypeChecker) arithmeticImplicitCast(node core.IBinaryOpNode) {
  r := self.integralPromotion(node.GetRight().GetType())
  l := self.integralPromotion(node.GetLeft().GetType())
  target := self.usualArithmeticConversion(l, r)
  if ! l.IsSameType(target) {
    typeNode := ast.NewTypeNode(node.GetLocation(), typesys.NewVoidTypeRef(node.GetLocation()))
    node.SetLeft(ast.NewCastNode(node.GetLocation(), typeNode, node.GetLeft()))
  }
  if ! r.IsSameType(target) {
    typeNode := ast.NewTypeNode(node.GetLocation(), typesys.NewVoidTypeRef(node.GetLocation()))
    node.SetLeft(ast.NewCastNode(node.GetLocation(), typeNode, node.GetRight()))
  }
  node.SetType(target)
}

func (self *TypeChecker) integralPromotion(t core.IType) core.IType {
  if ! t.IsInteger() {
    self.errorHandler.Errorf("integral promotion for %s\n", t)
  }
  intType := self.typeTable.SignedInt()
  if t.Size() < intType.Size() {
    return intType
  } else {
    return t
  }
}

func (self *TypeChecker) integralPromotedExpr(expr core.IExprNode) core.IExprNode {
  t := self.integralPromotion(expr.GetType())
  if t.IsSameType(expr.GetType()) {
    return expr
  } else {
    typeNode := ast.NewTypeNode(expr.GetLocation(), typesys.NewVoidTypeRef(expr.GetLocation()))
    return ast.NewCastNode(expr.GetLocation(), typeNode, expr)
  }
}

func (self *TypeChecker) usualArithmeticConversion(l core.IType, r core.IType) core.IType {
  s_int := self.typeTable.SignedInt()
  u_int := self.typeTable.UnsignedInt()
  s_long := self.typeTable.SignedLong()
  u_long := self.typeTable.UnsignedLong()
  if (l.IsSameType(u_int) && r.IsSameType(s_long)) || (r.IsSameType(u_int) && l.IsSameType(s_long)) {
    return u_long
  } else {
    if l.IsSameType(u_long) || r.IsSameType(u_long) {
      return u_long
    } else {
      if l.IsSameType(s_long) || r.IsSameType(s_long) {
        return s_long
      } else {
        if l.IsSameType(u_int) || r.IsSameType(u_int) {
          return u_int
        } else {
          return s_int
        }
      }
    }
  }
}

func (self *TypeChecker) expectsScalarLHS(node core.IUnaryArithmeticOpNode) {
  if node.GetExpr().IsParameter() {
    // parameter is always a scalar.
  } else {
    if node.GetExpr().GetType().IsArray() {
      self.errorHandler.Errorf("%s wrong operand type for %s: %s\n", node.GetLocation(), node.GetOperator(), node.GetExpr())
      return
    } else {
      self.mustBeScalar(node.GetExpr(), node.GetOperator())
    }
  }
  if node.GetExpr().GetType().IsInteger() {
    opType := self.integralPromotion(node.GetExpr().GetType())
    if ! node.GetExpr().GetType().IsSameType(opType) {
      node.SetOpType(opType)
    }
    node.SetAmount(1)
  } else {
    if node.GetExpr().GetType().IsPointer() {
      if node.GetExpr().GetType().GetBaseType().IsVoid() {
        self.errorHandler.Errorf("%s wrong operand type for %s: %s\n", node.GetLocation(), node.GetOperator(), node.GetExpr())
        return
      }
      node.SetAmount(node.GetExpr().GetType().GetBaseType().Size())
    } else {
      panic("must not happen")
    }
  }
}

func (self *TypeChecker) checkLHS(lhs core.IExprNode) bool {
  if lhs.IsParameter() {
    // parameter is always assignable.
    return true
  } else {
    if self.isInvalidLHSType(lhs.GetType()) {
      self.errorHandler.Errorf("%s invalid LHS expression type: %s\n", lhs.GetLocation(), lhs.GetType())
      return false
    } else {
      return true
    }
  }
}

func (self *TypeChecker) checkRHS(rhs core.IExprNode) bool {
  if self.isInvalidRHSType(rhs.GetType()) {
    self.errorHandler.Errorf("%s invalid RHS expression type: %s\n", rhs.GetLocation(), rhs.GetType())
    return false
  } else {
    return true
  }
}

func (self *TypeChecker) expectsSameIntegerOrPointerDiff(node core.IBinaryOpNode) {
  if node.GetLeft().IsPointer() && node.GetRight().IsPointer() {
    if node.GetOperator() == "+" {
      self.errorHandler.Errorf("%s invalid operation: pointer + pointer\n", node.GetLocation())
      return
    }
    node.SetType(self.typeTable.PtrDiffType())
  }
}

func (self *TypeChecker) expectsSameInteger(node core.IBinaryOpNode) {
  if ! self.mustBeInteger(node.GetLeft(), node.GetOperator()) {
    return
  }
  if ! self.mustBeInteger(node.GetRight(), node.GetOperator()) {
    return
  }
  self.arithmeticImplicitCast(node)
}

func (self *TypeChecker) VisitNode(unknown core.INode) {
  switch node := unknown.(type) {
    case *ast.BlockNode: {
      vars := node.GetVariables()
      for i := range vars {
        self.checkVariable(vars[i])
      }
      ast.VisitStmts(self, node.GetStmts())
    }
    case *ast.ExprStmtNode: {
      ast.VisitExpr(self, node.GetExpr())
      if self.isInvalidStatementType(node.GetExpr().GetType()) {
        self.errorHandler.Errorf("%s invalid statement type: %s\n", node.GetLocation(), node.GetExpr().GetType())
      }
    }
    case *ast.IfNode: {
      visitIfNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *ast.WhileNode: {
      visitWhileNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *ast.ForNode: {
      visitForNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *ast.SwitchNode: {
      visitSwitchNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *ast.ReturnNode: {
      visitReturnNode(self, node)
      if self.currentFunction.IsVoid() {
        if node.GetExpr() != nil {
          self.errorHandler.Errorf("%s returning value from void function\n", node.GetLocation())
        }
        if node.GetExpr().GetType().IsVoid() {
          self.errorHandler.Errorf("%s returning void\n", node.GetLocation())
        }
        node.SetExpr(self.implicitCast(self.currentFunction.GetReturnType(), node.GetExpr()))
      }
    }
    case *ast.AssignNode: {
      visitAssignNode(self, node)
      if ! self.checkLHS(node.GetLhs()) {
        return
      }
      if ! self.checkRHS(node.GetRhs()) {
        return
      }
      node.SetRhs(self.implicitCast(node.GetLhs().GetType(), node.GetRhs()))
    }
    case *ast.OpAssignNode: {
      visitOpAssignNode(self, node)
      if ! self.checkLHS(node.GetLhs()) {
        return
      }
      if ! self.checkRHS(node.GetRhs()) {
        return
      }
      if node.GetOperator() == "+" || node.GetOperator() == "-" {
        if node.GetLhs().GetType().IsPointer() {
          self.mustBeInteger(node.GetRhs(), node.GetOperator())
          node.SetRhs(self.integralPromotedExpr(node.GetRhs()))
          return
        }
      }
      if ! self.mustBeInteger(node.GetLhs(), node.GetOperator()) {
        return
      }
      if ! self.mustBeInteger(node.GetRhs(), node.GetOperator()) {
        return
      }
      l := self.integralPromotion(node.GetLhs().GetType())
      r := self.integralPromotion(node.GetRhs().GetType())
      opType := self.usualArithmeticConversion(l, r)
      if ! opType.IsCompatible(l) && self.isSafeIntegerCast(node.GetRhs(), opType) {
        self.errorHandler.Warnf("%s incompatible implicit cast from %s to %s\n", node.GetLocation(), opType, l)
      }
      if ! r.IsSameType(opType) {
        typeNode := ast.NewTypeNode(node.GetLocation(), typesys.NewVoidTypeRef(node.GetLocation()))
        node.SetRhs(ast.NewCastNode(node.GetLocation(), typeNode, node.GetRhs()))
      }
    }
    case *ast.CondExprNode: {
      visitCondExprNode(self, node)
      self.checkCond(node.GetCond())
      t := node.GetThenExpr().GetType()
      e := node.GetElseExpr().GetType()
      if ! t.IsSameType(e) {
        if t.IsCompatible(e) {
          // insert cast on thenBody
          typeNode := ast.NewTypeNode(node.GetLocation(), typesys.NewVoidTypeRef(node.GetLocation()))
          cast := ast.NewCastNode(node.GetLocation(), typeNode, node.GetThenExpr())
          node.SetThenExpr(cast)
        } else {
          if e.IsCompatible(t) {
            // insert cast on elseBody
            typeNode := ast.NewTypeNode(node.GetLocation(), typesys.NewVoidTypeRef(node.GetLocation()))
            cast := ast.NewCastNode(node.GetLocation(), typeNode, node.GetElseExpr())
            node.SetElseExpr(cast)
          } else {
            self.errorHandler.Errorf("%s invalid cast from %s to %s\n", node.GetLocation(), e, t)
          }
        }
      }
    }
    case *ast.BinaryOpNode: {
      visitBinaryOpNode(self, node)
      if node.GetOperator() == "+" || node.GetOperator() == "-" {
        self.expectsSameIntegerOrPointerDiff(node)
      } else {
        switch node.GetOperator() {
          case "*", "/", "%", "&", "|", "^", "<<", ">>": {
            self.expectsSameInteger(node)
          }
          case "==", "!=", "<", "<=", ">", ">=": {
            self.expectsComparableScalars(node)
          }
          default: {
            panic(fmt.Errorf("unknown binary operator: %s", node.GetOperator()))
          }
        }
      }
    }
    case *ast.LogicalAndNode: {
      visitLogicalAndNode(self, node)
      self.expectsComparableScalars(node)
    }
    case *ast.LogicalOrNode: {
      visitLogicalOrNode(self, node)
      self.expectsComparableScalars(node)
    }
    case *ast.UnaryOpNode: {
      visitUnaryOpNode(self, node)
      if node.GetOperator() == "!" {
        self.mustBeScalar(node.GetExpr(), node.GetOperator())
      } else {
        self.mustBeInteger(node.GetExpr(), node.GetOperator())
      }
    }
    case *ast.PrefixOpNode: {
      visitPrefixOpNode(self, node)
      self.expectsScalarLHS(node)
    }
    case *ast.SuffixOpNode: {
      visitSuffixOpNode(self, node)
      self.expectsScalarLHS(node)
    }
    case *ast.FuncallNode: {
      visitFuncallNode(self, node)
      t := node.GetFunctionType()
      if ! t.AcceptsArgc(node.NumArgs()) {
        self.errorHandler.Errorf("%s wrong number of arguments: %d\n", node.NumArgs())
        return
      }
      self.errorHandler.Infoln("FIXME: TypeChecker: implicit cast for function arguments have not yet implemented")
    }
    case *ast.ArefNode: {
      visitArefNode(self, node)
      self.mustBeInteger(node.GetIndex(), "[]")
    }
    case *ast.CastNode: {
      visitCastNode(self, node)
      if ! node.GetExpr().GetType().IsCastableTo(node.GetType()) {
        self.errorHandler.Errorf("%s invalid cast from %s to %s\n", node.GetLocation(), node.GetExpr().GetType(), node.GetType())
      }
    }
    default: {
      visitNode(self, unknown)
    }
  }
}
