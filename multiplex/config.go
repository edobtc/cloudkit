package multiplex

type Config struct {
	// Limit is the maximum number of messages to send
	// to the destination publisher before ending
	Limit int `json:"limit"`

	// Behavior is the behavior to use when sending messages
	Behavior Behavior `json:"behavior"`
}

func NewConfig() *Config {
	return &Config{}
}

func NewDefaultConfig() *Config {
	return &Config{
		Limit:    100,
		Behavior: ReadWrite,
	}
}
