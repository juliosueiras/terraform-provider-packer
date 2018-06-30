
package builders

import "github.com/hashicorp/terraform/helper/schema"

func OracleClassicResource() *schema.Resource {
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

"communicator": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_host": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_port": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_private_key_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_pty": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssh_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_agent_auth": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssh_disable_agent_forwarding": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssh_handshake_attempts": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_bastion_host": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_bastion_port": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_bastion_agent_auth": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssh_bastion_username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_bastion_password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_bastion_private_key_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_file_transfer_method": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_proxy_host": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_proxy_port": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_proxy_username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_proxy_password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_keep_alive_interval": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_read_write_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"winrm_username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"winrm_password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"winrm_host": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"winrm_port": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"winrm_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"winrm_use_ssl": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"winrm_insecure": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"winrm_use_ntlm": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"identity_domain": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"api_endpoint": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"shape": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_image_list": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"snapshot_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"dest_image_list": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"attributes": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"attributes_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_description": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_source_list": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

