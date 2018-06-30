
package builders

import "github.com/hashicorp/terraform/helper/schema"

func CloudstackResource() *schema.Resource {
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

"http_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"http_port_min": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"http_port_max": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

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

"api_url": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"api_key": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"secret_key": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"async_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"http_get_only": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssl_no_verify": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"cidr_list": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"create_security_group": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"disk_offering": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"disk_size": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"expunge": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"hypervisor": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"instance_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"keypair": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"network": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"project": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"public_ip_address": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"security_groups": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"service_offering": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_iso": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_template": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"temporary_keypair_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"use_local_ip_address": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"user_data": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"user_data_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"zone": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"template_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"template_display_text": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"template_os": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"template_featured": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"template_public": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"template_password_enabled": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"template_requires_hvm": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"template_scalable": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"template_tag": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

