package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func DigitalOceanResource() *schema.Resource {
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
			"api_token": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The client TOKEN to use to access your account. It can also be specified via environment variable DIGITALOCEAN_API_TOKEN, if set.",
			},

			"api_url": &schema.Schema{
				Optional:    true,
				Description: "Non standard api endpoint URL. Set this if you are using a DigitalOcean API compatible service. It can also be specified via environment variable DIGITALOCEAN_API_URL.",
				Type:        schema.TypeString,
			},

			"region": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The name (or slug) of the region to launch the droplet in. Consequently, this is the region where the snapshot will be available. See https://developers.digitalocean.com/documentation/v2/#list-all-regions for the accepted region names/slugs.",
			},

			"size": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The name (or slug) of the droplet size to use. See https://developers.digitalocean.com/documentation/v2/#list-all-sizes for the accepted size names/slugs.",
			},

			"image": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The name (or slug) of the base image to use. This is the image that will be used to launch a new droplet and provision it. See https://developers.digitalocean.com/documentation/v2/#list-all-images for details on how to get a list of the accepted image names/slugs.",
			},

			"private_networking": &schema.Schema{
				Optional:    true,
				Description: "Set to true to enable private networking for the droplet being created. This defaults to false, or not enabled.",
				Type:        schema.TypeBool,
			},

			"monitoring": &schema.Schema{
				Optional:    true,
				Description: "Set to true to enable monitoring for the droplet being created. This defaults to false, or not enabled.",
				Type:        schema.TypeBool,
			},

			"ipv6": &schema.Schema{
				Optional:    true,
				Description: "Set to true to enable ipv6 for the droplet being created. This defaults to false, or not enabled.",
				Type:        schema.TypeBool,
			},

			"snapshot_name": &schema.Schema{
				Optional:    true,
				Description: "The name of the resulting snapshot that will appear in your account. This must be unique. To help make this unique, use a function like timestamp (see configuration templates for more info)",
				Type:        schema.TypeString,
			},

			"snapshot_regions": &schema.Schema{
				Optional:    true,
				Description: "The regions of the resulting snapshot that will appear in your account.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"state_timeout": &schema.Schema{
				Optional:    true,
				Description: "The time to wait, as a duration string, for a droplet to enter a desired state (such as \"active\") before timing out. The default state timeout is \"6m\".",
				Type:        schema.TypeString,
			},

			"droplet_name": &schema.Schema{
				Optional:    true,
				Description: "The name assigned to the droplet. DigitalOcean sets the hostname of the machine to this value.",
				Type:        schema.TypeString,
			},

			"user_data": &schema.Schema{
				Optional:    true,
				Description: "User data to launch with the Droplet.",
				Type:        schema.TypeString,
			},

			"user_data_file": &schema.Schema{
				Optional:    true,
				Description: "Path to a file that will be used for the user data when launching the Droplet.",
				Type:        schema.TypeString,
			},
		},
	}
}
