
package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func SaltMasterlessResource() *schema.Resource {
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

"skip_bootstrap": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"bootstrap_args": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"disable_sudo": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"custom_state": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"minion_config": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"grains_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"local_state_tree": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"local_pillar_roots": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"remote_state_tree": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"remote_pillar_roots": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"temp_config_dir": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"no_exit_on_failure": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"log_level": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"salt_call_args": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"salt_bin_dir": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"guest_os_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

