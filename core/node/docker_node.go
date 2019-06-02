package node

import (
	"context"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

var dockerCli *client.Client
var mu sync.Mutex

// GetDockerCli get singleton docker client
func GetDockerCli() *client.Client {
	mu.Lock()
	defer mu.Unlock()

	if dockerCli == nil {
		var err error
		dockerCli, err = client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			panic(err)
		}
	}

	return dockerCli
}

// Node
type DockerNode struct {
	containerId  string
	networkId    string
	dockerClient *client.Client
}

func (node *DockerNode) Start() error {
	return node.dockerClient.ContainerStart(context.Background(), node.containerId, types.ContainerStartOptions{})
}

func (node *DockerNode) Stop() error {
	timeout := 3 * time.Second
	return node.dockerClient.ContainerStop(context.Background(), node.containerId, &timeout)
}

func (node *DockerNode) Pause() error {
	return node.dockerClient.ContainerPause(context.Background(), node.containerId)
}

func (node *DockerNode) Restart() error {
	timeout := 3 * time.Second
	return node.dockerClient.ContainerRestart(context.Background(), node.containerId, &timeout)
}

func (node *DockerNode) Connect() error {
	return node.dockerClient.NetworkConnect(context.Background(), node.networkId, node.containerId, &network.EndpointSettings{})
}

func (node *DockerNode) Disconnect() error {
	return node.dockerClient.NetworkDisconnect(context.Background(), node.networkId, node.containerId, true)
}

func (node *DockerNode) Update(r NodeResource) error {
	_, err := node.dockerClient.ContainerUpdate(context.Background(), node.containerId, container.UpdateConfig{
		Resources: container.Resources{
			Memory:   r.memory,
			CPUCount: r.cpu,
		},
	})
	return err
}

func NewDockerNode(containerId string, networkId string) *DockerNode {
	return &DockerNode{
		containerId:  containerId,
		networkId:    networkId,
		dockerClient: GetDockerCli(),
	}
}
