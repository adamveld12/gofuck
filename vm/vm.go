package vm

import _ "log"

const MEM_SIZE int = 30000

func Execute(program string) (chan<- rune, <-chan rune) {
	var memory [MEM_SIZE]int
	dataPtr := 0

	instructionPtr := 0
	instructionsLen := len(program)

	input, output := make(chan rune), make(chan rune)

	go func() {
		defer close(input)
		defer close(output)

		for instructionPtr < instructionsLen {
			switch program[instructionPtr] {
			case '+': // Increments the value at the current cell by one.
				memory[dataPtr] += 1
			case '-': // Decrements the value at the current cell by one.
				memory[dataPtr] -= 1
			case '>': // Moves the data pointer to the next cell (cell on the right).
				dataPtr++
			case '<': // Moves the data pointer to the previous cell (cell on the left).
				dataPtr--
			case '.': // Prints the ASCII value at the current cell (i.e. 65 = 'A').
				output <- rune(memory[dataPtr])
			case ',': // Reads a single input character into the current cell.
				memory[dataPtr] = int(<-input)
			case '[': // If the value at the current cell is zero, skips to the corresponding ]
				// Otherwise, move to the next instruction.
				value := memory[dataPtr]
				if value == 0 {
					nestDepth := 1
					for nestDepth > 0 {
						instructionPtr++
						symbol := program[instructionPtr]
						if symbol == '[' {
							nestDepth++
						} else if symbol == ']' {
							nestDepth--
						}
					}
					instructionPtr--
				}

			case ']': // If the value at the current cell is zero, move to the next instruction
				// Otherwise, move backwards in the instructions to the corresponding [ .
				value := memory[dataPtr]
				if value != 0 {
					nestDepth := 1
					for nestDepth > 0 {
						instructionPtr--
						symbol := program[instructionPtr]
						if symbol == ']' {
							nestDepth++
						} else if symbol == '[' {
							nestDepth--
						}
					}
					instructionPtr--
				}
			}
			instructionPtr++
		}
	}()

	return input, output
}
