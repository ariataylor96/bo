package main

import (
	"bufio"
	"fmt"
	"os"
)

func matching_left_brace(program []string, start int) int {
	relevant := program[:start+1]
	paren_count := 0

	for idx := len(relevant) - 1; idx >= 0; idx-- {
		val := relevant[idx]

		if val == "]" {
			paren_count++
		} else if val == "[" {
			paren_count--
		}

		if paren_count == 0 {
			return idx
		}
	}

	return 0
}

func matching_right_brace(program []string, start int) int {
	relevant := program[start:]
	paren_count := 0

	for idx, val := range relevant {
		if val == "[" {
			paren_count++
		} else if val == "]" {
			paren_count--
		}

		if paren_count == 0 {
			return idx
		}
	}

	return 0
}

func main() {
	handle, _ := os.Open(os.Args[1])
	defer handle.Close()

	data := make([]int, 0)
	program := Tokenize(handle)

	stdin_reader := bufio.NewReader(os.Stdin)

	data_pointer := 0
	instruction_pointer := 0

	for instruction_pointer < len(program) {
		// Allow for infinite data points
		for data_pointer >= len(data) {
			data = append(data, 0)
		}

		sym := program[instruction_pointer]

		switch sym {
		case ">":
			data_pointer++
		case "<":
			data_pointer--
		case "+":
			data[data_pointer]++
		case "-":
			data[data_pointer]--
		case ".":
			fmt.Printf("%c", rune(data[data_pointer]))
		case ",":
			next_token, size, _ := stdin_reader.ReadRune()

			if size == 0 {
				data[data_pointer] = 0
			} else {
				data[data_pointer] = int(next_token)
			}

		case "[":
			if data[data_pointer] == 0 {
				instruction_pointer = matching_right_brace(program, instruction_pointer)
			}
		case "]":
			if data[data_pointer] != 0 {
				instruction_pointer = matching_left_brace(program, instruction_pointer)
			}
		}

		instruction_pointer++
	}
}
