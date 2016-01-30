package main

import (
	"fmt"
	"os"

	"github.com/containership/go-containership.cloud/client"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone ContainerShip Plugin built at %s\n", buildDate)

	repo := drone.Repo{}
	vargs := Params{}

	plugin.Param("repo", &repo)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if vargs.Application == "" {
		vargs.Application = repo.Name
	}

	if vargs.ApiKey == "" {
		fmt.Println("Error: ContainerShip Cloud API Key is required!")
		os.Exit(1)
	}

	if vargs.ClusterId == "" {
		fmt.Println("Error: ContainerShip Cloud Cluster ID is required!")
		os.Exit(1)
	}

	if vargs.Image == "" {
		fmt.Println("Error: ContainerShip Application Image is required!")
		os.Exit(1)
	}

	if vargs.Organization == "" {
		fmt.Println("Error: ContainerShip Cloud Organization is required!")
		os.Exit(1)
	}

	ContainerShipCloudClient := client.NewContainerShipCloudClient(vargs.Organization, vargs.ApiKey)

	response, err := ContainerShipCloudClient.GetApplication(vargs.ClusterId, vargs.Application)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	application := map[string]interface{}{
		"id":    vargs.Application,
		"image": vargs.Image,
	}

	if response.StatusCode == 404 {
		create_response, err := ContainerShipCloudClient.CreateApplication(vargs.ClusterId, vargs.Application, application)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if create_response.StatusCode != 201 {
			fmt.Printf("Error: ContainerShip Cloud returned a %s response when creating %s\n", response.StatusCode, vargs.Application)
		} else {
			fmt.Printf("Success: ContainerShip Cloud created %s\n", vargs.Application)
		}
	} else if response.StatusCode == 200 {
		update_response, err := ContainerShipCloudClient.UpdateApplication(vargs.ClusterId, vargs.Application, application)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if update_response.StatusCode != 200 {
			fmt.Printf("Error: ContainerShip Cloud returned a %s response when updating %s\n", response.StatusCode, vargs.Application)
		} else {
			fmt.Printf("Success: ContainerShip Cloud updated %s\n", vargs.Application)
		}
	} else {
		fmt.Printf("Error: ContainerShip Cloud returned a %s response when fetching %s\n", response.StatusCode, vargs.Application)
		os.Exit(1)
	}
}
