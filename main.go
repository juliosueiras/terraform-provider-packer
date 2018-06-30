package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/juliosueiras/terraform-provider-packer/packer"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: packer.Provider,
	})
}
