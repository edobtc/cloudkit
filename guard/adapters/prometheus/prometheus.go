package prometheus

// Config is the collection of configuration options specific
// to the prometheus guard implementation
type Config struct {
	// TODO
}

// Guard is the prometheus guard
type Guard struct {
	Config Config
}

// Check is the method for performing the guard check
// that will be called by any experiment orchestrator
// to maintain ongoing checks of remote state/status, drift,
// change, etc...
func (g *Guard) Check() (bool, error) {
	return true, nil
}
