package events

import (
	"fmt"
	"github.com/docker/docker/client"
	"os"
)

const (
	defaultApiVersion = "1.22"
)

func NewDockerClient() (*client.Client, error) {
	host := fmt.Sprintf("tcp://%v:2375", os.Getenv("DEFAULT_GATEWAY"))
	return client.NewClient(host, defaultApiVersion, nil, nil)
}
