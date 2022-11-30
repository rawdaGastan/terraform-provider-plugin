package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/rawdaGastan/terraform-provider-plugin/internal/provider"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
