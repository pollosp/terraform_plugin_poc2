package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/pollosp/terraform_plugin_poc2/artifact"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: artifact.Provider})
}
