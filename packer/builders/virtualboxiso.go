
package builders

import "github.com/hashicorp/terraform/helper/schema"

func VirtualboxISOResource() *schema.Resource {
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

"format": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"export_opts": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"output_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"headless": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"vrdp_bind_address": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vrdp_port_min": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"vrdp_port_max": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"shutdown_command": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"shutdown_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"post_shutdown_delay": &schema.Schema{
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

"ssh_host_port_min": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_host_port_max": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_skip_nat_mapping": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssh_wait_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"vboxmanage": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"vboxmanage_post": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"virtualbox_version_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"disk_size": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"guest_additions_mode": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"guest_additions_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"guest_additions_sha256": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"guest_additions_url": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"guest_os_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"hard_drive_discard": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"hard_drive_interface": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"sata_port_count": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"hard_drive_nonrotational": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"iso_interface": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"keep_registered": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"skip_export": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"vm_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

