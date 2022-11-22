package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/rawdaGastan/terraform-provider-plugin/internal/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
