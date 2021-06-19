package mongotesting

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"testing"
)

func RunWithMongoInDocker(m *testing.M, mongoURI *string) int {
	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	res, err := c.ContainerCreate(
		ctx,
		&container.Config{
			Image: "mongo:latest",
			ExposedPorts: nat.PortSet{
				"": {},
			},
		},
		&container.HostConfig{
			PortBindings: nat.PortMap{
				"27017/tcp": []nat.PortBinding{
					{
						HostIP:   "127.0.0.1",
						HostPort: "0", // 自动选择端口
					},
				},
			},
		}, nil, nil, "",
	)
	if err != nil {
		panic(err)
	}

	err = c.ContainerStart(ctx, res.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	defer func() {
		err = c.ContainerRemove(ctx, res.ID, types.ContainerRemoveOptions{Force: true})
		if err != nil {
			panic(err)
		}
	}()

	inspect, err := c.ContainerInspect(ctx, res.ID)
	if err != nil {
		panic(err)
	}
	hostPort := inspect.NetworkSettings.Ports["27017/tcp"][0]
	*mongoURI = fmt.Sprintf("mongodb://%s:%s", hostPort.HostIP, hostPort.HostPort)
	return m.Run()
}
