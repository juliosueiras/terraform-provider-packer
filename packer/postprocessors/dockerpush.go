package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func DockerPushResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"only": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "For specifying a builder input",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"except": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "For excluding a builder input",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"keep_input_artifact": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"pipeline": &schema.Schema{
				Optional:    true,
				MaxItems:    1,
				Type:        schema.TypeList,
				Description: "Create a sequence definition",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"set": &schema.Schema{
							Required:    true,
							Type:        schema.TypeInt,
							Description: "The set(group) that is under",
						},
						"order": &schema.Schema{
							Required:    true,
							Type:        schema.TypeInt,
							Description: "The order to run in",
						},
					},
				},
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

			"login_username": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"login_password": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"login_server": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ecr_login": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"aws_access_key": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"aws_secret_key": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"aws_token": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"aws_profile": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}
