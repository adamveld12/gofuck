package main

import (
	"bufio"
	"fmt"
	"github.com/adamveld12/gofuck/vm"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var program string

	if len(os.Args) <= 1 {
		fmt.Println(os.Args)
		// maybe support REPL
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err.Error())
		}
		program = string(line)
	} else {
		program = getInputFile()
	}

	input, output := vm.Execute(program)
	go reader(input)
	printer(output)
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
		fmt.Print(<-output)
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
