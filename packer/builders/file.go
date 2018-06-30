
package builders

import "github.com/hashicorp/terraform/helper/schema"

func FileResource() *schema.Resource {
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

"source": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"target": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"content": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

