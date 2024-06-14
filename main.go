package main

import (
	"github.com/01-edu/z01"
	"os"
)

const memorySize = 2048

func main() {
	if len(os.Args) != 2 {
		return
	}

	code := os.Args[1]
	memory := make([]byte, memorySize)
	pointer := 0
	loopStack := []int{}

	// Parse and execute the Brainfuck code
	i := 0
	for i < len(code) {
		cmd := code[i]

		switch cmd {
		case '>':
			pointer++
			if pointer >= memorySize {
				pointer = 0 // handle wrap around, if needed
			}
		case '<':
			pointer--
			if pointer < 0 {
				pointer = memorySize - 1 // handle wrap around, if needed
			}
		case '+':
			memory[pointer]++
		case '-':
			memory[pointer]--
		case '.':
			z01.PrintRune(rune(memory[pointer]))
		case '[':
			if memory[pointer] == 0 {
				// find the corresponding ']'
				openBrackets := 1
				for openBrackets > 0 {
					i++
					if code[i] == '[' {
						openBrackets++
					} else if code[i] == ']' {
						openBrackets--
					}
				}
			} else {
				loopStack = append(loopStack, i)
			}
		case ']':
			if memory[pointer] != 0 {
				i = loopStack[len(loopStack)-1]
			} else {
				loopStack = loopStack[:len(loopStack)-1]
			}
		default:
			// Ignore unknown characters (comments)
		}

		i++
	}

	z01.PrintRune('\n')
}
