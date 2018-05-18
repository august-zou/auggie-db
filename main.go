package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

type InputBuffer struct {
	buffer string
	length int
}

func readInput(inputReader *bufio.Reader) *InputBuffer {
	var inputBuffer *InputBuffer = new(InputBuffer)

	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	inputLength := utf8.RuneCountInString(input)
	inputBuffer.buffer = SubString(input, 0, inputLength-1)
	inputBuffer.length = inputLength - 1
	return inputBuffer
}

func printPrompt() {
	fmt.Printf("auggie-db > ")
}

//SubString substring the source string
func SubString(str string, begin, length int) (substr string) {
	rs := []rune(str)
	lth := len(rs)

	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}

func main() {

	var inputReader *bufio.Reader = bufio.NewReader(os.Stdin)

	for {
		printPrompt()
		input := readInput(inputReader)
		if input.buffer == ".exit" {
			fmt.Println("bye.")
			return
		}
		if input.buffer != "" {
			fmt.Printf("Unrecognized command '%s'.\n", input.buffer)

		}
	}
}
