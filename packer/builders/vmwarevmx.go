package builders

import "github.com/hashicorp/terraform/helper/schema"

func VMWareVMXResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},

			"packer_build_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_builder_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_debug": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"packer_force": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"packer_on_error": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_user_variables": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},

			"http_directory": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"http_port_min": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"http_port_max": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"floppy_files": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"floppy_dirs": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"boot_wait": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"boot_command": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"disable_vnc": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"fusion_app_path": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"output_directory": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"headless": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"vnc_bind_address": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"vnc_port_min": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"vnc_port_max": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"vnc_disable_password": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"shutdown_command": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"shutdown_timeout": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"communicator": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_host": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_port": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"ssh_username": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_password": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_private_key_file": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_pty": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_timeout": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_agent_auth": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_disable_agent_forwarding": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_handshake_attempts": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"ssh_bastion_host": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_bastion_port": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"ssh_bastion_agent_auth": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_bastion_username": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_bastion_password": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_bastion_private_key_file": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_file_transfer_method": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_proxy_host": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_proxy_port": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"ssh_proxy_username": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_proxy_password": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_keep_alive_interval": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_read_write_timeout": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"winrm_username": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"winrm_password": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"winrm_host": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"winrm_port": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"winrm_timeout": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"winrm_use_ssl": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"winrm_insecure": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"winrm_use_ntlm": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_skip_request_pty": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_wait_timeout": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"tools_upload_flavor": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"tools_upload_path": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"vmx_data": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},

			"vmx_data_post": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},

			"vmx_remove_ethernet_interfaces": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"remote_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"skip_compaction": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"source_path": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"vm_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}
