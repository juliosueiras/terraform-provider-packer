
package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func ChefSoloResource() *schema.Resource {
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

"chef_environment": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"config_template": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"cookbook_paths": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"roles_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"data_bags_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"encrypted_data_bag_secret_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"environments_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"execute_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"install_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"remote_cookbook_paths": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"prevent_sudo": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"run_list": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"skip_install": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"staging_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"guest_os_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"version": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

