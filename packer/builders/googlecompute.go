
package builders

import "github.com/hashicorp/terraform/helper/schema"

func GoogleComputeResource() *schema.Resource {
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

"account_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"project_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"accelerator_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"accelerator_count": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"address": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"disable_default_service_account": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"disk_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"disk_size": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"disk_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_description": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_family": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_labels": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"image_licenses": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"instance_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"labels": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"machine_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"metadata": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"network": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"network_project_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"omit_external_ip": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"on_host_maintenance": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"preemptible": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"state_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"region": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"scopes": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"service_account_email": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_image": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_image_family": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"source_image_project_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"startup_script_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"subnetwork": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"tags": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"use_internal_ip": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"zone": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

