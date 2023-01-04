package droplet

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/edobtc/cloudkit/config"
	"github.com/edobtc/cloudkit/lnd"
	"github.com/pkg/sftp"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

const (
	StartCommand        = "chmod +x /var/app/lnd/docker/start-up.sh && /usr/bin/docker compose -f /var/app/lnd/docker/docker-compose.yml up -d"
	maxRetry            = 8
	provisionerInterval = 8 * time.Second
)

func RetryProvisioner(ip string) ([]byte, error) {
	time.Sleep(provisionerInterval)

	retries := 1
	var cert []byte
	var err error

	for {
		cert, err = Provision(ip)
		if err != nil {
			if retries >= maxRetry {
				log.Error(err)
				break
			} else {
				log.Infof("attempt %d at provisioning node failed, retrying in 10 seconds...", retries)
				retries++
				time.Sleep(provisionerInterval)
			}
		}
	}

	return cert, nil
}

// Provision uses ssh to connect to the new node
// place an lnd config with given parameters,
// start lnd
// and fetch the tls cert
func Provision(ip string) ([]byte, error) {
	// inject address with ssh port, eventually lets randomize and
	// make these ports configurable
	addr := fmt.Sprintf("%s:22", ip)

	client, session, err := Connect(addr, []byte(config.Read().SSHPrivKey))
	if err != nil {
		return nil, errors.New("failed to connect for provisioning")
	}

	defer client.Close()
	defer session.Close()

	sftpClient, err := SFTPClient(client)

	if err != nil {
		return nil, errors.New("failed to initialize sftp client")
	}
	defer sftpClient.Close()

	err = AddTemplate(sftpClient)
	if err != nil {
		return nil, err
	}

	data, err := Start(client, session)
	if err != nil {
		log.Info("start failed")
		return nil, err
	}

	log.Info(data)

	time.Sleep(10 & time.Second)

	cert, err := FetchCert(sftpClient)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func AddTemplate(client *sftp.Client) error {
	wf, err := client.Create(LNDConfigPath)
	if err != nil {
		return err
	}

	b, err := lnd.RenderTemplate()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = wf.Write(b)
	if err != nil {
		fmt.Println(err)
		return err
	}

	wf.Close()

	return nil
}

func Start(client *ssh.Client, session *ssh.Session) (io.Writer, error) {
	buf := bytes.Buffer{}
	session.Stdout = bufio.NewWriter(&buf)
	err := session.Run(StartCommand)
	return &buf, err
}

func FetchCert(client *sftp.Client) ([]byte, error) {
	path := fmt.Sprintf("%s/%s", LNDConfigBase, "tls.cert")
	ff, err := client.Open(path)
	if err != nil {
		return nil, err
	}

	defer ff.Close()

	return io.ReadAll(ff)
}

func ExtractCredentials(client *sftp.Client) (*lnd.Credentials, error) {
	creds := lnd.NewCredentials()

	for f, _ := range creds.Files {
		path := fmt.Sprintf("%s/%s", LNDConfigPath, f)
		ff, err := client.Open(path)
		if err != nil {
			creds.Errors = append(creds.Errors, err)
			continue
		}

		data, err := io.ReadAll(ff)
		if err != nil {
			creds.Errors = append(creds.Errors, err)
			continue
		}

		creds.Files[f] = data

		ff.Close()
	}

	return creds, nil
}
