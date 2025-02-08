package docker

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func GetContainerIPs() ([]string, error) {
	//client.WithHost("tcp://localhost:2375") -- Error getting IP containers: Cannot connect to the Docker daemon at tcp://localhost:2375. Is the docker daemon running?
	//client.WithHost("npipe:////./pipe/docker_engine")      -- Error getting IP containers: protocol not available
	//client.WithHost("npipe:////./pipe/dockerDesktopLinuxEngine") -- Error getting IP containers: protocol not available
	//(client.WithHost("npipe:////./pipe/dockerDesktopWindowsEngine") -- Error getting IP containers: protocol not available

	cli, err := client.NewClientWithOpts(client.WithHost("npipe:////./pipe/dockerDesktopWindowsEngine"))
	if err != nil {
		return nil, err
	}

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
