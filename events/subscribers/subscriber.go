package subscribers

type Subscriber interface {
	Start() chan bool
	Detach() error
	Listen() <-chan interface{}
}
