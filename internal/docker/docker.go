package docker

import (
	"context"
	"strings"

	"github.com/moby/moby/client"
)

type Service struct {
	ID      string
	Name    string
	Image   string
	State   string
	Health  string
	Link    string
	IconUrl string
}

func ListServices() ([]Service, error) {
	ctx := context.Background()
	cli, err := client.New(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	filterArgs := client.Filters{}
	filterArgs.Add("label", "dashboard.enable=true")

	containers, err := cli.ContainerList(ctx, client.ContainerListOptions{Filters: filterArgs})
	if err != nil {
		return nil, err
	}

	services := make([]Service, 0, len(containers.Items))

	for _, ctr := range containers.Items {
		info, err := cli.ContainerInspect(ctx, ctr.ID, client.ContainerInspectOptions{})
		if err != nil {
			return nil, err
		}
		service := Service{
			ID:     ctr.ID,
			Image:  ctr.Image,
			State:  string(info.Container.State.Status),
			Health: "none",
		}
		if len(ctr.Names) > 0 {
			service.Name = strings.TrimPrefix(ctr.Names[0], "/")
		}

		if info.Container.State.Health != nil {
			service.Health = string(info.Container.State.Health.Status)
		}

		services = append(services, service)
	}
	return services, nil
}
