package main

import (
	"fmt"
	"strings"
)

func Parse(textCommand string) Command {
	split := strings.Fields(textCommand)
	if split[0] == "print" {
		if len(split) != 2 {
			return InvalidArgsCount(split[0], 1, len(split)-1)
		} else {
			return &PrintCommand{
				arg: split[1],
			}
		}
	} else if split[0] == "cat" {
		if len(split) != 3 {
			return InvalidArgsCount(split[0], 1, len(split)-1)
		} else {
			return &CatCommand{
				arg1: split[1],
				arg2: split[2],
			}
		}
	} else {
		return &PrintCommand{
			arg: fmt.Sprintf("Invalid command \"%s\": command does not exist", split[0]),
		}
	}
}

func InvalidArgsCount(command string, desired int, actual int) Command {
	return &PrintCommand{
		arg: fmt.Sprintf("Invalid command \"%s\": Expected %d arguments, found %d", command, desired, actual),
	}
}
