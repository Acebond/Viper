package main

import "log"

type Interpreter struct {
	parser Parser
}

func (i *Interpreter) interpret() interface{} {
	tree := i.parser.parse()
	return tree.Visit(i)
	//return i.visit(tree)
}

//func (i *Interpreter) visit(node Node) interface{} {
//	return node.Visit(i)
//}

/*
Node Visit Methods
*/
var (
	GLOBAL_SCOPE = make(map[interface{}]interface{})
)

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

func (self *Compound) Visit(i *Interpreter) interface{} {
	for _, child := range self.children {
		child.Visit(i)
	}
	return nil
}

func (self *Assign) Visit(i *Interpreter) interface{} {
	varName := (self.left.(*Var)).token.Value
	GLOBAL_SCOPE[varName] = self.right.Visit(i)
	return nil
}

func (self *Var) Visit(i *Interpreter) interface{} {
	varName := self.token.Value
	if val, ok := GLOBAL_SCOPE[varName]; ok {
		return val
	} else {
		log.Panicf("Variable does not exists .. help\n")
		return nil
	}

}

func (self *NoOp) Visit(i *Interpreter) interface{} {
	return nil
}
