package ssh

import (
	"os"

	"github.com/edobtc/cloudkit/config"
)

func FetchCert() ([]byte, error) {
	key := config.Read().SSHPrivKey

	if key != "" {
		return []byte(key), nil
	}

	path := config.Read().SSHKeyPath

	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return []byte{}, err
		}

		return data, nil
	}

	return []byte{}, ErrNoSSHKeyDefined

}
