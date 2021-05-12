package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/thermeon/terraform-provider-bigv/bigv"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bigv.Provider,
	})
}
