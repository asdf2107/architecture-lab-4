package implementation

import (
	"fmt"

	"github.com/asdf2107/architecture-lab-4/src/interfaces"
)

type PrintCommand struct {
	Arg string
}

func (printCommand *PrintCommand) Execute(handler interfaces.Handler) {
	fmt.Println(printCommand.Arg)
}
