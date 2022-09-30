package main

import (
	"context"
	"log"

	"terraform-provider-scratch/scratch"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	// Example version string that can be overwritten by a release process
	version string = "dev"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/brendanthompson/scratch",
	}

	err := providerserver.Serve(context.Background(), scratch.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
