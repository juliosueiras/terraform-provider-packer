package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func ProfitbricksResource() *schema.Resource {
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
			"username": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "ProfitBricks username. This can be specified via environment variable `PROFITBRICKS_USERNAME', if provided. The value defined in the config has precedence over environemnt variable.",
			},

			"password": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "ProfitBricks password. This can be specified via environment variable `PROFITBRICKS_PASSWORD', if provided. The value defined in the config has precedence over environemnt variable.",
			},

			"url": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: `Endpoint for the ProfitBricks REST API. Default URL "https://api.profitbricks.com/rest/v2"`,
			},

			"location": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: `Defaults to "us/las".`,
			},

			"image": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "ProfitBricks volume image. Only Linux public images are supported. To obtain full list of available images you can use ProfitBricks CLI.",
			},

			"snapshot_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: `If snapshot name is not provided Packer will generate it`,
			},
			"snapshot_password": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: `Password for the snapshot`,
			},

			"disk_size": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: `Amount of disk space for this image in GB. Defaults to "50"`,
			},

			"disk_type": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: `Type of disk to use for this image. Defaults to "HDD".`,
			},

			"cores": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: `Amount of CPU cores to use for this build. Defaults to "4".`,
			},

			"ram": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: `Amount of RAM to use for this image. Defaults to "2048".`,
			},

			"retries": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: `Number of retries Packer will make status requests while waiting for the build to complete. Default value 120 seconds.`,
			},
		},
	}
}
