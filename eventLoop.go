package main

type EventLoop struct {
	queue   *Queue
	stopped bool
	busy    bool
	quit    chan bool
}

func (eventLoop *EventLoop) Post(cmd Command) {
	eventLoop.queue.enqueue(cmd)
	if eventLoop.busy && !eventLoop.stopped {
		eventLoop.startRoutine()
	}
}

func (eventLoop *EventLoop) startRoutine() {
	eventLoop.busy = false
	go func() {
		for {
			if eventLoop.queue.length() > 0 {
				cmd := eventLoop.queue.dequeue()
				cmd.Execute(eventLoop)
			} else if eventLoop.stopped {
				eventLoop.quit <- true
				return
			} else {
				eventLoop.busy = true
				return
			}
		}
	}()
}

func (eventLoop *EventLoop) Start() {
	eventLoop.quit = make(chan bool, 1)
	eventLoop.queue = &Queue{}
	eventLoop.startRoutine()
}

func (eventLoop *EventLoop) AwaitFinish() {
	eventLoop.stopped = true
	<-eventLoop.quit
}
