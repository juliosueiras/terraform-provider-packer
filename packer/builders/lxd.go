package builders

import "github.com/hashicorp/terraform/helper/schema"

func LXDResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},

			"packer_build_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_builder_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_debug": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"packer_force": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"packer_on_error": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_user_variables": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},

			"output_image": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"container_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"command_wrapper": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"profile": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"init_sleep": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"publish_properties": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},

			"launch_config": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},
		},
	}
}
