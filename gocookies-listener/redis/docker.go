package redis

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"net"
	"strconv"
)

func CreateRedisContainer(port int) (string, error) {
	// Create Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", fmt.Errorf("failed to create Docker client: %v", err)
	}

	// Check if port is available
	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return "", fmt.Errorf("port %d is already in use", port)
	}
	err = l.Close()
	if err != nil {
		return "", err
	}

	// Create Redis container config
	ctx := context.Background()
	hostConfig := &container.HostConfig{}
	containerConfig := &container.Config{
		Image: "redis:alpine3.18",
	}

	// Configure port mapping
	portMap := make(nat.PortMap)

	portMap["6379/tcp"] = []nat.PortBinding{{
		HostIP:   "0.0.0.0",
		HostPort: strconv.Itoa(port),
	}}

	hostConfig.PortBindings = portMap

	// Create container
	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("failed to create Redis container: %v", err)
	}

	// Start container
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", fmt.Errorf("failed to start Redis container: %v", err)
	}

	// Return container ID
	return resp.ID, nil
}
