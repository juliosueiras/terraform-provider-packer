
package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func DockerPushResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
"packer_build_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"packer_builder_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"packer_debug": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"packer_force": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"packer_on_error": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"packer_user_variables": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"login_username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"login_password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"login_server": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ecr_login": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"aws_access_key": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"aws_secret_key": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"aws_token": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"aws_profile": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

