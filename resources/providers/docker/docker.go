package docker

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/target"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
)

// Config holds allowed values for an implemented
// resource provider. Any value outside of this config
// is unable to be modified during an experiment
type Config struct {
	Target target.Target

	// Name is the name of the container
	Name string

	// Version is the version of the image or container
	Version string

	// Image is hte image name of the docker image
	Image string

	// Tag is the docker tag, defaults to :latest
	Tag string

	// ID is the id of the running container
	ID string
}

// Provider implements an docker Provider
type Provider struct {
	// Config holds our internal configuration options
	// for the instance of the Provider
	Config Config

	// RemoteConfig identifies the remote config
	RemoteConfig Config
}

// NewProvider initializes a Provider
// with defaults
func NewProvider(req *pb.CreateRequest) providers.Provider {
	return &Provider{Config: Config{
		Name:    req.Config.Name,
		Version: req.Config.Version,
		Tag:     "latest",
	}}
}

func NewProviderWithConfig(cfg Config) providers.Provider {
	return &Provider{Config: cfg}
}

// NewClient returns a new docker client
func NewClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return cli, nil
}

// Pull will pull an image from a registry, if the image
// is already pulled, it'll verify this and just return.
// TODO: add a 'force pull' and support options
func (p *Provider) Read() error {
	// check if already on machine here, pull if not, support
	// force pulling i think
	image, err := p.TaggedImage()

	if err != nil {
		return err
	}

	if image != nil {
		return nil
	}

	cli, err := NewClient()

	if err != nil {
		return err
	}

	options := types.ImagePullOptions{}

	out, err := cli.ImagePull(
		context.Background(),
		p.ImageName(),
		options,
	)

	if err != nil {
		return err
	}

	defer out.Close()

	if _, err := ioutil.ReadAll(out); err != nil {
		return err
	}

	return nil
}

func (p *Provider) images() ([]types.ImageSummary, error) {
	cli, err := NewClient()

	if err != nil {
		return []types.ImageSummary{}, err
	}

	return cli.ImageList(
		context.Background(),
		types.ImageListOptions{},
	)
}

// Select lists all pulled images of the repo
func (p *Provider) Select() (target.Selection, error) {
	filtered := target.Selection{}

	list, err := p.images()

	if err != nil {
		return filtered, err
	}

	for _, image := range list {
		for _, name := range image.RepoTags {
			if strings.Contains(name, p.Config.Image) {
				filtered.Data = append(filtered.Data, target.Resource{
					ID: image.ID,
				})
				break
			}
		}
	}

	return filtered, nil
}

// TaggedImage list any image that matches the image and version/tag
func (p *Provider) TaggedImage() (*types.ImageSummary, error) {
	list, err := p.images()

	if err != nil {
		return nil, err
	}

	for _, image := range list {
		for _, name := range image.RepoTags {
			if strings.Contains(name, p.ImageName()) {
				return &image, nil
			}
		}
	}

	return nil, nil
}

// Running returns a list of all running versions of a resource
func (p *Provider) Running() ([]types.Container, error) {
	cli, err := NewClient()

	if err != nil {
		return []types.Container{}, err
	}

	containers, err := cli.ContainerList(
		context.Background(),
		types.ContainerListOptions{},
	)

	if err != nil {
		return []types.Container{}, err
	}
	return containers, nil
}

// Start a specific resource
func (p *Provider) Start() (string, error) {
	cli, err := NewClient()

	if err != nil {
		return "", err
	}

	conf := container.Config{
		// Cmd:   []string{},
		Image: p.ImageName(),
	}
	hostConfig := container.HostConfig{}
	netConfig := network.NetworkingConfig{}

	c, err := cli.ContainerCreate(
		context.Background(),
		&conf,
		&hostConfig,
		&netConfig,
		nil,
		p.UniqueContainerName(),
	)

	if err != nil {
		return "", err
	}

	err = cli.ContainerStart(
		context.Background(),
		c.ID,
		types.ContainerStartOptions{},
	)

	if err != nil {
		return "", err
	}

	p.RemoteConfig.ID = c.ID

	return c.ID, nil
}

// Clone creates a modified variant
func (p *Provider) Clone() error {
	// launch container with fresh name/id

	return nil
}

// Cancel will abort and running or submitted Provider
func (p *Provider) Cancel() error { return nil }

// Stop will stop any running Provider
func (p *Provider) Stop() error { return nil }

// ProbeReadiness checks that the provisioned resource is available and
// ready to be included in a live experiment
func (p *Provider) ProbeReadiness() (bool, error) { return false, nil }

// AwaitReadiness should be implemented to detect
// when a Provider has finished setting up a variant
// and can begin using it in an experiment
func (p *Provider) AwaitReadiness() chan error { return make(chan error) }

// Teardown eradicates any resource that has been
// provisioned as part of a variant
func (p *Provider) Teardown() error { return nil }

// Annotate should implement applying labels or tags for a given resource type
func (p *Provider) Annotate(id string, l labels.Labels) error { return nil }

// General Helpers

// ImageName is the complete resource identifier
// to pull an image from a registry by name, and tag
func (p *Provider) ImageName() string {
	return fmt.Sprintf("%s:%s", p.Config.Image, p.Config.Tag)
}

// UniqueContainerName creates a unique name
// for running the image as a container
func (p *Provider) UniqueContainerName() string {
	return fmt.Sprintf("%d-%s", time.Now().Unix(), p.Config.Name)
}
