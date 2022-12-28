package subscriber

// WIP defining this
type Subscriber interface {
	Start() error
	Detach() error
	Listen() error
}
