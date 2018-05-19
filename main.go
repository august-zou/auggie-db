package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"unicode/utf8"
)

type MetaCommandResult int

const (
	META_COMMAND_SUCCESS = iota
	META_COMMAND_UNRECOGNIZED_COMMAND
)

type PrepareResult int

const (
	PREPARE_SUCCESS = iota
	PREPARE_UNRECOGNIZED_STATEMENT
)

type StatementType int

const (
	STATEMENT_INSERT = iota
	STATEMENT_SELECT
)

type Statement struct {
	sType StatementType
}

//InputBuffer store input buffer
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

func exitFunc() {
	fmt.Println("bye...")
	os.Exit(0)
}

//find the statement
func prepareStatement(input *InputBuffer, statement *Statement) PrepareResult {
	if input.buffer == "insert" {
		statement.sType = STATEMENT_INSERT
		return PREPARE_SUCCESS
	}
	if input.buffer == "select" {
		statement.sType = STATEMENT_SELECT
		return PREPARE_SUCCESS
	}
	return PREPARE_UNRECOGNIZED_STATEMENT
}

func doMetaCommand(input *InputBuffer) MetaCommandResult {
	if input.buffer == ".exit" {
		exitFunc()
	}
	return META_COMMAND_UNRECOGNIZED_COMMAND

}
func main() {

	// listen os signal ctrl+c kill
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println(s)
				exitFunc()
			default:
				fmt.Println("other", s)
			}
		}
	}()

	var inputReader *bufio.Reader = bufio.NewReader(os.Stdin)

	for {
		printPrompt()
		input := readInput(inputReader)
		if input.buffer == "" {
			continue
		}

		if input.buffer[0] == '.' {
			switch doMetaCommand(input) {
			case META_COMMAND_SUCCESS:
				continue
			case META_COMMAND_UNRECOGNIZED_COMMAND:
				fmt.Printf("this is meta command '%s'.\n", input.buffer)
				continue
			}
		}

		var statement Statement

		switch prepareStatement(input, &statement) {
		case PREPARE_SUCCESS:
			fmt.Printf("statement type %d'.\n", statement.sType)
			break
		case PREPARE_UNRECOGNIZED_STATEMENT:
			fmt.Printf("Unrecognized keyword at start of '%s'.\n", input.buffer)
		}
		// fmt.Printf("Unrecognized command '%s'.\n", input.buffer)

	}
}
