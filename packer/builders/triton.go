
package builders

import "github.com/hashicorp/terraform/helper/schema"

func TritonResource() *schema.Resource {
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

"triton_url": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"triton_account": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"triton_user": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"triton_key_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"triton_key_material": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"insecure_skip_tls_verify": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"source_machine_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_machine_package": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_machine_image": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_machine_networks": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"source_machine_metadata": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"source_machine_tags": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"source_machine_firewall_enabled": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"most_recent": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"image_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_version": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_description": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_homepage": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_eula_url": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_acls": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"image_tags": &schema.Schema{
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
			Type: schema.TypeInt,

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
			Type: schema.TypeInt,

		},

"ssh_read_write_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

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
			Type: schema.TypeInt,

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

	},
	}
}

