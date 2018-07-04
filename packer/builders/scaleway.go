package builders

import "github.com/hashicorp/terraform/helper/schema"

func ScalewayResource() *schema.Resource {
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

			"api_token": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"api_access_key": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"region": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"commercial_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"snapshot_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"server_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}
