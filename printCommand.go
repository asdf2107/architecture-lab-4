package main

import (
	"fmt"
)

type PrintCommand struct {
	arg string
}

func (printCommand *PrintCommand) Execute(handler Handler) {
	fmt.Println(printCommand.arg)
}
