package interfaces

// Command represents actions that can be performed
// in a single event loop iteration
type Command interface {
	Execute(handler Handler)
}
