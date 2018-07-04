package builders

import "github.com/hashicorp/terraform/helper/schema"

func LXCResource() *schema.Resource {
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

			"config_file": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"output_directory": &schema.Schema{
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

			"init_timeout": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"create_options": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"start_options": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"attach_options": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"template_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"template_parameters": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"template_environment_vars": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"target_runlevel": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},
		},
	}
}
