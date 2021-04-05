package main

import (
	"monkey/lexer"
	"monkey/token"
	"fmt"
)

func main() {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);`;

	l := lexer.New(input)
	for {
		tok := l.NextToken()
		fmt.Printf("%s %s\n", tok.Type, tok.Literal)
		if tok.Type == token.EOF {
			break
		}
	}
}
