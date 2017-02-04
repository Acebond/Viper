package main

import "log"

type Node interface {
	Visit(i *Interpreter) interface{}
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

func (self *BinOp) Visit(i *Interpreter) interface{} {
	//binOp := node.(BinOp)
	switch self.operator.Type {
	case PLUS:
		return self.left.Visit(i).(int) + self.right.Visit(i).(int)
		//return i.visit(self.left).(int) + i.visit(self.right).(int)
	case MINUS:
		return self.left.Visit(i).(int) - self.right.Visit(i).(int)
		//return i.visit(self.left).(int) - i.visit(self.right).(int)
	case MUL:
		return self.left.Visit(i).(int) * self.right.Visit(i).(int)
		//return i.visit(self.left).(int) * i.visit(self.right).(int)
	case DIV:
		return self.left.Visit(i).(int) / self.right.Visit(i).(int)
		//return i.visit(self.left).(int) / i.visit(self.right).(int)
	default:
		log.Printf("Visit Erro Not Type not valid")
		return nil
	}
}

func (self *Num) Visit(i *Interpreter) interface{} {
	return self.value
}

func (self *UnaryOp) Visit(i *Interpreter) interface{} {
	op := self.operator.Type
	if op == PLUS {
		return +(self.expr.Visit(i).(int))
		//return +(i.visit(self.expr).(int))
	} else if op == MINUS {
		return -(self.expr.Visit(i).(int))
		//return -(i.visit(self.expr).(int))
	} else {
		log.Panicf("Invalid Unary Type")
	}
	return 0
}
