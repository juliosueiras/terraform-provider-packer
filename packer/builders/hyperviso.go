
package builders

import "github.com/hashicorp/terraform/helper/schema"

func HyperVISOResource() *schema.Resource {
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

"iso_checksum": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"iso_checksum_url": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"iso_checksum_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"iso_urls": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"iso_target_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"iso_target_extension": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"iso_url": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"floppy_files": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"floppy_dirs": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"boot_wait": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"boot_command": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"output_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

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

"shutdown_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"shutdown_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"disk_size": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"disk_block_size": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ram_size": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"secondary_iso_images": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"guest_additions_mode": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"guest_additions_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vm_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"switch_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"switch_vlan_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"mac_address": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vlan_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"cpu": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"generation": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"enable_mac_spoofing": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"enable_dynamic_memory": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"enable_secure_boot": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"secure_boot_template": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"enable_virtualization_extensions": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"temp_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vhd_temp_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"disk_additional_size": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"skip_compaction": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"skip_export": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"differencing_disk": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"use_fixed_vhd_format": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"headless": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

	},
	}
}

