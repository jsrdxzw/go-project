package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"time"
)

func main() {
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
				"27017/tcp": {},
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
	time.Sleep(20 * time.Second)
	err = c.ContainerRemove(ctx, res.ID, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		panic(err)
	}
}
