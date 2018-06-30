
package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func VSphereTemplateResource() *schema.Resource {
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

"host": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"insecure": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"datacenter": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"folder": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

