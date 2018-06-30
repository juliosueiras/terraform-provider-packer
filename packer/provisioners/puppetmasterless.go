
package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func PuppetMasterlessResource() *schema.Resource {
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

"clean_staging_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"guest_os_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"execute_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"extra_arguments": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"hiera_config_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ignore_exit_codes": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"module_paths": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"manifest_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"manifest_dir": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"prevent_sudo": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"puppet_bin_dir": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"staging_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"working_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

