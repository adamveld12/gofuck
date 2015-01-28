package main

import (
	"bufio"
	"fmt"
	"github.com/adamveld12/gofuck/vm"
	"io/ioutil"
	"os"
)

func main() {
	var program string

	if len(os.Args) <= 1 {
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err.Error())
		}
		program = string(line)
	} else {
		program = getInputFile()
	}

	input, output := vm.Execute(program)
	go reader(input)
	printer(output)

	fmt.Println()
}

func reader(input chan<- rune) {
	reader := bufio.NewReader(os.Stdin)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			panic(err.Error())
		}
		input <- r
	}
}

func printer(output <-chan rune) {
	for {
		output, ok := <-output
		if ok {
			fmt.Print(string(output))
		} else {
			break
		}
	}
}

func getInputFile() string {
	fileArg := os.Args[1]
	file, err := ioutil.ReadFile(fileArg)
	if err != nil {
		panic(err)
	}

	return string(file)
}
