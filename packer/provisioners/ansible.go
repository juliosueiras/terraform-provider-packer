
package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func AnsibleResource() *schema.Resource {
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

"extra_arguments": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"ansible_env_vars": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"playbook_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"groups": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"empty_groups": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"host_alias": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"user": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"local_port": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_host_key_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_authorized_key_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"sftp_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"skip_version_check": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"use_sftp": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"inventory_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"inventory_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

