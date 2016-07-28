package main

import (
	"fmt"
	"log"
)

var symTab = map[Identifier]ValueType{}

type Statement interface {
	Execute()
}

type Expression interface {
	Evaluate() ValueType
}

type Condition interface {
	EvaluateCond() bool
}

type Lvalue interface {
	Evaluate() ValueType
	Assign(v ValueType)
}

func (v ValueType) Evaluate() ValueType {
	return v
}

func (e *BinExpr) Evaluate() ValueType {
	lhs := e.lhs.Evaluate()
	rhs := e.rhs.Evaluate()
	switch e.op {
	case '+':
		return lhs + rhs
	case '-':
		return lhs - rhs
	case '*':
		return lhs * rhs
	case '/':
		return lhs / rhs
	default:
		log.Fatalf("BinExpr.Evaluate: unknown numeric operator: %c\n", e.op)
	}
	return 0
}

func (e *BinExpr) EvaluateCond() bool {
	lhs := e.lhs.Evaluate()
	rhs := e.rhs.Evaluate()
	switch e.op {
	case '<':
		return lhs < rhs
	case '>':
		return lhs > rhs
	case '=':
		return lhs == rhs
	default:
		log.Fatalf("BinExpr.EvaluateCond: unknown relation operator: %c\n", e.op)
	}
	return false
}

func (e *LogicCond) EvaluateCond() bool {
	switch e.op {
	case AND:
		return e.lhs.EvaluateCond() && e.rhs.EvaluateCond()
	case OR:
		return e.lhs.EvaluateCond() || e.rhs.EvaluateCond()
	case NOT:
		return !e.lhs.EvaluateCond()
	default:
		log.Fatalf("LogicCond.EvaluateCond: unknown boolean operator: %d\n", e.op)
	}
	return false
}

func (s *AssignStmt) Execute() {
	s.lval.Assign(s.expr.Evaluate())
}

func (s *IfStmt) Execute() {
	if s.cond.EvaluateCond() {
		RunStmtBlock(s.true_clause)
	} else {
		RunStmtBlock(s.false_clause)
	}
}

func (s *WhileStmt) Execute() {
	for s.cond.EvaluateCond() {
		RunStmtBlock(s.body)
	}
}

func (s *PrintStmt) Execute() {
	args := make([]interface{}, len(s.arg_list))
	for i, expr := range s.arg_list {
		args[i] = expr.Evaluate()
	}
	fmt.Println(args...)
}

func (id Identifier) Assign(val ValueType) {
	symTab[id] = val
}

func (id Identifier) Evaluate() ValueType {
	val, ok := symTab[id]
	if !ok {
		log.Println("Identifier.Evaluate: symbol", id, "undefined")
	}
	return val
}

func RunStmtBlock(block []Statement) {
	for _, stmt := range block {
		stmt.Execute()
	}
}
