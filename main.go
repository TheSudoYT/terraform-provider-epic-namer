package main

import (
	"terraform-provider-epic-namer/epic"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: epic.Provider,
	})
}
