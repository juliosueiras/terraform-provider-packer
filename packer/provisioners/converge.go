
package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func ConvergeResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type: schema.TypeInt,
				Required: true,
			},
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

"bootstrap": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"version": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"bootstrap_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"prevent_bootstrap_sudo": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"module_dirs": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
			},

		},

"module": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"working_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"params": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"execute_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"prevent_sudo": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

	},
	}
}

