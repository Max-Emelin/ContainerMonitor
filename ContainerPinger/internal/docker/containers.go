package docker

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func GetContainerIPs() ([]string, error) {
	cli, err := client.NewClientWithOpts(client.WithHost("unix:///var/run/docker.sock"))
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var ips []string
	for _, container := range containers {
		containerJSON, err := cli.ContainerInspect(context.Background(), container.ID)
		if err != nil {
			log.Printf("Error inspecting container %s: %s", container.ID, err.Error())
			continue
		}

		if containerJSON.NetworkSettings != nil {
			for _, network := range containerJSON.NetworkSettings.Networks {
				ipAddress := network.IPAddress
				if ipAddress != "" {
					ips = append(ips, ipAddress)
				}
			}
		}
	}

	return ips, nil
}
