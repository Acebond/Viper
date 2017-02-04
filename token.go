package main

import (
	"fmt"
)

// Token Types

// EOF (end-of-file) token is used to indicate that
// there is no more input left for lexical analysis

// http://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go
const (
	INTEGER = iota
	PLUS    = iota
	MINUS   = iota
	MUL     = iota
	DIV     = iota
	LPAREN  = iota
	RPAREN  = iota
	ID      = iota
	ASSIGN  = iota
	LBRACE  = iota
	RBRACE  = iota
	SEMI    = iota
	EOF     = iota
)

// add lookup to function so LBRACE stores '{' and has a func pointer?

type Token struct {
	Type  int
	Value interface{}
}

func (t *Token) String() string {
	return fmt.Sprintf("Token (%d, %d)\n", t.Type, t.Value)
}
