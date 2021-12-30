package implementation

import (
	"github.com/asdf2107/architecture-lab-4/src/interfaces"
)

type CatCommand struct {
	Arg1 string
	Arg2 string
}

func (catCommand *CatCommand) Execute(handler interfaces.Handler) {
	res := catCommand.Arg1 + catCommand.Arg2
	handler.Post(&PrintCommand{Arg: res})
}
