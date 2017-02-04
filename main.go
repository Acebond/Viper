package main

import (
	"bufio"
	"fmt"
	"os"
)

func eval(text string) int {
	lexer := newLexer(text)
	parser := newParser(lexer)
	interpreter := Interpreter{parser}
	result := interpreter.interpret()
	return result.(int)
}

func unitTest() {
	tests := []struct {
		input string
		ans   int
	}{
		{"5+5", 10},
		{"5-5", 0},
		{"7 + 3 * (10 / (12 / (3 + 1) - 1)) / (2 + 3) - 5 - 3 + (8)", 10},
		{"-5", -5},
	}

	for _, test := range tests {
		if res := eval(test.input); res != test.ans {
			fmt.Printf("%s Failed! Got %d instead of %d\n", test.input, res, test.ans)
		}
	}
	fmt.Printf("Testing Complete\n")
}

func main() {
	if len(os.Args) > 1 {
		unitTest()
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Calc> ")
		text, _ := reader.ReadString('\n')
		fmt.Println(eval(text))
	}
}
