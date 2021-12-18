package main

// Handler allows to send commands to an event loop
// it’s associated with.
type Handler interface {
	Post(cmd Command)
}
