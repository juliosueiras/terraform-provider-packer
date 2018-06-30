
package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func VagrantResource() *schema.Resource {
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

"compression_level": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"include": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"output": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vagrantfile_template": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

