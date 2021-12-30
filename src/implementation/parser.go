package implementation

import (
	"fmt"
	"strings"

	. "github.com/asdf2107/architecture-lab-4/src/interfaces"
)

func Parse(textCommand string) Command {
	split := strings.Fields(textCommand)
	if split[0] == "print" {
		if len(split) != 2 {
			return InvalidArgsCount(split[0], 1, len(split)-1)
		} else {
			command := PrintCommand{
				Arg: split[1],
			}
			return &command
		}
	} else if split[0] == "cat" {
		if len(split) != 3 {
			return InvalidArgsCount(split[0], 1, len(split)-1)
		} else {
			command := CatCommand{
				Arg1: split[1],
				Arg2: split[2],
			}
			return &command
		}
	} else {
		command := PrintCommand{
			Arg: fmt.Sprintf("Invalid command \"%s\": command does not exist", split[0]),
		}
		return &command
	}
}

func InvalidArgsCount(command string, desired int, actual int) Command {
	result := &PrintCommand{
		Arg: fmt.Sprintf("Invalid command \"%s\": Expected %d arguments, found %d", command, desired, actual),
	}
	return result
}
