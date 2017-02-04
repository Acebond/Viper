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

func (p *Parser) factor() Node {
	// factor : (PLUS | MINUS) factor | INTEGER | LPAREN expr RPAREN
	token := p.currentToken
	switch token.Type {
	case PLUS:
		p.eat(PLUS)
		node := UnaryOp{operator: token, expr: p.factor()}
		return &node
	case MINUS:
		p.eat(MINUS)
		node := UnaryOp{operator: token, expr: p.factor()}
		return &node
	case INTEGER:
		p.eat(INTEGER)
		return &Num{token: token, value: token.Value}
	case LPAREN:
		p.eat(LPAREN)
		node := p.expr()
		p.eat(RPAREN)
		return node
	default:
		log.Panicf("Unknown Token for facotr %s\n", token)
	}
	return nil
}

func (p *Parser) term() Node {
	// term : factor ((MUL | DIV) factor)*
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
	// expr   : term ((PLUS | MINUS) term)*
	// term   : factor ((MUL | DIV) factor)*
	// factor : INTEGER | LPAREN expr RPAREN
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

func (p *Parser) parse() Node {
	return p.expr()
}
