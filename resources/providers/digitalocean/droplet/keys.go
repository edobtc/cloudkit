package droplet

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/digitalocean/godo"
	"github.com/edobtc/cloudkit/config"
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

func RemoveKey(id int) error {
	client := NewClient()
	ctx := context.TODO()

	_, err := client.Keys.DeleteByID(ctx, id)
	return err
}

type SshKeyOptions struct {
	Name string
	Key  []byte
}

// Select keys will list all keys
// and filter based on the default tags
func SelectKeys() ([]godo.Key, error) {
	client := NewClient()
	ctx := context.TODO()
	results := make([]godo.Key, 0)

	keys, _, err := client.Keys.List(ctx, nil)
	if err != nil {
		return results, err
	}

	for _, key := range keys {
		if strings.Contains(key.Name, "cloudkit") {
			results = append(results, key)
		}
	}
	return results, nil
}

func CloudKitSSHKey() (*godo.Key, error) {
	client := NewClient()
	ctx := context.TODO()

	keys, _, err := client.Keys.List(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		if key.Name == config.Read().SSHKeyName {
			return &key, nil
		}
	}
	return nil, errors.New("key not found")
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
