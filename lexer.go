package main

import (
	"log"
	"strconv"
	"unicode"
)

type Lexer struct {
	text        string
	position    int
	currentChar string
}

func newLexer(text string) Lexer {
	lexer := Lexer{text: text}
	lexer.advance(0)
	return lexer
}

func (i *Lexer) advance(steps int) {
	// Advance the 'pos' pointer and set the 'current_char' variable.
	i.position += steps
	if i.position > len(i.text)-1 {
		i.currentChar = ""
	} else {
		//i.text = strings.TrimSpace(i.text)
		i.currentChar = i.text[i.position:]
	}
}

func (i *Lexer) integer() int {
	intEndIndex := 0
	for intEndIndex < len(i.currentChar) && IsDigit(i.currentChar[intEndIndex]) {
		intEndIndex++
	}

	val, err := strconv.ParseInt(i.currentChar[0:intEndIndex], 10, 0)
	if err != nil {
		log.Panicf("Failed to parse integer\n")
	}
	i.advance(intEndIndex)
	return int(val)
}

func (i *Lexer) skipWhiteSpace() {
	for i.currentChar != "" && isWhiteSpace(i.currentChar[0]) {
		i.advance(1)
	}
}

func (i *Lexer) _id() Token {
	// Handle identifiers and reserved keywords
	idEndIndex := 0
	for i.currentChar != "" && unicode.IsLetter(rune(i.currentChar[idEndIndex])) {
		idEndIndex++
	}

	token := Token{Type: ID, Value: i.currentChar[0:idEndIndex]}
	i.advance(idEndIndex)
	return token
}

func (i *Lexer) getNextToken() Token {
	/*
		Lexical analyzer (also known as scanner or tokenizer)
		This method is responsible for breaking a sentence apart into tokens.
	*/
	//i.text = strings.TrimSpace(i.text) // we can comfirm there will be zero whitespace to the next Token

	//i.currentChar = i.text[i.position:] // put the next Token into currentChar

	for i.currentChar != "" { // while we still have something to read form text

		if isWhiteSpace(i.currentChar[0]) {
			i.skipWhiteSpace()
			continue
		}

		if unicode.IsLetter(rune(i.currentChar[0])) { //why is the string not already runes?
			return i._id()
		}

		if IsDigit(i.currentChar[0]) {
			return Token{INTEGER, i.integer()}
		}

		//currentChar := i.text[i.position]

		switch i.currentChar[0] {
		case '*':
			i.advance(1)
			return Token{MUL, MUL}
		case '/':
			i.advance(1)
			return Token{DIV, DIV}
		case '+':
			i.advance(1)
			return Token{PLUS, PLUS}
		case '-':
			i.advance(1)
			return Token{MINUS, MINUS}
		case '(':
			i.advance(1)
			return Token{LPAREN, LPAREN}
		case ')':
			i.advance(1)
			return Token{RPAREN, RPAREN}
		case '=':
			i.advance(1)
			return Token{ASSIGN, ASSIGN}
		case '{':
			i.advance(1)
			return Token{LBRACE, LBRACE}
		case '}':
			i.advance(1)
			return Token{RBRACE, RBRACE}
		case ';':
			i.advance(1)
			return Token{SEMI, SEMI}
		}

		log.Panicf("Unknwon Token: %s", i.currentChar)
	}
	return Token{EOF, 0}
}
