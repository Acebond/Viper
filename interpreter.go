package main

import (
	"log"
)

type Interpreter struct {
	parser Parser
}

func (i *Interpreter) visit(node Node) interface{} {
	switch v := node.(type) {
	case BinOp:
		return i.visitBinOp(v)
	case Num:
		return i.visitNum(v)
	case UnaryOp:
		return i.visitUnaryOp(v)
	default:
		log.Panicf("Visit errorn")
		return nil
	}
}

func (i *Interpreter) visitBinOp(binOp BinOp) interface{} {
	//binOp := node.(BinOp)
	switch binOp.operator.Type {
	case PLUS:
		return i.visit(binOp.left).(int) + i.visit(binOp.right).(int)
	case MINUS:
		return i.visit(binOp.left).(int) - i.visit(binOp.right).(int)
	case MUL:
		return i.visit(binOp.left).(int) * i.visit(binOp.right).(int)
	case DIV:
		return i.visit(binOp.left).(int) / i.visit(binOp.right).(int)
	default:
		log.Printf("Visit Erro Not Type not valid")
		return nil

	}
}

func (i *Interpreter) visitNum(num Num) int {
	return num.value
}

func (i *Interpreter) visitUnaryOp(unaryOp UnaryOp) int {
	op := unaryOp.operator.Type
	if op == PLUS {
		return +(i.visit(unaryOp.expr).(int))
	} else if op == MINUS {
		return -(i.visit(unaryOp.expr).(int))
	} else {
		log.Panicf("Invalid Unary Type")
	}
	return 0
}

func (i *Interpreter) interpret() interface{} {
	tree := i.parser.parse()
	return i.visit(tree)
}
