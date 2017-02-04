package main

type Interpreter struct {
	parser Parser
}

func (i *Interpreter) visit(node Node) interface{} {
	return node.Visit(i)
	/*switch v := node.(type) {
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
	*/
}

func (i *Interpreter) interpret() interface{} {
	tree := i.parser.parse()
	return i.visit(tree)
}
