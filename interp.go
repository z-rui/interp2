package main

import (
	"fmt"
	"log"
)

var symTab = map[Identifier]interface{}{}

func (v Number) Evaluate() interface{} {
	return v
}

func (e *ArithExpr) Evaluate() interface{} {
	lhs := e.lhs.Evaluate()
	rhs := e.rhs.Evaluate()

	if e.op == '+' {
		s1, ok1 := lhs.(string)
		s2, ok2 := rhs.(string)
		if ok1 && ok2 {
			return s1 + s2
		}
	}
	{
		lhs := lhs.(Number)
		rhs := rhs.(Number)
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
			panic("unreached")
		}
	}
}

func (e *RelExpr) Evaluate() interface{} {
	lhs := e.lhs.Evaluate()
	rhs := e.rhs.Evaluate()

	if lhs, ok := lhs.(string); ok {
		rhs := rhs.(string)
		switch e.op {
		case '<':
			return lhs < rhs
		case '>':
			return lhs > rhs
		case '=':
			return lhs == rhs
		default:
			panic("unreached")
		}
	}
	{
		lhs := lhs.(Number)
		rhs := rhs.(Number)
		switch e.op {
		case '<':
			return lhs < rhs
		case '>':
			return lhs > rhs
		case '=':
			return lhs == rhs
		default:
			panic("unreached")
		}
	}
}

func (e *LogicExpr) Evaluate() interface{} {
	lhs := e.lhs.Evaluate().(bool)

	switch e.op {
	case AND:
		return lhs && e.rhs.Evaluate().(bool)
	case OR:
		return lhs || e.rhs.Evaluate().(bool)
	case NOT:
		return !lhs
	default:
		panic("unreached")
	}
}

func (s *AssignStmt) Execute() {
	s.lval.Assign(s.expr.Evaluate())
}

func (s *IfStmt) Execute() {
	if s.cond.Evaluate().(bool) {
		RunStmtBlock(s.trueClause)
	} else {
		RunStmtBlock(s.falseClause)
	}
}

func (s *WhileStmt) Execute() {
	for s.cond.Evaluate().(bool) {
		RunStmtBlock(s.body)
	}
}

func (s *PrintStmt) Execute() {
	args := make([]interface{}, len(s.argList))
	for i, expr := range s.argList {
		args[i] = expr.Evaluate()
	}
	fmt.Println(args...)
}

func (id Identifier) Assign(val interface{}) {
	symTab[id] = val
}

func (id Identifier) Evaluate() interface{} {
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
