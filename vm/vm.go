package vm

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
		// sanitize input
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
			case '[': // If the value at the current cell is zero, skips to the corresponding ].
				if memory[dataPtr] == 0 {
					// skip to ]
					for program[instructionPtr] != ']' {
						instructionPtr++
					}
					// account for increment at end of the loop
					instructionPtr--
				}
			case ']': // If the value at the current cell is zero, move to the next instruction.
				if memory[dataPtr] != 0 {
					for program[instructionPtr] != '[' {
						instructionPtr--
					}
				}
			}

			instructionPtr++
		}
	}()

	return input, output
}
