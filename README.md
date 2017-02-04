# Viper
Simple Python interpreter written in Golang as a learning exercise.

[This](https://ruslanspivak.com/lsbasi-part1/) blog has helped in understand the concepts of AST and Symbol Tables. 

## Language Grammar
```
program : compound_statement

compound_statement : LBRACE statement_list RBRACE

statement_list : (statement SEMI)*

statement : compound_statement
            | assignment_statement
            | empty

assignment_statement : variable ASSIGN expr

empty :

expr: term ((PLUS | MINUS) term)*

term: factor ((MUL | DIV) factor)*

factor : PLUS factor
        | MINUS factor
        | INTEGER
        | LPAREN expr RPAREN
        | variable

variable: ID
```