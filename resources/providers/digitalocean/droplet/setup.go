package droplet

import (
	"fmt"
	"io"

	"github.com/edobtc/cloudkit/lnd"
	"github.com/pkg/sftp"
)

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
