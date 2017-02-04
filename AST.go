package main

type Node interface {
}

type BasicNode struct{}

type BinOp struct {
	left     Node
	operator Token
	right    Node
}

type Num struct {
	token Token
	value int
}

type UnaryOp struct {
	operator Token
	expr     Node
}
