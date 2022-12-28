package lambda

type Event struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Data        string `json:"data"`
}
