package implementation

import (
	"sync"

	. "github.com/asdf2107/architecture-lab-4/src/interfaces"
)

type Queue struct {
	sync.Mutex
	commands []Command
}

func (q *Queue) enqueue(cmd Command) {
	q.Lock()
	defer q.Unlock()
	q.commands = append(q.commands, cmd)
}

func (q *Queue) dequeue() Command {
	q.Lock()
	defer q.Unlock()
	res := q.commands[0]
	q.commands[0] = nil
	q.commands = q.commands[1:]
	return res
}

func (q *Queue) length() int {
	return len(q.commands)
}
