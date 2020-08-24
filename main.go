package main

import (
	"fmt"
)

type Brain struct {
	code   string
	ip     int
	memory [30000]int
	dp     int
}

func NewBrain(code string) *Brain {
	return &Brain{
		code: code,
	}
}
func main() {

	// Sample Input
	// Equivalent of 3 + 4 and placing in the result in Brain.memory[0]
	input := "++.+++>++++[<+>-]"

	b := NewBrain(input)

	fmt.Println("Code:: ", b.code)
	fmt.Println("Memory\t\t\tInstruction\tmem_pointer")

	for b.ip < len(b.code) {
		symbol := b.code[b.ip]
		switch symbol {
		case '.':
			fmt.Println(b.memory[b.dp])
		case '+':
			b.memory[b.dp]++
		case '-':
			b.memory[b.dp]--
		case '<':
			b.dp--
		case '>':
			b.dp++
		case '[':
			if b.memory[b.dp] == 0 {
				depth := 1
				for depth != 0 {
					b.ip++
					switch b.code[b.ip] {
					case ']':
						depth--
					case '[':
						depth++
					}

				}
			}
		case ']':
			if b.memory[b.dp] != 0 {
				depth := 1
				for depth != 0 {
					b.ip--
					switch b.code[b.ip] {
					case ']':
						depth++
					case '[':
						depth--
					}

				}
			}
		}
		fmt.Printf("%v\t%v\t\t%v\n", b.memory[0:10], string(symbol), b.dp)
		b.ip++
	}

}
