package provisioners

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func ChefClientResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The order for this provisioner to run in",
			},
			"chef_environment": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the chef_environment sent to the Chef server. By default this is empty and will not use an environment.",
			},

			"client_key": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to client key. If not set, this defaults to a file named client.pem in staging_directo",
			},

			"config_template": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to a template that will be used for the Chef configuration file. By default Packer only sets configuration it needs to match the settings set in the provisioner configuration. If you need to set configurations that the Packer provisioner doesn't support, then you should use a custom configuration template. See the dedicated \"Chef Configuration\" section below for more details.",
			},

			"encrypted_data_bag_secret_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to the file containing the secret for encrypted data bags. By default, this is empty, so no secret will be available.",
			},

			"execute_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command used to execute Chef. This has various configuration template variables available.",
			},

			"guest_os_type": &schema.Schema{
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"unix", "windows"}, true),
				Description:  "The target guest OS type, either \"unix\" or \"windows\". Setting this to \"windows\" will cause the provisioner to use Windows friendly paths and commands. By default, this is \"unix\".",
			},

			"install_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command used to install Chef. This has various configuration template variables available.",
			},
			"json": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "An arbitrary mapping of JSON that will be available as node attributes while running Chef.",
			},
			"knife_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command used to run Knife during node clean-up. This has various configuration template variables available",
			},

			"node_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the node to register with the Chef Server. This is optional and by default is packer-{{uuid}}.",
			},

			"policy_group": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of a policy group that exists on the Chef server. policy_name must also be specified.",
			},

			"policy_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of a policy, as identified by the name setting in a Policyfile.rb file. policy_group must also be specified.",
			},

			"prevent_sudo": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: " By default, the configured commands that are executed to install and run Chef are executed with sudo. If this is true, then the sudo will be omitted. This has no effect when guest_os_type is windows",
			},

			"run_list": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The run list for Chef. By default this is empty, and will use the run list sent down by the Chef Server.",
			},

			"server_url": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The URL to the Chef server. This is required.",
			},

			"skip_clean_client": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, Packer won't remove the client from the Chef server after it is done running. By default, this is false.",
			},

			"skip_clean_node": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, Packer won't remove the node from the Chef server after it is done running. By default, this is false.",
			},

			"skip_clean_staging_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, Packer won't remove the Chef staging directory from the machine after it is done running. By default, this is false.",
			},

			"skip_install": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, Chef will not automatically be installed on the machine using the Chef omnibus installers.",
			},

			"ssl_verify_mode": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Set to \"verify_none\" to skip validation of SSL certificates. If not set, this defaults to \"verify_peer\" which validates all SSL certifications.",
			},

			"trusted_certs_dir": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This is a directory that contains additional SSL certificates to trust. Any certificates in this directory will be added to whatever CA bundle ruby is using. Use this to add self-signed certs for your Chef Server or local HTTP file servers.",
			},

			"staging_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This is the directory where all the configuration of Chef by Packer will be placed. By default this is \"/tmp/packer-chef-client\" when guest_os_type unix and \"$env:TEMP/packer-chef-client\" when windows. This directory doesn't need to exist but must have proper permissions so that the user that Packer uses is able to create directories and write into this folder. By default the provisioner will create and chmod 0777 this directory.",
			},

			"validation_client_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Name of the validation client. If not set, this won't be set in the configuration and the default that Chef uses will be used.",
			},

			"validation_key_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to the validation key for communicating with the Chef Server. This will be uploaded to the remote machine. If this is NOT set, then it is your responsibility via other means (shell provisioner, etc.) to get a validation key to where Chef expects it.",
			},
		},
	}
}
