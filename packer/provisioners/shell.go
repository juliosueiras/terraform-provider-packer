
package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func ShellResource() *schema.Resource {
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

"inline_shebang": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"environment_vars": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"remote_folder": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"remote_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"remote_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"execute_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"start_retry_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"skip_clean": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"expect_disconnect": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

	},
	}
}

