package communicators

import "github.com/hashicorp/terraform/helper/schema"

func SSHCommunicatorResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ssh_username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username to connect to SSH with. Required if using SSH.",
			},
			"ssh_password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "A plaintext password to use to authenticate with SSH.",
			},
			"ssh_agent_auth": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, the local SSH agent will be used to authenticate connections to the remote host. Defaults to false.",
			},
			"ssh_bastion_agent_auth": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, the local SSH agent will be used to authenticate with the bastion host. Defaults to false.",
			},
			"ssh_bastion_host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A bastion host to use for the actual SSH connection.",
			},
			"ssh_bastion_password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password to use to authenticate with the bastion host.",
			},
			"ssh_bastion_port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port of the bastion host. Defaults to 1.",
			},
			"ssh_bastion_private_key_file": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A private key file to use to authenticate with the bastion host.",
			},
			"ssh_bastion_username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The username to connect to the bastion host.",
			},
			"ssh_disable_agent_forwarding": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, SSH agent forwarding will be disabled. Defaults to false.",
			},
			"ssh_file_transfer_method": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "How to transfer files, Secure copy (default) or SSH File Transfer Protocol.",
			},
			"ssh_handshake_attempts": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of handshakes to attempt with SSH once it can connect. This defaults to 10.",
			},
			"ssh_host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The address to SSH to. This usually is automatically configured by the builder.",
			},
			"ssh_keep_alive_interval": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "How often to send \"keep alive\" messages to the server. Set to a negative value (-1s) to disable. Example value: 10s. Defaults to 5s.",
			},
			"ssh_port": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port to connect to SSH. This defaults to 22.",
			},
			"ssh_private_key_file": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to a PEM encoded private key file to use to authenticate with SSH.",
			},
			"ssh_proxy_host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A SOCKS proxy host to use for SSH connection",
			},
			"ssh_proxy_password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password to use to authenticate with the proxy server. Optional.",
			},
			"ssh_proxy_port": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "A port of the SOCKS proxy. Defaults to 1080.",
			},
			"ssh_proxy_username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The username to authenticate with the proxy server. Optional.",
			},
			"ssh_pty": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, a PTY will be requested for the SSH connection. This defaults to false.",
			},
			"ssh_read_write_timeout": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The amount of time to wait for a remote command to end. This might be useful if, for example, packer hangs on a connection after a reboot. Example: 5m. Disabled by default.",
			},
			"ssh_timeout": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The time to wait for SSH to become available. Packer uses this to determine when the machine has booted so this is usually quite long. Example value: 10m.",
			},
		},
	}
}

func WinRMCommunicatorResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"winrm_host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The address for WinRM to connect to.",
			},
			"winrm_insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, do not check server certificate chain and host name.",
			},
			"winrm_password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The password to use to connect to WinRM.",
			},
			"winrm_port": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The WinRM port to connect to. This defaults to 5985 for plain unencrypted connection and 5986 for SSL when winrm_use_ssl is set to true.",
			},
			"winrm_timeout": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The amount of time to wait for WinRM to become available. This defaults to 30m since setting up a Windows machine generally takes a long time.",
			},
			"winrm_use_ntlm": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, NTLM authentication will be used for WinRM, rather than default (basic authentication), removing the requirement for basic authentication to be enabled within the target guest. Further reading for remote connection authentication can be found here.",
			},
			"winrm_use_ssl": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, use HTTPS for WinRM.",
			},
			"winrm_username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The username to use to connect to WinRM.",
			},
		},
	}
}
