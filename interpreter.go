package main

type Interpreter struct {
	parser Parser
}

//func (i *Interpreter) visit(node Node) interface{} {
//	return node.Visit(i)
//}

func (i *Interpreter) interpret() interface{} {
	tree := i.parser.parse()
	return tree.Visit(i)
	//return i.visit(tree)
}
