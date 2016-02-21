package main

import (
	"fmt"
	"os"

	"github.com/containership/go-containership.cloud/client"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	buildCommit string
)

func main() {
	fmt.Printf("Drone ContainerShip Plugin built from %s\n", buildCommit)

	repo := drone.Repo{}
	vargs := Params{}

	plugin.Param("repo", &repo)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if vargs.Application == "" {
		vargs.Application = repo.Name
	}

	if vargs.Image == "" {
		vargs.Image = repo.FullName + ":latest"
	}

	if vargs.ApiKey == "" {
		fmt.Println("Error: ContainerShip Cloud API Key is required!")
		os.Exit(1)
	}

	if vargs.ClusterId == "" {
		fmt.Println("Error: ContainerShip Cloud Cluster ID is required!")
		os.Exit(1)
	}

	if vargs.Organization == "" {
		fmt.Println("Error: ContainerShip Cloud Organization is required!")
		os.Exit(1)
	}

	client := client.NewContainerShipCloudClient(vargs.Organization, vargs.ApiKey)

	response, err := client.GetApplication(vargs.ClusterId, vargs.Application)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	application := map[string]interface{}{
		"id":    vargs.Application,
		"image": vargs.Image,
	}

	switch response.StatusCode {
	case 404:
		create_response, err := client.CreateApplication(vargs.ClusterId, vargs.Application, application)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch create_response.StatusCode {
		case 201:
			fmt.Printf("Success: ContainerShip Cloud created %s\n", vargs.Application)
		default:
			fmt.Printf("Error: ContainerShip Cloud returned a %d response when creating %s\n", create_response.StatusCode, vargs.Application)
		}
	case 200:
		update_response, err := client.UpdateApplication(vargs.ClusterId, vargs.Application, application)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch update_response.StatusCode {
		case 200:
			fmt.Printf("Success: ContainerShip Cloud updated %s\n", vargs.Application)
		default:
			fmt.Printf("Error: ContainerShip Cloud returned a %d response when updating %s\n", update_response.StatusCode, vargs.Application)
		}
	default:
		fmt.Printf("Error: ContainerShip Cloud returned a %d response when fetching %s\n", response.StatusCode, vargs.Application)
		os.Exit(1)
	}
}
