package droplet

import (
	"context"
	"fmt"
	"time"

	"github.com/digitalocean/godo"
	"github.com/edobtc/cloudkit/events/publishers/aws/sns"
	"github.com/sirupsen/logrus"
)

func Apply(name string) {
	ctx := context.TODO()

	images, err := FilterImages(ctx, "15.5")
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Infof("found version: %s", images[0].Name)

	key, err := CloudKitSSHKey()

	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Infof("found key: %s", key.Fingerprint)

	dp, err := CreateDroplet(ctx, &Config{
		Name:    name,
		ImageID: images[0].ID,
		SSHKey:  key.Fingerprint,
	})

	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Infof("creating droplet %d", dp.ID)

	droplet := <-Await(dp.ID)

	if droplet.Status == "active" {
		broadcast(droplet)
		return
	}

	logrus.Error(droplet.Status)
}

func Await(id int) chan *godo.Droplet {
	ch := make(chan *godo.Droplet)

	go func(c chan *godo.Droplet) {
		client := NewClient()
		defer close(c)

		for {
			time.Sleep(5 * time.Second)

			logrus.Info("awaiting droplet status...")

			dp, _, err := client.Droplets.Get(context.TODO(), id)
			if err != nil {
				logrus.Error(err)
			}

			if dp.Status == "active" {
				logrus.Info("droplet is active")
				c <- dp
				return
			}
		}
	}(ch)

	return ch
}

func broadcast(droplet *godo.Droplet) {
	fmt.Println(droplet.PublicIPv4())
	fmt.Println(droplet.PrivateIPv4())

	publisher := sns.NewPublisher()
	ip, _ := droplet.PublicIPv4()
	temp := fmt.Sprintf("new droplet IP %s", ip)
	err := publisher.Send([]byte(temp))
	if err != nil {
		logrus.Error(err)
	}
}
