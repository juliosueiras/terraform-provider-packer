
package builders

import "github.com/hashicorp/terraform/helper/schema"

func OpenstackResource() *schema.Resource {
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

"username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"user_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"identity_endpoint": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"tenant_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"tenant_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"domain_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"domain_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"insecure": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"region": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"endpoint_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"cacert": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"cert": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"key": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"token": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"cloud": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"metadata": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"image_visibility": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_members": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

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

"ssh_keypair_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"temporary_key_pair_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_interface": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_ip_version": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_image": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_image_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"flavor": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"availability_zone": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"rackconnect_wait": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"floating_ip_pool": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"floating_ip": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"reuse_ips": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"security_groups": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"networks": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"user_data": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"user_data_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"instance_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"instance_metadata": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"config_drive": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"openstack_provider": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"use_floating_ip": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

	},
	}
}

