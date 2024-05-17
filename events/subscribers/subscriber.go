package subscribers

type Subscriber interface {
	Start() chan bool
	Detach() error

	// we need to replace this return interface{}
	// with a specific type or a channel that will
	// return a structured piece of data eventually
	// (io.Reader or a Proto)
	Listen() <-chan interface{}
}
