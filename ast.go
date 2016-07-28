package main

type ValueType float64
type Identifier string

type BinExpr struct {
	op       int
	lhs, rhs Expression
}

type LogicCond struct {
	op       int
	lhs, rhs Condition
}

type AssignStmt struct {
	lval Lvalue
	expr Expression
}

type PrintStmt struct {
	argList []Expression
}

type WhileStmt struct {
	cond Condition
	body []Statement
}

type IfStmt struct {
	cond                    Condition
	trueClause, falseClause []Statement
}
