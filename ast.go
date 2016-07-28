package main

type Number float64
type String string
type Identifier string

type Statement interface {
	Execute()
}

type Expression interface {
	Evaluate() interface{}
}

type Lvalue interface {
	Evaluate() interface{}
	Assign(v interface{})
}

type BinExpr struct {
	op       int
	lhs, rhs Expression
}

type LogicExpr BinExpr
type ArithExpr BinExpr
type RelExpr BinExpr

type AssignStmt struct {
	lval Lvalue
	expr Expression
}

type PrintStmt struct {
	argList []Expression
}

type WhileStmt struct {
	cond Expression
	body []Statement
}

type IfStmt struct {
	cond                    Expression
	trueClause, falseClause []Statement
}
