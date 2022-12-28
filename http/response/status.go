package response

type Status struct {
	Additional string `json:"additional,omitempty"`
	Message    string `json:"message"`
}
