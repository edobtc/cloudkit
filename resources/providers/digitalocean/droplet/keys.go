package droplet

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	"github.com/edobtc/cloudkit/keys/ssh"
	"github.com/sirupsen/logrus"
)

func AddSSHKey(opt SshKeyOptions) (*godo.Key, error) {
	client := NewClient()
	ctx := context.TODO()

	req := &godo.KeyCreateRequest{
		Name:      opt.Name,
		PublicKey: string(opt.Key),
	}

	transfer, _, err := client.Keys.Create(ctx, req)
	return transfer, err
}

type SshKeyOptions struct {
	Name string
	Key  []byte
}

// GenerateAndAddSSHKey will generate
// a new ssh key and add it to the account
func GenerateAndAddSSHKey(name string) {
	s := ssh.NewSSHKey(ssh.DefaultOptions())

	err := s.Generate()
	if err != nil {
		logrus.Error(err)
		return
	}

	err = s.SaveToFile()
	if err != nil {
		logrus.Error(err)
		return
	}

	key, err := AddSSHKey(SshKeyOptions{
		Name: name,
		Key:  s.Pub,
	})

	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println(key)
}
