package main

import (
	"bufio"
	"fmt"
	"os"
	. "github.com/asdf2107/architecture-lab-4/src/commands"
)

func main() {
	inputFile := "input.txt"
	eventLoop := new(EventLoop)
	eventLoop.Start()

	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := Parse(commandLine) // parse the line to get a Command
			eventLoop.Post(cmd)
		}
	} else {
		eventLoop.Post(&PrintCommand{
			arg: fmt.Sprintf("Error opening file \"%s\"", err.Error()),
		})
	}
	eventLoop.AwaitFinish()
}
