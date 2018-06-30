
package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func ChefClientResource() *schema.Resource {
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

"client_key": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"config_template": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"encrypted_data_bag_secret_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"execute_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"guest_os_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"install_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"knife_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"node_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"policy_group": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"policy_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

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

"server_url": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"skip_clean_client": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"skip_clean_node": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"skip_clean_staging_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"skip_install": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssl_verify_mode": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"trusted_certs_dir": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"staging_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"validation_client_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"validation_key_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

