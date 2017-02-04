package main

import (
	"bufio"
	"fmt"
	"os"
)

func eval(text string) {
	lexer := newLexer(SpaceMap(text))
	parser := newParser(lexer)
	interpreter := Interpreter{parser}
	result := interpreter.interpret()
	fmt.Println(result)
}

func main() {
	if len(os.Args) > 1 { // User input
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Calc> ")
			text, _ := reader.ReadString('\n')
			eval(text)
		}
	}

	text := " 55  *  2"
	eval(text)

}
