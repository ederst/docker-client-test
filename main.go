package main

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/docker/docker/api/types"
	docker "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"

	log "github.com/sirupsen/logrus"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ctx := context.Background()
	name := "my-container"

	containerConfig := docker.Config{}
	hostConfig := docker.HostConfig{}
	networkingConfig := network.NetworkingConfig{}

	containerConfig.Cmd = []string{"env"}
	containerConfig.Image = "alpine"

	log.Infof("Create container '%s'", name)
	_, err = cli.ContainerCreate(ctx, &containerConfig, &hostConfig, &networkingConfig, nil, name)
	if err != nil {
		log.Errorf("Failed to create container '%s': %+v", name, err)
	}

	log.Infof("Start container '%s'", name)
	err = cli.ContainerStart(ctx, name, types.ContainerStartOptions{})
	if err != nil {
		log.Errorf("Failed to start container '%s': %+v", name, err)
	}

	log.Infof("Get logs of container '%s'", name)
	body, err := cli.ContainerLogs(ctx, name, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	defer body.Close()

	if err != nil {
		log.Errorf("Failed to read log from container '%s': %+v", name, err)
	}

	content, err := ioutil.ReadAll(body)

	log.Infof("Container '%s' logs:\n%s", name, string(content))

	log.Infof("Wait for container '%s' to stop", name)
	time.Sleep(1 * time.Second)

	log.Infof("Remove container '%s'", name)
	err = cli.ContainerRemove(ctx, name, types.ContainerRemoveOptions{})

	if err != nil {
		log.Errorf("Failed to remove container '%s': %+v", name, err)
	}
}
