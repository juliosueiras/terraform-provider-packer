
package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func CompressResource() *schema.Resource {
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

"output": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"format": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"compression_level": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"keep_input_artifact": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

	},
	}
}
