package droplet

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func Connect(addr string, key []byte) (*ssh.Client, *ssh.Session, error) {
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, nil, err
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}

func SFTPClient(client *ssh.Client) (*sftp.Client, error) {
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		return nil, err
	}

	return sftpClient, nil
}
