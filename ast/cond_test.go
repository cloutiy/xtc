package ast

import (
  "testing"
)

func TestCondExpr(t *testing.T) {
/*
  (n < 2) ? 1 : (f(n-1)+f(n-2))
 */
  x := CondExprNode(
    BinaryOpNode("<", VariableNode("n"), IntegerLiteralNode("2")),
    IntegerLiteralNode("1"),
    BinaryOpNode("+",
                 FuncallNode(VariableNode("f"), []IExprNode { BinaryOpNode("-", VariableNode("n"), IntegerLiteralNode("1")) }),
                 FuncallNode(VariableNode("f"), []IExprNode { BinaryOpNode("-", VariableNode("n"), IntegerLiteralNode("2")) })))
  assertEquals(t, x.String(), "(if (< n 2) 1 (+ (f (- n 1)) (f (- n 2))))")
}

func TestIf(t *testing.T) {
/*
  if (n % 2 == 0) {
    println("even");
  } else {
    println("odd");
  }
 */
  x := IfNode(
    BinaryOpNode("==", BinaryOpNode("%", VariableNode("n"), IntegerLiteralNode("2")), IntegerLiteralNode("0")),
    ExprStmtNode(FuncallNode(VariableNode("println"), []IExprNode { StringLiteralNode("\"even\"") })),
    ExprStmtNode(FuncallNode(VariableNode("println"), []IExprNode { StringLiteralNode("\"odd\"") })),
  )
  s := `
    (if (= (modulo n 2) 0)
      (println "even")
      (println "odd"))
  `
  assertEquals(t, x.String(), trimSpace(s))
}

func TestSwitch(t *testing.T) {
  /*
  switch (n) {
    case 1: println("one");
    case 2: println("two");
    default: println("plentiful")
  }
   */
  x := SwitchNode(
    VariableNode("n"),
    []IStmtNode {
      CaseNode(
        []IExprNode { IntegerLiteralNode("1") },
        ExprStmtNode(FuncallNode(VariableNode("println"), []IExprNode { StringLiteralNode("\"one\"") })),
      ),
      CaseNode(
        []IExprNode { IntegerLiteralNode("2") },
        ExprStmtNode(FuncallNode(VariableNode("println"), []IExprNode { StringLiteralNode("\"two\"") })),
      ),
      CaseNode(
        []IExprNode { },
        ExprStmtNode(FuncallNode(VariableNode("println"), []IExprNode { StringLiteralNode("\"plentiful\"") })),
      ),
    },
  )
  s := `
    (let ((switch-cond n))
      (cond
        ((= switch-cond 1) (println "one"))
        ((= switch-cond 2) (println "two"))
        (else (println "plentiful"))))
  `
  assertEquals(t, x.String(), trimSpace(s))
}
