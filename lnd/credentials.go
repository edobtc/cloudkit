package lnd

var (
	files = []string{
		"admin.macaroon",
		"tls.cert",
		"readonly.macaroon",
		"invoice.macaroon",
	}
)

type Credentials struct {
	Files  map[string][]byte
	Errors []error
}

func NewCredentials() *Credentials {
	lookup := make(map[string][]byte, len(files))

	for _, f := range files {
		lookup[f] = []byte{}
	}

	return &Credentials{
		Files:  lookup,
		Errors: make([]error, 0),
	}
}
