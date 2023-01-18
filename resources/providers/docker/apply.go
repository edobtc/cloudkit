package docker

import log "github.com/sirupsen/logrus"

// Apply runs the Provider end to end, so calls
// read and clone
func (p *Provider) Apply() error {
	err := p.Read()
	if err != nil {
		return err
	}

	id, err := p.Start()
	if err != nil {
		return err
	}

	log.Info(id)

	log.Info("done")

	return nil
}
