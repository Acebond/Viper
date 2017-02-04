package main

type Node interface {
	Visit(i *Interpreter) interface{}
}

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

type Compound struct {
	children []Node
}

type Assign struct {
	left      Node
	operation Token
	right     Node
}

type Var struct {
	token Token
	value int
}

type NoOp struct {
}
