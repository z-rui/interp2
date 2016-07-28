%{
package main

%}

%union{
	numval     float64
	ident      string
	lval       Lvalue
	expr       Expression
	stmt       Statement
	argList   []Expression
	stmtBlock []Statement
}

%start program

%token IF THEN ELSE END WHILE DO PRINT AND OR NOT

%token <numval> NUMBER
%token <ident> IDENTIFIER

%left '+' '-'
%left '*' '/'

%type <lval> var
%type <expr> expr condition rel_condition sub_condition
%type <stmt> statement if_stmt while_stmt assign_stmt print_stmt
%type <argList> arg_list
%type <stmtBlock> stmt_block program

%%

program	: stmt_block { yylex.(*Lexer).program = $1 }
	;

stmt_block
	: { $$ = []Statement{} }
	| stmt_block statement { $$ = append($1, $2) }
	;

statement
	: if_stmt { $$ = $1 }
	| while_stmt { $$ = $1 }
	| assign_stmt { $$ = $1 }
	| print_stmt { $$ = $1 }
	| error { $$ = nil }
	;

if_stmt	: IF condition THEN stmt_block ELSE stmt_block END IF { $$ = &IfStmt{$2, $4, $6} }
	| IF condition THEN stmt_block END IF { $$ = &IfStmt{$2, $4, nil} }
	;

while_stmt
	: WHILE condition DO stmt_block END WHILE { $$ = &WhileStmt{$2, $4} }
	;

assign_stmt
	: var '=' expr { $$ = &AssignStmt{$1, $3} }
	;

print_stmt
	: PRINT arg_list { $$ = &PrintStmt{$2} }
	;

arg_list
	: expr { $$ = []Expression{$1} }
	| arg_list ',' expr { $$ = append($1, $3) }
	;

condition
	: rel_condition { $$ = $1 }
	| sub_condition AND sub_condition { $$ = &LogicExpr{AND, $1, $3} }
	| sub_condition OR sub_condition { $$ = &LogicExpr{OR, $1, $3} }
	| NOT sub_condition { $$ = &LogicExpr{NOT, $2, nil} }

rel_condition
	: expr '=' expr { $$ = &RelExpr{'=', $1, $3} }
	| expr '<' expr { $$ = &RelExpr{'<', $1, $3} }
	| expr '>' expr { $$ = &RelExpr{'>', $1, $3} }
	;

sub_condition
	: rel_condition { $$ = $1 }
	| '(' condition ')' { $$ = $2 }
	;

var	: IDENTIFIER { $$ = Identifier($1) }
	;

expr	: var { $$ = $1 }
	| NUMBER { $$ = Number($1) }
	| '(' expr ')' { $$ = $2 }
	| expr '+' expr { $$ = &ArithExpr{'+', $1, $3} }
	| expr '-' expr { $$ = &ArithExpr{'-', $1, $3} }
	| expr '*' expr { $$ = &ArithExpr{'*', $1, $3} }
	| expr '/' expr { $$ = &ArithExpr{'/', $1, $3} }
	;

%%

