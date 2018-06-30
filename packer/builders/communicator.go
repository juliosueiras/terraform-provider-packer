package builders

import "github.com/hashicorp/terraform/helper/schema"

func sshCommunicatorResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ssh_username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ssh_password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"ssh_agent_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ssh_bastion_agent_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ssh_bastion_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_bastion_password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"ssh_bastion_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_bastion_private_key_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_bastion_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_disable_agent_forwarding": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ssh_file_transfer_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_handshake_attempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssh_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_keep_alive_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssh_private_key_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_proxy_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_proxy_password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"ssh_proxy_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssh_proxy_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_pty": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ssh_read_write_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func winrmCommunicatorResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"winrm_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"winrm_insecure": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"winrm_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"winrm_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"winrm_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"winrm_use_ntlm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"winrm_use_ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"winrm_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
