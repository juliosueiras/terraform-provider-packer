package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func FileResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
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
			"source": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"destination": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"direction": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"generated": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}
