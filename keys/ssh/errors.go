package ssh

import "errors"

var (
	ErrNoSSHKeyDefined = errors.New("no ssh key configured")
)
