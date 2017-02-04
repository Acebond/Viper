package main

import "log"

type Parser struct {
	lexer        Lexer
	currentToken Token
}

func newParser(lexer Lexer) Parser {
	return Parser{lexer: lexer, currentToken: lexer.getNextToken()}
}

func (p *Parser) eat(tokenType int) {
	if p.currentToken.Type == tokenType {
		p.currentToken = p.lexer.getNextToken()
	} else {
		log.Panicf("Current token type: %d and passed token type: %d differ\n", p.currentToken.Type, tokenType)
	}
}

func (p *Parser) parse() Node {
	//return p.expr()
	node := p.program()
	if p.currentToken.Type != EOF {
		log.Panicf("Should be at EOF\n")
	}
	return node
}

/*
GRAMMAR RULES
*/
func (p *Parser) factor() Node {
	/*
	   factor : PLUS factor
	   		| MINUS factor
	   		| INTEGER
	   		| LPAREN expr RPAREN
	   		| variable
	*/
	token := p.currentToken
	switch token.Type {
	case PLUS:
		p.eat(PLUS)
		node := &UnaryOp{operator: token, expr: p.factor()}
		return node
	case MINUS:
		p.eat(MINUS)
		node := &UnaryOp{operator: token, expr: p.factor()}
		return node
	case INTEGER:
		p.eat(INTEGER)
		return &Num{token: token, value: token.Value.(int)}
	case LPAREN:
		p.eat(LPAREN)
		node := p.expr()
		p.eat(RPAREN)
		return node
	default:
		node := p.variable()
		return node
		//log.Panicf("Unknown Token for facotr %s\n", token)
	}
	//return nil
}

func (p *Parser) term() Node {
	// term: factor ((MUL | DIV) factor)*
	node := p.factor()
	for p.currentToken.Type == MUL || p.currentToken.Type == DIV {
		token := p.currentToken
		if token.Type == MUL {
			p.eat(MUL)
		} else if token.Type == DIV {
			p.eat(DIV)
		}
		node = &BinOp{left: node, operator: token, right: p.factor()}
	}
	return node
}

func (p *Parser) expr() Node {
	// expr: term ((PLUS | MINUS) term)*
	node := p.term()
	for p.currentToken.Type == PLUS || p.currentToken.Type == MINUS {
		token := p.currentToken
		if token.Type == PLUS {
			p.eat(PLUS)
		} else if token.Type == MINUS {
			p.eat(MINUS)
		}
		node = &BinOp{left: node, operator: token, right: p.term()}
	}
	return node
}

func (p *Parser) program() Node {
	// program : compound_statement
	node := p.compoundStatment()
	return node
}

func (p *Parser) compoundStatment() *Compound {
	// compound_statement : LBRACE statement_list RBRACE
	p.eat(LBRACE)
	nodes := p.statementList()
	p.eat(RBRACE)

	root := &Compound{}
	for _, node := range nodes {
		root.children = append(root.children, node)
	}

	return root
}

func (p *Parser) statementList() []Node {
	// statement_list : (statement SEMI)*
	results := []Node{}
	results = append(results, p.statment())
	for p.currentToken.Type == SEMI {

		p.eat(SEMI)
		results = append(results, p.statment()) // fix to match grammer
	}

	if p.currentToken.Type == ID {
		log.Panicf("Got ID after statments")
	}

	return results
}

func (p *Parser) statment() Node {
	//statement : compound_statement | assignment_statement | empty
	switch p.currentToken.Type {
	case LBRACE:
		return p.compoundStatment()
	case ID:
		return p.assignmentStatement()
	default:
		return p.empty()
	}
}

func (p *Parser) assignmentStatement() Node {
	// assignment_statement : variable ASSIGN expr
	left := p.variable()
	token := p.currentToken
	p.eat(ASSIGN)
	right := p.expr()
	node := &Assign{left, token, right}
	return node
}

func (p *Parser) variable() Node {
	// variable: ID
	node := &Var{token: p.currentToken}
	p.eat(ID)
	return node
}

func (p *Parser) empty() Node {
	// empty :
	return &NoOp{}
}
