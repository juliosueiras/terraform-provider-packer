package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func OneandoneResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},
			"communicator": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"ssh": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     communicators.SSHCommunicatorResource(),
			},
			"winrm": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     communicators.WinRMCommunicatorResource(),
			},
			"token": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "1&1 REST API Token. This can be specified via environment variable ONEANDONE_TOKEN",
			},

			"url": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: `Endpoint for the 1&1 REST API. Default URL "https://cloudpanel-api.1and1.com/v1"`,
			},

			"image_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: `Resulting image. If "image_name" is not provided Packer will generate it`,
			},

			"data_center_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: `Name of virtual data center. Possible values "ES", "US", "GB", "DE". Default value "US"`,
			},

			"source_image_name": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "1&1 Server Appliance name of type IMAGE.",
			},

			"disk_size": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: `Amount of disk space for this image in GB. Defaults to "50"`,
			},

			"retries": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: `Number of retries Packer will make status requests while waiting for the build to complete. Default value "600".`,
			},
		},
	}
}
