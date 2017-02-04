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
	default:
		log.Panicf("Visit error]n")
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

func (i *Interpreter) interpret() interface{} {
	tree := i.parser.parse()
	return i.visit(tree)
}
