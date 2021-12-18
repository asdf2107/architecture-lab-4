package main

import (
	"fmt"
)

type CatCommand struct {
	arg1 string
	arg2 string
}

func (catCommand *CatCommand) Execute(handler Handler) {
	fmt.Println(catCommand.arg1 + catCommand.arg2)
}
