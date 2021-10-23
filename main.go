package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"

	"terraform-provider-scratch/scratch"
)

func main() {
	tfsdk.Serve(context.Background(), scratch.New, tfsdk.ServeOpts{
		Name: "scratch",
	})
}
